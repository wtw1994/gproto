// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wspb

import (
	context "context"
	model "github.com/DataWorkbench/gproto/pkg/model"
	request "github.com/DataWorkbench/gproto/pkg/request"
	response "github.com/DataWorkbench/gproto/pkg/response"
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
	// API of workspace manager.
	ListWorkspaces(ctx context.Context, in *request.ListWorkspaces, opts ...grpc.CallOption) (*response.ListWorkspaces, error)
	// DeleteWorkspaces delete the specified workspaces and its resources;
	// Resources includes
	//  - Members.
	//  - Audit(Operation Records).
	DeleteWorkspaces(ctx context.Context, in *request.DeleteWorkspaces, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	DisableWorkspaces(ctx context.Context, in *request.DisableWorkspaces, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	EnableWorkspaces(ctx context.Context, in *request.EnableWorkspaces, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	CreateWorkspace(ctx context.Context, in *request.CreateWorkspace, opts ...grpc.CallOption) (*response.CreateWorkspace, error)
	UpdateWorkspace(ctx context.Context, in *request.UpdateWorkspace, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	DescribeWorkspace(ctx context.Context, in *request.DescribeWorkspace, opts ...grpc.CallOption) (*response.DescribeWorkspace, error)
	// API of workspace operation audit log.
	AddAudit(ctx context.Context, in *request.AddAudit, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	ListAudits(ctx context.Context, in *request.ListAudits, opts ...grpc.CallOption) (*response.ListAudits, error)
	// API of workspace role.
	ListSystemRoles(ctx context.Context, in *request.ListSystemRoles, opts ...grpc.CallOption) (*response.ListSystemRoles, error)
	// API of workspace member.
	ListMembers(ctx context.Context, in *request.ListMembers, opts ...grpc.CallOption) (*response.ListMembers, error)
	UpsertMembers(ctx context.Context, in *request.UpsertMembers, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	DeleteMembers(ctx context.Context, in *request.DeleteMembers, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	// Permission Auth.
	CheckPermission(ctx context.Context, in *request.CheckPermission, opts ...grpc.CallOption) (*model.EmptyStruct, error)
}

type workspaceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceClient(cc grpc.ClientConnInterface) WorkspaceClient {
	return &workspaceClient{cc}
}

func (c *workspaceClient) ListWorkspaces(ctx context.Context, in *request.ListWorkspaces, opts ...grpc.CallOption) (*response.ListWorkspaces, error) {
	out := new(response.ListWorkspaces)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/ListWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) DeleteWorkspaces(ctx context.Context, in *request.DeleteWorkspaces, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/DeleteWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) DisableWorkspaces(ctx context.Context, in *request.DisableWorkspaces, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/DisableWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) EnableWorkspaces(ctx context.Context, in *request.EnableWorkspaces, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/EnableWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) CreateWorkspace(ctx context.Context, in *request.CreateWorkspace, opts ...grpc.CallOption) (*response.CreateWorkspace, error) {
	out := new(response.CreateWorkspace)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/CreateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) UpdateWorkspace(ctx context.Context, in *request.UpdateWorkspace, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/UpdateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) DescribeWorkspace(ctx context.Context, in *request.DescribeWorkspace, opts ...grpc.CallOption) (*response.DescribeWorkspace, error) {
	out := new(response.DescribeWorkspace)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/DescribeWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) AddAudit(ctx context.Context, in *request.AddAudit, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/AddAudit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) ListAudits(ctx context.Context, in *request.ListAudits, opts ...grpc.CallOption) (*response.ListAudits, error) {
	out := new(response.ListAudits)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/ListAudits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) ListSystemRoles(ctx context.Context, in *request.ListSystemRoles, opts ...grpc.CallOption) (*response.ListSystemRoles, error) {
	out := new(response.ListSystemRoles)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/ListSystemRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) ListMembers(ctx context.Context, in *request.ListMembers, opts ...grpc.CallOption) (*response.ListMembers, error) {
	out := new(response.ListMembers)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/ListMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) UpsertMembers(ctx context.Context, in *request.UpsertMembers, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/UpsertMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) DeleteMembers(ctx context.Context, in *request.DeleteMembers, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/DeleteMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceClient) CheckPermission(ctx context.Context, in *request.CheckPermission, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/wspb.Workspace/CheckPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServer is the server API for Workspace service.
// All implementations must embed UnimplementedWorkspaceServer
// for forward compatibility
type WorkspaceServer interface {
	// API of workspace manager.
	ListWorkspaces(context.Context, *request.ListWorkspaces) (*response.ListWorkspaces, error)
	// DeleteWorkspaces delete the specified workspaces and its resources;
	// Resources includes
	//  - Members.
	//  - Audit(Operation Records).
	DeleteWorkspaces(context.Context, *request.DeleteWorkspaces) (*model.EmptyStruct, error)
	DisableWorkspaces(context.Context, *request.DisableWorkspaces) (*model.EmptyStruct, error)
	EnableWorkspaces(context.Context, *request.EnableWorkspaces) (*model.EmptyStruct, error)
	CreateWorkspace(context.Context, *request.CreateWorkspace) (*response.CreateWorkspace, error)
	UpdateWorkspace(context.Context, *request.UpdateWorkspace) (*model.EmptyStruct, error)
	DescribeWorkspace(context.Context, *request.DescribeWorkspace) (*response.DescribeWorkspace, error)
	// API of workspace operation audit log.
	AddAudit(context.Context, *request.AddAudit) (*model.EmptyStruct, error)
	ListAudits(context.Context, *request.ListAudits) (*response.ListAudits, error)
	// API of workspace role.
	ListSystemRoles(context.Context, *request.ListSystemRoles) (*response.ListSystemRoles, error)
	// API of workspace member.
	ListMembers(context.Context, *request.ListMembers) (*response.ListMembers, error)
	UpsertMembers(context.Context, *request.UpsertMembers) (*model.EmptyStruct, error)
	DeleteMembers(context.Context, *request.DeleteMembers) (*model.EmptyStruct, error)
	// Permission Auth.
	CheckPermission(context.Context, *request.CheckPermission) (*model.EmptyStruct, error)
	mustEmbedUnimplementedWorkspaceServer()
}

// UnimplementedWorkspaceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServer struct {
}

func (UnimplementedWorkspaceServer) ListWorkspaces(context.Context, *request.ListWorkspaces) (*response.ListWorkspaces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaces not implemented")
}
func (UnimplementedWorkspaceServer) DeleteWorkspaces(context.Context, *request.DeleteWorkspaces) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspaces not implemented")
}
func (UnimplementedWorkspaceServer) DisableWorkspaces(context.Context, *request.DisableWorkspaces) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableWorkspaces not implemented")
}
func (UnimplementedWorkspaceServer) EnableWorkspaces(context.Context, *request.EnableWorkspaces) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableWorkspaces not implemented")
}
func (UnimplementedWorkspaceServer) CreateWorkspace(context.Context, *request.CreateWorkspace) (*response.CreateWorkspace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}
func (UnimplementedWorkspaceServer) UpdateWorkspace(context.Context, *request.UpdateWorkspace) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspace not implemented")
}
func (UnimplementedWorkspaceServer) DescribeWorkspace(context.Context, *request.DescribeWorkspace) (*response.DescribeWorkspace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeWorkspace not implemented")
}
func (UnimplementedWorkspaceServer) AddAudit(context.Context, *request.AddAudit) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAudit not implemented")
}
func (UnimplementedWorkspaceServer) ListAudits(context.Context, *request.ListAudits) (*response.ListAudits, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAudits not implemented")
}
func (UnimplementedWorkspaceServer) ListSystemRoles(context.Context, *request.ListSystemRoles) (*response.ListSystemRoles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSystemRoles not implemented")
}
func (UnimplementedWorkspaceServer) ListMembers(context.Context, *request.ListMembers) (*response.ListMembers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}
func (UnimplementedWorkspaceServer) UpsertMembers(context.Context, *request.UpsertMembers) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertMembers not implemented")
}
func (UnimplementedWorkspaceServer) DeleteMembers(context.Context, *request.DeleteMembers) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMembers not implemented")
}
func (UnimplementedWorkspaceServer) CheckPermission(context.Context, *request.CheckPermission) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
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

