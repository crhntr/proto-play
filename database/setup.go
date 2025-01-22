package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	TransactionFunc     func(q Querier) error
	WrapTransactionFunc func(ctx context.Context, options pgx.TxOptions, f TransactionFunc) error
)

func TransactionWrapper(pool *pgxpool.Pool) WrapTransactionFunc {
	return func(ctx context.Context, o pgx.TxOptions, f TransactionFunc) error {
		tx, err := pool.BeginTx(ctx, o)
		if err != nil {
			return err
		}
		defer func() { _ = tx.Rollback(ctx) }()

		q := New(tx)

		if err := f(q); err != nil {
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return err
		}
		return nil
	}
}

func Run(ctx context.Context, run func(ctx context.Context, pool *pgxpool.Pool) error) error {
	dbName, ok := os.LookupEnv("PGDATABASE")
	if !ok || dbName == "" {
		return errors.New("PGDATABASE environment variable not set")
	}
	dbURL, err := databaseURLFromEnv(dbName)
	if err != nil {
		return err
	}
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return err
	}
	if err := pool.Ping(ctx); err != nil {
		return err
	}
	defer pool.Close()
	return run(ctx, pool)
}

func databaseURLFromEnv(name string) (string, error) {
	user, ok := os.LookupEnv("PGUSER")
	if !ok || user == "" {
		return "", errors.New("PGUSER environment variable not set")
	}
	host, ok := os.LookupEnv("PGHOST")
	if !ok || host == "" {
		return "", errors.New("pgHost environment variable not set")
	}
	password, ok := os.LookupEnv("PGPASSWORD")
	if !ok || password == "" {
		return "", errors.New("PGPASSWORD environment variable not set")
	}
	port, ok := os.LookupEnv("PGPORT")
	if !ok || port == "" {
		return "", errors.New("PGPORT environment variable not set")
	}
	query, ok := os.LookupEnv("PGQUERYSTRING")
	if !ok || query == "" {
		return "", errors.New("PGQUERYSTRING environment variable not set")
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", user, password, host, port, name, query), nil
}
