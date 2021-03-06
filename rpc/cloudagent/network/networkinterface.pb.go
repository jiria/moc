// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkinterface.proto

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

type NetworkInterface_NetworkInterfaceType int32

const (
	NetworkInterface_Local  NetworkInterface_NetworkInterfaceType = 0
	NetworkInterface_Remote NetworkInterface_NetworkInterfaceType = 1
)

var NetworkInterface_NetworkInterfaceType_name = map[int32]string{
	0: "Local",
	1: "Remote",
}

var NetworkInterface_NetworkInterfaceType_value = map[string]int32{
	"Local":  0,
	"Remote": 1,
}

func (x NetworkInterface_NetworkInterfaceType) String() string {
	return proto.EnumName(NetworkInterface_NetworkInterfaceType_name, int32(x))
}

func (NetworkInterface_NetworkInterfaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_72a481e6e4e12e01, []int{3, 0}
}

type NetworkInterfaceRequest struct {
	NetworkInterfaces    []*NetworkInterface `protobuf:"bytes,1,rep,name=NetworkInterfaces,proto3" json:"NetworkInterfaces,omitempty"`
	OperationType        common.Operation    `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NetworkInterfaceRequest) Reset()         { *m = NetworkInterfaceRequest{} }
func (m *NetworkInterfaceRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaceRequest) ProtoMessage()    {}
func (*NetworkInterfaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72a481e6e4e12e01, []int{0}
}

func (m *NetworkInterfaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterfaceRequest.Unmarshal(m, b)
}
func (m *NetworkInterfaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterfaceRequest.Marshal(b, m, deterministic)
}
func (m *NetworkInterfaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterfaceRequest.Merge(m, src)
}
func (m *NetworkInterfaceRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkInterfaceRequest.Size(m)
}
func (m *NetworkInterfaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterfaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterfaceRequest proto.InternalMessageInfo

func (m *NetworkInterfaceRequest) GetNetworkInterfaces() []*NetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

func (m *NetworkInterfaceRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type NetworkInterfaceResponse struct {
	NetworkInterfaces    []*NetworkInterface `protobuf:"bytes,1,rep,name=NetworkInterfaces,proto3" json:"NetworkInterfaces,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NetworkInterfaceResponse) Reset()         { *m = NetworkInterfaceResponse{} }
func (m *NetworkInterfaceResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaceResponse) ProtoMessage()    {}
func (*NetworkInterfaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72a481e6e4e12e01, []int{1}
}

func (m *NetworkInterfaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterfaceResponse.Unmarshal(m, b)
}
func (m *NetworkInterfaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterfaceResponse.Marshal(b, m, deterministic)
}
func (m *NetworkInterfaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterfaceResponse.Merge(m, src)
}
func (m *NetworkInterfaceResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkInterfaceResponse.Size(m)
}
func (m *NetworkInterfaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterfaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterfaceResponse proto.InternalMessageInfo

func (m *NetworkInterfaceResponse) GetNetworkInterfaces() []*NetworkInterface {
	if m != nil {
		return m.NetworkInterfaces
	}
	return nil
}

func (m *NetworkInterfaceResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *NetworkInterfaceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type IpConfiguration struct {
	Ipaddress               string                    `protobuf:"bytes,1,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	Prefixlength            string                    `protobuf:"bytes,2,opt,name=prefixlength,proto3" json:"prefixlength,omitempty"`
	Subnetid                string                    `protobuf:"bytes,3,opt,name=subnetid,proto3" json:"subnetid,omitempty"`
	Primary                 bool                      `protobuf:"varint,4,opt,name=primary,proto3" json:"primary,omitempty"`
	Loadbalanceraddresspool []string                  `protobuf:"bytes,5,rep,name=loadbalanceraddresspool,proto3" json:"loadbalanceraddresspool,omitempty"`
	Allocation              common.IPAllocationMethod `protobuf:"varint,6,opt,name=allocation,proto3,enum=moc.IPAllocationMethod" json:"allocation,omitempty"`
	Gateway                 string                    `protobuf:"bytes,7,opt,name=gateway,proto3" json:"gateway,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                  `json:"-"`
	XXX_unrecognized        []byte                    `json:"-"`
	XXX_sizecache           int32                     `json:"-"`
}

func (m *IpConfiguration) Reset()         { *m = IpConfiguration{} }
func (m *IpConfiguration) String() string { return proto.CompactTextString(m) }
func (*IpConfiguration) ProtoMessage()    {}
func (*IpConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_72a481e6e4e12e01, []int{2}
}

func (m *IpConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpConfiguration.Unmarshal(m, b)
}
func (m *IpConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpConfiguration.Marshal(b, m, deterministic)
}
func (m *IpConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpConfiguration.Merge(m, src)
}
func (m *IpConfiguration) XXX_Size() int {
	return xxx_messageInfo_IpConfiguration.Size(m)
}
func (m *IpConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_IpConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_IpConfiguration proto.InternalMessageInfo

func (m *IpConfiguration) GetIpaddress() string {
	if m != nil {
		return m.Ipaddress
	}
	return ""
}

func (m *IpConfiguration) GetPrefixlength() string {
	if m != nil {
		return m.Prefixlength
	}
	return ""
}

func (m *IpConfiguration) GetSubnetid() string {
	if m != nil {
		return m.Subnetid
	}
	return ""
}

func (m *IpConfiguration) GetPrimary() bool {
	if m != nil {
		return m.Primary
	}
	return false
}

func (m *IpConfiguration) GetLoadbalanceraddresspool() []string {
	if m != nil {
		return m.Loadbalanceraddresspool
	}
	return nil
}

func (m *IpConfiguration) GetAllocation() common.IPAllocationMethod {
	if m != nil {
		return m.Allocation
	}
	return common.IPAllocationMethod_Invalid
}

func (m *IpConfiguration) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

type NetworkInterface struct {
	Name                 string                                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 NetworkInterface_NetworkInterfaceType `protobuf:"varint,3,opt,name=type,proto3,enum=moc.cloudagent.network.NetworkInterface_NetworkInterfaceType" json:"type,omitempty"`
	IpConfigurations     []*IpConfiguration                    `protobuf:"bytes,4,rep,name=ipConfigurations,proto3" json:"ipConfigurations,omitempty"`
	Macaddress           string                                `protobuf:"bytes,5,opt,name=macaddress,proto3" json:"macaddress,omitempty"`
	Dns                  *common.Dns                           `protobuf:"bytes,6,opt,name=dns,proto3" json:"dns,omitempty"`
	NodeName             string                                `protobuf:"bytes,7,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	GroupName            string                                `protobuf:"bytes,8,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName         string                                `protobuf:"bytes,9,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Status               *common.Status                        `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	VirtualMachineName   string                                `protobuf:"bytes,11,opt,name=virtualMachineName,proto3" json:"virtualMachineName,omitempty"`
	IovWeight            uint32                                `protobuf:"varint,12,opt,name=iovWeight,proto3" json:"iovWeight,omitempty"`
	Tags                 *common.Tags                          `protobuf:"bytes,13,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *NetworkInterface) Reset()         { *m = NetworkInterface{} }
func (m *NetworkInterface) String() string { return proto.CompactTextString(m) }
func (*NetworkInterface) ProtoMessage()    {}
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_72a481e6e4e12e01, []int{3}
}

func (m *NetworkInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterface.Unmarshal(m, b)
}
func (m *NetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterface.Marshal(b, m, deterministic)
}
func (m *NetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterface.Merge(m, src)
}
func (m *NetworkInterface) XXX_Size() int {
	return xxx_messageInfo_NetworkInterface.Size(m)
}
func (m *NetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterface proto.InternalMessageInfo

func (m *NetworkInterface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkInterface) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkInterface) GetType() NetworkInterface_NetworkInterfaceType {
	if m != nil {
		return m.Type
	}
	return NetworkInterface_Local
}

func (m *NetworkInterface) GetIpConfigurations() []*IpConfiguration {
	if m != nil {
		return m.IpConfigurations
	}
	return nil
}

func (m *NetworkInterface) GetMacaddress() string {
	if m != nil {
		return m.Macaddress
	}
	return ""
}

func (m *NetworkInterface) GetDns() *common.Dns {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *NetworkInterface) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NetworkInterface) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *NetworkInterface) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *NetworkInterface) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *NetworkInterface) GetVirtualMachineName() string {
	if m != nil {
		return m.VirtualMachineName
	}
	return ""
}

