// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtualmachinescaleset.proto

package compute

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
	network "github.com/microsoft/moc/rpc/nodeagent/network"
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

type VirtualMachineScaleSetRequest struct {
	VirtualMachineScaleSetSystems []*VirtualMachineScaleSet `protobuf:"bytes,1,rep,name=VirtualMachineScaleSetSystems,proto3" json:"VirtualMachineScaleSetSystems,omitempty"`
	OperationType                 common.Operation          `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                  `json:"-"`
	XXX_unrecognized              []byte                    `json:"-"`
	XXX_sizecache                 int32                     `json:"-"`
}

func (m *VirtualMachineScaleSetRequest) Reset()         { *m = VirtualMachineScaleSetRequest{} }
func (m *VirtualMachineScaleSetRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineScaleSetRequest) ProtoMessage()    {}
func (*VirtualMachineScaleSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{0}
}

func (m *VirtualMachineScaleSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineScaleSetRequest.Unmarshal(m, b)
}
func (m *VirtualMachineScaleSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineScaleSetRequest.Marshal(b, m, deterministic)
}
func (m *VirtualMachineScaleSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineScaleSetRequest.Merge(m, src)
}
func (m *VirtualMachineScaleSetRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineScaleSetRequest.Size(m)
}
func (m *VirtualMachineScaleSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineScaleSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineScaleSetRequest proto.InternalMessageInfo

func (m *VirtualMachineScaleSetRequest) GetVirtualMachineScaleSetSystems() []*VirtualMachineScaleSet {
	if m != nil {
		return m.VirtualMachineScaleSetSystems
	}
	return nil
}

func (m *VirtualMachineScaleSetRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualMachineScaleSetResponse struct {
	VirtualMachineScaleSetSystems []*VirtualMachineScaleSet `protobuf:"bytes,1,rep,name=VirtualMachineScaleSetSystems,proto3" json:"VirtualMachineScaleSetSystems,omitempty"`
	Result                        *wrappers.BoolValue       `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                         string                    `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                  `json:"-"`
	XXX_unrecognized              []byte                    `json:"-"`
	XXX_sizecache                 int32                     `json:"-"`
}

func (m *VirtualMachineScaleSetResponse) Reset()         { *m = VirtualMachineScaleSetResponse{} }
func (m *VirtualMachineScaleSetResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineScaleSetResponse) ProtoMessage()    {}
func (*VirtualMachineScaleSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{1}
}

func (m *VirtualMachineScaleSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineScaleSetResponse.Unmarshal(m, b)
}
func (m *VirtualMachineScaleSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineScaleSetResponse.Marshal(b, m, deterministic)
}
func (m *VirtualMachineScaleSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineScaleSetResponse.Merge(m, src)
}
func (m *VirtualMachineScaleSetResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineScaleSetResponse.Size(m)
}
func (m *VirtualMachineScaleSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineScaleSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineScaleSetResponse proto.InternalMessageInfo

func (m *VirtualMachineScaleSetResponse) GetVirtualMachineScaleSetSystems() []*VirtualMachineScaleSet {
	if m != nil {
		return m.VirtualMachineScaleSetSystems
	}
	return nil
}

func (m *VirtualMachineScaleSetResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualMachineScaleSetResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Sku struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             int64    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sku) Reset()         { *m = Sku{} }
func (m *Sku) String() string { return proto.CompactTextString(m) }
func (*Sku) ProtoMessage()    {}
func (*Sku) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{2}
}

func (m *Sku) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sku.Unmarshal(m, b)
}
func (m *Sku) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sku.Marshal(b, m, deterministic)
}
func (m *Sku) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sku.Merge(m, src)
}
func (m *Sku) XXX_Size() int {
	return xxx_messageInfo_Sku.Size(m)
}
func (m *Sku) XXX_DiscardUnknown() {
	xxx_messageInfo_Sku.DiscardUnknown(m)
}

var xxx_messageInfo_Sku proto.InternalMessageInfo

func (m *Sku) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sku) GetCapacity() int64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type NetworkConfigurationScaleSet struct {
	Interfaces           []*network.VirtualNetworkInterface `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *NetworkConfigurationScaleSet) Reset()         { *m = NetworkConfigurationScaleSet{} }
