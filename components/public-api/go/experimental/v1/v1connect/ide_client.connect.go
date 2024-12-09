// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: devtool/experimental/v1/ide_client.proto

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
	// IDEClientServiceName is the fully-qualified name of the IDEClientService service.
	IDEClientServiceName = "devtool.experimental.v1.IDEClientService"
)

// IDEClientServiceClient is a client for the devtool.experimental.v1.IDEClientService service.
type IDEClientServiceClient interface {
	// SendHeartbeat sends a clientheartbeat signal for a running workspace.
	SendHeartbeat(context.Context, *connect_go.Request[v1.SendHeartbeatRequest]) (*connect_go.Response[v1.SendHeartbeatResponse], error)
	// SendDidClose sends a client close signal for a running workspace.
	SendDidClose(context.Context, *connect_go.Request[v1.SendDidCloseRequest]) (*connect_go.Response[v1.SendDidCloseResponse], error)
	// UpdateGitStatus updates the status of a repository in a workspace.
	UpdateGitStatus(context.Context, *connect_go.Request[v1.UpdateGitStatusRequest]) (*connect_go.Response[v1.UpdateGitStatusResponse], error)
}

// NewIDEClientServiceClient constructs a client for the devtool.experimental.v1.IDEClientService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIDEClientServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IDEClientServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &iDEClientServiceClient{
		sendHeartbeat: connect_go.NewClient[v1.SendHeartbeatRequest, v1.SendHeartbeatResponse](
			httpClient,
			baseURL+"/devtool.experimental.v1.IDEClientService/SendHeartbeat",
			opts...,
		),
		sendDidClose: connect_go.NewClient[v1.SendDidCloseRequest, v1.SendDidCloseResponse](
			httpClient,
			baseURL+"/devtool.experimental.v1.IDEClientService/SendDidClose",
			opts...,
		),
		updateGitStatus: connect_go.NewClient[v1.UpdateGitStatusRequest, v1.UpdateGitStatusResponse](
			httpClient,
			baseURL+"/devtool.experimental.v1.IDEClientService/UpdateGitStatus",
			opts...,
		),
	}
}

// iDEClientServiceClient implements IDEClientServiceClient.
type iDEClientServiceClient struct {
	sendHeartbeat   *connect_go.Client[v1.SendHeartbeatRequest, v1.SendHeartbeatResponse]
	sendDidClose    *connect_go.Client[v1.SendDidCloseRequest, v1.SendDidCloseResponse]
	updateGitStatus *connect_go.Client[v1.UpdateGitStatusRequest, v1.UpdateGitStatusResponse]
}

// SendHeartbeat calls devtool.experimental.v1.IDEClientService.SendHeartbeat.
func (c *iDEClientServiceClient) SendHeartbeat(ctx context.Context, req *connect_go.Request[v1.SendHeartbeatRequest]) (*connect_go.Response[v1.SendHeartbeatResponse], error) {
	return c.sendHeartbeat.CallUnary(ctx, req)
}

// SendDidClose calls devtool.experimental.v1.IDEClientService.SendDidClose.
func (c *iDEClientServiceClient) SendDidClose(ctx context.Context, req *connect_go.Request[v1.SendDidCloseRequest]) (*connect_go.Response[v1.SendDidCloseResponse], error) {
	return c.sendDidClose.CallUnary(ctx, req)
}

// UpdateGitStatus calls devtool.experimental.v1.IDEClientService.UpdateGitStatus.
func (c *iDEClientServiceClient) UpdateGitStatus(ctx context.Context, req *connect_go.Request[v1.UpdateGitStatusRequest]) (*connect_go.Response[v1.UpdateGitStatusResponse], error) {
	return c.updateGitStatus.CallUnary(ctx, req)
}

// IDEClientServiceHandler is an implementation of the devtool.experimental.v1.IDEClientService
// service.
type IDEClientServiceHandler interface {
	// SendHeartbeat sends a clientheartbeat signal for a running workspace.
	SendHeartbeat(context.Context, *connect_go.Request[v1.SendHeartbeatRequest]) (*connect_go.Response[v1.SendHeartbeatResponse], error)
	// SendDidClose sends a client close signal for a running workspace.
	SendDidClose(context.Context, *connect_go.Request[v1.SendDidCloseRequest]) (*connect_go.Response[v1.SendDidCloseResponse], error)
	// UpdateGitStatus updates the status of a repository in a workspace.
	UpdateGitStatus(context.Context, *connect_go.Request[v1.UpdateGitStatusRequest]) (*connect_go.Response[v1.UpdateGitStatusResponse], error)
}

// NewIDEClientServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIDEClientServiceHandler(svc IDEClientServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/devtool.experimental.v1.IDEClientService/SendHeartbeat", connect_go.NewUnaryHandler(
		"/devtool.experimental.v1.IDEClientService/SendHeartbeat",
		svc.SendHeartbeat,
		opts...,
	))
	mux.Handle("/devtool.experimental.v1.IDEClientService/SendDidClose", connect_go.NewUnaryHandler(
		"/devtool.experimental.v1.IDEClientService/SendDidClose",
		svc.SendDidClose,
		opts...,
	))
	mux.Handle("/devtool.experimental.v1.IDEClientService/UpdateGitStatus", connect_go.NewUnaryHandler(
		"/devtool.experimental.v1.IDEClientService/UpdateGitStatus",
		svc.UpdateGitStatus,
		opts...,
	))
	return "/devtool.experimental.v1.IDEClientService/", mux
}

// UnimplementedIDEClientServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIDEClientServiceHandler struct{}

func (UnimplementedIDEClientServiceHandler) SendHeartbeat(context.Context, *connect_go.Request[v1.SendHeartbeatRequest]) (*connect_go.Response[v1.SendHeartbeatResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("devtool.experimental.v1.IDEClientService.SendHeartbeat is not implemented"))
}

func (UnimplementedIDEClientServiceHandler) SendDidClose(context.Context, *connect_go.Request[v1.SendDidCloseRequest]) (*connect_go.Response[v1.SendDidCloseResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("devtool.experimental.v1.IDEClientService.SendDidClose is not implemented"))
}

func (UnimplementedIDEClientServiceHandler) UpdateGitStatus(context.Context, *connect_go.Request[v1.UpdateGitStatusRequest]) (*connect_go.Response[v1.UpdateGitStatusResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("devtool.experimental.v1.IDEClientService.UpdateGitStatus is not implemented"))
}
