// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: template.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PipelineTemplateCreateRequest) Validate() error {
	return nil
}
func (this *PipelineTemplateCreateResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *PipelineTemplateApplyRequest) Validate() error {
	return nil
}
func (this *PipelineTemplateQueryRequest) Validate() error {
	return nil
}
func (this *PipelineTemplateQueryResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *PipelineTemplateQueryResponseData) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *PipelineTemplateVersionGetRequest) Validate() error {
	return nil
}
func (this *PipelineTemplateVersionGetResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *PipelineTemplateVersionQueryRequest) Validate() error {
	return nil
}
func (this *PipelineTemplateVersionQueryResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *PipelineTemplateRenderRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *PipelineTemplateRenderSpecRequest) Validate() error {
	if this.Spec != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Spec); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Spec", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *PipelineTemplateRenderResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *PipelineTemplateSearchResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *PipelineTemplateRender) Validate() error {
	if this.Version != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Version); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Version", err)
		}
	}
	for _, item := range this.Outputs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Outputs", err)
			}
		}
	}
	return nil
}
func (this *SnippetFormatOutputs) Validate() error {
	return nil
}
func (this *PipelineTemplate) Validate() error {
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
	return nil
}
func (this *PipelineTemplateVersion) Validate() error {
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
	return nil
}
func (this *PipelineTemplateSpec) Validate() error {
	for _, item := range this.Params {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
			}
		}
	}
	for _, item := range this.Outputs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Outputs", err)
			}
		}
	}
	return nil
}
func (this *PipelineTemplateSpecOutput) Validate() error {
	return nil
}
func (this *QuerySnippetYmlRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *QuerySnippetYmlResponse) Validate() error {
	return nil
}
func (this *PipelineParam) Validate() error {
	if this.Default != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Default); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Default", err)
		}
	}
	return nil
}
func (this *PipelineOutput) Validate() error {
	return nil
}
