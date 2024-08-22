// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: kewr/wms/query.proto

package wms

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
	Query_Params_FullMethodName       = "/kewr.wms.Query/Params"
	Query_Warehouse_FullMethodName    = "/kewr.wms.Query/Warehouse"
	Query_WarehouseAll_FullMethodName = "/kewr.wms.Query/WarehouseAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Warehouse items.
	Warehouse(ctx context.Context, in *QueryGetWarehouseRequest, opts ...grpc.CallOption) (*QueryGetWarehouseResponse, error)
	WarehouseAll(ctx context.Context, in *QueryAllWarehouseRequest, opts ...grpc.CallOption) (*QueryAllWarehouseResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Warehouse(ctx context.Context, in *QueryGetWarehouseRequest, opts ...grpc.CallOption) (*QueryGetWarehouseResponse, error) {
	out := new(QueryGetWarehouseResponse)
	err := c.cc.Invoke(ctx, Query_Warehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) WarehouseAll(ctx context.Context, in *QueryAllWarehouseRequest, opts ...grpc.CallOption) (*QueryAllWarehouseResponse, error) {
	out := new(QueryAllWarehouseResponse)
	err := c.cc.Invoke(ctx, Query_WarehouseAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Warehouse items.
	Warehouse(context.Context, *QueryGetWarehouseRequest) (*QueryGetWarehouseResponse, error)
	WarehouseAll(context.Context, *QueryAllWarehouseRequest) (*QueryAllWarehouseResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Warehouse(context.Context, *QueryGetWarehouseRequest) (*QueryGetWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Warehouse not implemented")
}
func (UnimplementedQueryServer) WarehouseAll(context.Context, *QueryAllWarehouseRequest) (*QueryAllWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WarehouseAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Warehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Warehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Warehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Warehouse(ctx, req.(*QueryGetWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_WarehouseAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllWarehouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WarehouseAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_WarehouseAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WarehouseAll(ctx, req.(*QueryAllWarehouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kewr.wms.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Warehouse",
			Handler:    _Query_Warehouse_Handler,
		},
		{
			MethodName: "WarehouseAll",
			Handler:    _Query_WarehouseAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kewr/wms/query.proto",
}
