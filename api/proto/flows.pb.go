// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flows.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
	proto2 "www.velocidex.com/golang/velociraptor/crypto/proto"
	proto1 "www.velocidex.com/golang/velociraptor/flows/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum values here correspond to Flow.State values.
type ApiFlow_State int32

const (
	ApiFlow_RUNNING        ApiFlow_State = 0
	ApiFlow_TERMINATED     ApiFlow_State = 1
	ApiFlow_ERROR          ApiFlow_State = 3
	ApiFlow_CLIENT_CRASHED ApiFlow_State = 4
)

var ApiFlow_State_name = map[int32]string{
	0: "RUNNING",
	1: "TERMINATED",
	3: "ERROR",
	4: "CLIENT_CRASHED",
}

var ApiFlow_State_value = map[string]int32{
	"RUNNING":        0,
	"TERMINATED":     1,
	"ERROR":          3,
	"CLIENT_CRASHED": 4,
}

func (x ApiFlow_State) String() string {
	return proto.EnumName(ApiFlow_State_name, int32(x))
}

func (ApiFlow_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{3, 0}
}

type AvailableDownloadFile struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Complete             bool     `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
	Size                 uint64   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Date                 string   `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvailableDownloadFile) Reset()         { *m = AvailableDownloadFile{} }
func (m *AvailableDownloadFile) String() string { return proto.CompactTextString(m) }
func (*AvailableDownloadFile) ProtoMessage()    {}
func (*AvailableDownloadFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{0}
}

func (m *AvailableDownloadFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailableDownloadFile.Unmarshal(m, b)
}
func (m *AvailableDownloadFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailableDownloadFile.Marshal(b, m, deterministic)
}
func (m *AvailableDownloadFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableDownloadFile.Merge(m, src)
}
func (m *AvailableDownloadFile) XXX_Size() int {
	return xxx_messageInfo_AvailableDownloadFile.Size(m)
}
func (m *AvailableDownloadFile) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableDownloadFile.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableDownloadFile proto.InternalMessageInfo

func (m *AvailableDownloadFile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AvailableDownloadFile) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *AvailableDownloadFile) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func (m *AvailableDownloadFile) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *AvailableDownloadFile) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type AvailableDownloads struct {
	Files                []*AvailableDownloadFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AvailableDownloads) Reset()         { *m = AvailableDownloads{} }
func (m *AvailableDownloads) String() string { return proto.CompactTextString(m) }
func (*AvailableDownloads) ProtoMessage()    {}
func (*AvailableDownloads) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{1}
}

func (m *AvailableDownloads) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailableDownloads.Unmarshal(m, b)
}
func (m *AvailableDownloads) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailableDownloads.Marshal(b, m, deterministic)
}
func (m *AvailableDownloads) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableDownloads.Merge(m, src)
}
func (m *AvailableDownloads) XXX_Size() int {
	return xxx_messageInfo_AvailableDownloads.Size(m)
}
func (m *AvailableDownloads) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableDownloads.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableDownloads proto.InternalMessageInfo

func (m *AvailableDownloads) GetFiles() []*AvailableDownloadFile {
	if m != nil {
		return m.Files
	}
	return nil
}

