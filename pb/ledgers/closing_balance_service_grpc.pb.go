// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: ledgers/closing_balance_service.proto

package ledgers

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

// ClosingBalanceServiceClient is the client API for ClosingBalanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClosingBalanceServiceClient interface {
	Closing(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*BoolMessage, error)
	ViewAccountBalance(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*AccountBalance, error)
	ViewDebtBalance(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*DebtBalance, error)
	ViewArBalance(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*ArBalance, error)
	ListAccountBalance(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*AccountBalances, error)
	ListDebtBalance(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*DebtBalances, error)
	ListArBalance(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*ArBalances, error)
}

type closingBalanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClosingBalanceServiceClient(cc grpc.ClientConnInterface) ClosingBalanceServiceClient {
	return &closingBalanceServiceClient{cc}
}

func (c *closingBalanceServiceClient) Closing(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*BoolMessage, error) {
	out := new(BoolMessage)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ClosingBalanceService/Closing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *closingBalanceServiceClient) ViewAccountBalance(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*AccountBalance, error) {
	out := new(AccountBalance)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ClosingBalanceService/ViewAccountBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *closingBalanceServiceClient) ViewDebtBalance(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*DebtBalance, error) {
	out := new(DebtBalance)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ClosingBalanceService/ViewDebtBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *closingBalanceServiceClient) ViewArBalance(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*ArBalance, error) {
	out := new(ArBalance)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ClosingBalanceService/ViewArBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *closingBalanceServiceClient) ListAccountBalance(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*AccountBalances, error) {
	out := new(AccountBalances)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ClosingBalanceService/ListAccountBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *closingBalanceServiceClient) ListDebtBalance(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*DebtBalances, error) {
	out := new(DebtBalances)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ClosingBalanceService/ListDebtBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *closingBalanceServiceClient) ListArBalance(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*ArBalances, error) {
	out := new(ArBalances)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ClosingBalanceService/ListArBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClosingBalanceServiceServer is the server API for ClosingBalanceService service.
// All implementations must embed UnimplementedClosingBalanceServiceServer
// for forward compatibility
type ClosingBalanceServiceServer interface {
	Closing(context.Context, *StringMessage) (*BoolMessage, error)
	ViewAccountBalance(context.Context, *StringMessage) (*AccountBalance, error)
	ViewDebtBalance(context.Context, *StringMessage) (*DebtBalance, error)
	ViewArBalance(context.Context, *StringMessage) (*ArBalance, error)
	ListAccountBalance(context.Context, *EmptyMessage) (*AccountBalances, error)
	ListDebtBalance(context.Context, *EmptyMessage) (*DebtBalances, error)
	ListArBalance(context.Context, *EmptyMessage) (*ArBalances, error)
	mustEmbedUnimplementedClosingBalanceServiceServer()
}

// UnimplementedClosingBalanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClosingBalanceServiceServer struct {
}

func (UnimplementedClosingBalanceServiceServer) Closing(context.Context, *StringMessage) (*BoolMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Closing not implemented")
}
func (UnimplementedClosingBalanceServiceServer) ViewAccountBalance(context.Context, *StringMessage) (*AccountBalance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewAccountBalance not implemented")
}
func (UnimplementedClosingBalanceServiceServer) ViewDebtBalance(context.Context, *StringMessage) (*DebtBalance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewDebtBalance not implemented")
}
func (UnimplementedClosingBalanceServiceServer) ViewArBalance(context.Context, *StringMessage) (*ArBalance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewArBalance not implemented")
}
func (UnimplementedClosingBalanceServiceServer) ListAccountBalance(context.Context, *EmptyMessage) (*AccountBalances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccountBalance not implemented")
}
func (UnimplementedClosingBalanceServiceServer) ListDebtBalance(context.Context, *EmptyMessage) (*DebtBalances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDebtBalance not implemented")
}
func (UnimplementedClosingBalanceServiceServer) ListArBalance(context.Context, *EmptyMessage) (*ArBalances, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArBalance not implemented")
}
func (UnimplementedClosingBalanceServiceServer) mustEmbedUnimplementedClosingBalanceServiceServer() {}

// UnsafeClosingBalanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClosingBalanceServiceServer will
// result in compilation errors.
type UnsafeClosingBalanceServiceServer interface {
	mustEmbedUnimplementedClosingBalanceServiceServer()
}

func RegisterClosingBalanceServiceServer(s grpc.ServiceRegistrar, srv ClosingBalanceServiceServer) {
	s.RegisterService(&ClosingBalanceService_ServiceDesc, srv)
}

func _ClosingBalanceService_Closing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClosingBalanceServiceServer).Closing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ClosingBalanceService/Closing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClosingBalanceServiceServer).Closing(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClosingBalanceService_ViewAccountBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClosingBalanceServiceServer).ViewAccountBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ClosingBalanceService/ViewAccountBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClosingBalanceServiceServer).ViewAccountBalance(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClosingBalanceService_ViewDebtBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClosingBalanceServiceServer).ViewDebtBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ClosingBalanceService/ViewDebtBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClosingBalanceServiceServer).ViewDebtBalance(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClosingBalanceService_ViewArBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClosingBalanceServiceServer).ViewArBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ClosingBalanceService/ViewArBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClosingBalanceServiceServer).ViewArBalance(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClosingBalanceService_ListAccountBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClosingBalanceServiceServer).ListAccountBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ClosingBalanceService/ListAccountBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClosingBalanceServiceServer).ListAccountBalance(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClosingBalanceService_ListDebtBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClosingBalanceServiceServer).ListDebtBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ClosingBalanceService/ListDebtBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClosingBalanceServiceServer).ListDebtBalance(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClosingBalanceService_ListArBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClosingBalanceServiceServer).ListArBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ClosingBalanceService/ListArBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClosingBalanceServiceServer).ListArBalance(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ClosingBalanceService_ServiceDesc is the grpc.ServiceDesc for ClosingBalanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClosingBalanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.ledgers.ClosingBalanceService",
	HandlerType: (*ClosingBalanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Closing",
			Handler:    _ClosingBalanceService_Closing_Handler,
		},
		{
			MethodName: "ViewAccountBalance",
			Handler:    _ClosingBalanceService_ViewAccountBalance_Handler,
		},
		{
			MethodName: "ViewDebtBalance",
			Handler:    _ClosingBalanceService_ViewDebtBalance_Handler,
		},
		{
			MethodName: "ViewArBalance",
			Handler:    _ClosingBalanceService_ViewArBalance_Handler,
		},
		{
			MethodName: "ListAccountBalance",
			Handler:    _ClosingBalanceService_ListAccountBalance_Handler,
		},
		{
			MethodName: "ListDebtBalance",
			Handler:    _ClosingBalanceService_ListDebtBalance_Handler,
		},
		{
			MethodName: "ListArBalance",
			Handler:    _ClosingBalanceService_ListArBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ledgers/closing_balance_service.proto",
}
