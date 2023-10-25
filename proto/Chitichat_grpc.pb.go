// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/Chitichat.proto

package Homework03_Tr_lsstemning

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
	ServerConnection_SendMessages_FullMethodName = "/proto.ServerConnection/SendMessages"
)

// ServerConnectionClient is the client API for ServerConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerConnectionClient interface {
	SendMessages(ctx context.Context, opts ...grpc.CallOption) (ServerConnection_SendMessagesClient, error)
}

type serverConnectionClient struct {
	cc grpc.ClientConnInterface
}

func NewServerConnectionClient(cc grpc.ClientConnInterface) ServerConnectionClient {
	return &serverConnectionClient{cc}
}

func (c *serverConnectionClient) SendMessages(ctx context.Context, opts ...grpc.CallOption) (ServerConnection_SendMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerConnection_ServiceDesc.Streams[0], ServerConnection_SendMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &serverConnectionSendMessagesClient{stream}
	return x, nil
}

type ServerConnection_SendMessagesClient interface {
	Send(*ClientMessage) error
	Recv() (*ServerBroadcast, error)
	grpc.ClientStream
}

type serverConnectionSendMessagesClient struct {
	grpc.ClientStream
}

func (x *serverConnectionSendMessagesClient) Send(m *ClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serverConnectionSendMessagesClient) Recv() (*ServerBroadcast, error) {
	m := new(ServerBroadcast)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerConnectionServer is the server API for ServerConnection service.
// All implementations must embed UnimplementedServerConnectionServer
// for forward compatibility
type ServerConnectionServer interface {
	SendMessages(ServerConnection_SendMessagesServer) error
	mustEmbedUnimplementedServerConnectionServer()
}

// UnimplementedServerConnectionServer must be embedded to have forward compatible implementations.
type UnimplementedServerConnectionServer struct {
}

func (UnimplementedServerConnectionServer) SendMessages(ServerConnection_SendMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessages not implemented")
}
func (UnimplementedServerConnectionServer) mustEmbedUnimplementedServerConnectionServer() {}

// UnsafeServerConnectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerConnectionServer will
// result in compilation errors.
type UnsafeServerConnectionServer interface {
	mustEmbedUnimplementedServerConnectionServer()
}

func RegisterServerConnectionServer(s grpc.ServiceRegistrar, srv ServerConnectionServer) {
	s.RegisterService(&ServerConnection_ServiceDesc, srv)
}

func _ServerConnection_SendMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerConnectionServer).SendMessages(&serverConnectionSendMessagesServer{stream})
}

type ServerConnection_SendMessagesServer interface {
	Send(*ServerBroadcast) error
	Recv() (*ClientMessage, error)
	grpc.ServerStream
}

type serverConnectionSendMessagesServer struct {
	grpc.ServerStream
}

func (x *serverConnectionSendMessagesServer) Send(m *ServerBroadcast) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serverConnectionSendMessagesServer) Recv() (*ClientMessage, error) {
	m := new(ClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerConnection_ServiceDesc is the grpc.ServiceDesc for ServerConnection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerConnection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServerConnection",
	HandlerType: (*ServerConnectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessages",
			Handler:       _ServerConnection_SendMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/Chitichat.proto",
}
