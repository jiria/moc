// Code generated by protoc-gen-go. DO NOT EDIT.
// source: loadbalancer.proto

package network

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

type LoadBalancerRequest struct {
	LoadBalancers        []*LoadBalancer  `protobuf:"bytes,1,rep,name=LoadBalancers,proto3" json:"LoadBalancers,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LoadBalancerRequest) Reset()         { *m = LoadBalancerRequest{} }
func (m *LoadBalancerRequest) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerRequest) ProtoMessage()    {}
func (*LoadBalancerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60ee00bafda1595, []int{0}
}

func (m *LoadBalancerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerRequest.Unmarshal(m, b)
}
func (m *LoadBalancerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerRequest.Marshal(b, m, deterministic)
}
func (m *LoadBalancerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerRequest.Merge(m, src)
}
func (m *LoadBalancerRequest) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerRequest.Size(m)
}
func (m *LoadBalancerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerRequest proto.InternalMessageInfo

func (m *LoadBalancerRequest) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

func (m *LoadBalancerRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type LoadBalancerResponse struct {
	LoadBalancers        []*LoadBalancer     `protobuf:"bytes,1,rep,name=LoadBalancers,proto3" json:"LoadBalancers,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LoadBalancerResponse) Reset()         { *m = LoadBalancerResponse{} }
func (m *LoadBalancerResponse) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerResponse) ProtoMessage()    {}
func (*LoadBalancerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60ee00bafda1595, []int{1}
}

func (m *LoadBalancerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerResponse.Unmarshal(m, b)
}
func (m *LoadBalancerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerResponse.Marshal(b, m, deterministic)
}
func (m *LoadBalancerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerResponse.Merge(m, src)
}
func (m *LoadBalancerResponse) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerResponse.Size(m)
}
func (m *LoadBalancerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerResponse proto.InternalMessageInfo

func (m *LoadBalancerResponse) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

func (m *LoadBalancerResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *LoadBalancerResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LoadBalancer struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Frontendip           string         `protobuf:"bytes,4,opt,name=frontendip,proto3" json:"frontendip,omitempty"`
	Backendpoolname      string         `protobuf:"bytes,5,opt,name=backendpoolname,proto3" json:"backendpoolname,omitempty"`
	Networkid            string         `protobuf:"bytes,6,opt,name=networkid,proto3" json:"networkid,omitempty"`
	Status               *common.Status `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity `protobuf:"bytes,9,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags   `protobuf:"bytes,10,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LoadBalancer) Reset()         { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()    {}
func (*LoadBalancer) Descriptor() ([]byte, []int) {
	return fileDescriptor_d60ee00bafda1595, []int{2}
}

func (m *LoadBalancer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancer.Unmarshal(m, b)
}
func (m *LoadBalancer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancer.Marshal(b, m, deterministic)
}
func (m *LoadBalancer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancer.Merge(m, src)
}
func (m *LoadBalancer) XXX_Size() int {
	return xxx_messageInfo_LoadBalancer.Size(m)
}
func (m *LoadBalancer) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancer.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancer proto.InternalMessageInfo

func (m *LoadBalancer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoadBalancer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LoadBalancer) GetFrontendip() string {
	if m != nil {
		return m.Frontendip
	}
	return ""
}

func (m *LoadBalancer) GetBackendpoolname() string {
	if m != nil {
		return m.Backendpoolname
	}
	return ""
}

func (m *LoadBalancer) GetNetworkid() string {
	if m != nil {
		return m.Networkid
	}
	return ""
}

func (m *LoadBalancer) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LoadBalancer) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *LoadBalancer) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadBalancerRequest)(nil), "moc.nodeagent.network.LoadBalancerRequest")
	proto.RegisterType((*LoadBalancerResponse)(nil), "moc.nodeagent.network.LoadBalancerResponse")
	proto.RegisterType((*LoadBalancer)(nil), "moc.nodeagent.network.LoadBalancer")
}

func init() { proto.RegisterFile("loadbalancer.proto", fileDescriptor_d60ee00bafda1595) }

