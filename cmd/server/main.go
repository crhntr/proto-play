package main

import (
	"cmp"
	"context"
	"errors"
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
	err := database.Run(ctx, func(ctx context.Context, pool *pgxpool.Pool) error {
		p := database.Use(pool)

		srv := play.New(slog.Default(), p)

		path, handler := v1connect.NewStoreServiceHandler(srv)

		mux := http.NewServeMux()
		mux.Handle(path, handler)

		server := &http.Server{
			Addr: "localhost:" + cmp.Or(os.Getenv("PORT"), "8767"),
			// Use h2c so we can serve HTTP/2 without TLS.
			Handler: h2c.NewHandler(mux, &http2.Server{}),
		}

		eg := new(errgroup.Group)
		eg.Go(func() error {
			return server.ListenAndServe()
		})
		eg.Go(func() error {
			<-ctx.Done()
			return server.Close()
		})
		if err := eg.Wait(); err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
