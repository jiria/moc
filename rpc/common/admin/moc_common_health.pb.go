// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin/health/moc_common_health.proto

package admin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HealthRequest struct {
	TimeoutSeconds       uint32   `protobuf:"varint,1,opt,name=timeoutSeconds,proto3" json:"timeoutSeconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d110a1beda4379f, []int{0}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

func (m *HealthRequest) GetTimeoutSeconds() uint32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

type HealthResponse struct {
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	State                common.HealthState  `protobuf:"varint,4,opt,name=State,proto3,enum=moc.HealthState" json:"State,omitempty"`
	Rebooted             bool                `protobuf:"varint,5,opt,name=Rebooted,proto3" json:"Rebooted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d110a1beda4379f, []int{1}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *HealthResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *HealthResponse) GetState() common.HealthState {
	if m != nil {
		return m.State
	}
	return common.HealthState_NOTKNOWN
}

func (m *HealthResponse) GetRebooted() bool {
	if m != nil {
		return m.Rebooted
	}
	return false
}

type AgentInfoResponse struct {
	Node                 *common.NodeInfo    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AgentInfoResponse) Reset()         { *m = AgentInfoResponse{} }
func (m *AgentInfoResponse) String() string { return proto.CompactTextString(m) }
func (*AgentInfoResponse) ProtoMessage()    {}
func (*AgentInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d110a1beda4379f, []int{2}
}

func (m *AgentInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentInfoResponse.Unmarshal(m, b)
}
func (m *AgentInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentInfoResponse.Marshal(b, m, deterministic)
}
func (m *AgentInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentInfoResponse.Merge(m, src)
}
func (m *AgentInfoResponse) XXX_Size() int {
	return xxx_messageInfo_AgentInfoResponse.Size(m)
}
func (m *AgentInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AgentInfoResponse proto.InternalMessageInfo

func (m *AgentInfoResponse) GetNode() *common.NodeInfo {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *AgentInfoResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *AgentInfoResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthRequest)(nil), "moc.common.admin.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "moc.common.admin.HealthResponse")
	proto.RegisterType((*AgentInfoResponse)(nil), "moc.common.admin.AgentInfoResponse")
}

func init() {
	proto.RegisterFile("admin/health/moc_common_health.proto", fileDescriptor_4d110a1beda4379f)
}

