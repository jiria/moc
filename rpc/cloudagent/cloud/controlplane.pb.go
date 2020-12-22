// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controlplane.proto

package cloud

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ControlPlaneState int32

const (
	ControlPlaneState_NotLeader ControlPlaneState = 0
	ControlPlaneState_Leader    ControlPlaneState = 1
)

var ControlPlaneState_name = map[int32]string{
	0: "NotLeader",
	1: "Leader",
}

var ControlPlaneState_value = map[string]int32{
	"NotLeader": 0,
	"Leader":    1,
}

func (x ControlPlaneState) String() string {
	return proto.EnumName(ControlPlaneState_name, int32(x))
}

func (ControlPlaneState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e5f0814968a0baa7, []int{0}
}

type ControlPlaneRequest struct {
	ControlPlanes        []*ControlPlane  `protobuf:"bytes,1,rep,name=ControlPlanes,json=controlPlanes,proto3" json:"ControlPlanes,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,json=operationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ControlPlaneRequest) Reset()         { *m = ControlPlaneRequest{} }
func (m *ControlPlaneRequest) String() string { return proto.CompactTextString(m) }
func (*ControlPlaneRequest) ProtoMessage()    {}
func (*ControlPlaneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5f0814968a0baa7, []int{0}
}

func (m *ControlPlaneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlPlaneRequest.Unmarshal(m, b)
}
func (m *ControlPlaneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlPlaneRequest.Marshal(b, m, deterministic)
}
func (m *ControlPlaneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlPlaneRequest.Merge(m, src)
}
func (m *ControlPlaneRequest) XXX_Size() int {
	return xxx_messageInfo_ControlPlaneRequest.Size(m)
}
func (m *ControlPlaneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlPlaneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ControlPlaneRequest proto.InternalMessageInfo

func (m *ControlPlaneRequest) GetControlPlanes() []*ControlPlane {
	if m != nil {
		return m.ControlPlanes
	}
	return nil
}

func (m *ControlPlaneRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type ControlPlaneResponse struct {
	ControlPlanes        []*ControlPlane     `protobuf:"bytes,1,rep,name=ControlPlanes,json=controlPlanes,proto3" json:"ControlPlanes,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,json=result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,json=error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ControlPlaneResponse) Reset()         { *m = ControlPlaneResponse{} }
func (m *ControlPlaneResponse) String() string { return proto.CompactTextString(m) }
func (*ControlPlaneResponse) ProtoMessage()    {}
func (*ControlPlaneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5f0814968a0baa7, []int{1}
}

func (m *ControlPlaneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlPlaneResponse.Unmarshal(m, b)
}
func (m *ControlPlaneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlPlaneResponse.Marshal(b, m, deterministic)
}
func (m *ControlPlaneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlPlaneResponse.Merge(m, src)
}
func (m *ControlPlaneResponse) XXX_Size() int {
	return xxx_messageInfo_ControlPlaneResponse.Size(m)
}
func (m *ControlPlaneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlPlaneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ControlPlaneResponse proto.InternalMessageInfo

func (m *ControlPlaneResponse) GetControlPlanes() []*ControlPlane {
	if m != nil {
		return m.ControlPlanes
	}
	return nil
}

func (m *ControlPlaneResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ControlPlaneResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ControlPlane struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LocationName         string            `protobuf:"bytes,3,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Fqdn                 string            `protobuf:"bytes,4,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Port                 int32             `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	EtcdClientPort       uint32            `protobuf:"varint,6,opt,name=etcdClientPort,proto3" json:"etcdClientPort,omitempty"`
	EtcdPeerPort         uint32            `protobuf:"varint,7,opt,name=etcdPeerPort,proto3" json:"etcdPeerPort,omitempty"`
	EtcdCaCertificate    string            `protobuf:"bytes,8,opt,name=etcdCaCertificate,proto3" json:"etcdCaCertificate,omitempty"`
	EtcdCaKey            string            `protobuf:"bytes,9,opt,name=etcdCaKey,proto3" json:"etcdCaKey,omitempty"`
	Status               *common.Status    `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	State                ControlPlaneState `protobuf:"varint,11,opt,name=state,proto3,enum=moc.cloudagent.controlplane.ControlPlaneState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ControlPlane) Reset()         { *m = ControlPlane{} }
func (m *ControlPlane) String() string { return proto.CompactTextString(m) }
func (*ControlPlane) ProtoMessage()    {}
func (*ControlPlane) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5f0814968a0baa7, []int{2}
}

func (m *ControlPlane) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlPlane.Unmarshal(m, b)
}
func (m *ControlPlane) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlPlane.Marshal(b, m, deterministic)
}
func (m *ControlPlane) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlPlane.Merge(m, src)
}
func (m *ControlPlane) XXX_Size() int {
	return xxx_messageInfo_ControlPlane.Size(m)
}
func (m *ControlPlane) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlPlane.DiscardUnknown(m)
}

var xxx_messageInfo_ControlPlane proto.InternalMessageInfo

