// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/request.proto

package request

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

// ListAudits used as a request parameters for RPC and HTTP(based on URL-Query)
type ListAudits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Limit the maximum number of entries returned this time.
	// Not required, Max 100, default 100.
	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty" params:"limit" form:"limit" default:"100" binding:"gt=0,lte=100" minimum:"0" maximum:"100"`
	// The offset position.
	// Not required, default 0.
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty" params:"offset" form:"offset" default:"0" binding:"gte=0"`
	// The used_id fixed to request user id.
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" params:"-" form:"-" binding:"-" swaggerignore:"true"`
	// Querying conditions. Not required.
	SpaceId string `protobuf:"bytes,4,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty" params:"space_id" form:"space_id" binding:"-"`
	// Querying conditions. Not required.
	Type model.OpAudit_Type `protobuf:"varint,5,opt,name=type,proto3,enum=model.OpAudit_Type" json:"type,omitempty" params:"type" form:"type" binding:"-"`
	// Querying conditions. Not required.
	Action string `protobuf:"bytes,6,opt,name=action,proto3" json:"action,omitempty" params:"action" form:"action" binding:"-"`
	// Querying conditions. Not required.
	State model.OpAudit_State `protobuf:"varint,7,opt,name=state,proto3,enum=model.OpAudit_State" json:"state,omitempty" params:"state" form:"state" binding:"-"`
	// Querying conditions. Not required.
	// Desc: Timestamp of start time.
	Started int64 `protobuf:"varint,8,opt,name=started,proto3" json:"started,omitempty" params:"started" form:"started" binding:"-"`
	// Querying conditions. Not required.
	// Desc: Timestamp of end time.
	Ended int64 `protobuf:"varint,9,opt,name=ended,proto3" json:"ended,omitempty" params:"ended" form:"ended" binding:"-"`
}

func (x *ListAudits) Reset() {
	*x = ListAudits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAudits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAudits) ProtoMessage() {}

func (x *ListAudits) ProtoReflect() protoreflect.Message {
	mi := &file_proto_request_proto_msgTypes[0]
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
	return file_proto_request_proto_rawDescGZIP(), []int{0}
}

func (x *ListAudits) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListAudits) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListAudits) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListAudits) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *ListAudits) GetType() model.OpAudit_Type {
	if x != nil {
		return x.Type
	}
	return model.OpAudit__
}

func (x *ListAudits) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ListAudits) GetState() model.OpAudit_State {
	if x != nil {
		return x.State
	}
	return model.OpAudit___
}

func (x *ListAudits) GetStarted() int64 {
	if x != nil {
		return x.Started
	}
	return 0
}

func (x *ListAudits) GetEnded() int64 {
	if x != nil {
		return x.Ended
	}
	return 0
}

var File_proto_request_proto protoreflect.FileDescriptor

var file_proto_request_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b,
	0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x73, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00, 0x18, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x78, 0x41, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x70, 0x41, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0xdf,
	0x1f, 0x00, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4f, 0x70, 0x41, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xe2,
	0xdf, 0x1f, 0x00, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x05,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xe2, 0xdf, 0x1f,
	0x00, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62,
	0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_request_proto_rawDescOnce sync.Once
	file_proto_request_proto_rawDescData = file_proto_request_proto_rawDesc
)

func file_proto_request_proto_rawDescGZIP() []byte {
	file_proto_request_proto_rawDescOnce.Do(func() {
		file_proto_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_request_proto_rawDescData)
	})
	return file_proto_request_proto_rawDescData
}

var file_proto_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_request_proto_goTypes = []interface{}{
	(*ListAudits)(nil),       // 0: request.ListAudits
	(model.OpAudit_Type)(0),  // 1: model.OpAudit.Type
	(model.OpAudit_State)(0), // 2: model.OpAudit.State
}
var file_proto_request_proto_depIdxs = []int32{
	1, // 0: request.ListAudits.type:type_name -> model.OpAudit.Type
	2, // 1: request.ListAudits.state:type_name -> model.OpAudit.State
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_request_proto_init() }
func file_proto_request_proto_init() {
	if File_proto_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_request_proto_goTypes,
		DependencyIndexes: file_proto_request_proto_depIdxs,
		MessageInfos:      file_proto_request_proto_msgTypes,
	}.Build()
	File_proto_request_proto = out.File
	file_proto_request_proto_rawDesc = nil
	file_proto_request_proto_goTypes = nil
	file_proto_request_proto_depIdxs = nil
}
