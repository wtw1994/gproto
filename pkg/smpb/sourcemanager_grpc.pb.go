// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package smpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SourcemanagerClient is the client API for Sourcemanager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourcemanagerClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	List(ctx context.Context, in *ListsRequest, opts ...grpc.CallOption) (*ListsReply, error)
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*InfoReply, error)
	PingSource(ctx context.Context, in *PingSourceRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	SotCreate(ctx context.Context, in *SotCreateRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	SotUpdate(ctx context.Context, in *SotUpdateRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	SotDelete(ctx context.Context, in *SotDeleteRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	SotList(ctx context.Context, in *SotListsRequest, opts ...grpc.CallOption) (*SotListsReply, error)
	SotDescribe(ctx context.Context, in *SotDescribeRequest, opts ...grpc.CallOption) (*SotInfoReply, error)
	EngineMap(ctx context.Context, in *EngingMapRequest, opts ...grpc.CallOption) (*EngineMapReply, error)
}

type sourcemanagerClient struct {
	cc grpc.ClientConnInterface
}

func NewSourcemanagerClient(cc grpc.ClientConnInterface) SourcemanagerClient {
	return &sourcemanagerClient{cc}
}

func (c *sourcemanagerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) List(ctx context.Context, in *ListsRequest, opts ...grpc.CallOption) (*ListsReply, error) {
	out := new(ListsReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/Describe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) PingSource(ctx context.Context, in *PingSourceRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/PingSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) SotCreate(ctx context.Context, in *SotCreateRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/SotCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) SotUpdate(ctx context.Context, in *SotUpdateRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/SotUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) SotDelete(ctx context.Context, in *SotDeleteRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/SotDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) SotList(ctx context.Context, in *SotListsRequest, opts ...grpc.CallOption) (*SotListsReply, error) {
	out := new(SotListsReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/SotList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) SotDescribe(ctx context.Context, in *SotDescribeRequest, opts ...grpc.CallOption) (*SotInfoReply, error) {
	out := new(SotInfoReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/SotDescribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcemanagerClient) EngineMap(ctx context.Context, in *EngingMapRequest, opts ...grpc.CallOption) (*EngineMapReply, error) {
	out := new(EngineMapReply)
	err := c.cc.Invoke(ctx, "/smpb.Sourcemanager/EngineMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourcemanagerServer is the server API for Sourcemanager service.
// All implementations must embed UnimplementedSourcemanagerServer
// for forward compatibility
type SourcemanagerServer interface {
	Create(context.Context, *CreateRequest) (*EmptyReply, error)
	Update(context.Context, *UpdateRequest) (*EmptyReply, error)
	Delete(context.Context, *DeleteRequest) (*EmptyReply, error)
	List(context.Context, *ListsRequest) (*ListsReply, error)
	Describe(context.Context, *DescribeRequest) (*InfoReply, error)
	PingSource(context.Context, *PingSourceRequest) (*EmptyReply, error)
	SotCreate(context.Context, *SotCreateRequest) (*EmptyReply, error)
	SotUpdate(context.Context, *SotUpdateRequest) (*EmptyReply, error)
	SotDelete(context.Context, *SotDeleteRequest) (*EmptyReply, error)
	SotList(context.Context, *SotListsRequest) (*SotListsReply, error)
	SotDescribe(context.Context, *SotDescribeRequest) (*SotInfoReply, error)
	EngineMap(context.Context, *EngingMapRequest) (*EngineMapReply, error)
	mustEmbedUnimplementedSourcemanagerServer()
}

// UnimplementedSourcemanagerServer must be embedded to have forward compatible implementations.
type UnimplementedSourcemanagerServer struct {
}

func (UnimplementedSourcemanagerServer) Create(context.Context, *CreateRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSourcemanagerServer) Update(context.Context, *UpdateRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSourcemanagerServer) Delete(context.Context, *DeleteRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSourcemanagerServer) List(context.Context, *ListsRequest) (*ListsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSourcemanagerServer) Describe(context.Context, *DescribeRequest) (*InfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Describe not implemented")
}
func (UnimplementedSourcemanagerServer) PingSource(context.Context, *PingSourceRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingSource not implemented")
}
func (UnimplementedSourcemanagerServer) SotCreate(context.Context, *SotCreateRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SotCreate not implemented")
}
func (UnimplementedSourcemanagerServer) SotUpdate(context.Context, *SotUpdateRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SotUpdate not implemented")
}
func (UnimplementedSourcemanagerServer) SotDelete(context.Context, *SotDeleteRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SotDelete not implemented")
}
func (UnimplementedSourcemanagerServer) SotList(context.Context, *SotListsRequest) (*SotListsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SotList not implemented")
}
func (UnimplementedSourcemanagerServer) SotDescribe(context.Context, *SotDescribeRequest) (*SotInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SotDescribe not implemented")
}
func (UnimplementedSourcemanagerServer) EngineMap(context.Context, *EngingMapRequest) (*EngineMapReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EngineMap not implemented")
}
func (UnimplementedSourcemanagerServer) mustEmbedUnimplementedSourcemanagerServer() {}

// UnsafeSourcemanagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourcemanagerServer will
// result in compilation errors.
type UnsafeSourcemanagerServer interface {
	mustEmbedUnimplementedSourcemanagerServer()
}

func RegisterSourcemanagerServer(s grpc.ServiceRegistrar, srv SourcemanagerServer) {
	s.RegisterService(&_Sourcemanager_serviceDesc, srv)
}

func _Sourcemanager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).List(ctx, req.(*ListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_PingSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).PingSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/PingSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).PingSource(ctx, req.(*PingSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_SotCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SotCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).SotCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/SotCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).SotCreate(ctx, req.(*SotCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_SotUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SotUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).SotUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/SotUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).SotUpdate(ctx, req.(*SotUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_SotDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SotDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).SotDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/SotDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).SotDelete(ctx, req.(*SotDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_SotList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SotListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).SotList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/SotList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).SotList(ctx, req.(*SotListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_SotDescribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SotDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).SotDescribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/SotDescribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).SotDescribe(ctx, req.(*SotDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sourcemanager_EngineMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EngingMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcemanagerServer).EngineMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smpb.Sourcemanager/EngineMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcemanagerServer).EngineMap(ctx, req.(*EngingMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sourcemanager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smpb.Sourcemanager",
	HandlerType: (*SourcemanagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Sourcemanager_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Sourcemanager_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Sourcemanager_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Sourcemanager_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Sourcemanager_Describe_Handler,
		},
		{
			MethodName: "PingSource",
			Handler:    _Sourcemanager_PingSource_Handler,
		},
		{
			MethodName: "SotCreate",
			Handler:    _Sourcemanager_SotCreate_Handler,
		},
		{
			MethodName: "SotUpdate",
			Handler:    _Sourcemanager_SotUpdate_Handler,
		},
		{
			MethodName: "SotDelete",
			Handler:    _Sourcemanager_SotDelete_Handler,
		},
		{
			MethodName: "SotList",
			Handler:    _Sourcemanager_SotList_Handler,
		},
		{
			MethodName: "SotDescribe",
			Handler:    _Sourcemanager_SotDescribe_Handler,
		},
		{
			MethodName: "EngineMap",
			Handler:    _Sourcemanager_EngineMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sourcemanager.proto",
}