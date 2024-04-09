// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: global.proto

package pb

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

// GlobalClient is the client API for Global service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalClient interface {
	StreamSnap(ctx context.Context, in *StreamSnapRequest, opts ...grpc.CallOption) (*StreamSnapShot, error)
}

type globalClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalClient(cc grpc.ClientConnInterface) GlobalClient {
	return &globalClient{cc}
}

func (c *globalClient) StreamSnap(ctx context.Context, in *StreamSnapRequest, opts ...grpc.CallOption) (*StreamSnapShot, error) {
	out := new(StreamSnapShot)
	err := c.cc.Invoke(ctx, "/m7s.Global/StreamSnap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlobalServer is the server API for Global service.
// All implementations must embed UnimplementedGlobalServer
// for forward compatibility
type GlobalServer interface {
	StreamSnap(context.Context, *StreamSnapRequest) (*StreamSnapShot, error)
	mustEmbedUnimplementedGlobalServer()
}

// UnimplementedGlobalServer must be embedded to have forward compatible implementations.
type UnimplementedGlobalServer struct {
}

func (UnimplementedGlobalServer) StreamSnap(context.Context, *StreamSnapRequest) (*StreamSnapShot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamSnap not implemented")
}
func (UnimplementedGlobalServer) mustEmbedUnimplementedGlobalServer() {}

// UnsafeGlobalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalServer will
// result in compilation errors.
type UnsafeGlobalServer interface {
	mustEmbedUnimplementedGlobalServer()
}

func RegisterGlobalServer(s grpc.ServiceRegistrar, srv GlobalServer) {
	s.RegisterService(&Global_ServiceDesc, srv)
}

func _Global_StreamSnap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamSnapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalServer).StreamSnap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m7s.Global/StreamSnap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalServer).StreamSnap(ctx, req.(*StreamSnapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Global_ServiceDesc is the grpc.ServiceDesc for Global service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Global_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "m7s.Global",
	HandlerType: (*GlobalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamSnap",
			Handler:    _Global_StreamSnap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "global.proto",
}
