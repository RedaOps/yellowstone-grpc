// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: geyser.proto

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

// GeyserClient is the client API for Geyser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeyserClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Geyser_SubscribeClient, error)
}

type geyserClient struct {
	cc grpc.ClientConnInterface
}

func NewGeyserClient(cc grpc.ClientConnInterface) GeyserClient {
	return &geyserClient{cc}
}

func (c *geyserClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Geyser_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Geyser_ServiceDesc.Streams[0], "/geyser.Geyser/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &geyserSubscribeClient{stream}
	return x, nil
}

type Geyser_SubscribeClient interface {
	Send(*SubscribeRequest) error
	Recv() (*SubscribeUpdate, error)
	grpc.ClientStream
}

type geyserSubscribeClient struct {
	grpc.ClientStream
}

func (x *geyserSubscribeClient) Send(m *SubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *geyserSubscribeClient) Recv() (*SubscribeUpdate, error) {
	m := new(SubscribeUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GeyserServer is the server API for Geyser service.
// All implementations must embed UnimplementedGeyserServer
// for forward compatibility
type GeyserServer interface {
	Subscribe(Geyser_SubscribeServer) error
	mustEmbedUnimplementedGeyserServer()
}

// UnimplementedGeyserServer must be embedded to have forward compatible implementations.
type UnimplementedGeyserServer struct {
}

func (UnimplementedGeyserServer) Subscribe(Geyser_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedGeyserServer) mustEmbedUnimplementedGeyserServer() {}

// UnsafeGeyserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeyserServer will
// result in compilation errors.
type UnsafeGeyserServer interface {
	mustEmbedUnimplementedGeyserServer()
}

func RegisterGeyserServer(s grpc.ServiceRegistrar, srv GeyserServer) {
	s.RegisterService(&Geyser_ServiceDesc, srv)
}

func _Geyser_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GeyserServer).Subscribe(&geyserSubscribeServer{stream})
}

type Geyser_SubscribeServer interface {
	Send(*SubscribeUpdate) error
	Recv() (*SubscribeRequest, error)
	grpc.ServerStream
}

type geyserSubscribeServer struct {
	grpc.ServerStream
}

func (x *geyserSubscribeServer) Send(m *SubscribeUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *geyserSubscribeServer) Recv() (*SubscribeRequest, error) {
	m := new(SubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Geyser_ServiceDesc is the grpc.ServiceDesc for Geyser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Geyser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geyser.Geyser",
	HandlerType: (*GeyserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Geyser_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "geyser.proto",
}
