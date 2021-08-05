// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/response.proto

package response

import (
	model "github.com/DataWorkbench/gproto/pkg/model"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// ListWorkspaces used as reply parameters in RPC or response body in HTTP.
type ListWorkspaces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos   []*model.Workspace `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
	HasMore bool               `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
}

func (x *ListWorkspaces) Reset() {
	*x = ListWorkspaces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkspaces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkspaces) ProtoMessage() {}

func (x *ListWorkspaces) ProtoReflect() protoreflect.Message {
	mi := &file_proto_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkspaces.ProtoReflect.Descriptor instead.
func (*ListWorkspaces) Descriptor() ([]byte, []int) {
	return file_proto_response_proto_rawDescGZIP(), []int{0}
}

func (x *ListWorkspaces) GetInfos() []*model.Workspace {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *ListWorkspaces) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

// ListAudits used as reply parameters in RPC or response body in HTTP.
type ListAudits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos   []*model.OpAudit `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
	HasMore bool             `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
}

func (x *ListAudits) Reset() {
	*x = ListAudits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAudits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAudits) ProtoMessage() {}

func (x *ListAudits) ProtoReflect() protoreflect.Message {
	mi := &file_proto_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAudits.ProtoReflect.Descriptor instead.
func (*ListAudits) Descriptor() ([]byte, []int) {
	return file_proto_response_proto_rawDescGZIP(), []int{1}
}

func (x *ListAudits) GetInfos() []*model.OpAudit {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *ListAudits) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

// CreateWorkspace used as reply parameters in RPC or response body in HTTP.
type CreateWorkspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workspace id that generated by system.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateWorkspace) Reset() {
	*x = CreateWorkspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspace) ProtoMessage() {}

func (x *CreateWorkspace) ProtoReflect() protoreflect.Message {
	mi := &file_proto_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkspace.ProtoReflect.Descriptor instead.
func (*CreateWorkspace) Descriptor() ([]byte, []int) {
	return file_proto_response_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWorkspace) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DescribeWorkspace used as reply parameters in RPC, And model.Workspace used as response body in HTTP.
type DescribeWorkspace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *model.Workspace `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *DescribeWorkspace) Reset() {
	*x = DescribeWorkspace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeWorkspace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeWorkspace) ProtoMessage() {}

func (x *DescribeWorkspace) ProtoReflect() protoreflect.Message {
	mi := &file_proto_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeWorkspace.ProtoReflect.Descriptor instead.
func (*DescribeWorkspace) Descriptor() ([]byte, []int) {
	return file_proto_response_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeWorkspace) GetInfo() *model.Workspace {
	if x != nil {
		return x.Info
	}
	return nil
}

// ListSystemRoles used as reply parameters in RPC or response body in HTTP.
type ListMembers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*model.Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	// map => "member_id": RoleList
	Roles map[string]*ListMembers_RoleList `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListMembers) Reset() {
	*x = ListMembers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMembers) ProtoMessage() {}

func (x *ListMembers) ProtoReflect() protoreflect.Message {
	mi := &file_proto_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMembers.ProtoReflect.Descriptor instead.
func (*ListMembers) Descriptor() ([]byte, []int) {
	return file_proto_response_proto_rawDescGZIP(), []int{4}
}

func (x *ListMembers) GetMembers() []*model.Member {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *ListMembers) GetRoles() map[string]*ListMembers_RoleList {
	if x != nil {
		return x.Roles
	}
	return nil
}

// ListSystemRoles used as reply parameters in RPC or response body in HTTP.
type ListSystemRoles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos   []*model.Role `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
	HasMore bool          `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
}

func (x *ListSystemRoles) Reset() {
	*x = ListSystemRoles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_response_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSystemRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSystemRoles) ProtoMessage() {}

