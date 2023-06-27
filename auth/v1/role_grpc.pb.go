// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: auth/v1/role.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RoleService_All_FullMethodName                   = "/auth.RoleService/All"
	RoleService_List_FullMethodName                  = "/auth.RoleService/List"
	RoleService_Get_FullMethodName                   = "/auth.RoleService/Get"
	RoleService_Create_FullMethodName                = "/auth.RoleService/Create"
	RoleService_Update_FullMethodName                = "/auth.RoleService/Update"
	RoleService_Delete_FullMethodName                = "/auth.RoleService/Delete"
	RoleService_UpdateRolePermissions_FullMethodName = "/auth.RoleService/UpdateRolePermissions"
)

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	All(ctx context.Context, in *RoleAllRequest, opts ...grpc.CallOption) (*RoleAllResponse, error)
	List(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error)
	Get(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	Create(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateRolePermissions(ctx context.Context, in *UpdateRolePermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) All(ctx context.Context, in *RoleAllRequest, opts ...grpc.CallOption) (*RoleAllResponse, error) {
	out := new(RoleAllResponse)
	err := c.cc.Invoke(ctx, RoleService_All_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) List(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error) {
	out := new(RoleListResponse)
	err := c.cc.Invoke(ctx, RoleService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Get(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, RoleService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Create(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RoleService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Update(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RoleService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Delete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RoleService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpdateRolePermissions(ctx context.Context, in *UpdateRolePermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RoleService_UpdateRolePermissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations must embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	All(context.Context, *RoleAllRequest) (*RoleAllResponse, error)
	List(context.Context, *RoleListRequest) (*RoleListResponse, error)
	Get(context.Context, *RoleGetRequest) (*RoleResponse, error)
	Create(context.Context, *RoleCreateRequest) (*emptypb.Empty, error)
	Update(context.Context, *RoleUpdateRequest) (*emptypb.Empty, error)
	Delete(context.Context, *RoleDeleteRequest) (*emptypb.Empty, error)
	UpdateRolePermissions(context.Context, *UpdateRolePermissionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRoleServiceServer()
}

// UnimplementedRoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) All(context.Context, *RoleAllRequest) (*RoleAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedRoleServiceServer) List(context.Context, *RoleListRequest) (*RoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRoleServiceServer) Get(context.Context, *RoleGetRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRoleServiceServer) Create(context.Context, *RoleCreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRoleServiceServer) Update(context.Context, *RoleUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRoleServiceServer) Delete(context.Context, *RoleDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRoleServiceServer) UpdateRolePermissions(context.Context, *UpdateRolePermissionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRolePermissions not implemented")
}
func (UnimplementedRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_All_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).All(ctx, req.(*RoleAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).List(ctx, req.(*RoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Get(ctx, req.(*RoleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Create(ctx, req.(*RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Update(ctx, req.(*RoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Delete(ctx, req.(*RoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpdateRolePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRolePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpdateRolePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoleService_UpdateRolePermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpdateRolePermissions(ctx, req.(*UpdateRolePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "All",
			Handler:    _RoleService_All_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RoleService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RoleService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RoleService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RoleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RoleService_Delete_Handler,
		},
		{
			MethodName: "UpdateRolePermissions",
			Handler:    _RoleService_UpdateRolePermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/role.proto",
}
