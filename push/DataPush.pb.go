// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push/DataPush.proto

package push

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DataCategory int32

const (
	DataCategory_DATA_CATEGORY_SERVICE             DataCategory = 0
	DataCategory_DATA_CATEGORY_CLUSTER             DataCategory = 1
	DataCategory_DATA_CATEGORY_ENDPOINT            DataCategory = 2
	DataCategory_DATA_CATEGORY_RULE                DataCategory = 3
	DataCategory_DATA_CATEGORY_AGENT_META          DataCategory = 4
	DataCategory_DATA_CATEGORY_AGENT_UPLOAD_STATUS DataCategory = 5
)

var DataCategory_name = map[int32]string{
	0: "DATA_CATEGORY_SERVICE",
	1: "DATA_CATEGORY_CLUSTER",
	2: "DATA_CATEGORY_ENDPOINT",
	3: "DATA_CATEGORY_RULE",
	4: "DATA_CATEGORY_AGENT_META",
	5: "DATA_CATEGORY_AGENT_UPLOAD_STATUS",
}

var DataCategory_value = map[string]int32{
	"DATA_CATEGORY_SERVICE":             0,
	"DATA_CATEGORY_CLUSTER":             1,
	"DATA_CATEGORY_ENDPOINT":            2,
	"DATA_CATEGORY_RULE":                3,
	"DATA_CATEGORY_AGENT_META":          4,
	"DATA_CATEGORY_AGENT_UPLOAD_STATUS": 5,
}

func (x DataCategory) String() string {
	return proto.EnumName(DataCategory_name, int32(x))
}

func (DataCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d761a3aab7a6026e, []int{0}
}

type PushStatus int32

const (
	PushStatus_PUSH_STATUS_OK   PushStatus = 0
	PushStatus_PUSH_STATUS_FAIL PushStatus = 1
)

var PushStatus_name = map[int32]string{
	0: "PUSH_STATUS_OK",
	1: "PUSH_STATUS_FAIL",
}

var PushStatus_value = map[string]int32{
	"PUSH_STATUS_OK":   0,
	"PUSH_STATUS_FAIL": 1,
}

func (x PushStatus) String() string {
	return proto.EnumName(PushStatus_name, int32(x))
}

func (PushStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d761a3aab7a6026e, []int{1}
}

// 推送动作
type PushAction int32

const (
	PushAction_PUSH_ACTION_FULL      PushAction = 0
	PushAction_PUSH_ACTION_INCREMENT PushAction = 1
)

var PushAction_name = map[int32]string{
	0: "PUSH_ACTION_FULL",
	1: "PUSH_ACTION_INCREMENT",
}

var PushAction_value = map[string]int32{
	"PUSH_ACTION_FULL":      0,
	"PUSH_ACTION_INCREMENT": 1,
}

func (x PushAction) String() string {
	return proto.EnumName(PushAction_name, int32(x))
}

func (PushAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d761a3aab7a6026e, []int{2}
}

// 推送请求
type DataPushRequest struct {
	Category             DataCategory `protobuf:"varint,1,opt,name=category,proto3,enum=lhotse.networking.v1alpha1.DataCategory" json:"category,omitempty"`
	Data                 *any.Any     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DataPushRequest) Reset()         { *m = DataPushRequest{} }
func (m *DataPushRequest) String() string { return proto.CompactTextString(m) }
func (*DataPushRequest) ProtoMessage()    {}
func (*DataPushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d761a3aab7a6026e, []int{0}
}

