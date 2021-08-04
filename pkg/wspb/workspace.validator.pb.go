// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/workspace.proto

package wspb

import (
	fmt "fmt"
	_ "github.com/DataWorkbench/gproto/pkg/model"
	_ "github.com/DataWorkbench/gproto/pkg/request"
	_ "github.com/DataWorkbench/gproto/pkg/response"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListWorkspacesRequest) Validate() error {
	if this.Owner == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Owner", fmt.Errorf(`value '%v' must not be an empty string`, this.Owner))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !(this.Limit < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be less than '101'`, this.Limit))
	}
	if !(this.Offset > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Offset", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Offset))
	}
	return nil
}
func (this *ListWorkspacesReply) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *CreateWorkspaceRequest) Validate() error {
	if !(len(this.Owner) > 10) {
		return github_com_mwitkow_go_proto_validators.FieldError("Owner", fmt.Errorf(`value '%v' must have a length greater than '10'`, this.Owner))
	}
	if !(len(this.Owner) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Owner", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Owner))
	}
	if !(len(this.Name) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Name))
	}
	if !(len(this.Name) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Name))
	}
	if !(len(this.Desc) < 1025) {
		return github_com_mwitkow_go_proto_validators.FieldError("Desc", fmt.Errorf(`value '%v' must have a length smaller than '1025'`, this.Desc))
	}
	return nil
}
func (this *DeleteWorkspaceRequest) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	return nil
}
func (this *UpdateWorkspaceRequest) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	if !(len(this.Name) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Name))
	}
	if !(len(this.Desc) < 1025) {
		return github_com_mwitkow_go_proto_validators.FieldError("Desc", fmt.Errorf(`value '%v' must have a length smaller than '1025'`, this.Desc))
	}
	return nil
}
func (this *DescribeWorkspaceRequest) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	return nil
}
func (this *DescribeWorkspaceReply) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *DisableWorkspaceRequest) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	return nil
}
func (this *EnableWorkspaceRequest) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	return nil
}
func (this *AddAuditRequest) Validate() error {
	if nil == this.Info {
		return github_com_mwitkow_go_proto_validators.FieldError("Info", fmt.Errorf("message must exist"))
	}
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *ListSystemRolesReply) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *ListMembersReply) Validate() error {
	for _, item := range this.Members {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Members", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ListMembersReply_RoleList) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *ListMembersRequest) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !(this.Limit < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be less than '101'`, this.Limit))
	}
	if !(this.Offset > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Offset", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Offset))
	}
	return nil
}
func (this *AddMemberRequest) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.MemberId) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("MemberId", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.MemberId))
	}
	if !(len(this.RoleIds) < 256) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleIds", fmt.Errorf(`value '%v' must have a length smaller than '256'`, this.RoleIds))
	}
	return nil
}
func (this *RemoveMemberRequest) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.MemberId) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("MemberId", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.MemberId))
	}
	return nil
}
func (this *UpdateMemberRequest) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.MemberId) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("MemberId", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.MemberId))
	}
	if !(len(this.RoleIds) < 256) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleIds", fmt.Errorf(`value '%v' must have a length smaller than '256'`, this.RoleIds))
	}
	return nil
}
func (this *CheckPermissionRequest) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.ReqUserId) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("ReqUserId", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.ReqUserId))
	}
	if !(this.OpType > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("OpType", fmt.Errorf(`value '%v' must be greater than '0'`, this.OpType))
	}
	if !(this.OpType < 5) {
		return github_com_mwitkow_go_proto_validators.FieldError("OpType", fmt.Errorf(`value '%v' must be less than '5'`, this.OpType))
	}
	return nil
}
