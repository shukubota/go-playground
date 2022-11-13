// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: protobuf/hello.proto

package grpc_example

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloBiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloBiDirectionalStreamClient, error)
	SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hello.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloBiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloBiDirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/hello.Greeter/SayHelloBiDirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloBiDirectionalStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloBiDirectionalStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloBiDirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloBiDirectionalStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloBiDirectionalStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/hello.Greeter/SayHelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHelloServerStreamClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloServerStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloBiDirectionalStream(Greeter_SayHelloBiDirectionalStreamServer) error
	SayHelloServerStream(*HelloRequest, Greeter_SayHelloServerStreamServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) SayHelloBiDirectionalStream(Greeter_SayHelloBiDirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBiDirectionalStream not implemented")
}
func (UnimplementedGreeterServer) SayHelloServerStream(*HelloRequest, Greeter_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloBiDirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloBiDirectionalStream(&greeterSayHelloBiDirectionalStreamServer{stream})
}

type Greeter_SayHelloBiDirectionalStreamServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloBiDirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloBiDirectionalStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloBiDirectionalStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayHelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloServerStream(m, &greeterSayHelloServerStreamServer{stream})
}

type Greeter_SayHelloServerStreamServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterSayHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloServerStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloBiDirectionalStream",
			Handler:       _Greeter_SayHelloBiDirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloServerStream",
			Handler:       _Greeter_SayHelloServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf/hello.proto",
}
