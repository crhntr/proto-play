// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/playground/v1/playground.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/crhntr/proto-play/api/playground/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// StoreServiceName is the fully-qualified name of the StoreService service.
	StoreServiceName = "api.playground.v1.StoreService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StoreServiceListByIDKeyProcedure is the fully-qualified name of the StoreService's ListByIDKey
	// RPC.
	StoreServiceListByIDKeyProcedure = "/api.playground.v1.StoreService/ListByIDKey"
	// StoreServiceExistsProcedure is the fully-qualified name of the StoreService's Exists RPC.
	StoreServiceExistsProcedure = "/api.playground.v1.StoreService/Exists"
)

// StoreServiceClient is a client for the api.playground.v1.StoreService service.
type StoreServiceClient interface {
	ListByIDKey(context.Context, *connect.Request[v1.ListByIDKeyRequest]) (*connect.Response[v1.ListByIDKeyResponse], error)
	Exists(context.Context, *connect.Request[v1.ExistsRequest]) (*connect.Response[v1.ExistsResponse], error)
}

// NewStoreServiceClient constructs a client for the api.playground.v1.StoreService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStoreServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) StoreServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	storeServiceMethods := v1.File_api_playground_v1_playground_proto.Services().ByName("StoreService").Methods()
	return &storeServiceClient{
		listByIDKey: connect.NewClient[v1.ListByIDKeyRequest, v1.ListByIDKeyResponse](
			httpClient,
			baseURL+StoreServiceListByIDKeyProcedure,
			connect.WithSchema(storeServiceMethods.ByName("ListByIDKey")),
			connect.WithClientOptions(opts...),
		),
		exists: connect.NewClient[v1.ExistsRequest, v1.ExistsResponse](
			httpClient,
			baseURL+StoreServiceExistsProcedure,
			connect.WithSchema(storeServiceMethods.ByName("Exists")),
			connect.WithClientOptions(opts...),
		),
	}
}

// storeServiceClient implements StoreServiceClient.
type storeServiceClient struct {
	listByIDKey *connect.Client[v1.ListByIDKeyRequest, v1.ListByIDKeyResponse]
	exists      *connect.Client[v1.ExistsRequest, v1.ExistsResponse]
}

// ListByIDKey calls api.playground.v1.StoreService.ListByIDKey.
func (c *storeServiceClient) ListByIDKey(ctx context.Context, req *connect.Request[v1.ListByIDKeyRequest]) (*connect.Response[v1.ListByIDKeyResponse], error) {
	return c.listByIDKey.CallUnary(ctx, req)
}

// Exists calls api.playground.v1.StoreService.Exists.
func (c *storeServiceClient) Exists(ctx context.Context, req *connect.Request[v1.ExistsRequest]) (*connect.Response[v1.ExistsResponse], error) {
	return c.exists.CallUnary(ctx, req)
}

// StoreServiceHandler is an implementation of the api.playground.v1.StoreService service.
type StoreServiceHandler interface {
	ListByIDKey(context.Context, *connect.Request[v1.ListByIDKeyRequest]) (*connect.Response[v1.ListByIDKeyResponse], error)
	Exists(context.Context, *connect.Request[v1.ExistsRequest]) (*connect.Response[v1.ExistsResponse], error)
}

// NewStoreServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStoreServiceHandler(svc StoreServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	storeServiceMethods := v1.File_api_playground_v1_playground_proto.Services().ByName("StoreService").Methods()
	storeServiceListByIDKeyHandler := connect.NewUnaryHandler(
		StoreServiceListByIDKeyProcedure,
		svc.ListByIDKey,
		connect.WithSchema(storeServiceMethods.ByName("ListByIDKey")),
		connect.WithHandlerOptions(opts...),
	)
	storeServiceExistsHandler := connect.NewUnaryHandler(
		StoreServiceExistsProcedure,
		svc.Exists,
		connect.WithSchema(storeServiceMethods.ByName("Exists")),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.playground.v1.StoreService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StoreServiceListByIDKeyProcedure:
			storeServiceListByIDKeyHandler.ServeHTTP(w, r)
		case StoreServiceExistsProcedure:
			storeServiceExistsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStoreServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStoreServiceHandler struct{}

func (UnimplementedStoreServiceHandler) ListByIDKey(context.Context, *connect.Request[v1.ListByIDKeyRequest]) (*connect.Response[v1.ListByIDKeyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.playground.v1.StoreService.ListByIDKey is not implemented"))
}

func (UnimplementedStoreServiceHandler) Exists(context.Context, *connect.Request[v1.ExistsRequest]) (*connect.Response[v1.ExistsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.playground.v1.StoreService.Exists is not implemented"))
}