func (m *DataPushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataPushRequest.Unmarshal(m, b)
}
func (m *DataPushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataPushRequest.Marshal(b, m, deterministic)
}
func (m *DataPushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPushRequest.Merge(m, src)
}
func (m *DataPushRequest) XXX_Size() int {
	return xxx_messageInfo_DataPushRequest.Size(m)
}
func (m *DataPushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataPushRequest proto.InternalMessageInfo

func (m *DataPushRequest) GetCategory() DataCategory {
	if m != nil {
		return m.Category
	}
	return DataCategory_DATA_CATEGORY_SERVICE
}

func (m *DataPushRequest) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type ServerDataPushRequest struct {
	Action               PushAction `protobuf:"varint,1,opt,name=action,proto3,enum=lhotse.networking.v1alpha1.PushAction" json:"action,omitempty"`
	Checksum             string     `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Revision             string     `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	Data                 *any.Any   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ServerDataPushRequest) Reset()         { *m = ServerDataPushRequest{} }
func (m *ServerDataPushRequest) String() string { return proto.CompactTextString(m) }
func (*ServerDataPushRequest) ProtoMessage()    {}
func (*ServerDataPushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d761a3aab7a6026e, []int{1}
}

func (m *ServerDataPushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerDataPushRequest.Unmarshal(m, b)
}
func (m *ServerDataPushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerDataPushRequest.Marshal(b, m, deterministic)
}
func (m *ServerDataPushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerDataPushRequest.Merge(m, src)
}
func (m *ServerDataPushRequest) XXX_Size() int {
	return xxx_messageInfo_ServerDataPushRequest.Size(m)
}
func (m *ServerDataPushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerDataPushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerDataPushRequest proto.InternalMessageInfo

func (m *ServerDataPushRequest) GetAction() PushAction {
	if m != nil {
		return m.Action
	}
	return PushAction_PUSH_ACTION_FULL
}

func (m *ServerDataPushRequest) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *ServerDataPushRequest) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ServerDataPushRequest) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

// 推送结果
type DataPushResponse struct {
	Status               PushStatus `protobuf:"varint,1,opt,name=status,proto3,enum=lhotse.networking.v1alpha1.PushStatus" json:"status,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DataPushResponse) Reset()         { *m = DataPushResponse{} }
func (m *DataPushResponse) String() string { return proto.CompactTextString(m) }
func (*DataPushResponse) ProtoMessage()    {}
func (*DataPushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d761a3aab7a6026e, []int{2}
}

func (m *DataPushResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataPushResponse.Unmarshal(m, b)
}
func (m *DataPushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataPushResponse.Marshal(b, m, deterministic)
}
func (m *DataPushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPushResponse.Merge(m, src)
}
func (m *DataPushResponse) XXX_Size() int {
	return xxx_messageInfo_DataPushResponse.Size(m)
}
func (m *DataPushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataPushResponse proto.InternalMessageInfo

func (m *DataPushResponse) GetStatus() PushStatus {
	if m != nil {
		return m.Status
	}
	return PushStatus_PUSH_STATUS_OK
}

func (m *DataPushResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Agent启动上报身份信息
type AgentRegisterRequest struct {
	AgentId              string   `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentRegisterRequest) Reset()         { *m = AgentRegisterRequest{} }
func (m *AgentRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*AgentRegisterRequest) ProtoMessage()    {}
func (*AgentRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d761a3aab7a6026e, []int{3}
}

func (m *AgentRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentRegisterRequest.Unmarshal(m, b)
}
func (m *AgentRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentRegisterRequest.Marshal(b, m, deterministic)
}
func (m *AgentRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentRegisterRequest.Merge(m, src)
}
func (m *AgentRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_AgentRegisterRequest.Size(m)
}
func (m *AgentRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AgentRegisterRequest proto.InternalMessageInfo

func (m *AgentRegisterRequest) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *AgentRegisterRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterEnum("lhotse.networking.v1alpha1.DataCategory", DataCategory_name, DataCategory_value)
	proto.RegisterEnum("lhotse.networking.v1alpha1.PushStatus", PushStatus_name, PushStatus_value)
	proto.RegisterEnum("lhotse.networking.v1alpha1.PushAction", PushAction_name, PushAction_value)
	proto.RegisterType((*DataPushRequest)(nil), "lhotse.networking.v1alpha1.DataPushRequest")
	proto.RegisterType((*ServerDataPushRequest)(nil), "lhotse.networking.v1alpha1.ServerDataPushRequest")
	proto.RegisterType((*DataPushResponse)(nil), "lhotse.networking.v1alpha1.DataPushResponse")
	proto.RegisterType((*AgentRegisterRequest)(nil), "lhotse.networking.v1alpha1.AgentRegisterRequest")
}