func (m *NetworkInterface) GetIovWeight() uint32 {
	if m != nil {
		return m.IovWeight
	}
	return 0
}

func (m *NetworkInterface) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.cloudagent.network.NetworkInterface_NetworkInterfaceType", NetworkInterface_NetworkInterfaceType_name, NetworkInterface_NetworkInterfaceType_value)
	proto.RegisterType((*NetworkInterfaceRequest)(nil), "moc.cloudagent.network.NetworkInterfaceRequest")
	proto.RegisterType((*NetworkInterfaceResponse)(nil), "moc.cloudagent.network.NetworkInterfaceResponse")
	proto.RegisterType((*IpConfiguration)(nil), "moc.cloudagent.network.IpConfiguration")
	proto.RegisterType((*NetworkInterface)(nil), "moc.cloudagent.network.NetworkInterface")
}

func init() { proto.RegisterFile("networkinterface.proto", fileDescriptor_72a481e6e4e12e01) }

var fileDescriptor_72a481e6e4e12e01 = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xf3, 0xd7, 0x64, 0xd2, 0x96, 0xb0, 0x94, 0xd6, 0x8a, 0xa0, 0x8a, 0xc2, 0x81, 0x5c,
	0xb0, 0x21, 0x20, 0xc1, 0x85, 0x43, 0x0b, 0x1c, 0x22, 0xd1, 0x02, 0xdb, 0xaa, 0x48, 0xdc, 0x36,
	0xf6, 0xc6, 0x59, 0xd5, 0xde, 0x31, 0xbb, 0xeb, 0x96, 0xbe, 0x00, 0x8f, 0xc0, 0x2b, 0xf0, 0x0c,
	0x48, 0x3c, 0x1c, 0xf2, 0xda, 0x49, 0xd3, 0xa4, 0x95, 0x7a, 0xe1, 0xe6, 0xf9, 0x66, 0xf6, 0x9b,
	0xcf, 0xdf, 0xcc, 0x2e, 0xec, 0x48, 0x6e, 0x2e, 0x50, 0x9d, 0x09, 0x69, 0xb8, 0x9a, 0xb0, 0x80,
	0x7b, 0xa9, 0x42, 0x83, 0x64, 0x27, 0xc1, 0xc0, 0x0b, 0x62, 0xcc, 0x42, 0x16, 0x71, 0x69, 0xbc,
	0xb2, 0xac, 0xbb, 0x17, 0x21, 0x46, 0x31, 0xf7, 0x6d, 0xd5, 0x38, 0x9b, 0xf8, 0x17, 0x8a, 0xa5,
	0x29, 0x57, 0xba, 0x38, 0xd7, 0xdd, 0x08, 0x30, 0x49, 0x50, 0x96, 0xd1, 0x83, 0xf2, 0xd8, 0x22,
	0xd8, 0xff, 0xed, 0xc0, 0xee, 0x51, 0x81, 0x8f, 0x66, 0x5d, 0x29, 0xff, 0x9e, 0x71, 0x6d, 0xc8,
	0x29, 0xdc, 0x5f, 0x4e, 0x69, 0xd7, 0xe9, 0x55, 0x07, 0xed, 0xe1, 0xc0, 0xbb, 0x59, 0x92, 0xb7,
	0xc2, 0xb5, 0x4a, 0x41, 0x5e, 0xc1, 0xe6, 0xa7, 0x94, 0x2b, 0x66, 0x04, 0xca, 0x93, 0xcb, 0x94,
	0xbb, 0x95, 0x9e, 0x33, 0xd8, 0x1a, 0x6e, 0x59, 0xce, 0x79, 0x86, 0x5e, 0x2f, 0xea, 0xff, 0x75,
	0xc0, 0x5d, 0x55, 0xaa, 0x53, 0x94, 0x9a, 0xff, 0x37, 0xa9, 0x43, 0x68, 0x50, 0xae, 0xb3, 0xd8,
	0x58, 0x8d, 0xed, 0x61, 0xd7, 0x2b, 0x2c, 0xf7, 0x66, 0x96, 0x7b, 0x07, 0x88, 0xf1, 0x29, 0x8b,
	0x33, 0x4e, 0xcb, 0x4a, 0xb2, 0x0d, 0xf5, 0x0f, 0x4a, 0xa1, 0x72, 0xab, 0x3d, 0x67, 0xd0, 0xa2,
	0x45, 0xd0, 0xff, 0x55, 0x81, 0x7b, 0xa3, 0xf4, 0x1d, 0xca, 0x89, 0x88, 0xb2, 0xe2, 0xb7, 0xc8,
	0x23, 0x68, 0x89, 0x94, 0x85, 0xa1, 0xe2, 0x3a, 0x57, 0x9b, 0x57, 0x5f, 0x01, 0xa4, 0x0f, 0x1b,
	0xa9, 0xe2, 0x13, 0xf1, 0x23, 0xe6, 0x32, 0x32, 0x53, 0xab, 0xa0, 0x45, 0xaf, 0x61, 0xa4, 0x0b,
	0x4d, 0x9d, 0x8d, 0x25, 0x37, 0x22, 0x2c, 0xdb, 0xcd, 0x63, 0xe2, 0xc2, 0x7a, 0xaa, 0x44, 0xc2,
	0xd4, 0xa5, 0x5b, 0xeb, 0x39, 0x83, 0x26, 0x9d, 0x85, 0xe4, 0x0d, 0xec, 0xc6, 0xc8, 0xc2, 0x31,
	0x8b, 0x99, 0x0c, 0xb8, 0x2a, 0x1b, 0xa6, 0x88, 0xb1, 0x5b, 0xef, 0x55, 0x07, 0x2d, 0x7a, 0x5b,
	0x9a, 0xbc, 0x06, 0x60, 0x71, 0x8c, 0x81, 0xd5, 0xef, 0x36, 0xec, 0xdc, 0x76, 0xad, 0xc1, 0xa3,
	0xcf, 0xfb, 0xf3, 0xc4, 0x21, 0x37, 0x53, 0x0c, 0xe9, 0x42, 0x69, 0x2e, 0x26, 0x62, 0x86, 0x5f,
	0xb0, 0x4b, 0x77, 0xdd, 0xea, 0x9c, 0x85, 0xfd, 0x3f, 0x35, 0xe8, 0x2c, 0x1b, 0x4f, 0x08, 0xd4,
	0x24, 0x4b, 0x78, 0x69, 0x8a, 0xfd, 0x26, 0x5b, 0x50, 0x11, 0x61, 0xe9, 0x42, 0x45, 0x84, 0xe4,
	0x0b, 0xd4, 0x4c, 0xbe, 0x3d, 0x55, 0xab, 0xe2, 0xed, 0x5d, 0xc7, 0xbc, 0x02, 0xe4, 0xdb, 0x45,
	0x2d, 0x15, 0x39, 0x86, 0x8e, 0xb8, 0x3e, 0x23, 0xed, 0xd6, 0xec, 0x16, 0x3d, 0xbd, 0x8d, 0x7e,
	0x69, 0xa6, 0x74, 0x85, 0x80, 0xec, 0x01, 0x24, 0x2c, 0x98, 0x8d, 0xb9, 0x6e, 0xf5, 0x2f, 0x20,
	0xa4, 0x0b, 0xd5, 0x50, 0x6a, 0x6b, 0x66, 0x7b, 0xd8, 0xb4, 0x7d, 0xde, 0x4b, 0x4d, 0x73, 0x30,
	0x9f, 0xaf, 0xc4, 0x90, 0x1f, 0xe5, 0x5e, 0x14, 0xbe, 0xcd, 0xe3, 0x7c, 0x7b, 0x22, 0x85, 0x59,
	0x6a, 0x93, 0xcd, 0x62, 0x7b, 0xe6, 0x40, 0xbe, 0x3d, 0x33, 0xf3, 0x6d, 0x41, 0xab, 0xd8, 0x9e,
	0x45, 0x8c, 0x3c, 0x81, 0x86, 0x36, 0xcc, 0x64, 0xda, 0x05, 0xdb, 0xbc, 0x6d, 0x9b, 0x1f, 0x5b,
	0x88, 0x96, 0x29, 0xe2, 0x01, 0x39, 0x17, 0xca, 0x64, 0x2c, 0x3e, 0x64, 0xc1, 0x54, 0xc8, 0x42,
	0x4c, 0xdb, 0xd2, 0xdd, 0x90, 0xb1, 0x4b, 0x8d, 0xe7, 0x5f, 0xb9, 0x88, 0xa6, 0xc6, 0xdd, 0xe8,
	0x39, 0x83, 0x4d, 0x7a, 0x05, 0x90, 0xc7, 0x50, 0x33, 0x2c, 0xd2, 0xee, 0xa6, 0x6d, 0xd8, 0xb2,
	0x0d, 0x4f, 0x58, 0xa4, 0xa9, 0x85, 0xfb, 0xcf, 0x60, 0xfb, 0xa6, 0xf1, 0x90, 0x16, 0xd4, 0x3f,
	0x62, 0xc0, 0xe2, 0xce, 0x1a, 0x81, 0xfc, 0x4a, 0x26, 0x68, 0x78, 0xc7, 0x19, 0xfe, 0x74, 0xe0,
	0xe1, 0x72, 0xfd, 0x7e, 0x3e, 0x1e, 0x92, 0x40, 0x63, 0x24, 0xcf, 0xf1, 0x8c, 0x13, 0xff, 0xce,
	0xf7, 0xbf, 0x78, 0xf6, 0xba, 0xcf, 0xef, 0x7e, 0xa0, 0x78, 0x7d, 0xfa, 0x6b, 0x07, 0x2f, 0xbe,
	0xf9, 0x91, 0x30, 0xd3, 0x6c, 0xec, 0x05, 0x98, 0xf8, 0x89, 0x08, 0x14, 0x6a, 0x9c, 0x18, 0x3f,
	0xc1, 0xc0, 0x57, 0x69, 0xe0, 0x5f, 0xb1, 0xf9, 0x25, 0xdb, 0xb8, 0x61, 0x9f, 0x90, 0x97, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xaf, 0x0a, 0xea, 0xf5, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkInterfaceAgentClient is the client API for NetworkInterfaceAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkInterfaceAgentClient interface {
	Invoke(ctx context.Context, in *NetworkInterfaceRequest, opts ...grpc.CallOption) (*NetworkInterfaceResponse, error)
}

