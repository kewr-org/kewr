// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: kewr/wms/tx.proto

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
	Msg_UpdateParams_FullMethodName    = "/kewr.wms.Msg/UpdateParams"
	Msg_CreateWarehouse_FullMethodName = "/kewr.wms.Msg/CreateWarehouse"
	Msg_UpdateWarehouse_FullMethodName = "/kewr.wms.Msg/UpdateWarehouse"
	Msg_DeleteWarehouse_FullMethodName = "/kewr.wms.Msg/DeleteWarehouse"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateWarehouse(ctx context.Context, in *MsgCreateWarehouse, opts ...grpc.CallOption) (*MsgCreateWarehouseResponse, error)
	UpdateWarehouse(ctx context.Context, in *MsgUpdateWarehouse, opts ...grpc.CallOption) (*MsgUpdateWarehouseResponse, error)
	DeleteWarehouse(ctx context.Context, in *MsgDeleteWarehouse, opts ...grpc.CallOption) (*MsgDeleteWarehouseResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateWarehouse(ctx context.Context, in *MsgCreateWarehouse, opts ...grpc.CallOption) (*MsgCreateWarehouseResponse, error) {
	out := new(MsgCreateWarehouseResponse)
	err := c.cc.Invoke(ctx, Msg_CreateWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateWarehouse(ctx context.Context, in *MsgUpdateWarehouse, opts ...grpc.CallOption) (*MsgUpdateWarehouseResponse, error) {
	out := new(MsgUpdateWarehouseResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteWarehouse(ctx context.Context, in *MsgDeleteWarehouse, opts ...grpc.CallOption) (*MsgDeleteWarehouseResponse, error) {
	out := new(MsgDeleteWarehouseResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteWarehouse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateWarehouse(context.Context, *MsgCreateWarehouse) (*MsgCreateWarehouseResponse, error)
	UpdateWarehouse(context.Context, *MsgUpdateWarehouse) (*MsgUpdateWarehouseResponse, error)
	DeleteWarehouse(context.Context, *MsgDeleteWarehouse) (*MsgDeleteWarehouseResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateWarehouse(context.Context, *MsgCreateWarehouse) (*MsgCreateWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWarehouse not implemented")
}
func (UnimplementedMsgServer) UpdateWarehouse(context.Context, *MsgUpdateWarehouse) (*MsgUpdateWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWarehouse not implemented")
}
func (UnimplementedMsgServer) DeleteWarehouse(context.Context, *MsgDeleteWarehouse) (*MsgDeleteWarehouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWarehouse not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateWarehouse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateWarehouse(ctx, req.(*MsgCreateWarehouse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateWarehouse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateWarehouse(ctx, req.(*MsgUpdateWarehouse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteWarehouse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteWarehouse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteWarehouse(ctx, req.(*MsgDeleteWarehouse))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kewr.wms.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateWarehouse",
			Handler:    _Msg_CreateWarehouse_Handler,
		},
		{
			MethodName: "UpdateWarehouse",
			Handler:    _Msg_UpdateWarehouse_Handler,
		},
		{
			MethodName: "DeleteWarehouse",
			Handler:    _Msg_DeleteWarehouse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kewr/wms/tx.proto",
}
