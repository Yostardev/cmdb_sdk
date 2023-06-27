// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: model/v1/model_group.proto

package model

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
	ModelGroupService_All_FullMethodName    = "/model.ModelGroupService/All"
	ModelGroupService_List_FullMethodName   = "/model.ModelGroupService/List"
	ModelGroupService_Get_FullMethodName    = "/model.ModelGroupService/Get"
	ModelGroupService_Create_FullMethodName = "/model.ModelGroupService/Create"
	ModelGroupService_Update_FullMethodName = "/model.ModelGroupService/Update"
	ModelGroupService_Delete_FullMethodName = "/model.ModelGroupService/Delete"
)

// ModelGroupServiceClient is the client API for ModelGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelGroupServiceClient interface {
	All(ctx context.Context, in *ModelGroupAllRequest, opts ...grpc.CallOption) (*ModelGroupAllResponse, error)
	List(ctx context.Context, in *ModelGroupListRequest, opts ...grpc.CallOption) (*ModelGroupListResponse, error)
	Get(ctx context.Context, in *ModelGroupGetRequest, opts ...grpc.CallOption) (*ModelGroupResponse, error)
	Create(ctx context.Context, in *ModelGroupCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *ModelGroupUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *ModelGroupDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type modelGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelGroupServiceClient(cc grpc.ClientConnInterface) ModelGroupServiceClient {
	return &modelGroupServiceClient{cc}
}

func (c *modelGroupServiceClient) All(ctx context.Context, in *ModelGroupAllRequest, opts ...grpc.CallOption) (*ModelGroupAllResponse, error) {
	out := new(ModelGroupAllResponse)
	err := c.cc.Invoke(ctx, ModelGroupService_All_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelGroupServiceClient) List(ctx context.Context, in *ModelGroupListRequest, opts ...grpc.CallOption) (*ModelGroupListResponse, error) {
	out := new(ModelGroupListResponse)
	err := c.cc.Invoke(ctx, ModelGroupService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelGroupServiceClient) Get(ctx context.Context, in *ModelGroupGetRequest, opts ...grpc.CallOption) (*ModelGroupResponse, error) {
	out := new(ModelGroupResponse)
	err := c.cc.Invoke(ctx, ModelGroupService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelGroupServiceClient) Create(ctx context.Context, in *ModelGroupCreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ModelGroupService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelGroupServiceClient) Update(ctx context.Context, in *ModelGroupUpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ModelGroupService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelGroupServiceClient) Delete(ctx context.Context, in *ModelGroupDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ModelGroupService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelGroupServiceServer is the server API for ModelGroupService service.
// All implementations must embed UnimplementedModelGroupServiceServer
// for forward compatibility
type ModelGroupServiceServer interface {
	All(context.Context, *ModelGroupAllRequest) (*ModelGroupAllResponse, error)
	List(context.Context, *ModelGroupListRequest) (*ModelGroupListResponse, error)
	Get(context.Context, *ModelGroupGetRequest) (*ModelGroupResponse, error)
	Create(context.Context, *ModelGroupCreateRequest) (*emptypb.Empty, error)
	Update(context.Context, *ModelGroupUpdateRequest) (*emptypb.Empty, error)
	Delete(context.Context, *ModelGroupDeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedModelGroupServiceServer()
}

// UnimplementedModelGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModelGroupServiceServer struct {
}

func (UnimplementedModelGroupServiceServer) All(context.Context, *ModelGroupAllRequest) (*ModelGroupAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedModelGroupServiceServer) List(context.Context, *ModelGroupListRequest) (*ModelGroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedModelGroupServiceServer) Get(context.Context, *ModelGroupGetRequest) (*ModelGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedModelGroupServiceServer) Create(context.Context, *ModelGroupCreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedModelGroupServiceServer) Update(context.Context, *ModelGroupUpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedModelGroupServiceServer) Delete(context.Context, *ModelGroupDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedModelGroupServiceServer) mustEmbedUnimplementedModelGroupServiceServer() {}

// UnsafeModelGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelGroupServiceServer will
// result in compilation errors.
type UnsafeModelGroupServiceServer interface {
	mustEmbedUnimplementedModelGroupServiceServer()
}

func RegisterModelGroupServiceServer(s grpc.ServiceRegistrar, srv ModelGroupServiceServer) {
	s.RegisterService(&ModelGroupService_ServiceDesc, srv)
}

func _ModelGroupService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelGroupAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelGroupServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelGroupService_All_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelGroupServiceServer).All(ctx, req.(*ModelGroupAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelGroupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelGroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelGroupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelGroupService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelGroupServiceServer).List(ctx, req.(*ModelGroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelGroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelGroupGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelGroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelGroupService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelGroupServiceServer).Get(ctx, req.(*ModelGroupGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelGroupCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelGroupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelGroupServiceServer).Create(ctx, req.(*ModelGroupCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelGroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelGroupUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelGroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelGroupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelGroupServiceServer).Update(ctx, req.(*ModelGroupUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelGroupDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelGroupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelGroupServiceServer).Delete(ctx, req.(*ModelGroupDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelGroupService_ServiceDesc is the grpc.ServiceDesc for ModelGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.ModelGroupService",
	HandlerType: (*ModelGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "All",
			Handler:    _ModelGroupService_All_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ModelGroupService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ModelGroupService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ModelGroupService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ModelGroupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ModelGroupService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/v1/model_group.proto",
}