var fileDescriptor_d60ee00bafda1595 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0xc9, 0x76, 0x1b, 0x91, 0xd9, 0x76, 0x11, 0xa6, 0x48, 0xd1, 0x0a, 0xaa, 0xd5, 0xf6,
	0xb2, 0x02, 0xc9, 0x41, 0x81, 0x17, 0x60, 0xa5, 0x1e, 0x2a, 0x21, 0x21, 0x85, 0x8a, 0x03, 0x37,
	0x27, 0x99, 0x0d, 0xd1, 0x26, 0x1e, 0x63, 0x3b, 0x94, 0x9e, 0x79, 0x07, 0x5e, 0x82, 0x97, 0x44,
	0x9d, 0x84, 0x92, 0x45, 0x1c, 0xf6, 0xc0, 0xcd, 0xfe, 0xe7, 0xf3, 0xef, 0xf1, 0x3f, 0x06, 0xd1,
	0x90, 0x2a, 0x73, 0xd5, 0x28, 0x5d, 0xa0, 0x95, 0xc6, 0x92, 0x27, 0xf1, 0xb4, 0xa5, 0x42, 0x6a,
	0x2a, 0x51, 0x55, 0xa8, 0xbd, 0xd4, 0xe8, 0x6f, 0xc8, 0xee, 0x16, 0xe7, 0x15, 0x51, 0xd5, 0x60,
	0xc2, 0x50, 0xde, 0x6d, 0x93, 0x1b, 0xab, 0x8c, 0x41, 0xeb, 0xfa, 0x63, 0x8b, 0x93, 0x82, 0xda,
	0x96, 0x74, 0xbf, 0x5b, 0xfd, 0x08, 0xe0, 0xc9, 0x3b, 0x52, 0xe5, 0x66, 0xf0, 0xce, 0xf0, 0x4b,
	0x87, 0xce, 0x8b, 0x2b, 0x38, 0x1d, 0xcb, 0x2e, 0x0e, 0x96, 0x47, 0xeb, 0x59, 0x7a, 0x21, 0xff,
	0x79, 0xa9, 0xdc, 0xb3, 0xd8, 0x3f, 0x29, 0xde, 0xc0, 0xe9, 0x7b, 0x83, 0x56, 0xf9, 0x9a, 0xf4,
	0xf5, 0xad, 0xc1, 0x78, 0xb2, 0x0c, 0xd6, 0xf3, 0x74, 0xce, 0x56, 0xf7, 0x95, 0x6c, 0x1f, 0x5a,
	0xfd, 0x0c, 0xe0, 0x6c, 0xbf, 0x31, 0x67, 0x48, 0x3b, 0xfc, 0x9f, 0x9d, 0xa5, 0x10, 0x66, 0xe8,
	0xba, 0xc6, 0x73, 0x4b, 0xb3, 0x74, 0x21, 0xfb, 0xec, 0xe4, 0xef, 0xec, 0xe4, 0x86, 0xa8, 0xf9,
	0xa8, 0x9a, 0x0e, 0xb3, 0x81, 0x14, 0x67, 0x70, 0x7c, 0x69, 0x2d, 0xd9, 0xf8, 0x68, 0x19, 0xac,
	0xa3, 0xac, 0xdf, 0xac, 0xbe, 0x4f, 0xe0, 0x64, 0xec, 0x2d, 0x04, 0x4c, 0xb5, 0x6a, 0x31, 0x0e,
	0x98, 0xe2, 0xb5, 0x98, 0xc3, 0xa4, 0x2e, 0xf9, 0xaa, 0x28, 0x9b, 0xd4, 0xa5, 0x38, 0x07, 0xd8,
	0x5a, 0xd2, 0x1e, 0x75, 0x59, 0x9b, 0x78, 0xca, 0xfa, 0x48, 0x11, 0x6b, 0x78, 0x94, 0xab, 0x62,
	0x87, 0xba, 0x34, 0x44, 0x0d, 0xdb, 0x1d, 0x33, 0xf4, 0xb7, 0x2c, 0x9e, 0x41, 0x34, 0xbc, 0xb7,
	0x2e, 0xe3, 0x90, 0x99, 0x3f, 0x82, 0xb8, 0x80, 0xd0, 0x79, 0xe5, 0x3b, 0x17, 0x3f, 0xe4, 0x67,
	0xce, 0x38, 0xaa, 0x0f, 0x2c, 0x65, 0x43, 0xe9, 0x0e, 0x42, 0xed, 0x6b, 0x7f, 0x1b, 0x47, 0x23,
	0xe8, 0x92, 0xa5, 0x6c, 0x28, 0x89, 0xe7, 0x30, 0xf5, 0xaa, 0x72, 0x31, 0x30, 0x12, 0x31, 0x72,
	0xad, 0x2a, 0x97, 0xb1, 0x9c, 0x7e, 0x83, 0xc7, 0xe3, 0x10, 0xde, 0xde, 0x0d, 0x42, 0x14, 0x10,
	0x5e, 0xe9, 0xaf, 0xb4, 0x43, 0xf1, 0xe2, 0x90, 0x11, 0xf5, 0xff, 0x6f, 0xf1, 0xf2, 0x20, 0xb6,
	0xff, 0x12, 0xab, 0x07, 0x9b, 0x57, 0x9f, 0x64, 0x55, 0xfb, 0xcf, 0x5d, 0x2e, 0x0b, 0x6a, 0x93,
	0xb6, 0x2e, 0x2c, 0x39, 0xda, 0xfa, 0xa4, 0xa5, 0x22, 0xb1, 0xa6, 0x48, 0xee, 0x8d, 0x92, 0xc1,
	0x28, 0x0f, 0x79, 0xc6, 0xaf, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x72, 0xe9, 0x12, 0xb5, 0x5a,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadBalancerAgentClient is the client API for LoadBalancerAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadBalancerAgentClient interface {
	Invoke(ctx context.Context, in *LoadBalancerRequest, opts ...grpc.CallOption) (*LoadBalancerResponse, error)
}

type loadBalancerAgentClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancerAgentClient(cc *grpc.ClientConn) LoadBalancerAgentClient {
	return &loadBalancerAgentClient{cc}
}

func (c *loadBalancerAgentClient) Invoke(ctx context.Context, in *LoadBalancerRequest, opts ...grpc.CallOption) (*LoadBalancerResponse, error) {
	out := new(LoadBalancerResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.network.LoadBalancerAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadBalancerAgentServer is the server API for LoadBalancerAgent service.
type LoadBalancerAgentServer interface {
	Invoke(context.Context, *LoadBalancerRequest) (*LoadBalancerResponse, error)
}

// UnimplementedLoadBalancerAgentServer can be embedded to have forward compatible implementations.
type UnimplementedLoadBalancerAgentServer struct {
}

func (*UnimplementedLoadBalancerAgentServer) Invoke(ctx context.Context, req *LoadBalancerRequest) (*LoadBalancerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterLoadBalancerAgentServer(s *grpc.Server, srv LoadBalancerAgentServer) {
	s.RegisterService(&_LoadBalancerAgent_serviceDesc, srv)
}

func _LoadBalancerAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.network.LoadBalancerAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerAgentServer).Invoke(ctx, req.(*LoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancerAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.network.LoadBalancerAgent",
	HandlerType: (*LoadBalancerAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _LoadBalancerAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loadbalancer.proto",
}
