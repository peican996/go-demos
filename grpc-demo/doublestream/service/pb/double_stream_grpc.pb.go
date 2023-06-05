// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: double_stream.proto

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

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	MyMethod(ctx context.Context, opts ...grpc.CallOption) (MyService_MyMethodClient, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) MyMethod(ctx context.Context, opts ...grpc.CallOption) (MyService_MyMethodClient, error) {
	stream, err := c.cc.NewStream(ctx, &MyService_ServiceDesc.Streams[0], "/MyService/MyMethod", opts...)
	if err != nil {
		return nil, err
	}
	x := &myServiceMyMethodClient{stream}
	return x, nil
}

type MyService_MyMethodClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type myServiceMyMethodClient struct {
	grpc.ClientStream
}

func (x *myServiceMyMethodClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myServiceMyMethodClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility
type MyServiceServer interface {
	MyMethod(MyService_MyMethodServer) error
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (UnimplementedMyServiceServer) MyMethod(MyService_MyMethodServer) error {
	return status.Errorf(codes.Unimplemented, "method MyMethod not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_MyMethod_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyServiceServer).MyMethod(&myServiceMyMethodServer{stream})
}

type MyService_MyMethodServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type myServiceMyMethodServer struct {
	grpc.ServerStream
}

func (x *myServiceMyMethodServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myServiceMyMethodServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MyMethod",
			Handler:       _MyService_MyMethod_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "double_stream.proto",
}
