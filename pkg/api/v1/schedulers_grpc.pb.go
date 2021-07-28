// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// SchedulersClient is the client API for Schedulers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulersClient interface {
	// Lists all schedulers.
	ListSchedulers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListSchedulersReply, error)
}

type schedulersClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulersClient(cc grpc.ClientConnInterface) SchedulersClient {
	return &schedulersClient{cc}
}

func (c *schedulersClient) ListSchedulers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListSchedulersReply, error) {
	out := new(ListSchedulersReply)
	err := c.cc.Invoke(ctx, "/api.v1.Schedulers/ListSchedulers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulersServer is the server API for Schedulers service.
// All implementations must embed UnimplementedSchedulersServer
// for forward compatibility
type SchedulersServer interface {
	// Lists all schedulers.
	ListSchedulers(context.Context, *EmptyRequest) (*ListSchedulersReply, error)
	mustEmbedUnimplementedSchedulersServer()
}

// UnimplementedSchedulersServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulersServer struct {
}

func (UnimplementedSchedulersServer) ListSchedulers(context.Context, *EmptyRequest) (*ListSchedulersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchedulers not implemented")
}
func (UnimplementedSchedulersServer) mustEmbedUnimplementedSchedulersServer() {}

// UnsafeSchedulersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulersServer will
// result in compilation errors.
type UnsafeSchedulersServer interface {
	mustEmbedUnimplementedSchedulersServer()
}

func RegisterSchedulersServer(s grpc.ServiceRegistrar, srv SchedulersServer) {
	s.RegisterService(&Schedulers_ServiceDesc, srv)
}

func _Schedulers_ListSchedulers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulersServer).ListSchedulers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Schedulers/ListSchedulers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulersServer).ListSchedulers(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Schedulers_ServiceDesc is the grpc.ServiceDesc for Schedulers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Schedulers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.Schedulers",
	HandlerType: (*SchedulersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSchedulers",
			Handler:    _Schedulers_ListSchedulers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/schedulers.proto",
}
