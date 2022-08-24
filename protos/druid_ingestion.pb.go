// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
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

type DruidIngestionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If not blank, indicates the message is a type of record to be stored into Druid.
	DataSourceName    string `protobuf:"bytes,1,opt,name=data_source_name,json=dataSourceName,proto3" json:"data_source_name,omitempty"`
	UseFieldDiscovery *bool  `protobuf:"varint,2,opt,name=use_field_discovery,json=useFieldDiscovery,proto3,oneof" json:"use_field_discovery,omitempty"`
}

func (x *DruidIngestionOptions) Reset() {
	*x = DruidIngestionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_druid_ingestion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DruidIngestionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DruidIngestionOptions) ProtoMessage() {}

func (x *DruidIngestionOptions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DruidIngestionOptions.ProtoReflect.Descriptor instead.
func (*DruidIngestionOptions) Descriptor() ([]byte, []int) {
	return file_druid_ingestion_proto_rawDescGZIP(), []int{0}
}

func (x *DruidIngestionOptions) GetDataSourceName() string {
	if x != nil {
		return x.DataSourceName
	}
	return ""
}

func (x *DruidIngestionOptions) GetUseFieldDiscovery() bool {
	if x != nil && x.UseFieldDiscovery != nil {
		return *x.UseFieldDiscovery
	}
	return false
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
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x15, 0x44, 0x72,
	0x75, 0x69, 0x64, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a,
	0x13, 0x75, 0x73, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x11, 0x75, 0x73,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x88,
	0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x3a, 0x66, 0x0a, 0x0a, 0x64, 0x72,
	0x75, 0x69, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe5, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x67, 0x65, 0x6e, 0x5f, 0x64, 0x72, 0x75, 0x69, 0x64, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x44, 0x72, 0x75, 0x69, 0x64, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x64, 0x72, 0x75, 0x69, 0x64, 0x4f, 0x70,
	0x74, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x73, 0x69, 0x6e, 0x68, 0x74, 0x75, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x64, 0x72, 0x75, 0x69, 0x64, 0x2d, 0x73,
	0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_druid_ingestion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_druid_ingestion_proto_goTypes = []interface{}{
	(*DruidIngestionOptions)(nil),     // 0: gen_druid_spec.DruidIngestionOptions
	(*descriptor.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
}
var file_druid_ingestion_proto_depIdxs = []int32{
	1, // 0: gen_druid_spec.druid_opts:extendee -> google.protobuf.MessageOptions
	0, // 1: gen_druid_spec.druid_opts:type_name -> gen_druid_spec.DruidIngestionOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_druid_ingestion_proto_init() }
func file_druid_ingestion_proto_init() {
	if File_druid_ingestion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_druid_ingestion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	file_druid_ingestion_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_druid_ingestion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
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
