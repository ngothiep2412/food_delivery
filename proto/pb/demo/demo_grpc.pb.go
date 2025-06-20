// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: demo.proto

package demo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HelloService_Hello_FullMethodName                 = "/demo.HelloService/Hello"
	HelloService_HelloStream_FullMethodName           = "/demo.HelloService/HelloStream"
	HelloService_GetRestaurantLikeStat_FullMethodName = "/demo.HelloService/GetRestaurantLikeStat"
)

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error)
	GetRestaurantLikeStat(ctx context.Context, in *RestaurantLikeStatRequest, opts ...grpc.CallOption) (*RestaurantLikeStatResponse, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, HelloService_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) HelloStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], HelloService_HelloStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HelloService_HelloStreamClient = grpc.ServerStreamingClient[HelloResponse]

func (c *helloServiceClient) GetRestaurantLikeStat(ctx context.Context, in *RestaurantLikeStatRequest, opts ...grpc.CallOption) (*RestaurantLikeStatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RestaurantLikeStatResponse)
	err := c.cc.Invoke(ctx, HelloService_GetRestaurantLikeStat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility.
type HelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloStream(*HelloRequest, grpc.ServerStreamingServer[HelloResponse]) error
	GetRestaurantLikeStat(context.Context, *RestaurantLikeStatRequest) (*RestaurantLikeStatResponse, error)
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHelloServiceServer struct{}

func (UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedHelloServiceServer) HelloStream(*HelloRequest, grpc.ServerStreamingServer[HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method HelloStream not implemented")
}
func (UnimplementedHelloServiceServer) GetRestaurantLikeStat(context.Context, *RestaurantLikeStatRequest) (*RestaurantLikeStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurantLikeStat not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}
func (UnimplementedHelloServiceServer) testEmbeddedByValue()                      {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	// If the following call pancis, it indicates UnimplementedHelloServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_HelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).HelloStream(m, &grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HelloService_HelloStreamServer = grpc.ServerStreamingServer[HelloResponse]

func _HelloService_GetRestaurantLikeStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantLikeStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).GetRestaurantLikeStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_GetRestaurantLikeStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).GetRestaurantLikeStat(ctx, req.(*RestaurantLikeStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
		{
			MethodName: "GetRestaurantLikeStat",
			Handler:    _HelloService_GetRestaurantLikeStat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloStream",
			Handler:       _HelloService_HelloStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "demo.proto",
}
