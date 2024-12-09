// Copyright (c) 2023 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: content.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentServiceClient interface {
	// DeleteUserContent deletes all content associated with a user.
	DeleteUserContent(ctx context.Context, in *DeleteUserContentRequest, opts ...grpc.CallOption) (*DeleteUserContentResponse, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) DeleteUserContent(ctx context.Context, in *DeleteUserContentRequest, opts ...grpc.CallOption) (*DeleteUserContentResponse, error) {
	out := new(DeleteUserContentResponse)
	err := c.cc.Invoke(ctx, "/contentservice.ContentService/DeleteUserContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
// All implementations must embed UnimplementedContentServiceServer
// for forward compatibility
type ContentServiceServer interface {
	// DeleteUserContent deletes all content associated with a user.
	DeleteUserContent(context.Context, *DeleteUserContentRequest) (*DeleteUserContentResponse, error)
	mustEmbedUnimplementedContentServiceServer()
}

// UnimplementedContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (UnimplementedContentServiceServer) DeleteUserContent(context.Context, *DeleteUserContentRequest) (*DeleteUserContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserContent not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}

// UnsafeContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServiceServer will
// result in compilation errors.
type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {
	s.RegisterService(&ContentService_ServiceDesc, srv)
}

func _ContentService_DeleteUserContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteUserContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contentservice.ContentService/DeleteUserContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteUserContent(ctx, req.(*DeleteUserContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentService_ServiceDesc is the grpc.ServiceDesc for ContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contentservice.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteUserContent",
			Handler:    _ContentService_DeleteUserContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}