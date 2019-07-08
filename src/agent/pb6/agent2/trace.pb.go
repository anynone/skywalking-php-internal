// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent-v2/trace.proto

package agent2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	common "pb6/common"
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

type SegmentObject struct {
	TraceSegmentId       *common.UniqueId `protobuf:"bytes,1,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	Spans                []*SpanObjectV2  `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	ServiceId            int32            `protobuf:"varint,3,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	ServiceInstanceId    int32            `protobuf:"varint,4,opt,name=serviceInstanceId,proto3" json:"serviceInstanceId,omitempty"`
	IsSizeLimited        bool             `protobuf:"varint,5,opt,name=isSizeLimited,proto3" json:"isSizeLimited,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SegmentObject) Reset()         { *m = SegmentObject{} }
func (m *SegmentObject) String() string { return proto.CompactTextString(m) }
func (*SegmentObject) ProtoMessage()    {}
func (*SegmentObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{0}
}

func (m *SegmentObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentObject.Unmarshal(m, b)
}
func (m *SegmentObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentObject.Marshal(b, m, deterministic)
}
func (m *SegmentObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentObject.Merge(m, src)
}
func (m *SegmentObject) XXX_Size() int {
	return xxx_messageInfo_SegmentObject.Size(m)
}
func (m *SegmentObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentObject.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentObject proto.InternalMessageInfo

func (m *SegmentObject) GetTraceSegmentId() *common.UniqueId {
	if m != nil {
		return m.TraceSegmentId
	}
	return nil
}

func (m *SegmentObject) GetSpans() []*SpanObjectV2 {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *SegmentObject) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *SegmentObject) GetServiceInstanceId() int32 {
	if m != nil {
		return m.ServiceInstanceId
	}
	return 0
}

func (m *SegmentObject) GetIsSizeLimited() bool {
	if m != nil {
		return m.IsSizeLimited
	}
	return false
}

type SegmentReference struct {
	RefType                 common.RefType   `protobuf:"varint,1,opt,name=refType,proto3,enum=RefType" json:"refType,omitempty"`
	ParentTraceSegmentId    *common.UniqueId `protobuf:"bytes,2,opt,name=parentTraceSegmentId,proto3" json:"parentTraceSegmentId,omitempty"`
	ParentSpanId            int32            `protobuf:"varint,3,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	ParentServiceInstanceId int32            `protobuf:"varint,4,opt,name=parentServiceInstanceId,proto3" json:"parentServiceInstanceId,omitempty"`
	NetworkAddress          string           `protobuf:"bytes,5,opt,name=networkAddress,proto3" json:"networkAddress,omitempty"`
	NetworkAddressId        int32            `protobuf:"varint,6,opt,name=networkAddressId,proto3" json:"networkAddressId,omitempty"`
	EntryServiceInstanceId  int32            `protobuf:"varint,7,opt,name=entryServiceInstanceId,proto3" json:"entryServiceInstanceId,omitempty"`
	EntryEndpoint           string           `protobuf:"bytes,8,opt,name=entryEndpoint,proto3" json:"entryEndpoint,omitempty"`
	EntryEndpointId         int32            `protobuf:"varint,9,opt,name=entryEndpointId,proto3" json:"entryEndpointId,omitempty"`
	ParentEndpoint          string           `protobuf:"bytes,10,opt,name=parentEndpoint,proto3" json:"parentEndpoint,omitempty"`
	ParentEndpointId        int32            `protobuf:"varint,11,opt,name=parentEndpointId,proto3" json:"parentEndpointId,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}         `json:"-"`
	XXX_unrecognized        []byte           `json:"-"`
	XXX_sizecache           int32            `json:"-"`
}

func (m *SegmentReference) Reset()         { *m = SegmentReference{} }
func (m *SegmentReference) String() string { return proto.CompactTextString(m) }
func (*SegmentReference) ProtoMessage()    {}
func (*SegmentReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{1}
}

func (m *SegmentReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentReference.Unmarshal(m, b)
}
func (m *SegmentReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentReference.Marshal(b, m, deterministic)
}
func (m *SegmentReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentReference.Merge(m, src)
}
func (m *SegmentReference) XXX_Size() int {
	return xxx_messageInfo_SegmentReference.Size(m)
}
func (m *SegmentReference) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentReference.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentReference proto.InternalMessageInfo

func (m *SegmentReference) GetRefType() common.RefType {
	if m != nil {
		return m.RefType
	}
	return common.RefType_CrossProcess
}

func (m *SegmentReference) GetParentTraceSegmentId() *common.UniqueId {
	if m != nil {
		return m.ParentTraceSegmentId
	}
	return nil
}

func (m *SegmentReference) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *SegmentReference) GetParentServiceInstanceId() int32 {
	if m != nil {
		return m.ParentServiceInstanceId
	}
	return 0
}