func (m *NetworkConfigurationScaleSet) String() string { return proto.CompactTextString(m) }
func (*NetworkConfigurationScaleSet) ProtoMessage()    {}
func (*NetworkConfigurationScaleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{3}
}

func (m *NetworkConfigurationScaleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkConfigurationScaleSet.Unmarshal(m, b)
}
func (m *NetworkConfigurationScaleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkConfigurationScaleSet.Marshal(b, m, deterministic)
}
func (m *NetworkConfigurationScaleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConfigurationScaleSet.Merge(m, src)
}
func (m *NetworkConfigurationScaleSet) XXX_Size() int {
	return xxx_messageInfo_NetworkConfigurationScaleSet.Size(m)
}
func (m *NetworkConfigurationScaleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConfigurationScaleSet.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConfigurationScaleSet proto.InternalMessageInfo

func (m *NetworkConfigurationScaleSet) GetInterfaces() []*network.VirtualNetworkInterface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type VirtualMachineProfile struct {
	Vmprefix             string                        `protobuf:"bytes,1,opt,name=vmprefix,proto3" json:"vmprefix,omitempty"`
	Network              *NetworkConfigurationScaleSet `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	Storage              *StorageConfiguration         `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Os                   *OperatingSystemConfiguration `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	Hardware             *HardwareConfiguration        `protobuf:"bytes,5,opt,name=hardware,proto3" json:"hardware,omitempty"`
	Security             *SecurityConfiguration        `protobuf:"bytes,6,opt,name=security,proto3" json:"security,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *VirtualMachineProfile) Reset()         { *m = VirtualMachineProfile{} }
func (m *VirtualMachineProfile) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineProfile) ProtoMessage()    {}
func (*VirtualMachineProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{4}
}

func (m *VirtualMachineProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineProfile.Unmarshal(m, b)
}
func (m *VirtualMachineProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineProfile.Marshal(b, m, deterministic)
}
func (m *VirtualMachineProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineProfile.Merge(m, src)
}
func (m *VirtualMachineProfile) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineProfile.Size(m)
}
func (m *VirtualMachineProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineProfile.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineProfile proto.InternalMessageInfo

func (m *VirtualMachineProfile) GetVmprefix() string {
	if m != nil {
		return m.Vmprefix
	}
	return ""
}

func (m *VirtualMachineProfile) GetNetwork() *NetworkConfigurationScaleSet {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *VirtualMachineProfile) GetStorage() *StorageConfiguration {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *VirtualMachineProfile) GetOs() *OperatingSystemConfiguration {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *VirtualMachineProfile) GetHardware() *HardwareConfiguration {
	if m != nil {
		return m.Hardware
	}
	return nil
}

func (m *VirtualMachineProfile) GetSecurity() *SecurityConfiguration {
	if m != nil {
		return m.Security
	}
	return nil
}

type VirtualMachineScaleSet struct {
	Name                    string                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                      string                       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Sku                     *Sku                         `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Virtualmachineprofile   *VirtualMachineProfile       `protobuf:"bytes,4,opt,name=virtualmachineprofile,proto3" json:"virtualmachineprofile,omitempty"`
	VirtualMachineSystems   []*VirtualMachine            `protobuf:"bytes,5,rep,name=VirtualMachineSystems,proto3" json:"VirtualMachineSystems,omitempty"`
	Status                  *common.Status               `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	DisableHighAvailability bool                         `protobuf:"varint,8,opt,name=DisableHighAvailability,proto3" json:"DisableHighAvailability,omitempty"`
	AllowedOwnerNodes       []string                     `protobuf:"bytes,9,rep,name=allowedOwnerNodes,proto3" json:"allowedOwnerNodes,omitempty"`
	Entity                  *common.Entity               `protobuf:"bytes,10,opt,name=entity,proto3" json:"entity,omitempty"`
	HighAvailabilityState   common.HighAvailabilityState `protobuf:"varint,11,opt,name=highAvailabilityState,proto3,enum=moc.HighAvailabilityState" json:"highAvailabilityState,omitempty"`
	Tags                    *common.Tags                 `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                     `json:"-"`
	XXX_unrecognized        []byte                       `json:"-"`
	XXX_sizecache           int32                        `json:"-"`
}

func (m *VirtualMachineScaleSet) Reset()         { *m = VirtualMachineScaleSet{} }
func (m *VirtualMachineScaleSet) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineScaleSet) ProtoMessage()    {}
func (*VirtualMachineScaleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ef1a2fddc3eb6a, []int{5}
}

func (m *VirtualMachineScaleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineScaleSet.Unmarshal(m, b)
}
func (m *VirtualMachineScaleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineScaleSet.Marshal(b, m, deterministic)
}
func (m *VirtualMachineScaleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineScaleSet.Merge(m, src)
}
func (m *VirtualMachineScaleSet) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineScaleSet.Size(m)
}
func (m *VirtualMachineScaleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineScaleSet.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineScaleSet proto.InternalMessageInfo

func (m *VirtualMachineScaleSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetSku() *Sku {
	if m != nil {
		return m.Sku
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetVirtualmachineprofile() *VirtualMachineProfile {
	if m != nil {
		return m.Virtualmachineprofile
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetVirtualMachineSystems() []*VirtualMachine {
	if m != nil {
		return m.VirtualMachineSystems
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetDisableHighAvailability() bool {
	if m != nil {
		return m.DisableHighAvailability
	}
	return false
}

func (m *VirtualMachineScaleSet) GetAllowedOwnerNodes() []string {
	if m != nil {
		return m.AllowedOwnerNodes
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetHighAvailabilityState() common.HighAvailabilityState {
	if m != nil {
		return m.HighAvailabilityState
	}
	return common.HighAvailabilityState_UNKNOWN_HA_STATE
}

func (m *VirtualMachineScaleSet) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualMachineScaleSetRequest)(nil), "moc.nodeagent.compute.VirtualMachineScaleSetRequest")
	proto.RegisterType((*VirtualMachineScaleSetResponse)(nil), "moc.nodeagent.compute.VirtualMachineScaleSetResponse")
	proto.RegisterType((*Sku)(nil), "moc.nodeagent.compute.Sku")
	proto.RegisterType((*NetworkConfigurationScaleSet)(nil), "moc.nodeagent.compute.NetworkConfigurationScaleSet")
	proto.RegisterType((*VirtualMachineProfile)(nil), "moc.nodeagent.compute.VirtualMachineProfile")
	proto.RegisterType((*VirtualMachineScaleSet)(nil), "moc.nodeagent.compute.VirtualMachineScaleSet")
}

func init() { proto.RegisterFile("virtualmachinescaleset.proto", fileDescriptor_d2ef1a2fddc3eb6a) }

var fileDescriptor_d2ef1a2fddc3eb6a = []byte{
	// 743 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x95, 0xcd, 0x6e, 0xe2, 0x48,
	0x10, 0xc7, 0xd7, 0x40, 0x08, 0x34, 0xd9, 0x48, 0xdb, 0x0a, 0xbb, 0x16, 0x9b, 0x44, 0x88, 0xd5,
	0x4a, 0x48, 0xcb, 0x9a, 0x11, 0x49, 0xa4, 0xb9, 0x26, 0x99, 0x48, 0xe4, 0x90, 0x0f, 0x35, 0x51,
	0x0e, 0x33, 0xa7, 0xc6, 0x14, 0xa6, 0x85, 0xed, 0xf6, 0x74, 0xb7, 0x61, 0xf2, 0x22, 0xf3, 0x00,
	0xf3, 0x38, 0x73, 0x98, 0x57, 0x98, 0x57, 0x19, 0xb9, 0xdd, 0x46, 0x81, 0xb1, 0x51, 0x72, 0x9a,
	0x1b, 0xdd, 0xf5, 0xaf, 0x5f, 0x55, 0x57, 0x95, 0x0b, 0x74, 0xb8, 0x60, 0x42, 0xc5, 0xd4, 0x0f,
	0xa8, 0x3b, 0x63, 0x21, 0x48, 0x97, 0xfa, 0x20, 0x41, 0x39, 0x91, 0xe0, 0x8a, 0xe3, 0x66, 0xc0,
	0x5d, 0x27, 0xe4, 0x13, 0xa0, 0x1e, 0x84, 0xca, 0x71, 0x79, 0x10, 0xc5, 0x0a, 0x5a, 0xc7, 0x1e,
	0xe7, 0x9e, 0x0f, 0x7d, 0x2d, 0x1a, 0xc7, 0xd3, 0xfe, 0x52, 0xd0, 0x28, 0x02, 0x21, 0x53, 0xb7,
	0xd6, 0xc1, 0x3a, 0xd4, 0xdc, 0xee, 0xb9, 0x3c, 0x08, 0x78, 0x68, 0x4e, 0x47, 0x46, 0x13, 0x82,
	0x5a, 0x72, 0x31, 0x67, 0xa1, 0x02, 0x31, 0xa5, 0xae, 0x11, 0x77, 0xbe, 0x5a, 0xe8, 0xe8, 0x31,
	0x55, 0xdc, 0xa4, 0x94, 0x51, 0x92, 0xda, 0x08, 0x14, 0x81, 0x8f, 0x31, 0x48, 0x85, 0x65, 0x91,
	0x60, 0xf4, 0x24, 0x15, 0x04, 0xd2, 0xb6, 0xda, 0xe5, 0x6e, 0x63, 0xf0, 0xbf, 0x93, 0xfb, 0x06,
	0xa7, 0x00, 0xbe, 0x9d, 0x89, 0x4f, 0xd1, 0xef, 0x77, 0x11, 0x08, 0xaa, 0x18, 0x0f, 0x1f, 0x9e,
	0x22, 0xb0, 0x4b, 0x6d, 0xab, 0xbb, 0x3f, 0xd8, 0xd7, 0x41, 0x56, 0x16, 0xb2, 0x2e, 0xea, 0x7c,
	0xb7, 0xd0, 0x71, 0xd1, 0x63, 0x64, 0xc4, 0x43, 0x09, 0xbf, 0xe6, 0x35, 0x03, 0x54, 0x25, 0x20,
	0x63, 0x5f, 0xe9, 0x67, 0x34, 0x06, 0x2d, 0x27, 0x6d, 0xac, 0x93, 0x35, 0xd6, 0xb9, 0xe0, 0xdc,
	0x7f, 0xa4, 0x7e, 0x0c, 0xc4, 0x28, 0xf1, 0x01, 0xda, 0xb9, 0x12, 0x82, 0x0b, 0xbb, 0xdc, 0xb6,
	0xba, 0x75, 0x92, 0x1e, 0x3a, 0x67, 0xa8, 0x3c, 0x9a, 0xc7, 0x18, 0xa3, 0x4a, 0x48, 0x03, 0xb0,
	0x2d, 0x6d, 0xd3, 0xbf, 0x71, 0x0b, 0xd5, 0x5c, 0x1a, 0x51, 0x97, 0xa9, 0x27, 0x1d, 0xa6, 0x4c,
	0x56, 0xe7, 0x4e, 0x88, 0x0e, 0x6f, 0xd3, 0xfe, 0x5f, 0xf2, 0x70, 0xca, 0xbc, 0x38, 0x2d, 0x5a,
	0x96, 0x27, 0xbe, 0x45, 0x68, 0x35, 0x18, 0x59, 0x09, 0x9c, 0x8d, 0x12, 0x98, 0x01, 0xca, 0x4a,
	0x60, 0x78, 0xd7, 0x99, 0x1b, 0x79, 0x46, 0xe8, 0x7c, 0x29, 0xa3, 0xe6, 0x7a, 0x49, 0xee, 0x05,
	0x9f, 0x32, 0x5f, 0x67, 0xb9, 0x08, 0x22, 0x01, 0x53, 0xf6, 0xc9, 0x64, 0xbf, 0x3a, 0xe3, 0x1b,
	0xb4, 0x6b, 0x82, 0x98, 0x3a, 0x9d, 0x14, 0x74, 0x61, 0xdb, 0x5b, 0x48, 0xc6, 0xc0, 0x57, 0x68,
	0x57, 0x2a, 0x2e, 0xa8, 0x07, 0xba, 0x86, 0x8d, 0xc1, 0x7f, 0x05, 0xb8, 0x51, 0xaa, 0x5a, 0xc3,
	0x91, 0xcc, 0x17, 0x5f, 0xa2, 0x12, 0x97, 0x76, 0x65, 0x6b, 0x42, 0x66, 0x0c, 0x43, 0x2f, 0xed,
	0xf8, 0x3a, 0xa9, 0xc4, 0x25, 0x1e, 0xa2, 0xda, 0x8c, 0x8a, 0xc9, 0x92, 0x0a, 0xb0, 0x77, 0x34,
	0xaa, 0x57, 0x80, 0x1a, 0x1a, 0xd9, 0x3a, 0x63, 0xe5, 0x9d, 0x90, 0x24, 0xb8, 0xb1, 0x48, 0xda,
	0x5c, 0xdd, 0x4a, 0x1a, 0x19, 0xd9, 0x06, 0x29, 0xf3, 0xee, 0x7c, 0xab, 0xa0, 0x3f, 0xf3, 0xe7,
	0x36, 0x77, 0xbe, 0xf6, 0x51, 0x89, 0x4d, 0x74, 0x63, 0xea, 0xa4, 0xc4, 0x26, 0xb8, 0x87, 0xca,
	0x72, 0x1e, 0x9b, 0xd2, 0xb6, 0x8a, 0x72, 0x98, 0xc7, 0x24, 0x91, 0xe1, 0x31, 0x6a, 0xae, 0x2f,
	0xab, 0x28, 0x1d, 0x08, 0x53, 0xd8, 0xde, 0x8b, 0xbe, 0x37, 0x33, 0x44, 0x24, 0x1f, 0x85, 0x3f,
	0x6c, 0x0e, 0x5d, 0xf6, 0x4d, 0xef, 0xe8, 0x81, 0xfe, 0xf7, 0x45, 0x31, 0x48, 0x3e, 0x03, 0xff,
	0x83, 0xaa, 0x52, 0x51, 0x15, 0x4b, 0x7b, 0x57, 0x67, 0xdc, 0xd0, 0xb4, 0x91, 0xbe, 0x22, 0xc6,
	0x84, 0xdf, 0xa2, 0xbf, 0xde, 0x31, 0x49, 0xc7, 0x3e, 0x0c, 0x99, 0x37, 0x3b, 0x5f, 0x50, 0xe6,
	0xd3, 0x31, 0xf3, 0x93, 0x5e, 0xd5, 0xda, 0x56, 0xb7, 0x46, 0x8a, 0xcc, 0xb8, 0x87, 0xfe, 0xa0,
	0xbe, 0xcf, 0x97, 0x30, 0xb9, 0x5b, 0x86, 0x20, 0x6e, 0xf9, 0x04, 0xa4, 0x5d, 0x6f, 0x97, 0xbb,
	0x75, 0xf2, 0xb3, 0x21, 0x49, 0x06, 0x42, 0x95, 0x60, 0xd1, 0xb3, 0x64, 0xae, 0xf4, 0x15, 0x31,
	0x26, 0x7c, 0x8f, 0x9a, 0xb3, 0x8d, 0x30, 0x49, 0xba, 0x60, 0x37, 0xf4, 0x2e, 0x4d, 0x5b, 0x36,
	0xcc, 0x53, 0x90, 0x7c, 0x47, 0x7c, 0x84, 0x2a, 0x8a, 0x7a, 0xd2, 0xde, 0xd3, 0x41, 0xeb, 0x1a,
	0xf0, 0x40, 0x3d, 0x49, 0xf4, 0xf5, 0xe0, 0xb3, 0x85, 0xfe, 0xce, 0x1f, 0xa8, 0xf3, 0xa4, 0xe0,
	0x78, 0x89, 0xaa, 0xd7, 0xe1, 0x82, 0xcf, 0x01, 0x9f, 0xbe, 0x6e, 0xbd, 0xa6, 0xff, 0x44, 0xad,
	0xb3, 0x57, 0x7a, 0xa5, 0x2b, 0xbf, 0xf3, 0xdb, 0xc5, 0x9b, 0xf7, 0x8e, 0xc7, 0xd4, 0x2c, 0x1e,
	0x27, 0x2e, 0xfd, 0x80, 0xb9, 0x82, 0x4b, 0x3e, 0x55, 0xfd, 0x80, 0xbb, 0x7d, 0x11, 0xb9, 0xfd,
	0x15, 0xb2, 0x6f, 0x90, 0xe3, 0xaa, 0xde, 0xcc, 0x27, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1c,
	0x74, 0xd2, 0x38, 0xb7, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualMachineScaleSetAgentClient is the client API for VirtualMachineScaleSetAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualMachineScaleSetAgentClient interface {
	Invoke(ctx context.Context, in *VirtualMachineScaleSetRequest, opts ...grpc.CallOption) (*VirtualMachineScaleSetResponse, error)
}

type virtualMachineScaleSetAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualMachineScaleSetAgentClient(cc *grpc.ClientConn) VirtualMachineScaleSetAgentClient {
	return &virtualMachineScaleSetAgentClient{cc}
}

func (c *virtualMachineScaleSetAgentClient) Invoke(ctx context.Context, in *VirtualMachineScaleSetRequest, opts ...grpc.CallOption) (*VirtualMachineScaleSetResponse, error) {
	out := new(VirtualMachineScaleSetResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.compute.VirtualMachineScaleSetAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineScaleSetAgentServer is the server API for VirtualMachineScaleSetAgent service.
type VirtualMachineScaleSetAgentServer interface {
	Invoke(context.Context, *VirtualMachineScaleSetRequest) (*VirtualMachineScaleSetResponse, error)
}

// UnimplementedVirtualMachineScaleSetAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineScaleSetAgentServer struct {
}

func (*UnimplementedVirtualMachineScaleSetAgentServer) Invoke(ctx context.Context, req *VirtualMachineScaleSetRequest) (*VirtualMachineScaleSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterVirtualMachineScaleSetAgentServer(s *grpc.Server, srv VirtualMachineScaleSetAgentServer) {
	s.RegisterService(&_VirtualMachineScaleSetAgent_serviceDesc, srv)
}

func _VirtualMachineScaleSetAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineScaleSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineScaleSetAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.compute.VirtualMachineScaleSetAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineScaleSetAgentServer).Invoke(ctx, req.(*VirtualMachineScaleSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMachineScaleSetAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.compute.VirtualMachineScaleSetAgent",
	HandlerType: (*VirtualMachineScaleSetAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualMachineScaleSetAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtualmachinescaleset.proto",
}
