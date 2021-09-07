// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: definition.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/structpb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PipelineDefinitionProcessRequest) Validate() error {
	if this.SnippetConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SnippetConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SnippetConfig", err)
		}
	}
	if this.PipelineCreateRequest != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PipelineCreateRequest); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PipelineCreateRequest", err)
		}
	}
	return nil
}
func (this *PipelineDefinitionProcessResponse) Validate() error {
	return nil
}
func (this *PipelineDefinitionProcessVersionRequest) Validate() error {
	return nil
}
func (this *PipelineDefinitionProcessVersionResponse) Validate() error {
	return nil
}
