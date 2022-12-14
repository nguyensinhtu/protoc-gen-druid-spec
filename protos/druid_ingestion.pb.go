// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: druid_ingestion.proto

package protos

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DruidIOConfigMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic             string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	UseFieldDiscovery bool   `protobuf:"varint,2,opt,name=use_field_discovery,json=useFieldDiscovery,proto3" json:"use_field_discovery,omitempty"`
	BootstrapServers  string `protobuf:"bytes,3,opt,name=bootstrap_servers,json=bootstrapServers,proto3" json:"bootstrap_servers,omitempty"`
	UseEarliestOffset bool   `protobuf:"varint,4,opt,name=use_earliest_offset,json=useEarliestOffset,proto3" json:"use_earliest_offset,omitempty"`
	Type              string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *DruidIOConfigMessageOptions) Reset() {
	*x = DruidIOConfigMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_druid_ingestion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DruidIOConfigMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DruidIOConfigMessageOptions) ProtoMessage() {}

func (x *DruidIOConfigMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_druid_ingestion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DruidIOConfigMessageOptions.ProtoReflect.Descriptor instead.
func (*DruidIOConfigMessageOptions) Descriptor() ([]byte, []int) {
	return file_druid_ingestion_proto_rawDescGZIP(), []int{0}
}

func (x *DruidIOConfigMessageOptions) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *DruidIOConfigMessageOptions) GetUseFieldDiscovery() bool {
	if x != nil {
		return x.UseFieldDiscovery
	}
	return false
}

func (x *DruidIOConfigMessageOptions) GetBootstrapServers() string {
	if x != nil {
		return x.BootstrapServers
	}
	return ""
}

func (x *DruidIOConfigMessageOptions) GetUseEarliestOffset() bool {
	if x != nil {
		return x.UseEarliestOffset
	}
	return false
}

func (x *DruidIOConfigMessageOptions) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DruidGranularityMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SegmentGranularity string `protobuf:"bytes,1,opt,name=segment_granularity,json=segmentGranularity,proto3" json:"segment_granularity,omitempty"`
	QueryGranularity   string `protobuf:"bytes,2,opt,name=query_granularity,json=queryGranularity,proto3" json:"query_granularity,omitempty"`
	Rollup             bool   `protobuf:"varint,3,opt,name=rollup,proto3" json:"rollup,omitempty"`
}

func (x *DruidGranularityMessageOptions) Reset() {
	*x = DruidGranularityMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_druid_ingestion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DruidGranularityMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DruidGranularityMessageOptions) ProtoMessage() {}

func (x *DruidGranularityMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_druid_ingestion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DruidGranularityMessageOptions.ProtoReflect.Descriptor instead.
func (*DruidGranularityMessageOptions) Descriptor() ([]byte, []int) {
	return file_druid_ingestion_proto_rawDescGZIP(), []int{1}
}

func (x *DruidGranularityMessageOptions) GetSegmentGranularity() string {
	if x != nil {
		return x.SegmentGranularity
	}
	return ""
}

func (x *DruidGranularityMessageOptions) GetQueryGranularity() string {
	if x != nil {
		return x.QueryGranularity
	}
	return ""
}

func (x *DruidGranularityMessageOptions) GetRollup() bool {
	if x != nil {
		return x.Rollup
	}
	return false
}

type DruidIngestionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If not blank, indicates the message is a type of record to be stored into Druid.
	DataSourceName string                          `protobuf:"bytes,1,opt,name=data_source_name,json=dataSourceName,proto3" json:"data_source_name,omitempty"`
	IngestionType  string                          `protobuf:"bytes,2,opt,name=ingestion_type,json=ingestionType,proto3" json:"ingestion_type,omitempty"`
	Granularity    *DruidGranularityMessageOptions `protobuf:"bytes,3,opt,name=granularity,proto3" json:"granularity,omitempty"`
	IoConfig       *DruidIOConfigMessageOptions    `protobuf:"bytes,4,opt,name=io_config,json=ioConfig,proto3" json:"io_config,omitempty"`
}

func (x *DruidIngestionOptions) Reset() {
	*x = DruidIngestionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_druid_ingestion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DruidIngestionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DruidIngestionOptions) ProtoMessage() {}

func (x *DruidIngestionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_druid_ingestion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DruidIngestionOptions.ProtoReflect.Descriptor instead.
func (*DruidIngestionOptions) Descriptor() ([]byte, []int) {
	return file_druid_ingestion_proto_rawDescGZIP(), []int{2}
}

func (x *DruidIngestionOptions) GetDataSourceName() string {
	if x != nil {
		return x.DataSourceName
	}
	return ""
}

func (x *DruidIngestionOptions) GetIngestionType() string {
	if x != nil {
		return x.IngestionType
	}
	return ""
}

func (x *DruidIngestionOptions) GetGranularity() *DruidGranularityMessageOptions {
	if x != nil {
		return x.Granularity
	}
	return nil
}

func (x *DruidIngestionOptions) GetIoConfig() *DruidIOConfigMessageOptions {
	if x != nil {
		return x.IoConfig
	}
	return nil
}

var file_druid_ingestion_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*DruidIngestionOptions)(nil),
		Field:         2021,
		Name:          "gen_druid_spec.druid_opts",
		Tag:           "bytes,2021,opt,name=druid_opts",
		Filename:      "druid_ingestion.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional gen_druid_spec.DruidIngestionOptions druid_opts = 2021;
	E_DruidOpts = &file_druid_ingestion_proto_extTypes[0]
)