func _Workspace_ListWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).ListWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/ListWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).ListWorkspaces(ctx, req.(*request.ListWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_DeleteWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).DeleteWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/DeleteWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).DeleteWorkspaces(ctx, req.(*request.DeleteWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_DisableWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DisableWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).DisableWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/DisableWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).DisableWorkspaces(ctx, req.(*request.DisableWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_EnableWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EnableWorkspaces)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).EnableWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/EnableWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).EnableWorkspaces(ctx, req.(*request.EnableWorkspaces))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateWorkspace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/CreateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).CreateWorkspace(ctx, req.(*request.CreateWorkspace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_UpdateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateWorkspace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).UpdateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/UpdateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).UpdateWorkspace(ctx, req.(*request.UpdateWorkspace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_DescribeWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DescribeWorkspace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).DescribeWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/DescribeWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).DescribeWorkspace(ctx, req.(*request.DescribeWorkspace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_AddAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.AddAudit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).AddAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/AddAudit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).AddAudit(ctx, req.(*request.AddAudit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_ListAudits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListAudits)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).ListAudits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/ListAudits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).ListAudits(ctx, req.(*request.ListAudits))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_ListSystemRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListSystemRoles)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).ListSystemRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/ListSystemRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).ListSystemRoles(ctx, req.(*request.ListSystemRoles))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListMembers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/ListMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).ListMembers(ctx, req.(*request.ListMembers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_UpsertMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpsertMembers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).UpsertMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/UpsertMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).UpsertMembers(ctx, req.(*request.UpsertMembers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_DeleteMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteMembers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).DeleteMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/DeleteMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).DeleteMembers(ctx, req.(*request.DeleteMembers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workspace_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CheckPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wspb.Workspace/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).CheckPermission(ctx, req.(*request.CheckPermission))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workspace_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wspb.Workspace",
	HandlerType: (*WorkspaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWorkspaces",
			Handler:    _Workspace_ListWorkspaces_Handler,
		},
		{
			MethodName: "DeleteWorkspaces",
			Handler:    _Workspace_DeleteWorkspaces_Handler,
		},
		{
			MethodName: "DisableWorkspaces",
			Handler:    _Workspace_DisableWorkspaces_Handler,
		},
		{
			MethodName: "EnableWorkspaces",
			Handler:    _Workspace_EnableWorkspaces_Handler,
		},
		{
			MethodName: "CreateWorkspace",
			Handler:    _Workspace_CreateWorkspace_Handler,
		},
		{
			MethodName: "UpdateWorkspace",
			Handler:    _Workspace_UpdateWorkspace_Handler,
		},
		{
			MethodName: "DescribeWorkspace",
			Handler:    _Workspace_DescribeWorkspace_Handler,
		},
		{
			MethodName: "AddAudit",
			Handler:    _Workspace_AddAudit_Handler,
		},
		{
			MethodName: "ListAudits",
			Handler:    _Workspace_ListAudits_Handler,
		},
		{
			MethodName: "ListSystemRoles",
			Handler:    _Workspace_ListSystemRoles_Handler,
		},
		{
			MethodName: "ListMembers",
			Handler:    _Workspace_ListMembers_Handler,
		},
		{
			MethodName: "UpsertMembers",
			Handler:    _Workspace_UpsertMembers_Handler,
		},
		{
			MethodName: "DeleteMembers",
			Handler:    _Workspace_DeleteMembers_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _Workspace_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/workspace.proto",
}
