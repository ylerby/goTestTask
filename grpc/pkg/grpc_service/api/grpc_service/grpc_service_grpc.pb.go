// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: api/grpc_service/grpc_service.proto

package grpc_service

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

const (
	FullUrl_GetFullUrl_FullMethodName = "/grpc_service.FullUrl/GetFullUrl"
)

// FullUrlClient is the client API for FullUrl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FullUrlClient interface {
	GetFullUrl(ctx context.Context, in *GetFullUrlRequest, opts ...grpc.CallOption) (*GetFullUrlResponse, error)
}

type fullUrlClient struct {
	cc grpc.ClientConnInterface
}

func NewFullUrlClient(cc grpc.ClientConnInterface) FullUrlClient {
	return &fullUrlClient{cc}
}

func (c *fullUrlClient) GetFullUrl(ctx context.Context, in *GetFullUrlRequest, opts ...grpc.CallOption) (*GetFullUrlResponse, error) {
	out := new(GetFullUrlResponse)
	err := c.cc.Invoke(ctx, FullUrl_GetFullUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FullUrlServer is the server API for FullUrl service.
// All implementations must embed UnimplementedFullUrlServer
// for forward compatibility
type FullUrlServer interface {
	GetFullUrl(context.Context, *GetFullUrlRequest) (*GetFullUrlResponse, error)
	mustEmbedUnimplementedFullUrlServer()
}

// UnimplementedFullUrlServer must be embedded to have forward compatible implementations.
type UnimplementedFullUrlServer struct {
}

func (UnimplementedFullUrlServer) GetFullUrl(context.Context, *GetFullUrlRequest) (*GetFullUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullUrl not implemented")
}
func (UnimplementedFullUrlServer) mustEmbedUnimplementedFullUrlServer() {}

// UnsafeFullUrlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FullUrlServer will
// result in compilation errors.
type UnsafeFullUrlServer interface {
	mustEmbedUnimplementedFullUrlServer()
}

func RegisterFullUrlServer(s grpc.ServiceRegistrar, srv FullUrlServer) {
	s.RegisterService(&FullUrl_ServiceDesc, srv)
}

func _FullUrl_GetFullUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFullUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FullUrlServer).GetFullUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FullUrl_GetFullUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FullUrlServer).GetFullUrl(ctx, req.(*GetFullUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FullUrl_ServiceDesc is the grpc.ServiceDesc for FullUrl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FullUrl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_service.FullUrl",
	HandlerType: (*FullUrlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFullUrl",
			Handler:    _FullUrl_GetFullUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc_service/grpc_service.proto",
}

const (
	ShortUrl_GetShortUrl_FullMethodName = "/grpc_service.ShortUrl/GetShortUrl"
)

// ShortUrlClient is the client API for ShortUrl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortUrlClient interface {
	GetShortUrl(ctx context.Context, in *GetShortUrlRequest, opts ...grpc.CallOption) (*GetShortUrlResponse, error)
}

type shortUrlClient struct {
	cc grpc.ClientConnInterface
}

func NewShortUrlClient(cc grpc.ClientConnInterface) ShortUrlClient {
	return &shortUrlClient{cc}
}

func (c *shortUrlClient) GetShortUrl(ctx context.Context, in *GetShortUrlRequest, opts ...grpc.CallOption) (*GetShortUrlResponse, error) {
	out := new(GetShortUrlResponse)
	err := c.cc.Invoke(ctx, ShortUrl_GetShortUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortUrlServer is the server API for ShortUrl service.
// All implementations must embed UnimplementedShortUrlServer
// for forward compatibility
type ShortUrlServer interface {
	GetShortUrl(context.Context, *GetShortUrlRequest) (*GetShortUrlResponse, error)
	mustEmbedUnimplementedShortUrlServer()
}

// UnimplementedShortUrlServer must be embedded to have forward compatible implementations.
type UnimplementedShortUrlServer struct {
}

func (UnimplementedShortUrlServer) GetShortUrl(context.Context, *GetShortUrlRequest) (*GetShortUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortUrl not implemented")
}
func (UnimplementedShortUrlServer) mustEmbedUnimplementedShortUrlServer() {}

// UnsafeShortUrlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortUrlServer will
// result in compilation errors.
type UnsafeShortUrlServer interface {
	mustEmbedUnimplementedShortUrlServer()
}

func RegisterShortUrlServer(s grpc.ServiceRegistrar, srv ShortUrlServer) {
	s.RegisterService(&ShortUrl_ServiceDesc, srv)
}

func _ShortUrl_GetShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortUrlServer).GetShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortUrl_GetShortUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortUrlServer).GetShortUrl(ctx, req.(*GetShortUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortUrl_ServiceDesc is the grpc.ServiceDesc for ShortUrl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortUrl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_service.ShortUrl",
	HandlerType: (*ShortUrlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShortUrl",
			Handler:    _ShortUrl_GetShortUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc_service/grpc_service.proto",
}