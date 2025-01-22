package play

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	v1 "github.com/crhntr/proto-play/api/playground/v1"
	"github.com/crhntr/proto-play/api/playground/v1/v1connect"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

var _ v1connect.StoreServiceHandler = (*Server)(nil)

func (s Server) ListByIDKey(ctx context.Context, c *connect.Request[v1.ListByIDKeyRequest]) (*connect.Response[v1.ListByIDKeyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, fmt.Errorf("ListByIDKey is not implemented"))
}

func (s Server) Exists(ctx context.Context, c *connect.Request[v1.ExistsRequest]) (*connect.Response[v1.ExistsResponse], error) {
	res := connect.NewResponse(&v1.ExistsResponse{})
	return res, nil
}