type FlowDetails struct {
	Context              *proto1.ArtifactCollectorContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	AvailableDownloads   *AvailableDownloads              `protobuf:"bytes,16,opt,name=available_downloads,json=availableDownloads,proto3" json:"available_downloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *FlowDetails) Reset()         { *m = FlowDetails{} }
func (m *FlowDetails) String() string { return proto.CompactTextString(m) }
func (*FlowDetails) ProtoMessage()    {}
func (*FlowDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{2}
}

func (m *FlowDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowDetails.Unmarshal(m, b)
}
func (m *FlowDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowDetails.Marshal(b, m, deterministic)
}
func (m *FlowDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowDetails.Merge(m, src)
}
func (m *FlowDetails) XXX_Size() int {
	return xxx_messageInfo_FlowDetails.Size(m)
}
func (m *FlowDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowDetails.DiscardUnknown(m)
}

var xxx_messageInfo_FlowDetails proto.InternalMessageInfo

func (m *FlowDetails) GetContext() *proto1.ArtifactCollectorContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *FlowDetails) GetAvailableDownloads() *AvailableDownloads {
	if m != nil {
		return m.AvailableDownloads
	}
	return nil
}

// Next id: 16
type ApiFlow struct {
	Urn                  string                 `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	FlowId               string                 `protobuf:"bytes,12,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	ClientId             string                 `protobuf:"bytes,15,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Args                 *any.Any               `protobuf:"bytes,3,opt,name=args,proto3" json:"args,omitempty"`
	RunnerArgs           *proto1.FlowRunnerArgs `protobuf:"bytes,4,opt,name=runner_args,json=runnerArgs,proto3" json:"runner_args,omitempty"`
	State                ApiFlow_State          `protobuf:"varint,5,opt,name=state,proto3,enum=proto.ApiFlow_State" json:"state,omitempty"`
	StartedAt            uint64                 `protobuf:"varint,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	LastActiveAt         uint64                 `protobuf:"varint,7,opt,name=last_active_at,json=lastActiveAt,proto3" json:"last_active_at,omitempty"`
	Creator              string                 `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	Context              *proto1.FlowContext    `protobuf:"bytes,11,opt,name=context,proto3" json:"context,omitempty"`
	NestedFlows          []*ApiFlow             `protobuf:"bytes,9,rep,name=nested_flows,json=nestedFlows,proto3" json:"nested_flows,omitempty"`
	InternalError        string                 `protobuf:"bytes,14,opt,name=internal_error,json=internalError,proto3" json:"internal_error,omitempty"`
	AvailableDownloads   *AvailableDownloads    `protobuf:"bytes,16,opt,name=available_downloads,json=availableDownloads,proto3" json:"available_downloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ApiFlow) Reset()         { *m = ApiFlow{} }
func (m *ApiFlow) String() string { return proto.CompactTextString(m) }
func (*ApiFlow) ProtoMessage()    {}
func (*ApiFlow) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{3}
}

func (m *ApiFlow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiFlow.Unmarshal(m, b)
}
func (m *ApiFlow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiFlow.Marshal(b, m, deterministic)
}
func (m *ApiFlow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiFlow.Merge(m, src)
}
func (m *ApiFlow) XXX_Size() int {
	return xxx_messageInfo_ApiFlow.Size(m)
}
func (m *ApiFlow) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiFlow.DiscardUnknown(m)
}

var xxx_messageInfo_ApiFlow proto.InternalMessageInfo

func (m *ApiFlow) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *ApiFlow) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *ApiFlow) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ApiFlow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApiFlow) GetArgs() *any.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ApiFlow) GetRunnerArgs() *proto1.FlowRunnerArgs {
	if m != nil {
		return m.RunnerArgs
	}
	return nil
}

func (m *ApiFlow) GetState() ApiFlow_State {
	if m != nil {
		return m.State
	}
	return ApiFlow_RUNNING
}

func (m *ApiFlow) GetStartedAt() uint64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

func (m *ApiFlow) GetLastActiveAt() uint64 {
	if m != nil {
		return m.LastActiveAt
	}
	return 0
}

func (m *ApiFlow) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ApiFlow) GetContext() *proto1.FlowContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *ApiFlow) GetNestedFlows() []*ApiFlow {
	if m != nil {
		return m.NestedFlows
	}
	return nil
}

func (m *ApiFlow) GetInternalError() string {
	if m != nil {
		return m.InternalError
	}
	return ""
}

func (m *ApiFlow) GetAvailableDownloads() *AvailableDownloads {
	if m != nil {
		return m.AvailableDownloads
	}
	return nil
}

// Since Velociraptor processes the responses immediately they are not
// stored.
type ApiFlowRequestDetails struct {
	Items                []*proto2.GrrMessage `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ApiFlowRequestDetails) Reset()         { *m = ApiFlowRequestDetails{} }
func (m *ApiFlowRequestDetails) String() string { return proto.CompactTextString(m) }
func (*ApiFlowRequestDetails) ProtoMessage()    {}
func (*ApiFlowRequestDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{4}
}