var File_druid_ingestion_proto protoreflect.FileDescriptor

var file_druid_ingestion_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x72, 0x75, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x65, 0x6e, 0x5f, 0x64, 0x72, 0x75,
	0x69, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x1b, 0x44, 0x72,
	0x75, 0x69, 0x64, 0x49, 0x4f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x2e, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x75, 0x73,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12,
	0x2b, 0x0a, 0x11, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x13,
	0x75, 0x73, 0x65, 0x5f, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x75, 0x73, 0x65, 0x45, 0x61,
	0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x96, 0x01, 0x0a, 0x1e, 0x44, 0x72, 0x75, 0x69, 0x64, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c,
	0x61, 0x72, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67,
	0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x67, 0x72,
	0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x71, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x6c, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x6c, 0x75, 0x70, 0x22, 0x84, 0x02, 0x0a, 0x15, 0x44, 0x72,
	0x75, 0x69, 0x64, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x65, 0x6e, 0x5f,
	0x64, 0x72, 0x75, 0x69, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x44, 0x72, 0x75, 0x69, 0x64,
	0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75,
	0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x09, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x65, 0x6e, 0x5f,
	0x64, 0x72, 0x75, 0x69, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x44, 0x72, 0x75, 0x69, 0x64,
	0x49, 0x4f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x08, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x3a, 0x66, 0x0a, 0x0a, 0x64, 0x72, 0x75, 0x69, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xe5, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x65, 0x6e, 0x5f, 0x64, 0x72, 0x75,
	0x69, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x44, 0x72, 0x75, 0x69, 0x64, 0x49, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x64,
	0x72, 0x75, 0x69, 0x64, 0x4f, 0x70, 0x74, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x73, 0x69, 0x6e,
	0x68, 0x74, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x64,
	0x72, 0x75, 0x69, 0x64, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_druid_ingestion_proto_rawDescOnce sync.Once
	file_druid_ingestion_proto_rawDescData = file_druid_ingestion_proto_rawDesc
)

func file_druid_ingestion_proto_rawDescGZIP() []byte {
	file_druid_ingestion_proto_rawDescOnce.Do(func() {
		file_druid_ingestion_proto_rawDescData = protoimpl.X.CompressGZIP(file_druid_ingestion_proto_rawDescData)
	})
	return file_druid_ingestion_proto_rawDescData
}

var file_druid_ingestion_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_druid_ingestion_proto_goTypes = []interface{}{
	(*DruidIOConfigMessageOptions)(nil),    // 0: gen_druid_spec.DruidIOConfigMessageOptions
	(*DruidGranularityMessageOptions)(nil), // 1: gen_druid_spec.DruidGranularityMessageOptions
	(*DruidIngestionOptions)(nil),          // 2: gen_druid_spec.DruidIngestionOptions
	(*descriptor.MessageOptions)(nil),      // 3: google.protobuf.MessageOptions
}
var file_druid_ingestion_proto_depIdxs = []int32{
	1, // 0: gen_druid_spec.DruidIngestionOptions.granularity:type_name -> gen_druid_spec.DruidGranularityMessageOptions
	0, // 1: gen_druid_spec.DruidIngestionOptions.io_config:type_name -> gen_druid_spec.DruidIOConfigMessageOptions
	3, // 2: gen_druid_spec.druid_opts:extendee -> google.protobuf.MessageOptions
	2, // 3: gen_druid_spec.druid_opts:type_name -> gen_druid_spec.DruidIngestionOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_druid_ingestion_proto_init() }
func file_druid_ingestion_proto_init() {
	if File_druid_ingestion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_druid_ingestion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DruidIOConfigMessageOptions); i {
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
		file_druid_ingestion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DruidGranularityMessageOptions); i {
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
		file_druid_ingestion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DruidIngestionOptions); i {
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
			RawDescriptor: file_druid_ingestion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_druid_ingestion_proto_goTypes,
		DependencyIndexes: file_druid_ingestion_proto_depIdxs,
		MessageInfos:      file_druid_ingestion_proto_msgTypes,
		ExtensionInfos:    file_druid_ingestion_proto_extTypes,
	}.Build()
	File_druid_ingestion_proto = out.File
	file_druid_ingestion_proto_rawDesc = nil
	file_druid_ingestion_proto_goTypes = nil
	file_druid_ingestion_proto_depIdxs = nil
}