func (m *SegmentReference) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *SegmentReference) GetNetworkAddressId() int32 {
	if m != nil {
		return m.NetworkAddressId
	}
	return 0
}

func (m *SegmentReference) GetEntryServiceInstanceId() int32 {
	if m != nil {
		return m.EntryServiceInstanceId
	}
	return 0
}

func (m *SegmentReference) GetEntryEndpoint() string {
	if m != nil {
		return m.EntryEndpoint
	}
	return ""
}

func (m *SegmentReference) GetEntryEndpointId() int32 {
	if m != nil {
		return m.EntryEndpointId
	}
	return 0
}

func (m *SegmentReference) GetParentEndpoint() string {
	if m != nil {
		return m.ParentEndpoint
	}
	return ""
}

func (m *SegmentReference) GetParentEndpointId() int32 {
	if m != nil {
		return m.ParentEndpointId
	}
	return 0
}

type SpanObjectV2 struct {
	SpanId               int32                        `protobuf:"varint,1,opt,name=spanId,proto3" json:"spanId,omitempty"`
	ParentSpanId         int32                        `protobuf:"varint,2,opt,name=parentSpanId,proto3" json:"parentSpanId,omitempty"`
	StartTime            int64                        `protobuf:"varint,3,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64                        `protobuf:"varint,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Refs                 []*SegmentReference          `protobuf:"bytes,5,rep,name=refs,proto3" json:"refs,omitempty"`
	OperationNameId      int32                        `protobuf:"varint,6,opt,name=operationNameId,proto3" json:"operationNameId,omitempty"`
	OperationName        string                       `protobuf:"bytes,7,opt,name=operationName,proto3" json:"operationName,omitempty"`
	PeerId               int32                        `protobuf:"varint,8,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Peer                 string                       `protobuf:"bytes,9,opt,name=peer,proto3" json:"peer,omitempty"`
	SpanType             common.SpanType              `protobuf:"varint,10,opt,name=spanType,proto3,enum=SpanType" json:"spanType,omitempty"`
	SpanLayer            common.SpanLayer             `protobuf:"varint,11,opt,name=spanLayer,proto3,enum=SpanLayer" json:"spanLayer,omitempty"`
	ComponentId          int32                        `protobuf:"varint,12,opt,name=componentId,proto3" json:"componentId,omitempty"`
	Component            string                       `protobuf:"bytes,13,opt,name=component,proto3" json:"component,omitempty"`
	IsError              bool                         `protobuf:"varint,14,opt,name=isError,proto3" json:"isError,omitempty"`
	Tags                 []*common.KeyStringValuePair `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Logs                 []*Log                       `protobuf:"bytes,16,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SpanObjectV2) Reset()         { *m = SpanObjectV2{} }
func (m *SpanObjectV2) String() string { return proto.CompactTextString(m) }
func (*SpanObjectV2) ProtoMessage()    {}
func (*SpanObjectV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{2}
}

func (m *SpanObjectV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanObjectV2.Unmarshal(m, b)
}
func (m *SpanObjectV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanObjectV2.Marshal(b, m, deterministic)
}
func (m *SpanObjectV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanObjectV2.Merge(m, src)
}
func (m *SpanObjectV2) XXX_Size() int {
	return xxx_messageInfo_SpanObjectV2.Size(m)
}
func (m *SpanObjectV2) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanObjectV2.DiscardUnknown(m)
}

var xxx_messageInfo_SpanObjectV2 proto.InternalMessageInfo

func (m *SpanObjectV2) GetSpanId() int32 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *SpanObjectV2) GetParentSpanId() int32 {
	if m != nil {
		return m.ParentSpanId
	}
	return 0
}

func (m *SpanObjectV2) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *SpanObjectV2) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *SpanObjectV2) GetRefs() []*SegmentReference {
	if m != nil {
		return m.Refs
	}
	return nil
}

func (m *SpanObjectV2) GetOperationNameId() int32 {
	if m != nil {
		return m.OperationNameId
	}
	return 0
}

func (m *SpanObjectV2) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *SpanObjectV2) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *SpanObjectV2) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *SpanObjectV2) GetSpanType() common.SpanType {
	if m != nil {
		return m.SpanType
	}
	return common.SpanType_Entry
}

func (m *SpanObjectV2) GetSpanLayer() common.SpanLayer {
	if m != nil {
		return m.SpanLayer
	}
	return common.SpanLayer_Unknown
}

func (m *SpanObjectV2) GetComponentId() int32 {
	if m != nil {
		return m.ComponentId
	}
	return 0
}

func (m *SpanObjectV2) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *SpanObjectV2) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *SpanObjectV2) GetTags() []*common.KeyStringValuePair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SpanObjectV2) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type Log struct {
	Time                 int64                        `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Data                 []*common.KeyStringValuePair `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_8124ab206744863a, []int{3}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Log) GetData() []*common.KeyStringValuePair {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SegmentObject)(nil), "SegmentObject")
	proto.RegisterType((*SegmentReference)(nil), "SegmentReference")
	proto.RegisterType((*SpanObjectV2)(nil), "SpanObjectV2")
	proto.RegisterType((*Log)(nil), "Log")
}

func init() { proto.RegisterFile("language-agent-v2/trace.proto", fileDescriptor_8124ab206744863a) }

var fileDescriptor_8124ab206744863a = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x6e, 0xeb, 0x44,
	0x10, 0xc6, 0x71, 0xe2, 0x34, 0xf1, 0xa4, 0xc9, 0xc9, 0xd9, 0x83, 0x0e, 0x3e, 0x15, 0x48, 0x51,
	0xe0, 0x40, 0x74, 0x44, 0x5d, 0xe1, 0x4a, 0x15, 0x37, 0x5c, 0x50, 0x54, 0xa1, 0x88, 0xa8, 0x54,
	0x9b, 0xb6, 0x48, 0x5c, 0xb1, 0xb5, 0xa7, 0xc6, 0x24, 0xde, 0x35, 0xeb, 0x6d, 0xab, 0x70, 0xcb,
	0x23, 0xf0, 0x16, 0xbc, 0x11, 0x6f, 0x83, 0x76, 0xec, 0xfc, 0x71, 0xd2, 0x5e, 0x79, 0xe7, 0xf7,
	0xcd, 0xae, 0x76, 0xe6, 0xf3, 0x2c, 0x7c, 0xb6, 0x10, 0x32, 0x79, 0x10, 0x09, 0x1e, 0x8b, 0x04,
	0xa5, 0x39, 0x7e, 0x0c, 0x4f, 0x8c, 0x16, 0x11, 0x06, 0xb9, 0x56, 0x46, 0x1d, 0xbd, 0x89, 0x54,
	0x96, 0x29, 0x79, 0x52, 0x7e, 0x2a, 0xf8, 0xae, 0x82, 0x94, 0x78, 0xbc, 0x2d, 0x8d, 0xfe, 0x73,
	0xa0, 0x37, 0xc3, 0x24, 0x43, 0x69, 0x7e, 0xbe, 0xfb, 0x03, 0x23, 0xc3, 0xbe, 0x81, 0x3e, 0xe5,
	0x55, 0x74, 0x12, 0xfb, 0xce, 0xd0, 0x19, 0x77, 0x43, 0x2f, 0xb8, 0x91, 0xe9, 0x9f, 0x0f, 0x38,
	0x89, 0xf9, 0x4e, 0x02, 0xfb, 0x1c, 0x5a, 0x45, 0x2e, 0x64, 0xe1, 0x37, 0x86, 0xcd, 0x71, 0x37,
	0xec, 0x05, 0xb3, 0x5c, 0xc8, 0xf2, 0xb8, 0xdb, 0x90, 0x97, 0x1a, 0xfb, 0x14, 0xbc, 0x02, 0xf5,
	0x63, 0x1a, 0xe1, 0x24, 0xf6, 0x9b, 0x43, 0x67, 0xdc, 0xe2, 0x1b, 0xc0, 0xbe, 0x86, 0xd7, 0xab,
	0x40, 0x16, 0x46, 0x48, 0xca, 0x72, 0x29, 0x6b, 0x5f, 0x60, 0x5f, 0x40, 0x2f, 0x2d, 0x66, 0xe9,
	0x5f, 0x38, 0x4d, 0xb3, 0xd4, 0x60, 0xec, 0xb7, 0x86, 0xce, 0xb8, 0xc3, 0xeb, 0x70, 0xf4, 0xb7,
	0x0b, 0x83, 0xea, 0x92, 0x1c, 0xef, 0x51, 0xa3, 0x8c, 0x90, 0x8d, 0xa0, 0xad, 0xf1, 0xfe, 0x7a,
	0x99, 0x23, 0xd5, 0xd5, 0x0f, 0x3b, 0x01, 0x2f, 0x63, 0xbe, 0x12, 0xd8, 0x77, 0xf0, 0x71, 0x2e,
	0x34, 0x4a, 0x73, 0x5d, 0x6f, 0x44, 0x63, 0xb7, 0x11, 0xcf, 0xa6, 0xb1, 0x11, 0x1c, 0x96, 0xdc,
	0xb6, 0x61, 0x5d, 0x6c, 0x8d, 0xb1, 0x6f, 0xe1, 0x93, 0x2a, 0x7e, 0xa1, 0xea, 0x97, 0x64, 0xf6,
	0x25, 0xf4, 0x25, 0x9a, 0x27, 0xa5, 0xe7, 0xdf, 0xc7, 0xb1, 0xc6, 0xa2, 0xa0, 0xe2, 0x3d, 0xbe,
	0x43, 0xd9, 0x07, 0x18, 0xd4, 0xc9, 0x24, 0xf6, 0x0f, 0xe8, 0xe8, 0x3d, 0xce, 0xce, 0xe0, 0x2d,
	0x4a, 0xa3, 0x97, 0xfb, 0x97, 0x69, 0xd3, 0x8e, 0x17, 0x54, 0xeb, 0x03, 0x29, 0x17, 0x32, 0xce,
	0x55, 0x2a, 0x8d, 0xdf, 0xa1, 0xab, 0xd4, 0x21, 0x1b, 0xc3, 0xab, 0x1a, 0x98, 0xc4, 0xbe, 0x47,
	0xc7, 0xee, 0x62, 0x5b, 0x5b, 0x59, 0xf6, 0xfa, 0x40, 0x28, 0x6b, 0xab, 0x53, 0x5b, 0x5b, 0x9d,
	0x4c, 0x62, 0xbf, 0x5b, 0xd6, 0xb6, 0xcb, 0x47, 0xff, 0xb8, 0x70, 0xb8, 0xfd, 0x3f, 0xb2, 0xb7,
	0x70, 0x50, 0x94, 0xc6, 0x38, 0xb4, 0xa5, 0x8a, 0xf6, 0x6c, 0x6b, 0x3c, 0x63, 0x9b, 0xfd, 0x89,
	0x8d, 0xd0, 0xe6, 0x3a, 0xcd, 0x90, 0x7c, 0x6d, 0xf2, 0x0d, 0x60, 0x3e, 0xb4, 0x51, 0xc6, 0xa4,
	0xb9, 0xa4, 0xad, 0x42, 0xf6, 0x1e, 0x5c, 0x8d, 0xf7, 0xd6, 0x2a, 0x3b, 0x20, 0xaf, 0x83, 0xdd,
	0xdf, 0x92, 0x93, 0x6c, 0x3b, 0xa5, 0x72, 0xd4, 0xc2, 0xa4, 0x4a, 0x5e, 0x8a, 0x0c, 0xd7, 0x96,
	0xed, 0x62, 0xdb, 0xf9, 0x1a, 0x22, 0xa3, 0x3c, 0x5e, 0x87, 0xb6, 0xd4, 0x1c, 0x51, 0x4f, 0x62,
	0x32, 0xa6, 0xc5, 0xab, 0x88, 0x31, 0x70, 0xed, 0x8a, 0x6c, 0xf0, 0x38, 0xad, 0xd9, 0x7b, 0xe8,
	0xd8, 0x46, 0xd0, 0x64, 0x00, 0x4d, 0x86, 0x47, 0x73, 0x4c, 0xa3, 0xb1, 0x96, 0xd8, 0x18, 0x3c,
	0xbb, 0x9e, 0x8a, 0x25, 0x6a, 0xea, 0x79, 0x3f, 0x04, 0xca, 0x23, 0xc2, 0x37, 0x22, 0x1b, 0x42,
	0x37, 0x52, 0x59, 0xae, 0x64, 0x39, 0x3c, 0x87, 0x74, 0x83, 0x6d, 0x64, 0xbb, 0xb9, 0x0e, 0xfd,
	0x1e, 0xdd, 0x65, 0x03, 0x6c, 0x37, 0xd3, 0xe2, 0x42, 0x6b, 0xa5, 0xfd, 0x3e, 0x8d, 0xf7, 0x2a,
	0x64, 0x5f, 0x81, 0x6b, 0x44, 0x52, 0xf8, 0xaf, 0xa8, 0x9b, 0x6f, 0x82, 0x9f, 0x70, 0x39, 0x33,
	0x3a, 0x95, 0xc9, 0xad, 0x58, 0x3c, 0xe0, 0x95, 0x48, 0x35, 0xa7, 0x04, 0xe6, 0x83, 0xbb, 0x50,
	0x49, 0xe1, 0x0f, 0x28, 0xd1, 0x0d, 0xa6, 0x2a, 0xe1, 0x44, 0x46, 0xe7, 0xd0, 0x9c, 0xaa, 0xc4,
	0x36, 0xc2, 0x58, 0xbb, 0x1c, 0xb2, 0x8b, 0xd6, 0xf6, 0xf4, 0x58, 0x18, 0x51, 0x3d, 0x66, 0xcf,
	0x9f, 0x6e, 0x13, 0xc2, 0x1f, 0xe1, 0xdd, 0xf6, 0xe4, 0x73, 0xcc, 0x95, 0x5e, 0x0d, 0x2c, 0xfb,
	0x00, 0xed, 0x48, 0x2d, 0x16, 0xf6, 0x45, 0x1d, 0x04, 0x37, 0x79, 0x61, 0x34, 0x8a, 0xac, 0xca,
	0x3c, 0xf2, 0x82, 0x1f, 0x54, 0x96, 0x09, 0x19, 0x17, 0xa3, 0x8f, 0xc6, 0xce, 0xf9, 0x6f, 0x70,
	0xaa, 0x74, 0x12, 0x88, 0x5c, 0x44, 0xbf, 0x63, 0x50, 0xcc, 0x97, 0x4f, 0x62, 0x31, 0x4f, 0xa5,
	0x25, 0x59, 0x50, 0x0d, 0x6b, 0xb0, 0x7a, 0xf7, 0x03, 0x7a, 0xf7, 0x83, 0xc7, 0xf0, 0xca, 0xf9,
	0x15, 0xf2, 0xbb, 0xb3, 0x13, 0x8a, 0xc3, 0x7f, 0x1b, 0x47, 0xb3, 0xf9, 0xf2, 0x97, 0x6a, 0xe3,
	0x65, 0xb9, 0xe9, 0xca, 0xbe, 0xf1, 0x91, 0x5a, 0xdc, 0x1d, 0xd0, 0x6b, 0x7f, 0xfa, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xed, 0xe8, 0x41, 0x38, 0x3e, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceSegmentReportServiceClient is the client API for TraceSegmentReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceSegmentReportServiceClient interface {
	Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentReportService_CollectClient, error)
}

type traceSegmentReportServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceSegmentReportServiceClient(cc *grpc.ClientConn) TraceSegmentReportServiceClient {
	return &traceSegmentReportServiceClient{cc}
}

func (c *traceSegmentReportServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentReportService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceSegmentReportService_serviceDesc.Streams[0], "/TraceSegmentReportService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceSegmentReportServiceCollectClient{stream}
	return x, nil
}

type TraceSegmentReportService_CollectClient interface {
	Send(*common.UpstreamSegment) error
	CloseAndRecv() (*common.Commands, error)
	grpc.ClientStream
}

type traceSegmentReportServiceCollectClient struct {
	grpc.ClientStream
}

func (x *traceSegmentReportServiceCollectClient) Send(m *common.UpstreamSegment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceSegmentReportServiceCollectClient) CloseAndRecv() (*common.Commands, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(common.Commands)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceSegmentReportServiceServer is the server API for TraceSegmentReportService service.
type TraceSegmentReportServiceServer interface {
	Collect(TraceSegmentReportService_CollectServer) error
}

// UnimplementedTraceSegmentReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceSegmentReportServiceServer struct {
}

func (*UnimplementedTraceSegmentReportServiceServer) Collect(srv TraceSegmentReportService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterTraceSegmentReportServiceServer(s *grpc.Server, srv TraceSegmentReportServiceServer) {
	s.RegisterService(&_TraceSegmentReportService_serviceDesc, srv)
}

func _TraceSegmentReportService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceSegmentReportServiceServer).Collect(&traceSegmentReportServiceCollectServer{stream})
}

type TraceSegmentReportService_CollectServer interface {
	SendAndClose(*common.Commands) error
	Recv() (*common.UpstreamSegment, error)
	grpc.ServerStream
}

type traceSegmentReportServiceCollectServer struct {
	grpc.ServerStream
}

func (x *traceSegmentReportServiceCollectServer) SendAndClose(m *common.Commands) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceSegmentReportServiceCollectServer) Recv() (*common.UpstreamSegment, error) {
	m := new(common.UpstreamSegment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceSegmentReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TraceSegmentReportService",
	HandlerType: (*TraceSegmentReportServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _TraceSegmentReportService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "language-agent-v2/trace.proto",
}