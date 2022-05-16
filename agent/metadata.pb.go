// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: agent/metadata.proto

package agent

import (
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

type AgentMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *AgentMeta) Reset() {
	*x = AgentMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentMeta) ProtoMessage() {}

func (x *AgentMeta) ProtoReflect() protoreflect.Message {
	mi := &file_agent_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentMeta.ProtoReflect.Descriptor instead.
func (*AgentMeta) Descriptor() ([]byte, []int) {
	return file_agent_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *AgentMeta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_agent_metadata_proto protoreflect.FileDescriptor

var file_agent_metadata_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x68, 0x6f, 0x74, 0x73, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x28, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x68, 0x6f, 0x74, 0x73, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5a, 0x10, 0x6c, 0x68, 0x6f, 0x74, 0x73, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_agent_metadata_proto_rawDescOnce sync.Once
	file_agent_metadata_proto_rawDescData = file_agent_metadata_proto_rawDesc
)

func file_agent_metadata_proto_rawDescGZIP() []byte {
	file_agent_metadata_proto_rawDescOnce.Do(func() {
		file_agent_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_metadata_proto_rawDescData)
	})
	return file_agent_metadata_proto_rawDescData
}

var file_agent_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_agent_metadata_proto_goTypes = []interface{}{
	(*AgentMeta)(nil), // 0: lhotse.api.agent.AgentMeta
}
var file_agent_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agent_metadata_proto_init() }
func file_agent_metadata_proto_init() {
	if File_agent_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentMeta); i {
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
			RawDescriptor: file_agent_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_agent_metadata_proto_goTypes,
		DependencyIndexes: file_agent_metadata_proto_depIdxs,
		MessageInfos:      file_agent_metadata_proto_msgTypes,
	}.Build()
	File_agent_metadata_proto = out.File
	file_agent_metadata_proto_rawDesc = nil
	file_agent_metadata_proto_goTypes = nil
	file_agent_metadata_proto_depIdxs = nil
}
