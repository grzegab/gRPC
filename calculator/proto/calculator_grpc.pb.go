// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: calculator.proto

package proto

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
	CalcOperation_Operation_FullMethodName = "/calculator.CalcOperation/Operation"
)

// CalcOperationClient is the client API for CalcOperation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcOperationClient interface {
	Operation(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OperationResponse], error)
}

type calcOperationClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcOperationClient(cc grpc.ClientConnInterface) CalcOperationClient {
	return &calcOperationClient{cc}
}

func (c *calcOperationClient) Operation(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OperationResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalcOperation_ServiceDesc.Streams[0], CalcOperation_Operation_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[OperationRequest, OperationResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalcOperation_OperationClient = grpc.ServerStreamingClient[OperationResponse]

// CalcOperationServer is the server API for CalcOperation service.
// All implementations must embed UnimplementedCalcOperationServer
// for forward compatibility.
type CalcOperationServer interface {
	Operation(*OperationRequest, grpc.ServerStreamingServer[OperationResponse]) error
	mustEmbedUnimplementedCalcOperationServer()
}

// UnimplementedCalcOperationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCalcOperationServer struct{}

func (UnimplementedCalcOperationServer) Operation(*OperationRequest, grpc.ServerStreamingServer[OperationResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Operation not implemented")
}
func (UnimplementedCalcOperationServer) mustEmbedUnimplementedCalcOperationServer() {}
func (UnimplementedCalcOperationServer) testEmbeddedByValue()                       {}

// UnsafeCalcOperationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcOperationServer will
// result in compilation errors.
type UnsafeCalcOperationServer interface {
	mustEmbedUnimplementedCalcOperationServer()
}

func RegisterCalcOperationServer(s grpc.ServiceRegistrar, srv CalcOperationServer) {
	// If the following call pancis, it indicates UnimplementedCalcOperationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CalcOperation_ServiceDesc, srv)
}

func _CalcOperation_Operation_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OperationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcOperationServer).Operation(m, &grpc.GenericServerStream[OperationRequest, OperationResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CalcOperation_OperationServer = grpc.ServerStreamingServer[OperationResponse]

// CalcOperation_ServiceDesc is the grpc.ServiceDesc for CalcOperation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcOperation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalcOperation",
	HandlerType: (*CalcOperationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Operation",
			Handler:       _CalcOperation_Operation_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator.proto",
}