func (m *ApiFlowRequestDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiFlowRequestDetails.Unmarshal(m, b)
}
func (m *ApiFlowRequestDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiFlowRequestDetails.Marshal(b, m, deterministic)
}
func (m *ApiFlowRequestDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiFlowRequestDetails.Merge(m, src)
}
func (m *ApiFlowRequestDetails) XXX_Size() int {
	return xxx_messageInfo_ApiFlowRequestDetails.Size(m)
}
func (m *ApiFlowRequestDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiFlowRequestDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ApiFlowRequestDetails proto.InternalMessageInfo

func (m *ApiFlowRequestDetails) GetItems() []*proto2.GrrMessage {
	if m != nil {
		return m.Items
	}
	return nil
}

type ApiFlowResultDetails struct {
	Items                []*proto2.GrrMessage `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ApiFlowResultDetails) Reset()         { *m = ApiFlowResultDetails{} }
func (m *ApiFlowResultDetails) String() string { return proto.CompactTextString(m) }
func (*ApiFlowResultDetails) ProtoMessage()    {}
func (*ApiFlowResultDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{5}
}

func (m *ApiFlowResultDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiFlowResultDetails.Unmarshal(m, b)
}
func (m *ApiFlowResultDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiFlowResultDetails.Marshal(b, m, deterministic)
}
func (m *ApiFlowResultDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiFlowResultDetails.Merge(m, src)
}
func (m *ApiFlowResultDetails) XXX_Size() int {
	return xxx_messageInfo_ApiFlowResultDetails.Size(m)
}
func (m *ApiFlowResultDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiFlowResultDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ApiFlowResultDetails proto.InternalMessageInfo

func (m *ApiFlowResultDetails) GetItems() []*proto2.GrrMessage {
	if m != nil {
		return m.Items
	}
	return nil
}

type ApiFlowLogDetails struct {
	Items                []*proto2.LogMessage `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ApiFlowLogDetails) Reset()         { *m = ApiFlowLogDetails{} }
func (m *ApiFlowLogDetails) String() string { return proto.CompactTextString(m) }
func (*ApiFlowLogDetails) ProtoMessage()    {}
func (*ApiFlowLogDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{6}
}

func (m *ApiFlowLogDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiFlowLogDetails.Unmarshal(m, b)
}
func (m *ApiFlowLogDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiFlowLogDetails.Marshal(b, m, deterministic)
}
func (m *ApiFlowLogDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiFlowLogDetails.Merge(m, src)
}
func (m *ApiFlowLogDetails) XXX_Size() int {
	return xxx_messageInfo_ApiFlowLogDetails.Size(m)
}
func (m *ApiFlowLogDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiFlowLogDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ApiFlowLogDetails proto.InternalMessageInfo

func (m *ApiFlowLogDetails) GetItems() []*proto2.LogMessage {
	if m != nil {
		return m.Items
	}
	return nil
}

type ApiFlowRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	FlowId               string   `protobuf:"bytes,2,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Offset               uint64   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                uint64   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	IncludeArchived      bool     `protobuf:"varint,5,opt,name=include_archived,json=includeArchived,proto3" json:"include_archived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiFlowRequest) Reset()         { *m = ApiFlowRequest{} }
func (m *ApiFlowRequest) String() string { return proto.CompactTextString(m) }
func (*ApiFlowRequest) ProtoMessage()    {}
func (*ApiFlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{7}
}

func (m *ApiFlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiFlowRequest.Unmarshal(m, b)
}
func (m *ApiFlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiFlowRequest.Marshal(b, m, deterministic)
}
func (m *ApiFlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiFlowRequest.Merge(m, src)
}
func (m *ApiFlowRequest) XXX_Size() int {
	return xxx_messageInfo_ApiFlowRequest.Size(m)
}
func (m *ApiFlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiFlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiFlowRequest proto.InternalMessageInfo

func (m *ApiFlowRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ApiFlowRequest) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *ApiFlowRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ApiFlowRequest) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ApiFlowRequest) GetIncludeArchived() bool {
	if m != nil {
		return m.IncludeArchived
	}
	return false
}

type ApiFlowResponse struct {
	Items                []*proto1.ArtifactCollectorContext `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ApiFlowResponse) Reset()         { *m = ApiFlowResponse{} }
func (m *ApiFlowResponse) String() string { return proto.CompactTextString(m) }
func (*ApiFlowResponse) ProtoMessage()    {}
func (*ApiFlowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{8}
}

func (m *ApiFlowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiFlowResponse.Unmarshal(m, b)
}
func (m *ApiFlowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiFlowResponse.Marshal(b, m, deterministic)
}
func (m *ApiFlowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiFlowResponse.Merge(m, src)
}
func (m *ApiFlowResponse) XXX_Size() int {
	return xxx_messageInfo_ApiFlowResponse.Size(m)
}
func (m *ApiFlowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiFlowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiFlowResponse proto.InternalMessageInfo

func (m *ApiFlowResponse) GetItems() []*proto1.ArtifactCollectorContext {
	if m != nil {
		return m.Items
	}
	return nil
}

type FlowDescriptors struct {
	Items                []*proto1.FlowDescriptor `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FlowDescriptors) Reset()         { *m = FlowDescriptors{} }
func (m *FlowDescriptors) String() string { return proto.CompactTextString(m) }
func (*FlowDescriptors) ProtoMessage()    {}
func (*FlowDescriptors) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6bd6ebccf1a7f18, []int{9}
}

