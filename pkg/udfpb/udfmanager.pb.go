// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/udfmanager.proto

package udfpb

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SpaceID     string `protobuf:"bytes,2,opt,name=SpaceID,proto3" json:"SpaceID,omitempty"`
	UdfType     string `protobuf:"bytes,3,opt,name=UdfType,proto3" json:"UdfType,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Comment     string `protobuf:"bytes,5,opt,name=Comment,proto3" json:"Comment,omitempty"`
	Define      string `protobuf:"bytes,6,opt,name=Define,proto3" json:"Define,omitempty"`
	UsageSample string `protobuf:"bytes,7,opt,name=UsageSample,proto3" json:"UsageSample,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_udfmanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_udfmanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_udfmanager_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CreateRequest) GetSpaceID() string {
	if x != nil {
		return x.SpaceID
	}
	return ""
}

func (x *CreateRequest) GetUdfType() string {
	if x != nil {
		return x.UdfType
	}
	return ""
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CreateRequest) GetDefine() string {
	if x != nil {
		return x.Define
	}
	return ""
}

func (x *CreateRequest) GetUsageSample() string {
	if x != nil {
		return x.UsageSample
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UdfType     string `protobuf:"bytes,2,opt,name=UdfType,proto3" json:"UdfType,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Comment     string `protobuf:"bytes,4,opt,name=Comment,proto3" json:"Comment,omitempty"`
	Define      string `protobuf:"bytes,5,opt,name=Define,proto3" json:"Define,omitempty"`
	UsageSample string `protobuf:"bytes,6,opt,name=UsageSample,proto3" json:"UsageSample,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_udfmanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_udfmanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_udfmanager_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateRequest) GetUdfType() string {
	if x != nil {
		return x.UdfType
	}
	return ""
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *UpdateRequest) GetDefine() string {
	if x != nil {
		return x.Define
	}
	return ""
}

func (x *UpdateRequest) GetUsageSample() string {
	if x != nil {
		return x.UsageSample
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_udfmanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_udfmanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_udfmanager_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceID string `protobuf:"bytes,1,opt,name=SpaceID,proto3" json:"SpaceID,omitempty"`
}

func (x *DeleteAllRequest) Reset() {
	*x = DeleteAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_udfmanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllRequest) ProtoMessage() {}

func (x *DeleteAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_udfmanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllRequest) Descriptor() ([]byte, []int) {
	return file_proto_udfmanager_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAllRequest) GetSpaceID() string {
	if x != nil {
		return x.SpaceID
	}
	return ""
}

type InfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SpaceID     string `protobuf:"bytes,2,opt,name=SpaceID,proto3" json:"SpaceID,omitempty"`
	UdfType     string `protobuf:"bytes,3,opt,name=UdfType,proto3" json:"UdfType,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Comment     string `protobuf:"bytes,5,opt,name=Comment,proto3" json:"Comment,omitempty"`
	Define      string `protobuf:"bytes,6,opt,name=Define,proto3" json:"Define,omitempty"`
	CreateTime  string `protobuf:"bytes,7,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime  string `protobuf:"bytes,8,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	UsageSample string `protobuf:"bytes,9,opt,name=UsageSample,proto3" json:"UsageSample,omitempty"`
}

func (x *InfoReply) Reset() {
	*x = InfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_udfmanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReply) ProtoMessage() {}

func (x *InfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_udfmanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReply.ProtoReflect.Descriptor instead.
func (*InfoReply) Descriptor() ([]byte, []int) {
	return file_proto_udfmanager_proto_rawDescGZIP(), []int{4}
}

func (x *InfoReply) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *InfoReply) GetSpaceID() string {
	if x != nil {
		return x.SpaceID
	}
	return ""
}

func (x *InfoReply) GetUdfType() string {
	if x != nil {
		return x.UdfType
	}
	return ""
}

func (x *InfoReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InfoReply) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *InfoReply) GetDefine() string {
	if x != nil {
		return x.Define
	}
	return ""
}

func (x *InfoReply) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *InfoReply) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *InfoReply) GetUsageSample() string {
	if x != nil {
		return x.UsageSample
	}
	return ""
}

type DescribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DescribeRequest) Reset() {
	*x = DescribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_udfmanager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRequest) ProtoMessage() {}

