// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: member.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/erda-project/erda-proto-go/common/pb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ScopeRoleAccessRequest) Validate() error {
	if this.Scope != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Scope); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Scope", err)
		}
	}
	return nil
}
func (this *Scope) Validate() error {
	return nil
}
func (this *ScopeRoleAccessResponse) Validate() error {
	for _, item := range this.PermissionList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PermissionList", err)
			}
		}
	}
	for _, item := range this.ResourceRoleList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ResourceRoleList", err)
			}
		}
	}
	return nil
}
func (this *ScopeResource) Validate() error {
	return nil
}
func (this *ListMemberRequest) Validate() error {
	return nil
}
func (this *ListMemberResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *MemberList) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *Member) Validate() error {
	if this.Scope != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Scope); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Scope", err)
		}
	}
	return nil
}
func (this *CreateOrUpdateMemberRequest) Validate() error {
	if this.Scope != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Scope); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Scope", err)
		}
	}
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	return nil
}
func (this *MemberAddOptions) Validate() error {
	return nil
}
func (this *CreateOrUpdateMemberResponse) Validate() error {
	return nil
}
func (this *DeleteMemberRequest) Validate() error {
	if this.Scope != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Scope); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Scope", err)
		}
	}
	return nil
}
func (this *DeleteMemberResponse) Validate() error {
	return nil
}
func (this *ListMemberRolesRequest) Validate() error {
	return nil
}
func (this *ListMemberRolesResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *RoleList) Validate() error {
	for _, item := range this.List {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("List", err)
			}
		}
	}
	return nil
}
func (this *RoleInfo) Validate() error {
	return nil
}
