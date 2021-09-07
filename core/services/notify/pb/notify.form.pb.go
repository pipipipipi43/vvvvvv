// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: notify.proto

package pb

import (
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	url "net/url"
	strconv "strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetAllNotifyTemplatesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAllNotifyTemplatesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Model)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Metadata)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Behavior)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Templates)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Render)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyTemplateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyTemplateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyRes)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateNotifyRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateNotifyResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteNotifyRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteNotifyResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateNotifyRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateNotifyResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetUserNotifyListRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetUserNotifyListResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotifyRes)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotifyTarget)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TargetValue)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotifyEnableRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotifyEnableResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateUserDefineNotifyTemplateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateUserDefineNotifyTemplateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyDetailRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetNotifyDetailResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotifyDetailResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAllGroupsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAllGroupsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetAllGroupData)(nil)

// GetAllNotifyTemplatesRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetAllNotifyTemplatesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetAllNotifyTemplatesResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetAllNotifyTemplatesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// Model implement urlenc.URLValuesUnmarshaler.
func (m *Model) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "metadata":
				if m.Metadata == nil {
					m.Metadata = &Metadata{}
				}
			case "metadata.name":
				if m.Metadata == nil {
					m.Metadata = &Metadata{}
				}
				m.Metadata.Name = vals[0]
			case "metadata.type":
				if m.Metadata == nil {
					m.Metadata = &Metadata{}
				}
				m.Metadata.Type = vals[0]
			case "metadata.module":
				if m.Metadata == nil {
					m.Metadata = &Metadata{}
				}
				m.Metadata.Module = vals[0]
			case "metadata.scope":
				if m.Metadata == nil {
					m.Metadata = &Metadata{}
				}
				m.Metadata.Scope = vals
			case "behavior":
				if m.Behavior == nil {
					m.Behavior = &Behavior{}
				}
			case "behavior.group":
				if m.Behavior == nil {
					m.Behavior = &Behavior{}
				}
				m.Behavior.Group = vals[0]
			}
		}
	}
	return nil
}

// Metadata implement urlenc.URLValuesUnmarshaler.
func (m *Metadata) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "type":
				m.Type = vals[0]
			case "module":
				m.Module = vals[0]
			case "scope":
				m.Scope = vals
			}
		}
	}
	return nil
}

// Behavior implement urlenc.URLValuesUnmarshaler.
func (m *Behavior) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "group":
				m.Group = vals[0]
			}
		}
	}
	return nil
}

// Templates implement urlenc.URLValuesUnmarshaler.
func (m *Templates) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "trigger":
				m.Trigger = vals
			case "targets":
				m.Targets = vals
			case "i18n":
				m.I18N = vals
			case "render":
				if m.Render == nil {
					m.Render = &Render{}
				}
			case "render.title":
				if m.Render == nil {
					m.Render = &Render{}
				}
				m.Render.Title = vals[0]
			case "render.template":
				if m.Render == nil {
					m.Render = &Render{}
				}
				m.Render.Template = vals[0]
			}
		}
	}
	return nil
}

// Render implement urlenc.URLValuesUnmarshaler.
func (m *Render) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "title":
				m.Title = vals[0]
			case "template":
				m.Template = vals[0]
			}
		}
	}
	return nil
}

// GetNotifyTemplateRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyTemplateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			case "name":
				m.Name = vals[0]
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// GetNotifyTemplateResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyTemplateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetNotifyRes implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyRes) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// CreateNotifyRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateNotifyRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scopeId":
				m.ScopeId = vals[0]
			case "scope":
				m.Scope = vals[0]
			case "templateId":
				m.TemplateId = vals
			case "notifyName":
				m.NotifyName = vals[0]
			case "notifyGroupId":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.NotifyGroupId = val
			case "channels":
				m.Channels = vals
			}
		}
	}
	return nil
}

// CreateNotifyResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateNotifyResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// DeleteNotifyRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteNotifyRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "scope":
				m.Scope = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			}
		}
	}
	return nil
}

// DeleteNotifyResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteNotifyResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// UpdateNotifyRequest implement urlenc.URLValuesUnmarshaler.
func (m *UpdateNotifyRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "scope":
				m.Scope = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			case "channels":
				m.Channels = vals
			case "notifyGroupId":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.NotifyGroupId = val
			case "templateId":
				m.TemplateId = vals
			}
		}
	}
	return nil
}

// UpdateNotifyResponse implement urlenc.URLValuesUnmarshaler.
func (m *UpdateNotifyResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// GetUserNotifyListRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetUserNotifyListRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			}
		}
	}
	return nil
}

// GetUserNotifyListResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetUserNotifyListResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// NotifyRes implement urlenc.URLValuesUnmarshaler.
func (m *NotifyRes) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "createdAt":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
			case "createdAt.seconds":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreatedAt.Seconds = val
			case "createdAt.nanos":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.CreatedAt.Nanos = int32(val)
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "notifyId":
				m.NotifyId = vals[0]
			case "notifyName":
				m.NotifyName = vals[0]
			case "target":
				m.Target = vals[0]
			case "enable":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Enable = val
			case "items":
				m.Items = vals
			}
		}
	}
	return nil
}

// NotifyTarget implement urlenc.URLValuesUnmarshaler.
func (m *NotifyTarget) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// TargetValue implement urlenc.URLValuesUnmarshaler.
func (m *TargetValue) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "receiver":
				m.Receiver = vals[0]
			case "secret":
				m.Secret = vals[0]
			}
		}
	}
	return nil
}

// NotifyEnableRequest implement urlenc.URLValuesUnmarshaler.
func (m *NotifyEnableRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "scopeId":
				m.ScopeId = vals[0]
			case "scope":
				m.Scope = vals[0]
			}
		}
	}
	return nil
}

// NotifyEnableResponse implement urlenc.URLValuesUnmarshaler.
func (m *NotifyEnableResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CreateUserDefineNotifyTemplateRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateUserDefineNotifyTemplateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "group":
				m.Group = vals[0]
			case "trigger":
				m.Trigger = vals
			case "title":
				m.Title = vals
			case "template":
				m.Template = vals
			case "scope":
				m.Scope = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			case "targets":
				m.Targets = vals
			}
		}
	}
	return nil
}

// CreateUserDefineNotifyTemplateResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateUserDefineNotifyTemplateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// GetNotifyDetailRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyDetailRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			}
		}
	}
	return nil
}

// GetNotifyDetailResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetNotifyDetailResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &NotifyDetailResponse{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &NotifyDetailResponse{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Id = val
			case "data.notifyId":
				if m.Data == nil {
					m.Data = &NotifyDetailResponse{}
				}
				m.Data.NotifyId = vals[0]
			case "data.notifyName":
				if m.Data == nil {
					m.Data = &NotifyDetailResponse{}
				}
				m.Data.NotifyName = vals[0]
			case "data.target":
				if m.Data == nil {
					m.Data = &NotifyDetailResponse{}
				}
				m.Data.Target = vals[0]
			case "data.groupType":
				if m.Data == nil {
					m.Data = &NotifyDetailResponse{}
				}
				m.Data.GroupType = vals[0]
			}
		}
	}
	return nil
}

// NotifyDetailResponse implement urlenc.URLValuesUnmarshaler.
func (m *NotifyDetailResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "notifyId":
				m.NotifyId = vals[0]
			case "notifyName":
				m.NotifyName = vals[0]
			case "target":
				m.Target = vals[0]
			case "groupType":
				m.GroupType = vals[0]
			}
		}
	}
	return nil
}

// GetAllGroupsRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetAllGroupsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			}
		}
	}
	return nil
}

// GetAllGroupsResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetAllGroupsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetAllGroupData implement urlenc.URLValuesUnmarshaler.
func (m *GetAllGroupData) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "value":
				m.Value = vals[0]
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}
