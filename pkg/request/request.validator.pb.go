// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/request.proto

package request

import (
	fmt "fmt"
	_ "github.com/DataWorkbench/gproto/pkg/model"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListAudits) Validate() error {
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !(this.Limit < 101) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be less than '101'`, this.Limit))
	}
	if !(this.Offset > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Offset", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Offset))
	}
	if !(len(this.UserId) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.UserId))
	}
	return nil
}