func (m *FlowDescriptors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowDescriptors.Unmarshal(m, b)
}
func (m *FlowDescriptors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowDescriptors.Marshal(b, m, deterministic)
}
func (m *FlowDescriptors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowDescriptors.Merge(m, src)
}
func (m *FlowDescriptors) XXX_Size() int {
	return xxx_messageInfo_FlowDescriptors.Size(m)
}
func (m *FlowDescriptors) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowDescriptors.DiscardUnknown(m)
}

var xxx_messageInfo_FlowDescriptors proto.InternalMessageInfo

func (m *FlowDescriptors) GetItems() []*proto1.FlowDescriptor {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.ApiFlow_State", ApiFlow_State_name, ApiFlow_State_value)
	proto.RegisterType((*AvailableDownloadFile)(nil), "proto.AvailableDownloadFile")
	proto.RegisterType((*AvailableDownloads)(nil), "proto.AvailableDownloads")
	proto.RegisterType((*FlowDetails)(nil), "proto.FlowDetails")
	proto.RegisterType((*ApiFlow)(nil), "proto.ApiFlow")
	proto.RegisterType((*ApiFlowRequestDetails)(nil), "proto.ApiFlowRequestDetails")
	proto.RegisterType((*ApiFlowResultDetails)(nil), "proto.ApiFlowResultDetails")
	proto.RegisterType((*ApiFlowLogDetails)(nil), "proto.ApiFlowLogDetails")
	proto.RegisterType((*ApiFlowRequest)(nil), "proto.ApiFlowRequest")
	proto.RegisterType((*ApiFlowResponse)(nil), "proto.ApiFlowResponse")
	proto.RegisterType((*FlowDescriptors)(nil), "proto.FlowDescriptors")
}

func init() { proto.RegisterFile("flows.proto", fileDescriptor_b6bd6ebccf1a7f18) }

