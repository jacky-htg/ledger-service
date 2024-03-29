// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: ledgers/chart_of_account_service.proto

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

// ChartOfAccountServiceClient is the client API for ChartOfAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChartOfAccountServiceClient interface {
	Create(ctx context.Context, in *ChartOfAccount, opts ...grpc.CallOption) (*ChartOfAccount, error)
	Update(ctx context.Context, in *ChartOfAccount, opts ...grpc.CallOption) (*ChartOfAccount, error)
	List(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*ChartOfAccounts, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*BoolMessage, error)
}

type chartOfAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChartOfAccountServiceClient(cc grpc.ClientConnInterface) ChartOfAccountServiceClient {
	return &chartOfAccountServiceClient{cc}
}

func (c *chartOfAccountServiceClient) Create(ctx context.Context, in *ChartOfAccount, opts ...grpc.CallOption) (*ChartOfAccount, error) {
	out := new(ChartOfAccount)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ChartOfAccountService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartOfAccountServiceClient) Update(ctx context.Context, in *ChartOfAccount, opts ...grpc.CallOption) (*ChartOfAccount, error) {
	out := new(ChartOfAccount)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ChartOfAccountService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartOfAccountServiceClient) List(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*ChartOfAccounts, error) {
	out := new(ChartOfAccounts)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ChartOfAccountService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartOfAccountServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*BoolMessage, error) {
	out := new(BoolMessage)
	err := c.cc.Invoke(ctx, "/wiradata.ledgers.ChartOfAccountService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChartOfAccountServiceServer is the server API for ChartOfAccountService service.
// All implementations must embed UnimplementedChartOfAccountServiceServer
// for forward compatibility
type ChartOfAccountServiceServer interface {
	Create(context.Context, *ChartOfAccount) (*ChartOfAccount, error)
	Update(context.Context, *ChartOfAccount) (*ChartOfAccount, error)
	List(context.Context, *EmptyMessage) (*ChartOfAccounts, error)
	Delete(context.Context, *Id) (*BoolMessage, error)
	mustEmbedUnimplementedChartOfAccountServiceServer()
}

// UnimplementedChartOfAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChartOfAccountServiceServer struct {
}

func (UnimplementedChartOfAccountServiceServer) Create(context.Context, *ChartOfAccount) (*ChartOfAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChartOfAccountServiceServer) Update(context.Context, *ChartOfAccount) (*ChartOfAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedChartOfAccountServiceServer) List(context.Context, *EmptyMessage) (*ChartOfAccounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedChartOfAccountServiceServer) Delete(context.Context, *Id) (*BoolMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedChartOfAccountServiceServer) mustEmbedUnimplementedChartOfAccountServiceServer() {}

// UnsafeChartOfAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChartOfAccountServiceServer will
// result in compilation errors.
type UnsafeChartOfAccountServiceServer interface {
	mustEmbedUnimplementedChartOfAccountServiceServer()
}

func RegisterChartOfAccountServiceServer(s grpc.ServiceRegistrar, srv ChartOfAccountServiceServer) {
	s.RegisterService(&ChartOfAccountService_ServiceDesc, srv)
}

func _ChartOfAccountService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChartOfAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartOfAccountServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ChartOfAccountService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartOfAccountServiceServer).Create(ctx, req.(*ChartOfAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartOfAccountService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChartOfAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartOfAccountServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ChartOfAccountService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartOfAccountServiceServer).Update(ctx, req.(*ChartOfAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartOfAccountService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartOfAccountServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ChartOfAccountService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartOfAccountServiceServer).List(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartOfAccountService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartOfAccountServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiradata.ledgers.ChartOfAccountService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartOfAccountServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// ChartOfAccountService_ServiceDesc is the grpc.ServiceDesc for ChartOfAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChartOfAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.ledgers.ChartOfAccountService",
	HandlerType: (*ChartOfAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChartOfAccountService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ChartOfAccountService_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ChartOfAccountService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ChartOfAccountService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ledgers/chart_of_account_service.proto",
}