var fileDescriptor_4d110a1beda4379f = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xc7, 0x6b, 0xe8, 0x56, 0xc5, 0xa1, 0xab, 0x62, 0x55, 0x10, 0x82, 0x04, 0x51, 0x40, 0x55,
	0x10, 0x92, 0x2d, 0x85, 0x03, 0x67, 0x8a, 0x2a, 0xe0, 0x00, 0x07, 0x57, 0xe2, 0xc0, 0x65, 0x95,
	0x8f, 0xc9, 0x87, 0x88, 0x3d, 0xc1, 0x71, 0x84, 0x78, 0x02, 0xde, 0x82, 0x67, 0xe0, 0x11, 0x51,
	0xec, 0xdd, 0xd5, 0xb2, 0x2b, 0x6e, 0x3d, 0x7a, 0x7e, 0x7f, 0xcf, 0xfc, 0xe7, 0x83, 0xbe, 0xc8,
	0x2b, 0xd5, 0x69, 0xd1, 0x42, 0xde, 0xdb, 0x56, 0x28, 0x2c, 0x57, 0x25, 0x2a, 0x85, 0x7a, 0xe5,
	0x23, 0x7c, 0x30, 0x68, 0x91, 0x9d, 0x2b, 0x2c, 0xb9, 0x07, 0xdc, 0x7d, 0x88, 0x9e, 0x34, 0x88,
	0x4d, 0x0f, 0xc2, 0xf1, 0x62, 0xaa, 0x05, 0xa8, 0xc1, 0xfe, 0xf4, 0xf2, 0xe8, 0xe9, 0x3e, 0xfc,
	0x61, 0xf2, 0x61, 0x00, 0x33, 0xae, 0xf9, 0xa3, 0x9d, 0x3a, 0xeb, 0xac, 0x1e, 0x3c, 0xde, 0x01,
	0x1a, 0x2b, 0xe8, 0x74, 0x8d, 0x1e, 0x25, 0x6f, 0xe8, 0xd9, 0x07, 0x67, 0x49, 0xc2, 0xf7, 0x09,
	0x46, 0xcb, 0x2e, 0xe9, 0xd2, 0x76, 0x0a, 0x70, 0xb2, 0x37, 0x50, 0xa2, 0xae, 0xc6, 0x90, 0xc4,
	0x24, 0x3d, 0x93, 0x7b, 0xd1, 0xe4, 0x37, 0xa1, 0xcb, 0xcd, 0xcf, 0x71, 0x40, 0x3d, 0x02, 0xcb,
	0xe8, 0x89, 0x84, 0x71, 0xea, 0x6d, 0x78, 0x27, 0x26, 0x69, 0x90, 0x45, 0xdc, 0x1b, 0xe6, 0x1b,
	0xc3, 0xfc, 0x0a, 0xb1, 0xff, 0x92, 0xf7, 0x13, 0xc8, 0xb5, 0x92, 0x5d, 0xd0, 0xc5, 0xb5, 0x31,
	0x68, 0xc2, 0xbb, 0x31, 0x49, 0xef, 0x49, 0xff, 0x60, 0x97, 0x74, 0x71, 0x63, 0x73, 0x0b, 0xe1,
	0x71, 0x4c, 0xd2, 0x65, 0x76, 0xce, 0xe7, 0x41, 0xf9, 0x6a, 0x2e, 0x2e, 0x3d, 0x66, 0x11, 0x3d,
	0x95, 0x50, 0x20, 0x5a, 0xa8, 0xc2, 0x45, 0x4c, 0xd2, 0x53, 0xb9, 0x7d, 0x27, 0xbf, 0x08, 0x7d,
	0xf0, 0xb6, 0x01, 0x6d, 0x3f, 0xea, 0x1a, 0xb7, 0x1e, 0x53, 0x7a, 0x3c, 0x4f, 0xc0, 0x35, 0x15,
	0x64, 0x17, 0x7c, 0x67, 0x03, 0x9f, 0xb1, 0x02, 0xa7, 0x75, 0x8a, 0xdb, 0xeb, 0x26, 0xfb, 0x43,
	0x68, 0xe0, 0xcd, 0x3b, 0x3f, 0x4c, 0xd2, 0xe0, 0x5d, 0x0b, 0xe5, 0x37, 0x1f, 0x63, 0xcf, 0xf8,
	0xfe, 0x19, 0xf0, 0x7f, 0x56, 0x12, 0xc5, 0xff, 0x17, 0xf8, 0xae, 0x92, 0x23, 0xf6, 0x89, 0xde,
	0x7f, 0x0f, 0x76, 0xdb, 0x2f, 0x7b, 0x78, 0xe0, 0xf6, 0x7a, 0xbe, 0xa4, 0xe8, 0xf9, 0x61, 0xae,
	0x83, 0x21, 0x25, 0x47, 0x57, 0xaf, 0xbe, 0xbe, 0x6c, 0x3a, 0xdb, 0x4e, 0xc5, 0x2c, 0x15, 0xaa,
	0x2b, 0x0d, 0x8e, 0x58, 0xdb, 0xf9, 0x92, 0x85, 0x19, 0x4a, 0xe1, 0x13, 0x08, 0x97, 0xa0, 0x38,
	0x71, 0x35, 0x5e, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x04, 0x4e, 0x17, 0xf5, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthAgentClient is the client API for HealthAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthAgentClient interface {
	CheckHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	GetAgentInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AgentInfoResponse, error)
}

type healthAgentClient struct {
	cc *grpc.ClientConn
}

func NewHealthAgentClient(cc *grpc.ClientConn) HealthAgentClient {
	return &healthAgentClient{cc}
}

func (c *healthAgentClient) CheckHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/moc.common.admin.HealthAgent/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAgentClient) GetAgentInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AgentInfoResponse, error) {
	out := new(AgentInfoResponse)
	err := c.cc.Invoke(ctx, "/moc.common.admin.HealthAgent/GetAgentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthAgentServer is the server API for HealthAgent service.
type HealthAgentServer interface {
	CheckHealth(context.Context, *HealthRequest) (*HealthResponse, error)
	GetAgentInfo(context.Context, *empty.Empty) (*AgentInfoResponse, error)
}

// UnimplementedHealthAgentServer can be embedded to have forward compatible implementations.
type UnimplementedHealthAgentServer struct {
}

func (*UnimplementedHealthAgentServer) CheckHealth(ctx context.Context, req *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (*UnimplementedHealthAgentServer) GetAgentInfo(ctx context.Context, req *empty.Empty) (*AgentInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentInfo not implemented")
}

func RegisterHealthAgentServer(s *grpc.Server, srv HealthAgentServer) {
	s.RegisterService(&_HealthAgent_serviceDesc, srv)
}

func _HealthAgent_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAgentServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.common.admin.HealthAgent/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAgentServer).CheckHealth(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAgent_GetAgentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAgentServer).GetAgentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.common.admin.HealthAgent/GetAgentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAgentServer).GetAgentInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.common.admin.HealthAgent",
	HandlerType: (*HealthAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHealth",
			Handler:    _HealthAgent_CheckHealth_Handler,
		},
		{
			MethodName: "GetAgentInfo",
			Handler:    _HealthAgent_GetAgentInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/health/moc_common_health.proto",
}