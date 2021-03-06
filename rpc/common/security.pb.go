// Code generated by protoc-gen-go. DO NOT EDIT.
// source: security.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Algorithm int32

const (
	Algorithm_A_UNKNOWN  Algorithm = 0
	Algorithm_RSA15      Algorithm = 1
	Algorithm_RSAOAEP    Algorithm = 2
	Algorithm_RSAOAEP256 Algorithm = 3
	Algorithm_A256KW     Algorithm = 4
)

var Algorithm_name = map[int32]string{
	0: "A_UNKNOWN",
	1: "RSA15",
	2: "RSAOAEP",
	3: "RSAOAEP256",
	4: "A256KW",
}

var Algorithm_value = map[string]int32{
	"A_UNKNOWN":  0,
	"RSA15":      1,
	"RSAOAEP":    2,
	"RSAOAEP256": 3,
	"A256KW":     4,
}

func (x Algorithm) String() string {
	return proto.EnumName(Algorithm_name, int32(x))
}

func (Algorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{0}
}

type KeyOperation int32

const (
	KeyOperation_ENCRYPT   KeyOperation = 0
	KeyOperation_DECRYPT   KeyOperation = 1
	KeyOperation_WRAPKEY   KeyOperation = 2
	KeyOperation_UNWRAPKEY KeyOperation = 3
)

var KeyOperation_name = map[int32]string{
	0: "ENCRYPT",
	1: "DECRYPT",
	2: "WRAPKEY",
	3: "UNWRAPKEY",
}

var KeyOperation_value = map[string]int32{
	"ENCRYPT":   0,
	"DECRYPT":   1,
	"WRAPKEY":   2,
	"UNWRAPKEY": 3,
}

func (x KeyOperation) String() string {
	return proto.EnumName(KeyOperation_name, int32(x))
}

func (KeyOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{1}
}

// https://docs.microsoft.com/en-us/rest/api/keyvault/createkey/createkey#jsonwebkeytype
type JsonWebKeyType int32

const (
	JsonWebKeyType_EC      JsonWebKeyType = 0
	JsonWebKeyType_EC_HSM  JsonWebKeyType = 1
	JsonWebKeyType_RSA     JsonWebKeyType = 2
	JsonWebKeyType_RSA_HSM JsonWebKeyType = 3
	JsonWebKeyType_OCT     JsonWebKeyType = 4
	JsonWebKeyType_AES     JsonWebKeyType = 5
)

var JsonWebKeyType_name = map[int32]string{
	0: "EC",
	1: "EC_HSM",
	2: "RSA",
	3: "RSA_HSM",
	4: "OCT",
	5: "AES",
}

var JsonWebKeyType_value = map[string]int32{
	"EC":      0,
	"EC_HSM":  1,
	"RSA":     2,
	"RSA_HSM": 3,
	"OCT":     4,
	"AES":     5,
}

func (x JsonWebKeyType) String() string {
	return proto.EnumName(JsonWebKeyType_name, int32(x))
}

func (JsonWebKeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{2}
}

type JsonWebKeyCurveName int32

const (
	JsonWebKeyCurveName_P_256  JsonWebKeyCurveName = 0
	JsonWebKeyCurveName_P_256K JsonWebKeyCurveName = 1
	JsonWebKeyCurveName_P_384  JsonWebKeyCurveName = 2
	JsonWebKeyCurveName_P_521  JsonWebKeyCurveName = 3
)

var JsonWebKeyCurveName_name = map[int32]string{
	0: "P_256",
	1: "P_256K",
	2: "P_384",
	3: "P_521",
}

var JsonWebKeyCurveName_value = map[string]int32{
	"P_256":  0,
	"P_256K": 1,
	"P_384":  2,
	"P_521":  3,
}

func (x JsonWebKeyCurveName) String() string {
	return proto.EnumName(JsonWebKeyCurveName_name, int32(x))
}

func (JsonWebKeyCurveName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{3}
}

type KeySize int32

const (
	KeySize_K_UNKNOWN KeySize = 0
	KeySize__256      KeySize = 1
	KeySize__2048     KeySize = 2
	KeySize__3072     KeySize = 3
	KeySize__4096     KeySize = 4
)

var KeySize_name = map[int32]string{
	0: "K_UNKNOWN",
	1: "_256",
	2: "_2048",
	3: "_3072",
	4: "_4096",
}

var KeySize_value = map[string]int32{
	"K_UNKNOWN": 0,
	"_256":      1,
	"_2048":     2,
	"_3072":     3,
	"_4096":     4,
}

func (x KeySize) String() string {
	return proto.EnumName(KeySize_name, int32(x))
}

func (KeySize) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{4}
}

func init() {
	proto.RegisterEnum("moc.Algorithm", Algorithm_name, Algorithm_value)
	proto.RegisterEnum("moc.KeyOperation", KeyOperation_name, KeyOperation_value)
	proto.RegisterEnum("moc.JsonWebKeyType", JsonWebKeyType_name, JsonWebKeyType_value)
	proto.RegisterEnum("moc.JsonWebKeyCurveName", JsonWebKeyCurveName_name, JsonWebKeyCurveName_value)
	proto.RegisterEnum("moc.KeySize", KeySize_name, KeySize_value)
}

