// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalculatorSerivceClient is the client API for CalculatorSerivce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorSerivceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorSerivce_PrimesClient, error)
	Avg(ctx context.Context, opts ...grpc.CallOption) (CalculatorSerivce_AvgClient, error)
	Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorSerivce_MaxClient, error)
	Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
}

type calculatorSerivceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorSerivceClient(cc grpc.ClientConnInterface) CalculatorSerivceClient {
	return &calculatorSerivceClient{cc}
}

func (c *calculatorSerivceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorSerivce/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorSerivceClient) Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorSerivce_PrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorSerivce_ServiceDesc.Streams[0], "/calculator.CalculatorSerivce/Primes", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorSerivcePrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorSerivce_PrimesClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type calculatorSerivcePrimesClient struct {
	grpc.ClientStream
}

func (x *calculatorSerivcePrimesClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorSerivceClient) Avg(ctx context.Context, opts ...grpc.CallOption) (CalculatorSerivce_AvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorSerivce_ServiceDesc.Streams[1], "/calculator.CalculatorSerivce/Avg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorSerivceAvgClient{stream}
	return x, nil
}

type CalculatorSerivce_AvgClient interface {
	Send(*AvgRequest) error
	CloseAndRecv() (*AvgResponse, error)
	grpc.ClientStream
}

type calculatorSerivceAvgClient struct {
	grpc.ClientStream
}

func (x *calculatorSerivceAvgClient) Send(m *AvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorSerivceAvgClient) CloseAndRecv() (*AvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorSerivceClient) Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorSerivce_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorSerivce_ServiceDesc.Streams[2], "/calculator.CalculatorSerivce/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorSerivceMaxClient{stream}
	return x, nil
}

type CalculatorSerivce_MaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type calculatorSerivceMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorSerivceMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorSerivceMaxClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorSerivceClient) Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorSerivce/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorSerivceServer is the server API for CalculatorSerivce service.
// All implementations must embed UnimplementedCalculatorSerivceServer
// for forward compatibility
type CalculatorSerivceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Primes(*PrimeRequest, CalculatorSerivce_PrimesServer) error
	Avg(CalculatorSerivce_AvgServer) error
	Max(CalculatorSerivce_MaxServer) error
	Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error)
	mustEmbedUnimplementedCalculatorSerivceServer()
}

// UnimplementedCalculatorSerivceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorSerivceServer struct {
}

func (UnimplementedCalculatorSerivceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorSerivceServer) Primes(*PrimeRequest, CalculatorSerivce_PrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method Primes not implemented")
}
func (UnimplementedCalculatorSerivceServer) Avg(CalculatorSerivce_AvgServer) error {
	return status.Errorf(codes.Unimplemented, "method Avg not implemented")
}
func (UnimplementedCalculatorSerivceServer) Max(CalculatorSerivce_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (UnimplementedCalculatorSerivceServer) Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}
func (UnimplementedCalculatorSerivceServer) mustEmbedUnimplementedCalculatorSerivceServer() {}

// UnsafeCalculatorSerivceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorSerivceServer will
// result in compilation errors.
type UnsafeCalculatorSerivceServer interface {
	mustEmbedUnimplementedCalculatorSerivceServer()
}

func RegisterCalculatorSerivceServer(s grpc.ServiceRegistrar, srv CalculatorSerivceServer) {
	s.RegisterService(&CalculatorSerivce_ServiceDesc, srv)
}

func _CalculatorSerivce_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorSerivceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorSerivce/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorSerivceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorSerivce_Primes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorSerivceServer).Primes(m, &calculatorSerivcePrimesServer{stream})
}

type CalculatorSerivce_PrimesServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type calculatorSerivcePrimesServer struct {
	grpc.ServerStream
}

func (x *calculatorSerivcePrimesServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorSerivce_Avg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorSerivceServer).Avg(&calculatorSerivceAvgServer{stream})
}

type CalculatorSerivce_AvgServer interface {
	SendAndClose(*AvgResponse) error
	Recv() (*AvgRequest, error)
	grpc.ServerStream
}

type calculatorSerivceAvgServer struct {
	grpc.ServerStream
}

func (x *calculatorSerivceAvgServer) SendAndClose(m *AvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorSerivceAvgServer) Recv() (*AvgRequest, error) {
	m := new(AvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorSerivce_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorSerivceServer).Max(&calculatorSerivceMaxServer{stream})
}

type CalculatorSerivce_MaxServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type calculatorSerivceMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorSerivceMaxServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorSerivceMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorSerivce_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorSerivceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorSerivce/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorSerivceServer).Sqrt(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorSerivce_ServiceDesc is the grpc.ServiceDesc for CalculatorSerivce service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorSerivce_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorSerivce",
	HandlerType: (*CalculatorSerivceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorSerivce_Sum_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _CalculatorSerivce_Sqrt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Primes",
			Handler:       _CalculatorSerivce_Primes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Avg",
			Handler:       _CalculatorSerivce_Avg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Max",
			Handler:       _CalculatorSerivce_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
