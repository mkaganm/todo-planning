// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: providergrpc/provider.proto

package providergrpc

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

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderClient interface {
	GetTasks(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) GetTasks(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/providergrpc.Provider/GetTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility
type ProviderServer interface {
	GetTasks(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServer struct {
}

func (UnimplementedProviderServer) GetTasks(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_GetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).GetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/providergrpc.Provider/GetTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).GetTasks(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Provider_ServiceDesc is the grpc.ServiceDesc for Provider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "providergrpc.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTasks",
			Handler:    _Provider_GetTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "providergrpc/provider.proto",
}