func (m *ControlPlane) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ControlPlane) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ControlPlane) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *ControlPlane) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *ControlPlane) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ControlPlane) GetEtcdClientPort() uint32 {
	if m != nil {
		return m.EtcdClientPort
	}
	return 0
}

func (m *ControlPlane) GetEtcdPeerPort() uint32 {
	if m != nil {
		return m.EtcdPeerPort
	}
	return 0
}

func (m *ControlPlane) GetEtcdCaCertificate() string {
	if m != nil {
		return m.EtcdCaCertificate
	}
	return ""
}

func (m *ControlPlane) GetEtcdCaKey() string {
	if m != nil {
		return m.EtcdCaKey
	}
	return ""
}

func (m *ControlPlane) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ControlPlane) GetState() ControlPlaneState {
	if m != nil {
		return m.State
	}
	return ControlPlaneState_NotLeader
}

func init() {
	proto.RegisterEnum("moc.cloudagent.controlplane.ControlPlaneState", ControlPlaneState_name, ControlPlaneState_value)
	proto.RegisterType((*ControlPlaneRequest)(nil), "moc.cloudagent.controlplane.ControlPlaneRequest")
	proto.RegisterType((*ControlPlaneResponse)(nil), "moc.cloudagent.controlplane.ControlPlaneResponse")
	proto.RegisterType((*ControlPlane)(nil), "moc.cloudagent.controlplane.ControlPlane")
}

func init() { proto.RegisterFile("controlplane.proto", fileDescriptor_e5f0814968a0baa7) }

