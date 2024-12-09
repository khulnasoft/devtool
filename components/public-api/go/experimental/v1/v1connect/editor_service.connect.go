// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: devtool/experimental/v1/editor_service.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/khulnasoft/devtool/components/public-api/go/experimental/v1"
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
	// EditorServiceName is the fully-qualified name of the EditorService service.
	EditorServiceName = "devtool.experimental.v1.EditorService"
)

// EditorServiceClient is a client for the devtool.experimental.v1.EditorService service.
type EditorServiceClient interface {
	ListEditorOptions(context.Context, *connect_go.Request[v1.ListEditorOptionsRequest]) (*connect_go.Response[v1.ListEditorOptionsResponse], error)
}

// NewEditorServiceClient constructs a client for the devtool.experimental.v1.EditorService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEditorServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EditorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &editorServiceClient{
		listEditorOptions: connect_go.NewClient[v1.ListEditorOptionsRequest, v1.ListEditorOptionsResponse](
			httpClient,
			baseURL+"/devtool.experimental.v1.EditorService/ListEditorOptions",
			opts...,
		),
	}
}

// editorServiceClient implements EditorServiceClient.
type editorServiceClient struct {
	listEditorOptions *connect_go.Client[v1.ListEditorOptionsRequest, v1.ListEditorOptionsResponse]
}

// ListEditorOptions calls devtool.experimental.v1.EditorService.ListEditorOptions.
func (c *editorServiceClient) ListEditorOptions(ctx context.Context, req *connect_go.Request[v1.ListEditorOptionsRequest]) (*connect_go.Response[v1.ListEditorOptionsResponse], error) {
	return c.listEditorOptions.CallUnary(ctx, req)
}

// EditorServiceHandler is an implementation of the devtool.experimental.v1.EditorService service.
type EditorServiceHandler interface {
	ListEditorOptions(context.Context, *connect_go.Request[v1.ListEditorOptionsRequest]) (*connect_go.Response[v1.ListEditorOptionsResponse], error)
}

// NewEditorServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEditorServiceHandler(svc EditorServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/devtool.experimental.v1.EditorService/ListEditorOptions", connect_go.NewUnaryHandler(
		"/devtool.experimental.v1.EditorService/ListEditorOptions",
		svc.ListEditorOptions,
		opts...,
	))
	return "/devtool.experimental.v1.EditorService/", mux
}

// UnimplementedEditorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEditorServiceHandler struct{}

func (UnimplementedEditorServiceHandler) ListEditorOptions(context.Context, *connect_go.Request[v1.ListEditorOptionsRequest]) (*connect_go.Response[v1.ListEditorOptionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("devtool.experimental.v1.EditorService.ListEditorOptions is not implemented"))
}
