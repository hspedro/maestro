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

// RoomsServiceClient is the client API for RoomsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomsServiceClient interface {
	// Updates a game room with ping data.
	UpdateRoomWithPing(ctx context.Context, in *UpdateRoomWithPingRequest, opts ...grpc.CallOption) (*UpdateRoomWithPingResponse, error)
}

type roomsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomsServiceClient(cc grpc.ClientConnInterface) RoomsServiceClient {
	return &roomsServiceClient{cc}
}

func (c *roomsServiceClient) UpdateRoomWithPing(ctx context.Context, in *UpdateRoomWithPingRequest, opts ...grpc.CallOption) (*UpdateRoomWithPingResponse, error) {
	out := new(UpdateRoomWithPingResponse)
	err := c.cc.Invoke(ctx, "/api.v1.RoomsService/UpdateRoomWithPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomsServiceServer is the server API for RoomsService service.
// All implementations must embed UnimplementedRoomsServiceServer
// for forward compatibility
type RoomsServiceServer interface {
	// Updates a game room with ping data.
	UpdateRoomWithPing(context.Context, *UpdateRoomWithPingRequest) (*UpdateRoomWithPingResponse, error)
	mustEmbedUnimplementedRoomsServiceServer()
}

// UnimplementedRoomsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoomsServiceServer struct {
}

func (UnimplementedRoomsServiceServer) UpdateRoomWithPing(context.Context, *UpdateRoomWithPingRequest) (*UpdateRoomWithPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoomWithPing not implemented")
}
func (UnimplementedRoomsServiceServer) mustEmbedUnimplementedRoomsServiceServer() {}

// UnsafeRoomsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomsServiceServer will
// result in compilation errors.
type UnsafeRoomsServiceServer interface {
	mustEmbedUnimplementedRoomsServiceServer()
}

func RegisterRoomsServiceServer(s grpc.ServiceRegistrar, srv RoomsServiceServer) {
	s.RegisterService(&RoomsService_ServiceDesc, srv)
}

func _RoomsService_UpdateRoomWithPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomWithPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomsServiceServer).UpdateRoomWithPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.RoomsService/UpdateRoomWithPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomsServiceServer).UpdateRoomWithPing(ctx, req.(*UpdateRoomWithPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoomsService_ServiceDesc is the grpc.ServiceDesc for RoomsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.RoomsService",
	HandlerType: (*RoomsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateRoomWithPing",
			Handler:    _RoomsService_UpdateRoomWithPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/rooms.proto",
}