var fileDescriptor_e5f0814968a0baa7 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x93, 0xc6, 0x90, 0x49, 0x1c, 0xd1, 0xa5, 0x07, 0x2b, 0x20, 0x14, 0x05, 0x09, 0x05,
	0x04, 0x6b, 0x08, 0xbc, 0x00, 0x0d, 0x1c, 0x10, 0xa8, 0x8d, 0x16, 0xc4, 0x81, 0xdb, 0x66, 0x3d,
	0x09, 0x16, 0xf6, 0x8e, 0xbb, 0x5e, 0x83, 0x7a, 0x44, 0xe2, 0x31, 0x78, 0x09, 0xde, 0x10, 0xed,
	0x6e, 0xa1, 0x09, 0x95, 0x50, 0x91, 0xb8, 0x7d, 0xfe, 0x7e, 0x66, 0x26, 0xb3, 0x13, 0x60, 0x8a,
	0xb4, 0x35, 0x54, 0xd6, 0xa5, 0xd4, 0xc8, 0x6b, 0x43, 0x96, 0xd8, 0xad, 0x8a, 0x14, 0x57, 0x25,
	0xb5, 0xb9, 0xdc, 0xa0, 0xb6, 0x7c, 0xdb, 0x32, 0xbe, 0xb3, 0x21, 0xda, 0x94, 0x98, 0x79, 0xeb,
	0xaa, 0x5d, 0x67, 0x5f, 0x8c, 0xac, 0x6b, 0x34, 0x4d, 0x08, 0x8f, 0x87, 0x8a, 0xaa, 0x8a, 0x74,
	0xf8, 0x9a, 0x7e, 0x8f, 0xe0, 0xe6, 0x22, 0xc4, 0x97, 0x2e, 0x2e, 0xf0, 0xb4, 0xc5, 0xc6, 0xb2,
	0x13, 0x48, 0xb6, 0xe9, 0x26, 0x8d, 0x26, 0xdd, 0xd9, 0x60, 0x7e, 0x9f, 0xff, 0xa5, 0x35, 0xdf,
	0x29, 0x94, 0xa8, 0xed, 0x3c, 0x7b, 0x06, 0xc9, 0x49, 0x8d, 0x46, 0xda, 0x82, 0xf4, 0xbb, 0xb3,
	0x1a, 0xd3, 0xce, 0x24, 0x9a, 0x8d, 0xe6, 0x23, 0x5f, 0xf0, 0xb7, 0x22, 0x12, 0xda, 0x36, 0x4d,
	0x7f, 0x44, 0x70, 0xb8, 0x3b, 0x5e, 0x53, 0x93, 0x6e, 0xf0, 0xff, 0xcf, 0x37, 0x87, 0x58, 0x60,
	0xd3, 0x96, 0xd6, 0x0f, 0x36, 0x98, 0x8f, 0x79, 0xd8, 0x23, 0xff, 0xb5, 0x47, 0x7e, 0x44, 0x54,
	0xbe, 0x97, 0x65, 0x8b, 0x22, 0x36, 0xde, 0xc9, 0x0e, 0xa1, 0xf7, 0xd2, 0x18, 0x32, 0x69, 0x77,
	0x12, 0xcd, 0xfa, 0xa2, 0x87, 0xee, 0x63, 0xfa, 0xb5, 0x0b, 0xc3, 0xed, 0x4e, 0x6c, 0x04, 0x9d,
	0x22, 0x4f, 0x23, 0xef, 0xe9, 0x14, 0x39, 0x63, 0xb0, 0xaf, 0x65, 0x15, 0x36, 0xd0, 0x17, 0x1e,
	0xb3, 0x29, 0x0c, 0x4b, 0x52, 0xfe, 0x87, 0x1f, 0x3b, 0x2d, 0x54, 0xdc, 0xe1, 0x5c, 0x6e, 0x7d,
	0x9a, 0xeb, 0x74, 0x3f, 0xe4, 0x1c, 0x76, 0x5c, 0x4d, 0xc6, 0xa6, 0xbd, 0x49, 0x34, 0xeb, 0x09,
	0x8f, 0xd9, 0x3d, 0x18, 0xa1, 0x55, 0xf9, 0xa2, 0x2c, 0x50, 0xdb, 0xa5, 0x53, 0xe3, 0x49, 0x34,
	0x4b, 0xc4, 0x1f, 0xac, 0xeb, 0xe9, 0x98, 0x25, 0xa2, 0xf1, 0xae, 0x6b, 0xde, 0xb5, 0xc3, 0xb1,
	0x87, 0x70, 0xe0, 0x53, 0x72, 0x81, 0xc6, 0x16, 0xeb, 0x42, 0x49, 0x8b, 0xe9, 0x75, 0x3f, 0xc0,
	0x65, 0x81, 0xdd, 0x86, 0x7e, 0x20, 0x5f, 0xe3, 0x59, 0xda, 0xf7, 0xae, 0x0b, 0x82, 0xdd, 0x85,
	0xb8, 0xb1, 0xd2, 0xb6, 0x4d, 0x0a, 0x7e, 0xc5, 0x03, 0xff, 0x58, 0x6f, 0x3d, 0x25, 0xce, 0x25,
	0xf6, 0x02, 0x7a, 0x0e, 0x61, 0x3a, 0xf0, 0xf7, 0xc1, 0xaf, 0xfc, 0xa0, 0xae, 0x0e, 0x8a, 0x10,
	0x7e, 0xc0, 0xe1, 0xe0, 0x92, 0xc6, 0x12, 0xe8, 0x1f, 0x93, 0x7d, 0x83, 0x32, 0x47, 0x73, 0x63,
	0x8f, 0x01, 0xc4, 0xe7, 0x38, 0x9a, 0x7f, 0x8b, 0x76, 0x03, 0xcf, 0x5d, 0x2f, 0x46, 0x10, 0xbf,
	0xd2, 0x9f, 0xe9, 0x13, 0xb2, 0xc7, 0x57, 0xbf, 0xab, 0xf0, 0x07, 0x1a, 0x3f, 0xf9, 0x87, 0x44,
	0xb8, 0xe9, 0xe9, 0xde, 0x51, 0xf6, 0xe1, 0xd1, 0xa6, 0xb0, 0x1f, 0xdb, 0x15, 0x57, 0x54, 0x65,
	0x55, 0xa1, 0x0c, 0x35, 0xb4, 0xb6, 0x59, 0x45, 0x2a, 0x33, 0xb5, 0xca, 0x2e, 0xca, 0x05, 0xb8,
	0x8a, 0xfd, 0x75, 0x3e, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0x18, 0x27, 0x06, 0x19, 0x26, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControlPlaneAgentClient is the client API for ControlPlaneAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlPlaneAgentClient interface {
	Invoke(ctx context.Context, in *ControlPlaneRequest, opts ...grpc.CallOption) (*ControlPlaneResponse, error)
}

type controlPlaneAgentClient struct {
	cc *grpc.ClientConn
}

func NewControlPlaneAgentClient(cc *grpc.ClientConn) ControlPlaneAgentClient {
	return &controlPlaneAgentClient{cc}
}

func (c *controlPlaneAgentClient) Invoke(ctx context.Context, in *ControlPlaneRequest, opts ...grpc.CallOption) (*ControlPlaneResponse, error) {
	out := new(ControlPlaneResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.controlplane.ControlPlaneAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlPlaneAgentServer is the server API for ControlPlaneAgent service.
type ControlPlaneAgentServer interface {
	Invoke(context.Context, *ControlPlaneRequest) (*ControlPlaneResponse, error)
}

// UnimplementedControlPlaneAgentServer can be embedded to have forward compatible implementations.
type UnimplementedControlPlaneAgentServer struct {
}

func (*UnimplementedControlPlaneAgentServer) Invoke(ctx context.Context, req *ControlPlaneRequest) (*ControlPlaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterControlPlaneAgentServer(s *grpc.Server, srv ControlPlaneAgentServer) {
	s.RegisterService(&_ControlPlaneAgent_serviceDesc, srv)
}

func _ControlPlaneAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlPlaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.controlplane.ControlPlaneAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneAgentServer).Invoke(ctx, req.(*ControlPlaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlPlaneAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.controlplane.ControlPlaneAgent",
	HandlerType: (*ControlPlaneAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _ControlPlaneAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane.proto",
}
