// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: devtool/v1/token.proto

package v1

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

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenServiceClient interface {
	// CreateUserToken creates a new temporary access token for the specified user.
	// +admin – only to be used by installation admins
	CreateTemporaryAccessToken(ctx context.Context, in *CreateTemporaryAccessTokenRequest, opts ...grpc.CallOption) (*CreateTemporaryAccessTokenResponse, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) CreateTemporaryAccessToken(ctx context.Context, in *CreateTemporaryAccessTokenRequest, opts ...grpc.CallOption) (*CreateTemporaryAccessTokenResponse, error) {
	out := new(CreateTemporaryAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/devtool.v1.TokenService/CreateTemporaryAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
// All implementations must embed UnimplementedTokenServiceServer
// for forward compatibility
type TokenServiceServer interface {
	// CreateUserToken creates a new temporary access token for the specified user.
	// +admin – only to be used by installation admins
	CreateTemporaryAccessToken(context.Context, *CreateTemporaryAccessTokenRequest) (*CreateTemporaryAccessTokenResponse, error)
	mustEmbedUnimplementedTokenServiceServer()
}

// UnimplementedTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (UnimplementedTokenServiceServer) CreateTemporaryAccessToken(context.Context, *CreateTemporaryAccessTokenRequest) (*CreateTemporaryAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemporaryAccessToken not implemented")
}
func (UnimplementedTokenServiceServer) mustEmbedUnimplementedTokenServiceServer() {}

// UnsafeTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServiceServer will
// result in compilation errors.
type UnsafeTokenServiceServer interface {
	mustEmbedUnimplementedTokenServiceServer()
}

func RegisterTokenServiceServer(s grpc.ServiceRegistrar, srv TokenServiceServer) {
	s.RegisterService(&TokenService_ServiceDesc, srv)
}

func _TokenService_CreateTemporaryAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemporaryAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).CreateTemporaryAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devtool.v1.TokenService/CreateTemporaryAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).CreateTemporaryAccessToken(ctx, req.(*CreateTemporaryAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenService_ServiceDesc is the grpc.ServiceDesc for TokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "devtool.v1.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemporaryAccessToken",
			Handler:    _TokenService_CreateTemporaryAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "devtool/v1/token.proto",
}
