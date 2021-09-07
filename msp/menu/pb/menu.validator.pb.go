// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: menu.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/erda-project/erda-proto-go/common/pb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetMenuRequest) Validate() error {
	if this.TenantId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantId", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantId))
	}
	if this.Type == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must not be an empty string`, this.Type))
	}
	return nil
}
func (this *GetMenuResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *MenuItem) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Children {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Children", err)
			}
		}
	}
	return nil
}
func (this *GetSettingRequest) Validate() error {
	if this.TenantGroup == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantGroup", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantGroup))
	}
	return nil
}
func (this *GetSettingResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *EngineSetting) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
