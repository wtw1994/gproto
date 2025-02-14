// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/workspace.proto

package wspb

import (
	model "github.com/DataWorkbench/gproto/pkg/model"
	request "github.com/DataWorkbench/gproto/pkg/request"
	response "github.com/DataWorkbench/gproto/pkg/response"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_proto_workspace_proto protoreflect.FileDescriptor

var file_proto_workspace_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x77, 0x73, 0x70, 0x62, 0x1a, 0x0b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb9, 0x07, 0x0a, 0x09, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x10, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x19,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12,
	0x48, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x19, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x11,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x1b, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08,
	0x41, 0x64, 0x64, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x75, 0x64, 0x69, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x73, 0x12,
	0x13, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x73, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x18, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x15, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x12, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x73, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_workspace_proto_goTypes = []interface{}{
	(*request.ListWorkspaces)(nil),     // 0: request.ListWorkspaces
	(*request.DeleteWorkspaces)(nil),   // 1: request.DeleteWorkspaces
	(*request.DisableWorkspaces)(nil),  // 2: request.DisableWorkspaces
	(*request.EnableWorkspaces)(nil),   // 3: request.EnableWorkspaces
	(*request.CreateWorkspace)(nil),    // 4: request.CreateWorkspace
	(*request.UpdateWorkspace)(nil),    // 5: request.UpdateWorkspace
	(*request.DescribeWorkspace)(nil),  // 6: request.DescribeWorkspace
	(*request.AddAudit)(nil),           // 7: request.AddAudit
	(*request.ListAudits)(nil),         // 8: request.ListAudits
	(*request.ListSystemRoles)(nil),    // 9: request.ListSystemRoles
	(*request.ListMembers)(nil),        // 10: request.ListMembers
	(*request.UpsertMembers)(nil),      // 11: request.UpsertMembers
	(*request.DeleteMembers)(nil),      // 12: request.DeleteMembers
	(*request.CheckPermission)(nil),    // 13: request.CheckPermission
	(*response.ListWorkspaces)(nil),    // 14: response.ListWorkspaces
	(*model.EmptyStruct)(nil),          // 15: model.EmptyStruct
	(*response.CreateWorkspace)(nil),   // 16: response.CreateWorkspace
	(*response.DescribeWorkspace)(nil), // 17: response.DescribeWorkspace
	(*response.ListAudits)(nil),        // 18: response.ListAudits
	(*response.ListSystemRoles)(nil),   // 19: response.ListSystemRoles
	(*response.ListMembers)(nil),       // 20: response.ListMembers
}
var file_proto_workspace_proto_depIdxs = []int32{
	0,  // 0: wspb.Workspace.ListWorkspaces:input_type -> request.ListWorkspaces
	1,  // 1: wspb.Workspace.DeleteWorkspaces:input_type -> request.DeleteWorkspaces
	2,  // 2: wspb.Workspace.DisableWorkspaces:input_type -> request.DisableWorkspaces
	3,  // 3: wspb.Workspace.EnableWorkspaces:input_type -> request.EnableWorkspaces
	4,  // 4: wspb.Workspace.CreateWorkspace:input_type -> request.CreateWorkspace
	5,  // 5: wspb.Workspace.UpdateWorkspace:input_type -> request.UpdateWorkspace
	6,  // 6: wspb.Workspace.DescribeWorkspace:input_type -> request.DescribeWorkspace
	7,  // 7: wspb.Workspace.AddAudit:input_type -> request.AddAudit
	8,  // 8: wspb.Workspace.ListAudits:input_type -> request.ListAudits
	9,  // 9: wspb.Workspace.ListSystemRoles:input_type -> request.ListSystemRoles
	10, // 10: wspb.Workspace.ListMembers:input_type -> request.ListMembers
	11, // 11: wspb.Workspace.UpsertMembers:input_type -> request.UpsertMembers
	12, // 12: wspb.Workspace.DeleteMembers:input_type -> request.DeleteMembers
	13, // 13: wspb.Workspace.CheckPermission:input_type -> request.CheckPermission
	14, // 14: wspb.Workspace.ListWorkspaces:output_type -> response.ListWorkspaces
	15, // 15: wspb.Workspace.DeleteWorkspaces:output_type -> model.EmptyStruct
	15, // 16: wspb.Workspace.DisableWorkspaces:output_type -> model.EmptyStruct
	15, // 17: wspb.Workspace.EnableWorkspaces:output_type -> model.EmptyStruct
	16, // 18: wspb.Workspace.CreateWorkspace:output_type -> response.CreateWorkspace
	15, // 19: wspb.Workspace.UpdateWorkspace:output_type -> model.EmptyStruct
	17, // 20: wspb.Workspace.DescribeWorkspace:output_type -> response.DescribeWorkspace
	15, // 21: wspb.Workspace.AddAudit:output_type -> model.EmptyStruct
	18, // 22: wspb.Workspace.ListAudits:output_type -> response.ListAudits
	19, // 23: wspb.Workspace.ListSystemRoles:output_type -> response.ListSystemRoles
	20, // 24: wspb.Workspace.ListMembers:output_type -> response.ListMembers
	15, // 25: wspb.Workspace.UpsertMembers:output_type -> model.EmptyStruct
	15, // 26: wspb.Workspace.DeleteMembers:output_type -> model.EmptyStruct
	15, // 27: wspb.Workspace.CheckPermission:output_type -> model.EmptyStruct
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_workspace_proto_init() }
func file_proto_workspace_proto_init() {
	if File_proto_workspace_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_workspace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_workspace_proto_goTypes,
		DependencyIndexes: file_proto_workspace_proto_depIdxs,
	}.Build()
	File_proto_workspace_proto = out.File
	file_proto_workspace_proto_rawDesc = nil
	file_proto_workspace_proto_goTypes = nil
	file_proto_workspace_proto_depIdxs = nil
}
