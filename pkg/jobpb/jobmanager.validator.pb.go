// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/jobmanager.proto

package jobpb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *JobReply) Validate() error {
	if !(this.Status > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Status", fmt.Errorf(`value '%v' must be greater than '0'`, this.Status))
	}
	if !(len(this.Message) < 4000) {
		return github_com_mwitkow_go_proto_validators.FieldError("Message", fmt.Errorf(`value '%v' must have a length smaller than '4000'`, this.Message))
	}
	return nil
}
func (this *EmptyReply) Validate() error {
	return nil
}
func (this *RunJobRequest) Validate() error {
	if !(len(this.ID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.ID))
	}
	if !(len(this.WorkspaceID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkspaceID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.WorkspaceID))
	}
	if !(this.NodeType > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeType", fmt.Errorf(`value '%v' must be greater than '0'`, this.NodeType))
	}
	if !(len(this.Depends) < 4000) {
		return github_com_mwitkow_go_proto_validators.FieldError("Depends", fmt.Errorf(`value '%v' must have a length smaller than '4000'`, this.Depends))
	}
	if !(len(this.MainRun) < 8000) {
		return github_com_mwitkow_go_proto_validators.FieldError("MainRun", fmt.Errorf(`value '%v' must have a length smaller than '8000'`, this.MainRun))
	}
	return nil
}
func (this *GetJobStatusRequest) Validate() error {
	if !(len(this.ID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.ID))
	}
	return nil
}
func (this *CancelJobRequest) Validate() error {
	if !(len(this.ID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.ID))
	}
	return nil
}
