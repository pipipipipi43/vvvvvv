// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: queue_usage.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueueUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InUseCPU          float64           `protobuf:"fixed64,1,opt,name=inUseCPU,proto3" json:"inUseCPU,omitempty"`
	InUseMemoryMB     float64           `protobuf:"fixed64,2,opt,name=inUseMemoryMB,proto3" json:"inUseMemoryMB,omitempty"`
	RemainingCPU      float64           `protobuf:"fixed64,3,opt,name=remainingCPU,proto3" json:"remainingCPU,omitempty"`
	RemainingMemoryMB float64           `protobuf:"fixed64,4,opt,name=remainingMemoryMB,proto3" json:"remainingMemoryMB,omitempty"`
	ProcessingCount   int64             `protobuf:"varint,5,opt,name=processingCount,proto3" json:"processingCount,omitempty"`
	PendingCount      int64             `protobuf:"varint,6,opt,name=pendingCount,proto3" json:"pendingCount,omitempty"`
	ProcessingDetails []*QueueUsageItem `protobuf:"bytes,7,rep,name=processingDetails,proto3" json:"processingDetails,omitempty"`
	PendingDetails    []*QueueUsageItem `protobuf:"bytes,8,rep,name=pendingDetails,proto3" json:"pendingDetails,omitempty"`
}

func (x *QueueUsage) Reset() {
	*x = QueueUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_usage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueUsage) ProtoMessage() {}

func (x *QueueUsage) ProtoReflect() protoreflect.Message {
	mi := &file_queue_usage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueUsage.ProtoReflect.Descriptor instead.
func (*QueueUsage) Descriptor() ([]byte, []int) {
	return file_queue_usage_proto_rawDescGZIP(), []int{0}
}

func (x *QueueUsage) GetInUseCPU() float64 {
	if x != nil {
		return x.InUseCPU
	}
	return 0
}

func (x *QueueUsage) GetInUseMemoryMB() float64 {
	if x != nil {
		return x.InUseMemoryMB
	}
	return 0
}

func (x *QueueUsage) GetRemainingCPU() float64 {
	if x != nil {
		return x.RemainingCPU
	}
	return 0
}

func (x *QueueUsage) GetRemainingMemoryMB() float64 {
	if x != nil {
		return x.RemainingMemoryMB
	}
	return 0
}

func (x *QueueUsage) GetProcessingCount() int64 {
	if x != nil {
		return x.ProcessingCount
	}
	return 0
}

func (x *QueueUsage) GetPendingCount() int64 {
	if x != nil {
		return x.PendingCount
	}
	return 0
}

func (x *QueueUsage) GetProcessingDetails() []*QueueUsageItem {
	if x != nil {
		return x.ProcessingDetails
	}
	return nil
}

func (x *QueueUsage) GetPendingDetails() []*QueueUsageItem {
	if x != nil {
		return x.PendingDetails
	}
	return nil
}

type QueueUsageItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineID       uint64                 `protobuf:"varint,1,opt,name=pipelineID,proto3" json:"pipelineID,omitempty"`
	RequestsCPU      float64                `protobuf:"fixed64,2,opt,name=requestsCPU,proto3" json:"requestsCPU,omitempty"`
	RequestsMemoryMB float64                `protobuf:"fixed64,3,opt,name=requestsMemoryMB,proto3" json:"requestsMemoryMB,omitempty"`
	Index            int64                  `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	Priority         int64                  `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	AddedTime        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=addedTime,proto3" json:"addedTime,omitempty"`
}

func (x *QueueUsageItem) Reset() {
	*x = QueueUsageItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_usage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueUsageItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueUsageItem) ProtoMessage() {}

func (x *QueueUsageItem) ProtoReflect() protoreflect.Message {
	mi := &file_queue_usage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueUsageItem.ProtoReflect.Descriptor instead.
func (*QueueUsageItem) Descriptor() ([]byte, []int) {
	return file_queue_usage_proto_rawDescGZIP(), []int{1}
}

func (x *QueueUsageItem) GetPipelineID() uint64 {
	if x != nil {
		return x.PipelineID
	}
	return 0
}

func (x *QueueUsageItem) GetRequestsCPU() float64 {
	if x != nil {
		return x.RequestsCPU
	}
	return 0
}

func (x *QueueUsageItem) GetRequestsMemoryMB() float64 {
	if x != nil {
		return x.RequestsMemoryMB
	}
	return 0
}

func (x *QueueUsageItem) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *QueueUsageItem) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *QueueUsageItem) GetAddedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AddedTime
	}
	return nil
}

var File_queue_usage_proto protoreflect.FileDescriptor

var file_queue_usage_proto_rawDesc = []byte{
	0x0a, 0x11, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x75, 0x65, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x50, 0x55, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x50, 0x55, 0x12, 0x24,
	0x0a, 0x0d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x42, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x4d, 0x42, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x43, 0x50, 0x55, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x50, 0x55, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x42, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x11, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x42, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x11,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x45, 0x0a, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xea, 0x01, 0x0a, 0x0e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x43, 0x50, 0x55, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x43, 0x50, 0x55, 0x12, 0x2a, 0x0a,
	0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d,
	0x42, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4d, 0x42, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x61, 0x64, 0x64, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_queue_usage_proto_rawDescOnce sync.Once
	file_queue_usage_proto_rawDescData = file_queue_usage_proto_rawDesc
)

func file_queue_usage_proto_rawDescGZIP() []byte {
	file_queue_usage_proto_rawDescOnce.Do(func() {
		file_queue_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_queue_usage_proto_rawDescData)
	})
	return file_queue_usage_proto_rawDescData
}

var file_queue_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_queue_usage_proto_goTypes = []interface{}{
	(*QueueUsage)(nil),            // 0: erda.pipeline.QueueUsage
	(*QueueUsageItem)(nil),        // 1: erda.pipeline.QueueUsageItem
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_queue_usage_proto_depIdxs = []int32{
	1, // 0: erda.pipeline.QueueUsage.processingDetails:type_name -> erda.pipeline.QueueUsageItem
	1, // 1: erda.pipeline.QueueUsage.pendingDetails:type_name -> erda.pipeline.QueueUsageItem
	2, // 2: erda.pipeline.QueueUsageItem.addedTime:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_queue_usage_proto_init() }
func file_queue_usage_proto_init() {
	if File_queue_usage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queue_usage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueUsage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_queue_usage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueUsageItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_queue_usage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_queue_usage_proto_goTypes,
		DependencyIndexes: file_queue_usage_proto_depIdxs,
		MessageInfos:      file_queue_usage_proto_msgTypes,
	}.Build()
	File_queue_usage_proto = out.File
	file_queue_usage_proto_rawDesc = nil
	file_queue_usage_proto_goTypes = nil
	file_queue_usage_proto_depIdxs = nil
}
