package main

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"

	play "github.com/crhntr/proto-play"
	"github.com/crhntr/proto-play/api/playground/v1/v1connect"
	"github.com/crhntr/proto-play/database"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := runWithDB(ctx, func(ctx context.Context, pool *pgxpool.Pool) error {
		srv := play.New(slog.Default(), func(ctx context.Context, f func(q database.Querier) error) error {
			tx, err := pool.Begin(ctx)
			if err != nil {
				return err
			}
			defer func() { _ = tx.Rollback(ctx) }()

			q := database.New(tx)

			if err := f(q); err != nil {
				return err
			}
			if err := tx.Commit(ctx); err != nil {
				return err
			}
			return nil
		})

		path, handler := v1connect.NewStoreServiceHandler(srv)

		mux := http.NewServeMux()
		mux.Handle(path, handler)

		server := &http.Server{
			Addr: "localhost:" + cmp.Or(os.Getenv("PORT"), "8767"),
			// Use h2c so we can serve HTTP/2 without TLS.
			Handler: h2c.NewHandler(mux, &http2.Server{}),
		}

		eg := errgroup.Group{}
		eg.Go(func() error {
			return server.ListenAndServe()
		})
		eg.Go(func() error {
			<-ctx.Done()
			return server.Close()
		})
		if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func runWithDB(ctx context.Context, run func(ctx context.Context, pool *pgxpool.Pool) error) error {
	dbName, ok := os.LookupEnv("PGDATABASE")
	if !ok || dbName == "" {
		return errors.New("PGDATABASE environment variable not set")
	}
	pool, err := createDatabase(ctx, dbName)
	if err != nil {
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

func createDatabase(ctx context.Context, name string) (*pgxpool.Pool, error) {
	//pgURL, err := databaseURLFromEnv("postgres")
	//if err != nil {
	//	return nil, err
	//}
	//pgPool, err := pgxpool.New(ctx, pgURL)
	//if err != nil {
	//	return nil, err
	//}
	//if err := pgPool.Ping(ctx); err != nil {
	//	return nil, err
	//}
	//_, _ = pgPool.Exec(ctx, fmt.Sprintf(`DROP DATABASE %s`, name))
	//if _, err := pgPool.Exec(ctx, fmt.Sprintf(`CREATE DATABASE %s`, name)); err != nil {
	//	return nil, err
	//}
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