func init() { proto.RegisterFile("security.proto", fileDescriptor_55a487c716a8b59c) }

var fileDescriptor_55a487c716a8b59c = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x51, 0x8f, 0x9a, 0x40,
	0x14, 0x85, 0x15, 0x50, 0xeb, 0x6d, 0x6b, 0x6e, 0xa6, 0x3f, 0xa1, 0xe9, 0x0b, 0x0f, 0xa2, 0x28,
	0xd6, 0x3e, 0x8e, 0x48, 0xd2, 0x74, 0x5a, 0x20, 0xa0, 0x21, 0xf6, 0x85, 0x28, 0x99, 0x2a, 0x49,
	0xc7, 0x21, 0x88, 0x9b, 0xb0, 0xbf, 0x7e, 0x33, 0x83, 0x9b, 0xdd, 0xb7, 0x73, 0x4e, 0x72, 0xbe,
	0x7b, 0xef, 0x0c, 0x4c, 0x6e, 0xbc, 0xb8, 0xd7, 0x65, 0xd3, 0x4e, 0xab, 0x5a, 0x36, 0x92, 0x98,
	0x42, 0x16, 0x76, 0x08, 0x63, 0xfa, 0xff, 0x2c, 0xeb, 0xb2, 0xb9, 0x08, 0xf2, 0x19, 0xc6, 0x34,
	0xdf, 0x87, 0x2c, 0x8c, 0xb2, 0x10, 0x7b, 0x64, 0x0c, 0x83, 0x24, 0xa5, 0x73, 0x0f, 0xfb, 0xe4,
	0x23, 0x8c, 0x92, 0x94, 0x46, 0x34, 0x88, 0xd1, 0x20, 0x13, 0x80, 0x87, 0x71, 0xbd, 0x15, 0x9a,
	0x04, 0x60, 0x48, 0x5d, 0x6f, 0xc5, 0x32, 0xb4, 0xec, 0x2d, 0x7c, 0x62, 0xbc, 0x8d, 0x2a, 0x5e,
	0x1f, 0x9b, 0x52, 0x5e, 0x55, 0x31, 0x08, 0xfd, 0xe4, 0x10, 0xef, 0xb0, 0xa7, 0xcc, 0x36, 0xe8,
	0x8c, 0x46, 0x66, 0x09, 0x8d, 0x59, 0x70, 0x40, 0x43, 0x4d, 0xde, 0x87, 0xaf, 0xd6, 0xb4, 0x7f,
	0xc3, 0xe4, 0xd7, 0x4d, 0x5e, 0x33, 0x7e, 0x62, 0xbc, 0xdd, 0xb5, 0x15, 0x27, 0x43, 0x30, 0x02,
	0x1f, 0x7b, 0x6a, 0x56, 0xe0, 0xe7, 0x3f, 0xd3, 0x3f, 0xd8, 0x27, 0x23, 0x30, 0x93, 0x94, 0xa2,
	0xf1, 0xd8, 0x4e, 0xa7, 0xa6, 0x4a, 0x23, 0x7f, 0x87, 0x96, 0x12, 0x34, 0x48, 0x71, 0x60, 0x6f,
	0xe0, 0xcb, 0x1b, 0xcd, 0xbf, 0xd7, 0x4f, 0x3c, 0x3c, 0x0a, 0xae, 0xce, 0x8b, 0x73, 0x75, 0x81,
	0xa6, 0x6a, 0xc9, 0xb0, 0xdf, 0xc5, 0x8b, 0xf5, 0x12, 0x8d, 0x4e, 0x7a, 0xee, 0x1c, 0x4d, 0xdb,
	0x87, 0x11, 0xe3, 0x6d, 0x5a, 0x3e, 0x73, 0xb5, 0x2b, 0x7b, 0xf7, 0x4a, 0x1f, 0xc0, 0xd2, 0x14,
	0xdd, 0xcc, 0xdd, 0xd9, 0x72, 0xdd, 0x35, 0xf3, 0xc5, 0xec, 0xbb, 0x8b, 0xa6, 0x96, 0xcb, 0xd9,
	0x8f, 0x15, 0x5a, 0x9b, 0x6f, 0x7f, 0xbf, 0x9e, 0xcb, 0xe6, 0x72, 0x3f, 0x4d, 0x0b, 0x29, 0x1c,
	0x51, 0x16, 0xb5, 0xbc, 0xc9, 0x7f, 0x8d, 0x23, 0x64, 0xe1, 0xd4, 0x55, 0xe1, 0x14, 0x52, 0x08,
	0x79, 0x3d, 0x0d, 0xf5, 0xff, 0x2c, 0x5e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x5f, 0x4d, 0xb4,
	0xb1, 0x01, 0x00, 0x00,
}
