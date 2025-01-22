package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	dbName, ok := os.LookupEnv("PGDATABASE")
	if !ok || dbName == "" {
		panic(errors.New("PGDATABASE environment variable not set"))
	}
	pool, err := createDatabase(ctx, dbName)
	if err != nil {
		panic(err)
	}
	defer pool.Close()
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

func createDatabase(ctx context.Context, name string) (*pgxpool.Pool, error) {
	pgURL, err := databaseURLFromEnv("postgres")
	if err != nil {
		return nil, err
	}
	pgPool, err := pgxpool.New(ctx, pgURL)
	if err != nil {
		return nil, err
	}
	if err := pgPool.Ping(ctx); err != nil {
		return nil, err
	}
	_, _ = pgPool.Exec(ctx, fmt.Sprintf(`DROP DATABASE %s`, name))
	if _, err := pgPool.Exec(ctx, fmt.Sprintf(`CREATE DATABASE %s`, name)); err != nil {
		return nil, err
	}
	dbURL, err := databaseURLFromEnv(name)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	return pool, nil
}