var fileDescriptor_b6bd6ebccf1a7f18 = []byte{
	// 1101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0xd9, 0xc4, 0x8e, 0xed, 0x71, 0x70, 0xdc, 0x69, 0x92, 0x6e, 0xd2, 0x42, 0x56, 0x16,
	0x85, 0xb4, 0x95, 0xd6, 0x52, 0x00, 0x21, 0x50, 0xf9, 0xb1, 0xb1, 0x9d, 0xc4, 0x25, 0x75, 0xa5,
	0x49, 0x42, 0x4f, 0xc8, 0x1a, 0xef, 0x8e, 0xed, 0x41, 0xeb, 0x19, 0x33, 0x33, 0x1b, 0x27, 0x3d,
	0x21, 0x71, 0xe4, 0xca, 0x9d, 0x2b, 0x7f, 0x08, 0x7f, 0x09, 0xfc, 0x1b, 0x1c, 0xd0, 0xbc, 0xdd,
	0x8d, 0x6d, 0x0a, 0x6a, 0x2b, 0xf5, 0xb4, 0x33, 0xef, 0xbd, 0xef, 0x67, 0xde, 0xbc, 0x99, 0x7d,
	0x83, 0xaa, 0xc3, 0x58, 0xce, 0xb4, 0x3f, 0x55, 0xd2, 0x48, 0x5c, 0x84, 0xcf, 0xee, 0x17, 0xb3,
	0xd9, 0xcc, 0xbf, 0x64, 0xb1, 0x0c, 0x79, 0xc4, 0xae, 0xfc, 0x50, 0x4e, 0x9a, 0x23, 0x19, 0x53,
	0x31, 0x6a, 0xa6, 0x46, 0x45, 0xa7, 0x46, 0xaa, 0x26, 0x04, 0x37, 0x35, 0x9b, 0x50, 0x61, 0x78,
	0x98, 0x22, 0x76, 0xbf, 0x7c, 0x3d, 0x2d, 0xac, 0x9a, 0x11, 0x16, 0x32, 0x78, 0x5d, 0x79, 0xa8,
	0xae, 0xa7, 0x46, 0x66, 0xfa, 0x1f, 0xe4, 0x20, 0x97, 0xef, 0x8c, 0xa4, 0x1c, 0xc5, 0x2c, 0x75,
	0x0c, 0x92, 0x61, 0x93, 0x8a, 0xeb, 0xcc, 0x15, 0xbc, 0x79, 0x62, 0x54, 0x19, 0x3e, 0xa4, 0xa1,
	0xc9, 0xe8, 0x8d, 0x9f, 0x1d, 0xb4, 0x15, 0x5c, 0x52, 0x1e, 0xd3, 0x41, 0xcc, 0xda, 0x72, 0x26,
	0x62, 0x49, 0xa3, 0x23, 0x1e, 0x33, 0x8c, 0x51, 0x41, 0xd0, 0x09, 0x73, 0x1d, 0xcf, 0xd9, 0xaf,
	0x10, 0x18, 0x5b, 0xdb, 0x94, 0x9a, 0xb1, 0x5b, 0x4c, 0x6d, 0x76, 0x8c, 0x77, 0x51, 0x39, 0x94,
	0x93, 0x69, 0xcc, 0x0c, 0x73, 0x57, 0x3c, 0x67, 0xbf, 0x4c, 0x6e, 0xe6, 0x36, 0x5e, 0xf3, 0x17,
	0xcc, 0x5d, 0xf5, 0x9c, 0xfd, 0x02, 0x81, 0xb1, 0xb5, 0x45, 0xd4, 0x30, 0xb7, 0x90, 0x32, 0xec,
	0xb8, 0x71, 0x82, 0xf0, 0x4b, 0x49, 0x68, 0x7c, 0x80, 0x8a, 0x43, 0x1e, 0x33, 0xed, 0x3a, 0xde,
	0xea, 0x7e, 0xf5, 0xe0, 0x5e, 0x9a, 0xb2, 0xff, 0x9f, 0xe9, 0x92, 0x34, 0xb4, 0xf1, 0xab, 0x83,
	0xaa, 0x47, 0xb1, 0x9c, 0xb5, 0x99, 0xa1, 0x3c, 0xd6, 0xf8, 0x73, 0x54, 0x0a, 0xa5, 0x30, 0xec,
	0xca, 0xc0, 0x46, 0xaa, 0x07, 0x7b, 0x39, 0x25, 0x2b, 0x44, 0x4b, 0xc6, 0x31, 0x0b, 0x8d, 0x54,
	0xad, 0x34, 0x8c, 0xe4, 0xf1, 0xf8, 0x09, 0xba, 0x4d, 0xf3, 0xa5, 0xfa, 0x51, 0x9e, 0x95, 0x5b,
	0x07, 0xcc, 0xce, 0xff, 0x25, 0xa3, 0x09, 0xa6, 0x2f, 0xd9, 0x1a, 0xbf, 0x54, 0x50, 0x29, 0x98,
	0x72, 0x9b, 0x19, 0xf6, 0xd1, 0x6a, 0xa2, 0x44, 0x5a, 0xd7, 0xc3, 0x7b, 0x7f, 0xfe, 0xfd, 0xd7,
	0x1f, 0xce, 0x36, 0xaa, 0x9c, 0x31, 0xad, 0xb9, 0x14, 0xdd, 0x36, 0xae, 0xd8, 0x28, 0xef, 0x82,
	0xf4, 0x7c, 0x62, 0x03, 0xf1, 0x27, 0xa8, 0x64, 0x4f, 0xb0, 0xcf, 0x23, 0x77, 0x1d, 0x34, 0x77,
	0x41, 0xb3, 0x85, 0x2a, 0x19, 0xb1, 0x1b, 0xe1, 0x32, 0x68, 0x78, 0xe4, 0x93, 0xb5, 0x61, 0x6a,
	0x79, 0x8c, 0x2a, 0x61, 0xcc, 0x99, 0x30, 0x56, 0xb7, 0x01, 0xba, 0x3d, 0xd0, 0xed, 0xa0, 0x6a,
	0x30, 0xe5, 0x2d, 0xf0, 0x75, 0x23, 0x8c, 0xd2, 0x11, 0x68, 0xcb, 0x61, 0x6e, 0xfd, 0x30, 0x3b,
	0xfc, 0x15, 0x10, 0x62, 0x10, 0xae, 0x63, 0x04, 0xab, 0x58, 0x87, 0x9f, 0x5d, 0x88, 0x67, 0xa8,
	0x40, 0xd5, 0x48, 0xc3, 0x01, 0x57, 0x0f, 0x36, 0xfd, 0xf4, 0xae, 0xfa, 0xf9, 0x5d, 0xf5, 0x03,
	0x71, 0x7d, 0x78, 0x1f, 0xd4, 0x7b, 0x78, 0x03, 0xd4, 0x54, 0x8d, 0x92, 0x09, 0x13, 0x46, 0xfb,
	0x0f, 0xd7, 0x8f, 0x99, 0x09, 0xd4, 0x48, 0xb7, 0x62, 0xaa, 0x35, 0x01, 0x10, 0xfe, 0x0e, 0x55,
	0x55, 0x22, 0x04, 0x53, 0x7d, 0xe0, 0x16, 0x80, 0xbb, 0x95, 0x15, 0xdb, 0x02, 0x08, 0x78, 0xad,
	0xec, 0xf0, 0x7d, 0x00, 0xbb, 0x78, 0x1b, 0xc0, 0xa9, 0x6c, 0x81, 0x4f, 0x90, 0xba, 0x89, 0xc5,
	0xc7, 0xa8, 0xa8, 0x8d, 0xbd, 0x76, 0xf6, 0xea, 0xd6, 0x0e, 0x36, 0xf3, 0xe3, 0x4b, 0x2b, 0xe8,
	0x9f, 0x59, 0x5f, 0x5e, 0x58, 0x7c, 0xbb, 0x95, 0x28, 0x65, 0x8b, 0x62, 0x6b, 0xe9, 0x81, 0xce,
	0x27, 0xa9, 0x1e, 0x3f, 0x41, 0x48, 0x1b, 0xaa, 0x0c, 0x8b, 0xfa, 0xd4, 0xb8, 0x6b, 0xf6, 0x62,
	0x1f, 0x3e, 0x02, 0xdd, 0x7d, 0x54, 0x25, 0xed, 0xa3, 0x36, 0x35, 0xcc, 0xf0, 0x09, 0xc3, 0xdb,
	0xcf, 0xc7, 0x4c, 0xa4, 0x84, 0x19, 0xd5, 0x5e, 0xa8, 0x18, 0x35, 0x2c, 0xf2, 0x49, 0x25, 0x93,
	0x07, 0x06, 0x5f, 0xa0, 0x5a, 0x4c, 0xb5, 0xe9, 0xd3, 0xd0, 0xf0, 0x4b, 0x66, 0x79, 0x25, 0xe0,
	0x35, 0x81, 0xf7, 0x60, 0x99, 0xb7, 0xbb, 0xcc, 0xb3, 0x42, 0x2f, 0x15, 0xfa, 0x64, 0xdd, 0xce,
	0x02, 0x98, 0x04, 0x06, 0x7f, 0x86, 0x4a, 0xb0, 0x9a, 0x54, 0x6e, 0x19, 0xce, 0xef, 0x3d, 0xe0,
	0xdd, 0xc1, 0x5b, 0xcf, 0xc7, 0xd2, 0xcb, 0x96, 0xf7, 0xcc, 0x98, 0x01, 0xcd, 0x27, 0x79, 0x34,
	0xfe, 0x76, 0xfe, 0xb3, 0x54, 0xa1, 0xf0, 0x78, 0xa1, 0xf0, 0xd9, 0xff, 0x31, 0x87, 0x2d, 0x15,
	0x29, 0xd3, 0xf9, 0xf3, 0xdf, 0xe7, 0x7b, 0xb4, 0x2e, 0x98, 0xb6, 0x75, 0x82, 0xfe, 0xe3, 0x56,
	0xe0, 0x27, 0xae, 0x2d, 0x17, 0xfe, 0xf0, 0x21, 0xd0, 0x3e, 0xc0, 0x8d, 0x1e, 0xc4, 0x02, 0x4c,
	0xdf, 0xe4, 0x38, 0xb8, 0xf6, 0xcc, 0x98, 0xeb, 0x2c, 0xcf, 0x6a, 0xca, 0xb3, 0x3a, 0x8d, 0x7f,
	0x77, 0x50, 0x8d, 0x0b, 0xc3, 0x94, 0xa0, 0x71, 0x9f, 0x29, 0x25, 0x95, 0x5b, 0x83, 0xcd, 0xfe,
	0xe4, 0x00, 0xf2, 0x05, 0xbe, 0x0a, 0x84, 0x07, 0x0e, 0xcf, 0x8c, 0xa9, 0xf1, 0xc6, 0x74, 0x3a,
	0x65, 0x82, 0x45, 0xde, 0x6c, 0xcc, 0x63, 0xe6, 0x29, 0x46, 0x23, 0x2e, 0x46, 0xf3, 0x32, 0x78,
	0xe7, 0x76, 0x25, 0xae, 0x3d, 0x21, 0x8d, 0x47, 0x73, 0x69, 0x1a, 0xcc, 0xae, 0x58, 0x98, 0x98,
	0xc5, 0x70, 0x6f, 0x90, 0x40, 0x54, 0x9e, 0x82, 0x17, 0x51, 0x43, 0x3d, 0xae, 0x75, 0xc2, 0x7c,
	0xf2, 0x6e, 0x6e, 0xed, 0x58, 0xc6, 0x5b, 0x6d, 0x24, 0x2d, 0x54, 0x84, 0xbb, 0x8a, 0xab, 0xa8,
	0x44, 0x2e, 0x7a, 0xbd, 0x6e, 0xef, 0xb8, 0xfe, 0x0e, 0xae, 0x21, 0x74, 0xde, 0x21, 0x4f, 0xbb,
	0xbd, 0xe0, 0xbc, 0xd3, 0xae, 0x3b, 0xb8, 0x82, 0x8a, 0x1d, 0x42, 0x9e, 0x91, 0xfa, 0x2a, 0xc6,
	0xa8, 0xd6, 0x3a, 0xed, 0x76, 0x7a, 0xe7, 0xfd, 0x16, 0x09, 0xce, 0x4e, 0x3a, 0xed, 0x7a, 0xa1,
	0xf1, 0x0d, 0xda, 0xca, 0xea, 0x4f, 0xd8, 0x8f, 0x09, 0xd3, 0x26, 0xef, 0x96, 0x1f, 0xa1, 0x22,
	0x37, 0x6c, 0x92, 0x77, 0xdc, 0x5b, 0x59, 0x6e, 0xc7, 0x4a, 0x3d, 0x65, 0x5a, 0xd3, 0x11, 0x23,
	0xa9, 0xbf, 0xf1, 0x35, 0xda, 0xbc, 0x21, 0xe8, 0x24, 0x7e, 0x73, 0xc0, 0x63, 0x74, 0x2b, 0x03,
	0x9c, 0xca, 0xd1, 0x2b, 0xd4, 0xa7, 0x72, 0xf4, 0x2f, 0xf5, 0x6f, 0x0e, 0xaa, 0x2d, 0xef, 0x00,
	0xdf, 0x5d, 0xec, 0x77, 0xe9, 0x9b, 0x35, 0x6f, 0x67, 0x77, 0xe6, 0x2d, 0x14, 0x3a, 0xda, 0x4d,
	0x97, 0xdc, 0x46, 0x6b, 0x72, 0x38, 0xd4, 0xcc, 0x64, 0x4f, 0x54, 0x36, 0xc3, 0x9b, 0xa8, 0x18,
	0xca, 0x44, 0x18, 0x68, 0x40, 0x05, 0x92, 0x4e, 0xf0, 0x03, 0x54, 0xe7, 0x22, 0x8c, 0x93, 0x88,
	0xf5, 0xa9, 0x0a, 0xc7, 0xfc, 0x92, 0x45, 0xd0, 0x4f, 0xca, 0x64, 0x23, 0xb3, 0x07, 0x99, 0xb9,
	0x71, 0x82, 0x36, 0xe6, 0x05, 0x9a, 0x4a, 0xa1, 0x19, 0xfe, 0x34, 0xdf, 0xdd, 0x0a, 0xec, 0xee,
	0x95, 0x0f, 0x51, 0xb6, 0xd7, 0xaf, 0xd0, 0x46, 0xfa, 0xa0, 0xe9, 0x50, 0x71, 0xfb, 0x9e, 0x6b,
	0xfc, 0x68, 0xb9, 0x4e, 0x8b, 0xed, 0x71, 0x1e, 0x96, 0xe9, 0x07, 0x6b, 0xe0, 0xfc, 0xf8, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x3f, 0x1e, 0xfd, 0x16, 0x09, 0x00, 0x00,
}
