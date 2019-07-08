// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent/JVMMetricsService.proto

package agent

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

type JVMMetrics struct {
	Metrics               []*common.JVMMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	ApplicationInstanceId int32               `protobuf:"varint,2,opt,name=applicationInstanceId,proto3" json:"applicationInstanceId,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}            `json:"-"`
	XXX_unrecognized      []byte              `json:"-"`
	XXX_sizecache         int32               `json:"-"`
}

func (m *JVMMetrics) Reset()         { *m = JVMMetrics{} }
func (m *JVMMetrics) String() string { return proto.CompactTextString(m) }
func (*JVMMetrics) ProtoMessage()    {}
func (*JVMMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_84f947edc73ca4ba, []int{0}
}

func (m *JVMMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMMetrics.Unmarshal(m, b)
}
func (m *JVMMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMMetrics.Marshal(b, m, deterministic)
}
func (m *JVMMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMMetrics.Merge(m, src)
}
func (m *JVMMetrics) XXX_Size() int {
	return xxx_messageInfo_JVMMetrics.Size(m)
}
func (m *JVMMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_JVMMetrics proto.InternalMessageInfo

func (m *JVMMetrics) GetMetrics() []*common.JVMMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *JVMMetrics) GetApplicationInstanceId() int32 {
	if m != nil {
		return m.ApplicationInstanceId
	}
	return 0
}

func init() {
	proto.RegisterType((*JVMMetrics)(nil), "JVMMetrics")
}

func init() {
	proto.RegisterFile("language-agent/JVMMetricsService.proto", fileDescriptor_84f947edc73ca4ba)
}

var fileDescriptor_84f947edc73ca4ba = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x15, 0x2d, 0x6e, 0x2e, 0x1a, 0x10, 0x4a, 0x2e, 0x96, 0xa2, 0xd2, 0x8b, 0x5b, 0xa9,
	0xe2, 0xc1, 0xa3, 0x78, 0xa9, 0x50, 0x29, 0x2d, 0x28, 0x88, 0x97, 0xe9, 0x3a, 0xa4, 0x21, 0xbb,
	0x33, 0xcb, 0xee, 0x6a, 0xe8, 0x2f, 0xf9, 0x95, 0x92, 0xa4, 0x35, 0xa0, 0xbd, 0xed, 0xce, 0x9b,
	0xf7, 0xe6, 0xbd, 0x27, 0x2e, 0x35, 0x50, 0xf6, 0x09, 0x19, 0x5e, 0x41, 0x86, 0x14, 0x46, 0x4f,
	0x2f, 0xd3, 0x29, 0x06, 0x97, 0x2b, 0xbf, 0x40, 0xf7, 0x95, 0x2b, 0x94, 0xd6, 0x71, 0xe0, 0xf4,
	0xec, 0xcf, 0xde, 0x23, 0x97, 0xe4, 0x83, 0x43, 0x30, 0x9b, 0x85, 0x63, 0xc5, 0xc6, 0x30, 0x55,
	0x02, 0xcd, 0x64, 0xb0, 0x12, 0xa2, 0x55, 0x4b, 0xce, 0x45, 0xd7, 0x34, 0xcf, 0x5e, 0xd4, 0xdf,
	0x1f, 0xc6, 0x63, 0x21, 0x7f, 0xd1, 0xf9, 0x16, 0x4a, 0x6e, 0xc5, 0x29, 0x58, 0xab, 0x73, 0x05,
	0x21, 0x67, 0x9a, 0x90, 0x0f, 0x40, 0x0a, 0x27, 0x1f, 0xbd, 0x4e, 0x3f, 0x1a, 0x1e, 0xcc, 0x77,
	0x83, 0xe3, 0x7b, 0x71, 0xf2, 0xcf, 0x77, 0x72, 0x21, 0xba, 0x8a, 0xb5, 0x46, 0x15, 0x92, 0xb8,
	0x3d, 0xe5, 0xd3, 0x58, 0xb6, 0xde, 0x07, 0x7b, 0x0f, 0xef, 0xe2, 0x9a, 0x5d, 0x26, 0xc1, 0x82,
	0x5a, 0xa1, 0xf4, 0xc5, 0xba, 0x04, 0x5d, 0xe4, 0x54, 0x4d, 0x8c, 0x24, 0x0c, 0x25, 0xbb, 0x42,
	0x6e, 0xc3, 0xcb, 0x3a, 0xfc, 0x2c, 0x7a, 0x3b, 0xb2, 0xcb, 0xbb, 0x51, 0xfd, 0xf9, 0xee, 0xa4,
	0x8b, 0x62, 0xfd, 0xba, 0x21, 0x3d, 0x37, 0x84, 0x59, 0x55, 0x80, 0x62, 0xbd, 0x3c, 0xac, 0xab,
	0xb8, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0x51, 0x8b, 0x4b, 0x4d, 0x67, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JVMMetricsServiceClient is the client API for JVMMetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JVMMetricsServiceClient interface {
	Collect(ctx context.Context, in *JVMMetrics, opts ...grpc.CallOption) (*Downstream, error)
}

type jVMMetricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewJVMMetricsServiceClient(cc *grpc.ClientConn) JVMMetricsServiceClient {
	return &jVMMetricsServiceClient{cc}
}

func (c *jVMMetricsServiceClient) Collect(ctx context.Context, in *JVMMetrics, opts ...grpc.CallOption) (*Downstream, error) {
	out := new(Downstream)
	err := c.cc.Invoke(ctx, "/JVMMetricsService/collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JVMMetricsServiceServer is the server API for JVMMetricsService service.
type JVMMetricsServiceServer interface {
	Collect(context.Context, *JVMMetrics) (*Downstream, error)
}

// UnimplementedJVMMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJVMMetricsServiceServer struct {
}

func (*UnimplementedJVMMetricsServiceServer) Collect(ctx context.Context, req *JVMMetrics) (*Downstream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}

func RegisterJVMMetricsServiceServer(s *grpc.Server, srv JVMMetricsServiceServer) {
	s.RegisterService(&_JVMMetricsService_serviceDesc, srv)
}

func _JVMMetricsService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JVMMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JVMMetricsServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JVMMetricsService/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JVMMetricsServiceServer).Collect(ctx, req.(*JVMMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

var _JVMMetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "JVMMetricsService",
	HandlerType: (*JVMMetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "collect",
			Handler:    _JVMMetricsService_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent/JVMMetricsService.proto",
}