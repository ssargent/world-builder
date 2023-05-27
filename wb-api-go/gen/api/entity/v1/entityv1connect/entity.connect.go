// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/entity/v1/entity.proto

package entityv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ssargent/world-builder/wb-api-go/gen/api/entity/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// EntityServiceName is the fully-qualified name of the EntityService service.
	EntityServiceName = "api.entity.v1.EntityService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EntityServiceGetEntityProcedure is the fully-qualified name of the EntityService's GetEntity RPC.
	EntityServiceGetEntityProcedure = "/api.entity.v1.EntityService/GetEntity"
)

// EntityServiceClient is a client for the api.entity.v1.EntityService service.
type EntityServiceClient interface {
	GetEntity(context.Context, *connect_go.Request[v1.GetEntityRequest]) (*connect_go.Response[v1.GetEntityResponse], error)
}

// NewEntityServiceClient constructs a client for the api.entity.v1.EntityService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEntityServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EntityServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &entityServiceClient{
		getEntity: connect_go.NewClient[v1.GetEntityRequest, v1.GetEntityResponse](
			httpClient,
			baseURL+EntityServiceGetEntityProcedure,
			opts...,
		),
	}
}

// entityServiceClient implements EntityServiceClient.
type entityServiceClient struct {
	getEntity *connect_go.Client[v1.GetEntityRequest, v1.GetEntityResponse]
}

// GetEntity calls api.entity.v1.EntityService.GetEntity.
func (c *entityServiceClient) GetEntity(ctx context.Context, req *connect_go.Request[v1.GetEntityRequest]) (*connect_go.Response[v1.GetEntityResponse], error) {
	return c.getEntity.CallUnary(ctx, req)
}

// EntityServiceHandler is an implementation of the api.entity.v1.EntityService service.
type EntityServiceHandler interface {
	GetEntity(context.Context, *connect_go.Request[v1.GetEntityRequest]) (*connect_go.Response[v1.GetEntityResponse], error)
}

// NewEntityServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEntityServiceHandler(svc EntityServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(EntityServiceGetEntityProcedure, connect_go.NewUnaryHandler(
		EntityServiceGetEntityProcedure,
		svc.GetEntity,
		opts...,
	))
	return "/api.entity.v1.EntityService/", mux
}

// UnimplementedEntityServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEntityServiceHandler struct{}

func (UnimplementedEntityServiceHandler) GetEntity(context.Context, *connect_go.Request[v1.GetEntityRequest]) (*connect_go.Response[v1.GetEntityResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.entity.v1.EntityService.GetEntity is not implemented"))
}
