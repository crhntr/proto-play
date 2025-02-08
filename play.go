package play

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "github.com/crhntr/proto-play/api/playground/v1"
	"github.com/crhntr/proto-play/api/playground/v1/v1connect"
	"github.com/crhntr/proto-play/database"
)

//go:generate buf generate
//go:generate sqlc generate

type Querier interface {
	Query(ctx context.Context, f database.QuerierFunc) error
}

type Server struct {
	log *slog.Logger
	db  Querier
}

func New(logger *slog.Logger, db Querier) *Server {
	return &Server{log: logger, db: db}
}

var _ v1connect.StoreServiceHandler = (*Server)(nil)

func (s *Server) ListByIDKey(ctx context.Context, req *connect.Request[v1.ListByIDKeyRequest]) (*connect.Response[v1.ListByIDKeyResponse], error) {
	queryJSON, err := protojson.Marshal(req.Msg.GetQuery())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	s.log.Debug("exists", "query_json", string(queryJSON))
	res := connect.NewResponse(new(v1.ListByIDKeyResponse))
	if err = s.db.Query(ctx, func(q database.Querier) error {
		_, err := q.Find(ctx, queryJSON)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return connect.NewError(connect.CodeNotFound, fmt.Errorf("query not found"))
			}
			return err
		}
		res.Msg.Messages = append(res.Msg.Messages)
		return nil
	}); err != nil {
		return nil, connectError(err)
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
	if err = s.db.Query(ctx, func(q database.Querier) error {
		_, err := q.Find(ctx, queryJSON)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return connect.NewError(connect.CodeNotFound, fmt.Errorf("query not found"))
			}
			return err
		}
		res.Msg.Found = true
		return nil
	}); err != nil {
		return nil, connectError(err)
	}
	return res, nil
}

func (s *Server) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	contentJSON, err := protojson.Marshal(req.Msg.GetContent())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	res := connect.NewResponse(new(v1.CreateResponse))
	if err = s.db.Query(ctx, func(q database.Querier) error {
		number, err := q.Add(ctx, contentJSON)
		if err != nil {
			return err
		}
		res.Msg.Number = number
		return nil
	}); err != nil {
		return nil, connectError(err)
	}
	return res, nil
}

func (s *Server) List(ctx context.Context, req *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error) {
	res := connect.NewResponse(new(v1.ListResponse))
	if err := s.db.Query(ctx, func(q database.Querier) error {
		messages, err := q.List(ctx)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return connect.NewError(connect.CodeNotFound, fmt.Errorf("query not found"))
			}
			return err
		}
		list := make([]v1.Identifier, len(messages))
		for i, message := range messages {
			if err := protojson.Unmarshal(message.Content, &list[i]); err != nil {
				return connect.NewError(connect.CodeInvalidArgument, err)
			}
			res.Msg.List = append(res.Msg.List, &list[i])
		}
		return nil
	}); err != nil {
		return nil, connectError(err)
	}
	return res, nil
}

func connectError(err error) error {
	var cErr *connect.Error
	if errors.As(err, &cErr) {
		return cErr
	}
	return connect.NewError(connect.CodeUnknown, err)
}
