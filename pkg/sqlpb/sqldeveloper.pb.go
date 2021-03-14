// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/sqldeveloper.proto

package sqlpb

import (
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

type EmptyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReply) Reset() {
	*x = EmptyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sqldeveloper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReply) ProtoMessage() {}

func (x *EmptyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sqldeveloper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReply.ProtoReflect.Descriptor instead.
func (*EmptyReply) Descriptor() ([]byte, []int) {
	return file_proto_sqldeveloper_proto_rawDescGZIP(), []int{0}
}

type DAG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dag string `protobuf:"bytes,1,opt,name=Dag,proto3" json:"Dag,omitempty"`
}

func (x *DAG) Reset() {
	*x = DAG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sqldeveloper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DAG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DAG) ProtoMessage() {}

func (x *DAG) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sqldeveloper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DAG.ProtoReflect.Descriptor instead.
func (*DAG) Descriptor() ([]byte, []int) {
	return file_proto_sqldeveloper_proto_rawDescGZIP(), []int{1}
}

func (x *DAG) GetDag() string {
	if x != nil {
		return x.Dag
	}
	return ""
}

type SQL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sql string `protobuf:"bytes,1,opt,name=Sql,proto3" json:"Sql,omitempty"`
}

func (x *SQL) Reset() {
	*x = SQL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sqldeveloper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQL) ProtoMessage() {}

func (x *SQL) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sqldeveloper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQL.ProtoReflect.Descriptor instead.
func (*SQL) Descriptor() ([]byte, []int) {
	return file_proto_sqldeveloper_proto_rawDescGZIP(), []int{2}
}

func (x *SQL) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

var File_proto_sqldeveloper_proto protoreflect.FileDescriptor

var file_proto_sqldeveloper_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x71, 0x6c, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x71, 0x6c, 0x70,
	0x62, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77,
	0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x0a, 0x03, 0x44, 0x41, 0x47, 0x12, 0x18,
	0x0a, 0x03, 0x44, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x70, 0x01, 0x52, 0x03, 0x44, 0x61, 0x67, 0x22, 0x1f, 0x0a, 0x03, 0x53, 0x51, 0x4c, 0x12,
	0x18, 0x0a, 0x03, 0x53, 0x71, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x70, 0x01, 0x52, 0x03, 0x53, 0x71, 0x6c, 0x32, 0x5a, 0x0a, 0x0c, 0x53, 0x71, 0x6c,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x08, 0x44, 0x41, 0x47,
	0x54, 0x6f, 0x53, 0x51, 0x4c, 0x12, 0x0a, 0x2e, 0x73, 0x71, 0x6c, 0x70, 0x62, 0x2e, 0x44, 0x41,
	0x47, 0x1a, 0x0a, 0x2e, 0x73, 0x71, 0x6c, 0x70, 0x62, 0x2e, 0x53, 0x51, 0x4c, 0x22, 0x00, 0x12,
	0x24, 0x0a, 0x08, 0x53, 0x51, 0x4c, 0x54, 0x6f, 0x44, 0x41, 0x47, 0x12, 0x0a, 0x2e, 0x73, 0x71,
	0x6c, 0x70, 0x62, 0x2e, 0x53, 0x51, 0x4c, 0x1a, 0x0a, 0x2e, 0x73, 0x71, 0x6c, 0x70, 0x62, 0x2e,
	0x44, 0x41, 0x47, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x71, 0x6c,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sqldeveloper_proto_rawDescOnce sync.Once
	file_proto_sqldeveloper_proto_rawDescData = file_proto_sqldeveloper_proto_rawDesc
)

func file_proto_sqldeveloper_proto_rawDescGZIP() []byte {
	file_proto_sqldeveloper_proto_rawDescOnce.Do(func() {
		file_proto_sqldeveloper_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sqldeveloper_proto_rawDescData)
	})
	return file_proto_sqldeveloper_proto_rawDescData
}

var file_proto_sqldeveloper_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_sqldeveloper_proto_goTypes = []interface{}{
	(*EmptyReply)(nil), // 0: sqlpb.EmptyReply
	(*DAG)(nil),        // 1: sqlpb.DAG
	(*SQL)(nil),        // 2: sqlpb.SQL
}
var file_proto_sqldeveloper_proto_depIdxs = []int32{
	1, // 0: sqlpb.Sqldeveloper.DAGToSQL:input_type -> sqlpb.DAG
	2, // 1: sqlpb.Sqldeveloper.SQLToDAG:input_type -> sqlpb.SQL
	2, // 2: sqlpb.Sqldeveloper.DAGToSQL:output_type -> sqlpb.SQL
	1, // 3: sqlpb.Sqldeveloper.SQLToDAG:output_type -> sqlpb.DAG
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_sqldeveloper_proto_init() }
func file_proto_sqldeveloper_proto_init() {
	if File_proto_sqldeveloper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sqldeveloper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReply); i {
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
		file_proto_sqldeveloper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DAG); i {
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
		file_proto_sqldeveloper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQL); i {
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
			RawDescriptor: file_proto_sqldeveloper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sqldeveloper_proto_goTypes,
		DependencyIndexes: file_proto_sqldeveloper_proto_depIdxs,
		MessageInfos:      file_proto_sqldeveloper_proto_msgTypes,
	}.Build()
	File_proto_sqldeveloper_proto = out.File
	file_proto_sqldeveloper_proto_rawDesc = nil
	file_proto_sqldeveloper_proto_goTypes = nil
	file_proto_sqldeveloper_proto_depIdxs = nil
}
