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

type ControlPlaneRequest struct {
	ControlPlanes        []*ControlPlane  `protobuf:"bytes,1,rep,name=ControlPlanes,proto3" json:"ControlPlanes,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
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

type ControlPlanePeerRequest struct {
	ControlPlanes        []*ControlPlane `protobuf:"bytes,1,rep,name=ControlPlanes,proto3" json:"ControlPlanes,omitempty"`
	Peer                 *Peer           `protobuf:"bytes,2,opt,name=Peer,proto3" json:"Peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ControlPlanePeerRequest) Reset()         { *m = ControlPlanePeerRequest{} }
func (m *ControlPlanePeerRequest) String() string { return proto.CompactTextString(m) }
func (*ControlPlanePeerRequest) ProtoMessage()    {}
func (*ControlPlanePeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5f0814968a0baa7, []int{1}
}

func (m *ControlPlanePeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlPlanePeerRequest.Unmarshal(m, b)
}
func (m *ControlPlanePeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlPlanePeerRequest.Marshal(b, m, deterministic)
}
func (m *ControlPlanePeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlPlanePeerRequest.Merge(m, src)
}
func (m *ControlPlanePeerRequest) XXX_Size() int {
	return xxx_messageInfo_ControlPlanePeerRequest.Size(m)
}
func (m *ControlPlanePeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlPlanePeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ControlPlanePeerRequest proto.InternalMessageInfo

func (m *ControlPlanePeerRequest) GetControlPlanes() []*ControlPlane {
	if m != nil {
		return m.ControlPlanes
	}
	return nil
}

func (m *ControlPlanePeerRequest) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type ControlPlaneResponse struct {
	ControlPlanes        []*ControlPlane     `protobuf:"bytes,1,rep,name=ControlPlanes,proto3" json:"ControlPlanes,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ControlPlaneResponse) Reset()         { *m = ControlPlaneResponse{} }
func (m *ControlPlaneResponse) String() string { return proto.CompactTextString(m) }
func (*ControlPlaneResponse) ProtoMessage()    {}
func (*ControlPlaneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5f0814968a0baa7, []int{2}
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
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LocationName         string         `protobuf:"bytes,2,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Peers                []*Peer        `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
	LeaderIp             string         `protobuf:"bytes,4,opt,name=leaderIp,proto3" json:"leaderIp,omitempty"`
	LeaderPort           string         `protobuf:"bytes,5,opt,name=leaderPort,proto3" json:"leaderPort,omitempty"`
	LeaderFqdn           string         `protobuf:"bytes,6,opt,name=leaderFqdn,proto3" json:"leaderFqdn,omitempty"`
	Status               *common.Status `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ControlPlane) Reset()         { *m = ControlPlane{} }
func (m *ControlPlane) String() string { return proto.CompactTextString(m) }
func (*ControlPlane) ProtoMessage()    {}
func (*ControlPlane) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5f0814968a0baa7, []int{3}
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

func (m *ControlPlane) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *ControlPlane) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *ControlPlane) GetLeaderIp() string {
	if m != nil {
		return m.LeaderIp
	}
	return ""
}

func (m *ControlPlane) GetLeaderPort() string {
	if m != nil {
		return m.LeaderPort
	}
	return ""
}

func (m *ControlPlane) GetLeaderFqdn() string {
	if m != nil {
		return m.LeaderFqdn
	}
	return ""
}

func (m *ControlPlane) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type Peer struct {
	Fqdn                 string   `protobuf:"bytes,1,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Port                 string   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5f0814968a0baa7, []int{4}
}

func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Peer) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func init() {
	proto.RegisterType((*ControlPlaneRequest)(nil), "moc.cloudagent.controlplane.ControlPlaneRequest")
	proto.RegisterType((*ControlPlanePeerRequest)(nil), "moc.cloudagent.controlplane.ControlPlanePeerRequest")
	proto.RegisterType((*ControlPlaneResponse)(nil), "moc.cloudagent.controlplane.ControlPlaneResponse")
	proto.RegisterType((*ControlPlane)(nil), "moc.cloudagent.controlplane.ControlPlane")
	proto.RegisterType((*Peer)(nil), "moc.cloudagent.controlplane.Peer")
}

func init() { proto.RegisterFile("controlplane.proto", fileDescriptor_e5f0814968a0baa7) }

var fileDescriptor_e5f0814968a0baa7 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x69, 0xe2, 0xd2, 0x49, 0x1a, 0x89, 0xa5, 0x12, 0x96, 0x91, 0xaa, 0x60, 0x2e, 0xe1,
	0xc0, 0x1a, 0x4c, 0x11, 0xe7, 0x16, 0x81, 0xd4, 0x0b, 0x8d, 0x0c, 0xe2, 0xc0, 0xcd, 0xb1, 0x27,
	0xc1, 0x62, 0xed, 0xd9, 0xee, 0xae, 0x5b, 0xf1, 0x2e, 0x5c, 0x78, 0x04, 0x9e, 0x8d, 0x2b, 0x07,
	0xe4, 0xdd, 0xd0, 0xda, 0x97, 0x2a, 0x20, 0xe0, 0x36, 0x3f, 0xdf, 0x7c, 0xf3, 0xcd, 0xcc, 0x2e,
	0xb0, 0x9c, 0x6a, 0xa3, 0x48, 0x48, 0x91, 0xd5, 0xc8, 0xa5, 0x22, 0x43, 0xec, 0x7e, 0x45, 0x39,
	0xcf, 0x05, 0x35, 0x45, 0xb6, 0xc6, 0xda, 0xf0, 0x2e, 0x24, 0x3c, 0x5c, 0x13, 0xad, 0x05, 0xc6,
	0x16, 0xba, 0x6c, 0x56, 0xf1, 0xa5, 0xca, 0xa4, 0x44, 0xa5, 0x5d, 0x71, 0x38, 0xc9, 0xa9, 0xaa,
	0xa8, 0x76, 0x5e, 0xf4, 0xc5, 0x83, 0xbb, 0x2f, 0x5d, 0xf9, 0xa2, 0x2d, 0x4f, 0xf1, 0xbc, 0x41,
	0x6d, 0xd8, 0x19, 0xec, 0x77, 0xc3, 0x3a, 0xf0, 0x66, 0x3b, 0xf3, 0x71, 0xf2, 0x88, 0xdf, 0xd0,
	0x9a, 0xf7, 0x88, 0xfa, 0xf5, 0xec, 0x08, 0xf6, 0xcf, 0x24, 0xaa, 0xcc, 0x94, 0x54, 0xbf, 0xfb,
	0x2c, 0x31, 0x18, 0xcc, 0xbc, 0xf9, 0x34, 0x99, 0x5a, 0xc2, 0xab, 0x4c, 0xda, 0x07, 0x45, 0x5f,
	0x3d, 0xb8, 0xd7, 0xe5, 0x59, 0x20, 0xaa, 0x7f, 0x26, 0xf1, 0x39, 0x0c, 0x5b, 0x7e, 0xab, 0x6c,
	0x9c, 0x3c, 0xb8, 0x91, 0xc7, 0x0a, 0xb1, 0xf0, 0xe8, 0x9b, 0x07, 0x07, 0xfd, 0x15, 0x6a, 0x49,
	0xb5, 0xc6, 0xbf, 0x2f, 0x30, 0x01, 0x3f, 0x45, 0xdd, 0x08, 0xb3, 0x91, 0x18, 0x72, 0x77, 0x6b,
	0xfe, 0xeb, 0xd6, 0xfc, 0x84, 0x48, 0xbc, 0xcf, 0x44, 0x83, 0xe9, 0x06, 0xc9, 0x0e, 0x60, 0xf4,
	0x4a, 0x29, 0x52, 0xc1, 0xce, 0xcc, 0x9b, 0xef, 0xa5, 0xce, 0x89, 0x7e, 0x78, 0x30, 0xe9, 0x72,
	0xb3, 0x29, 0x0c, 0xca, 0x22, 0xf0, 0x2c, 0x66, 0x50, 0x16, 0x2c, 0x82, 0x89, 0xa0, 0xdc, 0x1e,
	0xe2, 0x4d, 0x56, 0xb9, 0x6b, 0xed, 0xa5, 0xbd, 0x18, 0x7b, 0x01, 0x23, 0x89, 0xa8, 0x74, 0xb0,
	0x63, 0xe7, 0xda, 0x62, 0x61, 0x0e, 0xcf, 0x42, 0xb8, 0x2d, 0x30, 0x2b, 0x50, 0x9d, 0xca, 0x60,
	0x68, 0x89, 0xaf, 0x7c, 0x76, 0x08, 0xe0, 0xec, 0x05, 0x29, 0x13, 0x8c, 0x6c, 0xb6, 0x13, 0xb9,
	0xce, 0xbf, 0x3e, 0x2f, 0xea, 0xc0, 0xef, 0xe6, 0xdb, 0x08, 0x7b, 0x08, 0xbe, 0x36, 0x99, 0x69,
	0x74, 0xb0, 0x6b, 0x77, 0x34, 0xb6, 0xaa, 0xde, 0xda, 0x50, 0xba, 0x49, 0x45, 0xdc, 0x5d, 0x9a,
	0x31, 0x18, 0xae, 0x5a, 0x1a, 0x37, 0xb7, 0xb5, 0xdb, 0x98, 0x6c, 0x5b, 0xbb, 0x89, 0xad, 0x9d,
	0x7c, 0x1f, 0xc0, 0x9d, 0xee, 0xba, 0x8e, 0xdb, 0xf9, 0x18, 0x81, 0x7f, 0x5a, 0x5f, 0xd0, 0x27,
	0x64, 0x4f, 0xb6, 0x3f, 0xa9, 0x7b, 0xbc, 0xe1, 0xd3, 0xdf, 0xa8, 0x70, 0xcf, 0x29, 0xba, 0xc5,
	0x0c, 0xec, 0x1e, 0x17, 0x85, 0x55, 0x7e, 0xb4, 0x75, 0x7d, 0xe7, 0xcb, 0xfc, 0x59, 0xd7, 0x4b,
	0x80, 0x14, 0x2b, 0xba, 0xc0, 0xff, 0xdc, 0xf8, 0x24, 0xfe, 0xf0, 0x78, 0x5d, 0x9a, 0x8f, 0xcd,
	0x92, 0xe7, 0x54, 0xc5, 0x55, 0x99, 0x2b, 0xd2, 0xb4, 0x32, 0x71, 0x45, 0x79, 0xac, 0x64, 0x1e,
	0x5f, 0xd3, 0x39, 0x73, 0xe9, 0xdb, 0x7f, 0xf0, 0xec, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24,
	0x23, 0x8e, 0xd7, 0x34, 0x05, 0x00, 0x00,
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
	AddPeer(ctx context.Context, in *ControlPlanePeerRequest, opts ...grpc.CallOption) (*ControlPlaneResponse, error)
	RemovePeer(ctx context.Context, in *ControlPlanePeerRequest, opts ...grpc.CallOption) (*ControlPlaneResponse, error)
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

func (c *controlPlaneAgentClient) AddPeer(ctx context.Context, in *ControlPlanePeerRequest, opts ...grpc.CallOption) (*ControlPlaneResponse, error) {
	out := new(ControlPlaneResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.controlplane.ControlPlaneAgent/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneAgentClient) RemovePeer(ctx context.Context, in *ControlPlanePeerRequest, opts ...grpc.CallOption) (*ControlPlaneResponse, error) {
	out := new(ControlPlaneResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.controlplane.ControlPlaneAgent/RemovePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlPlaneAgentServer is the server API for ControlPlaneAgent service.
type ControlPlaneAgentServer interface {
	Invoke(context.Context, *ControlPlaneRequest) (*ControlPlaneResponse, error)
	AddPeer(context.Context, *ControlPlanePeerRequest) (*ControlPlaneResponse, error)
	RemovePeer(context.Context, *ControlPlanePeerRequest) (*ControlPlaneResponse, error)
}

// UnimplementedControlPlaneAgentServer can be embedded to have forward compatible implementations.
type UnimplementedControlPlaneAgentServer struct {
}

func (*UnimplementedControlPlaneAgentServer) Invoke(ctx context.Context, req *ControlPlaneRequest) (*ControlPlaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedControlPlaneAgentServer) AddPeer(ctx context.Context, req *ControlPlanePeerRequest) (*ControlPlaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (*UnimplementedControlPlaneAgentServer) RemovePeer(ctx context.Context, req *ControlPlanePeerRequest) (*ControlPlaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePeer not implemented")
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

func _ControlPlaneAgent_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlPlanePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneAgentServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.controlplane.ControlPlaneAgent/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneAgentServer).AddPeer(ctx, req.(*ControlPlanePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneAgent_RemovePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlPlanePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneAgentServer).RemovePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.controlplane.ControlPlaneAgent/RemovePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneAgentServer).RemovePeer(ctx, req.(*ControlPlanePeerRequest))
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
		{
			MethodName: "AddPeer",
			Handler:    _ControlPlaneAgent_AddPeer_Handler,
		},
		{
			MethodName: "RemovePeer",
			Handler:    _ControlPlaneAgent_RemovePeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane.proto",
}
