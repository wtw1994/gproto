// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/scheduler.proto

package shpb

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

// Request parameters used to Push.
type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId  string            `protobuf:"bytes,1,opt,name=SpaceId,proto3" json:"SpaceId,omitempty"`
	FlowId   string            `protobuf:"bytes,2,opt,name=FlowId,proto3" json:"FlowId,omitempty"`
	Version  int64             `protobuf:"varint,3,opt,name=Version,proto3" json:"Version,omitempty"`
	Nodes    []*model.NodeInfo `protobuf:"bytes,4,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
	Schedule *model.SchInfo    `protobuf:"bytes,5,opt,name=Schedule,proto3" json:"Schedule,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_proto_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *AddRequest) GetFlowId() string {
	if x != nil {
		return x.FlowId
	}
	return ""
}

func (x *AddRequest) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AddRequest) GetNodes() []*model.NodeInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *AddRequest) GetSchedule() *model.SchInfo {
	if x != nil {
		return x.Schedule
	}
	return nil
}

// Request parameters used to Remove.
type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlowId string `protobuf:"bytes,1,opt,name=FlowId,proto3" json:"FlowId,omitempty"`
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_proto_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveRequest) GetFlowId() string {
	if x != nil {
		return x.FlowId
	}
	return ""
}

var File_proto_scheduler_proto protoreflect.FileDescriptor

var file_proto_scheduler_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x68, 0x70, 0x62, 0x1a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x07, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x06, 0x46,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x05, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53,
	0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x08, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x30, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x46, 0x6c, 0x6f, 0x77, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14,
	0x52, 0x06, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x32, 0x6f, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x10, 0x2e, 0x73,
	0x68, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x13,
	0x2e, 0x73, 0x68, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x73, 0x68, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_scheduler_proto_rawDescOnce sync.Once
	file_proto_scheduler_proto_rawDescData = file_proto_scheduler_proto_rawDesc
)

func file_proto_scheduler_proto_rawDescGZIP() []byte {
	file_proto_scheduler_proto_rawDescOnce.Do(func() {
		file_proto_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_scheduler_proto_rawDescData)
	})
	return file_proto_scheduler_proto_rawDescData
}

var file_proto_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_scheduler_proto_goTypes = []interface{}{
	(*AddRequest)(nil),        // 0: shpb.AddRequest
	(*RemoveRequest)(nil),     // 1: shpb.RemoveRequest
	(*model.NodeInfo)(nil),    // 2: model.NodeInfo
	(*model.SchInfo)(nil),     // 3: model.SchInfo
	(*model.EmptyStruct)(nil), // 4: model.EmptyStruct
}
var file_proto_scheduler_proto_depIdxs = []int32{
	2, // 0: shpb.AddRequest.Nodes:type_name -> model.NodeInfo
	3, // 1: shpb.AddRequest.Schedule:type_name -> model.SchInfo
	0, // 2: shpb.Scheduler.Add:input_type -> shpb.AddRequest
	1, // 3: shpb.Scheduler.Remove:input_type -> shpb.RemoveRequest
	4, // 4: shpb.Scheduler.Add:output_type -> model.EmptyStruct
	4, // 5: shpb.Scheduler.Remove:output_type -> model.EmptyStruct
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_scheduler_proto_init() }
func file_proto_scheduler_proto_init() {
	if File_proto_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_proto_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRequest); i {
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
			RawDescriptor: file_proto_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_scheduler_proto_goTypes,
		DependencyIndexes: file_proto_scheduler_proto_depIdxs,
		MessageInfos:      file_proto_scheduler_proto_msgTypes,
	}.Build()
	File_proto_scheduler_proto = out.File
	file_proto_scheduler_proto_rawDesc = nil
	file_proto_scheduler_proto_goTypes = nil
	file_proto_scheduler_proto_depIdxs = nil
}