func (x *DescribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_udfmanager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRequest.ProtoReflect.Descriptor instead.
func (*DescribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_udfmanager_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit   int32  `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset  int32  `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	SpaceID string `protobuf:"bytes,3,opt,name=SpaceID,proto3" json:"SpaceID,omitempty"`
}

func (x *ListsRequest) Reset() {
	*x = ListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_udfmanager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListsRequest) ProtoMessage() {}

func (x *ListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_udfmanager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListsRequest.ProtoReflect.Descriptor instead.
func (*ListsRequest) Descriptor() ([]byte, []int) {
	return file_proto_udfmanager_proto_rawDescGZIP(), []int{6}
}

func (x *ListsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListsRequest) GetSpaceID() string {
	if x != nil {
		return x.SpaceID
	}
	return ""
}

// Reply parameters used to ListWorkspaces
type ListsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Infos []*InfoReply `protobuf:"bytes,2,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *ListsReply) Reset() {
	*x = ListsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_udfmanager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListsReply) ProtoMessage() {}

func (x *ListsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_udfmanager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListsReply.ProtoReflect.Descriptor instead.
func (*ListsReply) Descriptor() ([]byte, []int) {
	return file_proto_udfmanager_proto_rawDescGZIP(), []int{7}
}

func (x *ListsReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListsReply) GetInfos() []*InfoReply {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_proto_udfmanager_proto protoreflect.FileDescriptor

var file_proto_udfmanager_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x64, 0x66, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x64, 0x66, 0x70, 0x62, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x15, 0x52, 0x02, 0x49, 0x44, 0x12, 0x21,
	0x0a, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x07, 0x55, 0x64, 0x66, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78, 0x11, 0x52, 0x07, 0x55, 0x64,
	0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78, 0x41, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81, 0x02, 0x52, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x70, 0x00, 0x52, 0x06,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xe2, 0xdf, 0x1f,
	0x0b, 0x70, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x0b, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x15, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x07, 0x55, 0x64, 0x66, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78, 0x11, 0x52, 0x07,
	0x55, 0x64, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78, 0x41, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81, 0x02, 0x52,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x70, 0x00,
	0x52, 0x06, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xe2,
	0xdf, 0x1f, 0x0b, 0x70, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x0b,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x28, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01,
	0x14, 0x52, 0x02, 0x49, 0x44, 0x22, 0x35, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03,
	0x80, 0x01, 0x14, 0x52, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x22, 0xce, 0x02, 0x0a,
	0x09, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x15, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x21, 0x0a, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x07, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x07, 0x55, 0x64, 0x66, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78, 0x11,
	0x52, 0x07, 0x55, 0x64, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78,
	0x41, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81,
	0x02, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x70, 0x00, 0x52, 0x06, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x41, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x41, 0x52, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x70, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01,
	0x52, 0x0b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x2a, 0x0a,
	0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf,
	0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x49, 0x44, 0x22, 0x78, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00,
	0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x21, 0x0a, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x07, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x22, 0x63, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x25, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x64, 0x66, 0x70, 0x62, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20,
	0x00, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0xd4, 0x02, 0x0a, 0x0a, 0x55, 0x64, 0x66,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x2e, 0x75, 0x64, 0x66, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x75, 0x64, 0x66, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x75, 0x64, 0x66, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x75, 0x64, 0x66, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x16, 0x2e, 0x75, 0x64, 0x66, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x75, 0x64, 0x66, 0x70,
	0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x75, 0x64, 0x66, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x75, 0x64, 0x66,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61,
	0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x64, 0x66, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_udfmanager_proto_rawDescOnce sync.Once
	file_proto_udfmanager_proto_rawDescData = file_proto_udfmanager_proto_rawDesc
)

func file_proto_udfmanager_proto_rawDescGZIP() []byte {
	file_proto_udfmanager_proto_rawDescOnce.Do(func() {
		file_proto_udfmanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_udfmanager_proto_rawDescData)
	})
	return file_proto_udfmanager_proto_rawDescData
}

var file_proto_udfmanager_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_udfmanager_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),     // 0: udfpb.CreateRequest
	(*UpdateRequest)(nil),     // 1: udfpb.UpdateRequest
	(*DeleteRequest)(nil),     // 2: udfpb.DeleteRequest
	(*DeleteAllRequest)(nil),  // 3: udfpb.DeleteAllRequest
	(*InfoReply)(nil),         // 4: udfpb.InfoReply
	(*DescribeRequest)(nil),   // 5: udfpb.DescribeRequest
	(*ListsRequest)(nil),      // 6: udfpb.ListsRequest
	(*ListsReply)(nil),        // 7: udfpb.ListsReply
	(*model.EmptyStruct)(nil), // 8: model.EmptyStruct
}
var file_proto_udfmanager_proto_depIdxs = []int32{
	4, // 0: udfpb.ListsReply.Infos:type_name -> udfpb.InfoReply
	0, // 1: udfpb.Udfmanager.Create:input_type -> udfpb.CreateRequest
	1, // 2: udfpb.Udfmanager.Update:input_type -> udfpb.UpdateRequest
	2, // 3: udfpb.Udfmanager.Delete:input_type -> udfpb.DeleteRequest
	3, // 4: udfpb.Udfmanager.DeleteAll:input_type -> udfpb.DeleteAllRequest
	5, // 5: udfpb.Udfmanager.Describe:input_type -> udfpb.DescribeRequest
	6, // 6: udfpb.Udfmanager.List:input_type -> udfpb.ListsRequest
	8, // 7: udfpb.Udfmanager.Create:output_type -> model.EmptyStruct
	8, // 8: udfpb.Udfmanager.Update:output_type -> model.EmptyStruct
	8, // 9: udfpb.Udfmanager.Delete:output_type -> model.EmptyStruct
	8, // 10: udfpb.Udfmanager.DeleteAll:output_type -> model.EmptyStruct
	4, // 11: udfpb.Udfmanager.Describe:output_type -> udfpb.InfoReply
	7, // 12: udfpb.Udfmanager.List:output_type -> udfpb.ListsReply
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_udfmanager_proto_init() }
func file_proto_udfmanager_proto_init() {
	if File_proto_udfmanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_udfmanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_udfmanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_udfmanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_udfmanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllRequest); i {
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
		file_proto_udfmanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoReply); i {
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
		file_proto_udfmanager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRequest); i {
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
		file_proto_udfmanager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListsRequest); i {
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
		file_proto_udfmanager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListsReply); i {
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
			RawDescriptor: file_proto_udfmanager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_udfmanager_proto_goTypes,
		DependencyIndexes: file_proto_udfmanager_proto_depIdxs,
		MessageInfos:      file_proto_udfmanager_proto_msgTypes,
	}.Build()
	File_proto_udfmanager_proto = out.File
	file_proto_udfmanager_proto_rawDesc = nil
	file_proto_udfmanager_proto_goTypes = nil
	file_proto_udfmanager_proto_depIdxs = nil
}
