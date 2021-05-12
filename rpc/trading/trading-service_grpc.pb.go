// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trading

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

// TradingServiceClient is the client API for TradingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradingServiceClient interface {
	SetOrderStatusPaid(ctx context.Context, in *SetOrderStatusPaidRequest, opts ...grpc.CallOption) (*SetOrderStatusPaidReply, error)
	GetOrderInvoice(ctx context.Context, in *GetOrderInvoiceRequest, opts ...grpc.CallOption) (*GetOrderInvoiceReply, error)
}

type tradingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradingServiceClient(cc grpc.ClientConnInterface) TradingServiceClient {
	return &tradingServiceClient{cc}
}

func (c *tradingServiceClient) SetOrderStatusPaid(ctx context.Context, in *SetOrderStatusPaidRequest, opts ...grpc.CallOption) (*SetOrderStatusPaidReply, error) {
	out := new(SetOrderStatusPaidReply)
	err := c.cc.Invoke(ctx, "/TradingService/SetOrderStatusPaid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingServiceClient) GetOrderInvoice(ctx context.Context, in *GetOrderInvoiceRequest, opts ...grpc.CallOption) (*GetOrderInvoiceReply, error) {
	out := new(GetOrderInvoiceReply)
	err := c.cc.Invoke(ctx, "/TradingService/GetOrderInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradingServiceServer is the server API for TradingService service.
// All implementations must embed UnimplementedTradingServiceServer
// for forward compatibility
type TradingServiceServer interface {
	SetOrderStatusPaid(context.Context, *SetOrderStatusPaidRequest) (*SetOrderStatusPaidReply, error)
	GetOrderInvoice(context.Context, *GetOrderInvoiceRequest) (*GetOrderInvoiceReply, error)
	mustEmbedUnimplementedTradingServiceServer()
}

// UnimplementedTradingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTradingServiceServer struct {
}

func (UnimplementedTradingServiceServer) SetOrderStatusPaid(context.Context, *SetOrderStatusPaidRequest) (*SetOrderStatusPaidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOrderStatusPaid not implemented")
}
func (UnimplementedTradingServiceServer) GetOrderInvoice(context.Context, *GetOrderInvoiceRequest) (*GetOrderInvoiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderInvoice not implemented")
}
func (UnimplementedTradingServiceServer) mustEmbedUnimplementedTradingServiceServer() {}

// UnsafeTradingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradingServiceServer will
// result in compilation errors.
type UnsafeTradingServiceServer interface {
	mustEmbedUnimplementedTradingServiceServer()
}

func RegisterTradingServiceServer(s grpc.ServiceRegistrar, srv TradingServiceServer) {
	s.RegisterService(&TradingService_ServiceDesc, srv)
}

func _TradingService_SetOrderStatusPaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOrderStatusPaidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).SetOrderStatusPaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TradingService/SetOrderStatusPaid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).SetOrderStatusPaid(ctx, req.(*SetOrderStatusPaidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingService_GetOrderInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingServiceServer).GetOrderInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TradingService/GetOrderInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingServiceServer).GetOrderInvoice(ctx, req.(*GetOrderInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TradingService_ServiceDesc is the grpc.ServiceDesc for TradingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TradingService",
	HandlerType: (*TradingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetOrderStatusPaid",
			Handler:    _TradingService_SetOrderStatusPaid_Handler,
		},
		{
			MethodName: "GetOrderInvoice",
			Handler:    _TradingService_GetOrderInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trading-service.proto",
}
