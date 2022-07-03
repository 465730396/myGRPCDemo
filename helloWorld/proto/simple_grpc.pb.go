// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: helloWorld/proto/simple.proto

package proto

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

// SimpleClient is the client API for Simple service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleClient interface {
	Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	Conversations(ctx context.Context, opts ...grpc.CallOption) (Simple_ConversationsClient, error)
}

type simpleClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleClient(cc grpc.ClientConnInterface) SimpleClient {
	return &simpleClient{cc}
}

func (c *simpleClient) Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/proto.Simple/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleClient) Conversations(ctx context.Context, opts ...grpc.CallOption) (Simple_ConversationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Simple_ServiceDesc.Streams[0], "/proto.Simple/Conversations", opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleConversationsClient{stream}
	return x, nil
}

type Simple_ConversationsClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type simpleConversationsClient struct {
	grpc.ClientStream
}

func (x *simpleConversationsClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simpleConversationsClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleServer is the server API for Simple service.
// All implementations must embed UnimplementedSimpleServer
// for forward compatibility
type SimpleServer interface {
	Route(context.Context, *SimpleRequest) (*SimpleResponse, error)
	Conversations(Simple_ConversationsServer) error
	mustEmbedUnimplementedSimpleServer()
}

// UnimplementedSimpleServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleServer struct {
}

func (UnimplementedSimpleServer) Route(context.Context, *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedSimpleServer) Conversations(Simple_ConversationsServer) error {
	return status.Errorf(codes.Unimplemented, "method Conversations not implemented")
}
func (UnimplementedSimpleServer) mustEmbedUnimplementedSimpleServer() {}

// UnsafeSimpleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleServer will
// result in compilation errors.
type UnsafeSimpleServer interface {
	mustEmbedUnimplementedSimpleServer()
}

func RegisterSimpleServer(s grpc.ServiceRegistrar, srv SimpleServer) {
	s.RegisterService(&Simple_ServiceDesc, srv)
}

func _Simple_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Simple/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServer).Route(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simple_Conversations_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimpleServer).Conversations(&simpleConversationsServer{stream})
}

type Simple_ConversationsServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type simpleConversationsServer struct {
	grpc.ServerStream
}

func (x *simpleConversationsServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simpleConversationsServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Simple_ServiceDesc is the grpc.ServiceDesc for Simple service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Simple_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Simple",
	HandlerType: (*SimpleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _Simple_Route_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Conversations",
			Handler:       _Simple_Conversations_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloWorld/proto/simple.proto",
}