type networkInterfaceAgentClient struct {
	cc *grpc.ClientConn
}

func NewNetworkInterfaceAgentClient(cc *grpc.ClientConn) NetworkInterfaceAgentClient {
	return &networkInterfaceAgentClient{cc}
}

func (c *networkInterfaceAgentClient) Invoke(ctx context.Context, in *NetworkInterfaceRequest, opts ...grpc.CallOption) (*NetworkInterfaceResponse, error) {
	out := new(NetworkInterfaceResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.NetworkInterfaceAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkInterfaceAgentServer is the server API for NetworkInterfaceAgent service.
type NetworkInterfaceAgentServer interface {
	Invoke(context.Context, *NetworkInterfaceRequest) (*NetworkInterfaceResponse, error)
}

// UnimplementedNetworkInterfaceAgentServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkInterfaceAgentServer struct {
}

func (*UnimplementedNetworkInterfaceAgentServer) Invoke(ctx context.Context, req *NetworkInterfaceRequest) (*NetworkInterfaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterNetworkInterfaceAgentServer(s *grpc.Server, srv NetworkInterfaceAgentServer) {
	s.RegisterService(&_NetworkInterfaceAgent_serviceDesc, srv)
}

func _NetworkInterfaceAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkInterfaceAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.NetworkInterfaceAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkInterfaceAgentServer).Invoke(ctx, req.(*NetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkInterfaceAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.network.NetworkInterfaceAgent",
	HandlerType: (*NetworkInterfaceAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _NetworkInterfaceAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkinterface.proto",
}
