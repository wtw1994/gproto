// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/scheduler.proto

package shpb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/DataWorkbench/gproto/pkg/types"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *NodeInfo) Validate() error {
	if !(len(this.ID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.ID))
	}
	if !(this.Status > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Status", fmt.Errorf(`value '%v' must be greater than '0'`, this.Status))
	}
	if !(this.Status < 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("Status", fmt.Errorf(`value '%v' must be less than '3'`, this.Status))
	}
	if !(this.RetryLimit > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("RetryLimit", fmt.Errorf(`value '%v' must be greater than '-1'`, this.RetryLimit))
	}
	if !(this.RetryInterval > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("RetryInterval", fmt.Errorf(`value '%v' must be greater than '0'`, this.RetryInterval))
	}
	if !(this.Timeout > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Timeout", fmt.Errorf(`value '%v' must be greater than '0'`, this.Timeout))
	}
	if !(this.FailureStrategy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("FailureStrategy", fmt.Errorf(`value '%v' must be greater than '0'`, this.FailureStrategy))
	}
	if !(this.FailureStrategy < 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("FailureStrategy", fmt.Errorf(`value '%v' must be less than '3'`, this.FailureStrategy))
	}
	if !(this.Type > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must be greater than '0'`, this.Type))
	}
	if !(len(this.Defines) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Defines", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Defines))
	}
	return nil
}
func (this *ScheduleInfo) Validate() error {
	if !(this.Started > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Started", fmt.Errorf(`value '%v' must be greater than '0'`, this.Started))
	}
	if !(this.Ended > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Ended", fmt.Errorf(`value '%v' must be greater than '0'`, this.Ended))
	}
	if !(this.Priority > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Priority", fmt.Errorf(`value '%v' must be greater than '0'`, this.Priority))
	}
	if !(this.Priority < 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("Priority", fmt.Errorf(`value '%v' must be less than '3'`, this.Priority))
	}
	if !(this.FailureStrategy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("FailureStrategy", fmt.Errorf(`value '%v' must be greater than '0'`, this.FailureStrategy))
	}
	if !(this.FailureStrategy < 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("FailureStrategy", fmt.Errorf(`value '%v' must be less than '3'`, this.FailureStrategy))
	}
	if !(this.DependStrategy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("DependStrategy", fmt.Errorf(`value '%v' must be greater than '0'`, this.DependStrategy))
	}
	if !(this.DependStrategy < 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("DependStrategy", fmt.Errorf(`value '%v' must be less than '3'`, this.DependStrategy))
	}
	if !(this.ScheduleStrategy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleStrategy", fmt.Errorf(`value '%v' must be greater than '0'`, this.ScheduleStrategy))
	}
	if !(this.ScheduleStrategy < 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleStrategy", fmt.Errorf(`value '%v' must be less than '3'`, this.ScheduleStrategy))
	}
	if !(this.ScheduleLimit > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduleLimit", fmt.Errorf(`value '%v' must be greater than '-1'`, this.ScheduleLimit))
	}
	return nil
}
func (this *PushRequest) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.FlowId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.FlowId))
	}
	if !(this.Version > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must be greater than '0'`, this.Version))
	}
	for _, item := range this.Nodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
			}
		}
	}
	if this.Schedule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Schedule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Schedule", err)
		}
	}
	return nil
}
func (this *RemoveRequest) Validate() error {
	if !(len(this.FlowId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.FlowId))
	}
	return nil
}
func (this *ReportRequest) Validate() error {
	if !(len(this.InstanceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.InstanceId))
	}
	return nil
}
