// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WorkspaceClient is the client API for Workspace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceClient interface {
	Lists(ctx context.Context, in *ListsRequest, opts ...grpc.CallOption) (*ListsReply, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeReply, error)
	Disable(ctx context.Context, in *DisableRequest, opts ...grpc.CallOption) (*EmptyReply, error)
	Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EmptyReply, error)
}

type workspaceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceClient(cc grpc.ClientConnInterface) WorkspaceClient {
	return &workspaceClient{cc}
}

func (c *workspaceClient) Lists(ctx context.Context, in *ListsRequest, opts ...grpc.CallOption) (*ListsReply, error) {
	out := new(ListsReply)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/Lists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeReply, error) {
	out := new(DescribeReply)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/Describe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) Disable(ctx context.Context, in *DisableRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/Disable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/Enable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServer is the server API for Workspace service.
// All implementations must embed UnimplementedWorkspaceServer
// for forward compatibility
type WorkspaceServer interface {
	Lists(context.Context, *ListsRequest) (*ListsReply, error)
	Create(context.Context, *CreateRequest) (*EmptyReply, error)
	Delete(context.Context, *DeleteRequest) (*EmptyReply, error)
	Update(context.Context, *UpdateRequest) (*EmptyReply, error)
	Describe(context.Context, *DescribeRequest) (*DescribeReply, error)
	Disable(context.Context, *DisableRequest) (*EmptyReply, error)
	Enable(context.Context, *EnableRequest) (*EmptyReply, error)
	mustEmbedUnimplementedWorkspaceServer()
}

// UnimplementedWorkspaceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServer struct {
}

func (UnimplementedWorkspaceServer) Lists(context.Context, *ListsRequest) (*ListsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lists not implemented")
}
func (UnimplementedWorkspaceServer) Create(context.Context, *CreateRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWorkspaceServer) Delete(context.Context, *DeleteRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWorkspaceServer) Update(context.Context, *UpdateRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWorkspaceServer) Describe(context.Context, *DescribeRequest) (*DescribeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Describe not implemented")
}
func (UnimplementedWorkspaceServer) Disable(context.Context, *DisableRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedWorkspaceServer) Enable(context.Context, *EnableRequest) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedWorkspaceServer) mustEmbedUnimplementedWorkspaceServer() {}

// UnsafeWorkspaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServer will
// result in compilation errors.
type UnsafeWorkspaceServer interface {
	mustEmbedUnimplementedWorkspaceServer()
}

func RegisterWorkspaceServer(s grpc.ServiceRegistrar, srv WorkspaceServer) {
	s.RegisterService(&_Workspace_serviceDesc, srv)
}

func _Workspace_Lists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).Lists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/Lists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).Lists(ctx, req.(*ListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/Disable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).Disable(ctx, req.(*DisableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/Enable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).Enable(ctx, req.(*EnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workspace_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wspb.Workspace",
	HandlerType: (*WorkspaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lists",
			Handler:    _Workspace_Lists_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Workspace_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Workspace_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Workspace_Update_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Workspace_Describe_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _Workspace_Disable_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _Workspace_Enable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/workspace.proto",
}
