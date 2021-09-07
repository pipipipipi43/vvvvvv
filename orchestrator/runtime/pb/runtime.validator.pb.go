// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: runtime.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "github.com/erda-project/erda-proto-go/common/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetRuntimeRequest) Validate() error {
	return nil
}
func (this *Resources) Validate() error {
	return nil
}
func (this *Deployments) Validate() error {
	return nil
}
func (this *Service) Validate() error {
	if this.Deployments != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deployments); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deployments", err)
		}
	}
	if this.Resources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resources", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Errors {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Errors", err)
			}
		}
	}
	return nil
}
func (this *StatusMap) Validate() error {
	return nil
}
func (this *Extra) Validate() error {
	return nil
}
func (this *ErrorResponse) Validate() error {
	return nil
}
func (this *RuntimeInspect) Validate() error {
	if this.Resources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resources", err)
		}
	}
	if this.Extra != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Extra); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Extra", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	if this.TimeCreated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TimeCreated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TimeCreated", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	for _, item := range this.Errors {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Errors", err)
			}
		}
	}
	return nil
}
