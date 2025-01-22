package play

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "github.com/crhntr/proto-play/api/playground/v1"
	"github.com/crhntr/proto-play/api/playground/v1/v1connect"
	"github.com/crhntr/proto-play/database"
)

//go:generate buf generate
//go:generate sqlc generate

type TXFunc func(ctx context.Context, f func(q database.Querier) error) error

type Server struct {
	log         *slog.Logger
	transaction TXFunc
}

func New(logger *slog.Logger, tx TXFunc) *Server {
	return &Server{log: logger, transaction: tx}
}

var _ v1connect.StoreServiceHandler = (*Server)(nil)

func (s *Server) ListByIDKey(ctx context.Context, req *connect.Request[v1.ListByIDKeyRequest]) (*connect.Response[v1.ListByIDKeyResponse], error) {
	queryJSON, err := protojson.Marshal(req.Msg.GetQuery())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	s.log.Debug("exists", "query_json", string(queryJSON))
	res := connect.NewResponse(new(v1.ListByIDKeyResponse))
	if err := s.transaction(ctx, func(q database.Querier) error {
		_, err := q.Find(ctx, queryJSON)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil
			}
			return err
		}
		res.Msg.Messages = append(res.Msg.Messages)
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	return res, nil
}

func (s *Server) Exists(ctx context.Context, req *connect.Request[v1.ExistsRequest]) (*connect.Response[v1.ExistsResponse], error) {
	queryJSON, err := protojson.Marshal(req.Msg.GetQuery())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	s.log.Debug("exists", "query_json", string(queryJSON))
	res := connect.NewResponse(new(v1.ExistsResponse))
	if err := s.transaction(ctx, func(q database.Querier) error {
		_, err := q.Find(ctx, queryJSON)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil
			}
			return err
		}
		res.Msg.Found = true
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	return res, nil
}

func (s *Server) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	const constraint = "unique_idx_messages_content"
	contentJSON, err := protojson.Marshal(req.Msg.GetContent())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	res := connect.NewResponse(new(v1.CreateResponse))
	if err := s.transaction(ctx, func(q database.Querier) error {
		number, err := q.Add(ctx, contentJSON)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.ConstraintName != constraint {
					return err
				}
				return nil
			}
			return err
		}
		res.Msg.Number = number
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	return res, nil
}