func init() { proto.RegisterFile("push/DataPush.proto", fileDescriptor_d761a3aab7a6026e) }

var fileDescriptor_d761a3aab7a6026e = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xeb, 0xb6, 0xb4, 0xe9, 0x14, 0xb5, 0xd6, 0x36, 0xad, 0x92, 0x88, 0x43, 0x89, 0x44,
	0x15, 0x05, 0xe4, 0xd0, 0x20, 0x71, 0x03, 0x69, 0x71, 0xdc, 0x62, 0xd5, 0x75, 0xa2, 0xb5, 0x8d,
	0x04, 0x17, 0x6b, 0xeb, 0x2c, 0x8e, 0xd5, 0xc4, 0x36, 0xde, 0x75, 0x20, 0x47, 0x3e, 0x16, 0x07,
	0xbe, 0x1b, 0xf2, 0xbf, 0x34, 0xe1, 0x5f, 0x7b, 0x9c, 0x79, 0xef, 0xcd, 0xfe, 0x76, 0x33, 0x31,
	0x1c, 0xc5, 0x29, 0x9f, 0xf4, 0x06, 0x54, 0xd0, 0x51, 0xca, 0x27, 0x4a, 0x9c, 0x44, 0x22, 0x42,
	0xad, 0xe9, 0x24, 0x12, 0x9c, 0x29, 0x21, 0x13, 0x5f, 0xa3, 0xe4, 0x36, 0x08, 0x7d, 0x65, 0x7e,
	0x4e, 0xa7, 0xf1, 0x84, 0x9e, 0xb7, 0x9a, 0x7e, 0x14, 0xf9, 0x53, 0xd6, 0xcb, 0x9d, 0x37, 0xe9,
	0xe7, 0x1e, 0x0d, 0x17, 0x45, 0xac, 0xfd, 0x5d, 0x82, 0xc3, 0x6a, 0x12, 0x61, 0x5f, 0x52, 0xc6,
	0x05, 0x1a, 0x40, 0xcd, 0xa3, 0x82, 0xf9, 0x51, 0xb2, 0x68, 0x48, 0xa7, 0x52, 0xe7, 0xa0, 0xdf,
	0x51, 0xfe, 0x3d, 0x5d, 0xc9, 0xe2, 0x6a, 0xe9, 0x27, 0xcb, 0x24, 0xea, 0xc0, 0xf6, 0x98, 0x0a,
	0xda, 0xd8, 0x3c, 0x95, 0x3a, 0xfb, 0xfd, 0xba, 0x52, 0x30, 0x28, 0x15, 0x83, 0x82, 0xc3, 0x05,
	0xc9, 0x1d, 0xed, 0x1f, 0x12, 0x1c, 0x5b, 0x2c, 0x99, 0xb3, 0xe4, 0x77, 0x92, 0xb7, 0xb0, 0x43,
	0x3d, 0x11, 0x44, 0x61, 0xc9, 0x71, 0xf6, 0x3f, 0x8e, 0x2c, 0x88, 0x73, 0x37, 0x29, 0x53, 0xa8,
	0x05, 0x35, 0x6f, 0xc2, 0xbc, 0x5b, 0x9e, 0xce, 0x72, 0x8e, 0x3d, 0xb2, 0xac, 0x33, 0x2d, 0x61,
	0xf3, 0x80, 0x67, 0xd3, 0xb7, 0x0a, 0xad, 0xaa, 0x97, 0xec, 0xdb, 0xf7, 0xb2, 0x4f, 0x41, 0xbe,
	0x83, 0xe6, 0x71, 0x14, 0x72, 0x96, 0x51, 0x73, 0x41, 0x45, 0xca, 0x1f, 0x4a, 0x6d, 0xe5, 0x6e,
	0x52, 0xa6, 0x50, 0x03, 0x76, 0x67, 0x8c, 0x73, 0xea, 0xb3, 0x12, 0xba, 0x2a, 0xdb, 0x57, 0x50,
	0xc7, 0x3e, 0x0b, 0x05, 0x61, 0x7e, 0xc0, 0x05, 0x4b, 0xaa, 0x77, 0x6a, 0x42, 0x8d, 0x66, 0x7d,
	0x37, 0x18, 0xe7, 0x67, 0xee, 0x91, 0xdd, 0xbc, 0xd6, 0xc7, 0xd9, 0xb0, 0x39, 0x4b, 0xf2, 0x5b,
	0x96, 0xc3, 0xca, 0xb2, 0xfb, 0x53, 0x82, 0xc7, 0xab, 0xbf, 0x1d, 0x6a, 0xc2, 0xf1, 0x00, 0xdb,
	0xd8, 0x55, 0xb1, 0xad, 0x5d, 0x0e, 0xc9, 0x47, 0xd7, 0xd2, 0xc8, 0x07, 0x5d, 0xd5, 0xe4, 0x8d,
	0x3f, 0x25, 0xd5, 0x70, 0x2c, 0x5b, 0x23, 0xb2, 0x84, 0x5a, 0x70, 0xb2, 0x2e, 0x69, 0xe6, 0x60,
	0x34, 0xd4, 0x4d, 0x5b, 0xde, 0x44, 0x27, 0x80, 0xd6, 0x35, 0xe2, 0x18, 0x9a, 0xbc, 0x85, 0x9e,
	0x40, 0x63, 0xbd, 0x8f, 0x2f, 0x35, 0xd3, 0x76, 0xaf, 0x35, 0x1b, 0xcb, 0xdb, 0xe8, 0x19, 0x3c,
	0xfd, 0x9b, 0xea, 0x8c, 0x8c, 0x21, 0x1e, 0xb8, 0x96, 0x8d, 0x6d, 0xc7, 0x92, 0x1f, 0x75, 0x5f,
	0x03, 0xdc, 0x3d, 0x1e, 0x42, 0x70, 0x30, 0x72, 0xac, 0xf7, 0xa5, 0xec, 0x0e, 0xaf, 0xe4, 0x0d,
	0x54, 0x07, 0x79, 0xb5, 0x77, 0x81, 0x75, 0x43, 0x96, 0xba, 0x6f, 0x8a, 0x5c, 0xb1, 0x2a, 0x4b,
	0x0f, 0x56, 0x6d, 0x7d, 0x68, 0xba, 0x17, 0x8e, 0x61, 0x14, 0xf7, 0x5d, 0xed, 0xea, 0xa6, 0x4a,
	0xb4, 0x6b, 0xcd, 0xb4, 0x65, 0xa9, 0xff, 0x0d, 0xf6, 0xf3, 0x63, 0x59, 0x32, 0x0f, 0x3c, 0x86,
	0x02, 0xa8, 0x65, 0x65, 0xf6, 0x90, 0xe8, 0xf9, 0x7d, 0x7f, 0x93, 0x95, 0xdd, 0x6e, 0xbd, 0x78,
	0x98, 0xb9, 0xd8, 0xa9, 0x8e, 0xf4, 0x52, 0x7a, 0x77, 0x06, 0x47, 0x5e, 0x34, 0xab, 0x62, 0x34,
	0x0e, 0x94, 0xec, 0x43, 0xf0, 0xe9, 0xb0, 0x68, 0xf4, 0x68, 0x1c, 0xf4, 0xb2, 0xc6, 0xcd, 0x4e,
	0xbe, 0xa7, 0xaf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x32, 0x34, 0x08, 0x28, 0x04, 0x00,
	0x00,
}