func (x *ListSystemRoles) ProtoReflect() protoreflect.Message {
	mi := &file_proto_response_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSystemRoles.ProtoReflect.Descriptor instead.
func (*ListSystemRoles) Descriptor() ([]byte, []int) {
	return file_proto_response_proto_rawDescGZIP(), []int{5}
}

func (x *ListSystemRoles) GetInfos() []*model.Role {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *ListSystemRoles) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

type ListMembers_RoleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*model.Role `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *ListMembers_RoleList) Reset() {
	*x = ListMembers_RoleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_response_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMembers_RoleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMembers_RoleList) ProtoMessage() {}

func (x *ListMembers_RoleList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_response_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMembers_RoleList.ProtoReflect.Descriptor instead.
func (*ListMembers_RoleList) Descriptor() ([]byte, []int) {
	return file_proto_response_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ListMembers_RoleList) GetInfos() []*model.Role {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_proto_response_proto protoreflect.FileDescriptor

var file_proto_response_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69,
	0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x05,
	0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07, 0x68,
	0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x22, 0x59, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x70, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x1f, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72,
	0x65, 0x22, 0x2a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a,
	0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x1a, 0x2d, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x69,
	0x6e, 0x66, 0x6f, 0x73, 0x1a, 0x58, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x04, 0xe2,
	0xdf, 0x1f, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x68, 0x61,
	0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0xdf,
	0x1f, 0x00, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f,
	0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_response_proto_rawDescOnce sync.Once
	file_proto_response_proto_rawDescData = file_proto_response_proto_rawDesc
)

func file_proto_response_proto_rawDescGZIP() []byte {
	file_proto_response_proto_rawDescOnce.Do(func() {
		file_proto_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_response_proto_rawDescData)
	})
	return file_proto_response_proto_rawDescData
}

var file_proto_response_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_response_proto_goTypes = []interface{}{
	(*ListWorkspaces)(nil),       // 0: response.ListWorkspaces
	(*ListAudits)(nil),           // 1: response.ListAudits
	(*CreateWorkspace)(nil),      // 2: response.CreateWorkspace
	(*DescribeWorkspace)(nil),    // 3: response.DescribeWorkspace
	(*ListMembers)(nil),          // 4: response.ListMembers
	(*ListSystemRoles)(nil),      // 5: response.ListSystemRoles
	(*ListMembers_RoleList)(nil), // 6: response.ListMembers.RoleList
	nil,                          // 7: response.ListMembers.RolesEntry
	(*model.Workspace)(nil),      // 8: model.Workspace
	(*model.OpAudit)(nil),        // 9: model.OpAudit
	(*model.Member)(nil),         // 10: model.Member
	(*model.Role)(nil),           // 11: model.Role
}
var file_proto_response_proto_depIdxs = []int32{
	8,  // 0: response.ListWorkspaces.infos:type_name -> model.Workspace
	9,  // 1: response.ListAudits.infos:type_name -> model.OpAudit
	8,  // 2: response.DescribeWorkspace.info:type_name -> model.Workspace
	10, // 3: response.ListMembers.members:type_name -> model.Member
	7,  // 4: response.ListMembers.roles:type_name -> response.ListMembers.RolesEntry
	11, // 5: response.ListSystemRoles.infos:type_name -> model.Role
	11, // 6: response.ListMembers.RoleList.infos:type_name -> model.Role
	6,  // 7: response.ListMembers.RolesEntry.value:type_name -> response.ListMembers.RoleList
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_response_proto_init() }
func file_proto_response_proto_init() {
	if File_proto_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkspaces); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAudits); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkspace); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeWorkspace); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMembers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_response_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSystemRoles); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_response_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMembers_RoleList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_response_proto_goTypes,
		DependencyIndexes: file_proto_response_proto_depIdxs,
		MessageInfos:      file_proto_response_proto_msgTypes,
	}.Build()
	File_proto_response_proto = out.File
	file_proto_response_proto_rawDesc = nil
	file_proto_response_proto_goTypes = nil
	file_proto_response_proto_depIdxs = nil
}
