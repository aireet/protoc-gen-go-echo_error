// Code generated by protoc-gen-go. DO NOT EDIT.
// source: errors.proto

package go_micro_srv_test

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

type Errors_ErrorCode int32

const (
	Errors_Default      Errors_ErrorCode = 0
	Errors_UserNotFind  Errors_ErrorCode = 300001
	Errors_UserNotLogin Errors_ErrorCode = 300002
	Errors_InfraError   Errors_ErrorCode = 300003
)

var Errors_ErrorCode_name = map[int32]string{
	0:      "Default",
	300001: "UserNotFind",
	300002: "UserNotLogin",
	300003: "InfraError",
}

var Errors_ErrorCode_value = map[string]int32{
	"Default":      0,
	"UserNotFind":  300001,
	"UserNotLogin": 300002,
	"InfraError":   300003,
}

func (x Errors_ErrorCode) String() string {
	return proto.EnumName(Errors_ErrorCode_name, int32(x))
}

func (Errors_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0, 0}
}

type Errors struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Errors) Reset()         { *m = Errors{} }
func (m *Errors) String() string { return proto.CompactTextString(m) }
func (*Errors) ProtoMessage()    {}
func (*Errors) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0}
}

func (m *Errors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Errors.Unmarshal(m, b)
}
func (m *Errors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Errors.Marshal(b, m, deterministic)
}
func (m *Errors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Errors.Merge(m, src)
}
func (m *Errors) XXX_Size() int {
	return xxx_messageInfo_Errors.Size(m)
}
func (m *Errors) XXX_DiscardUnknown() {
	xxx_messageInfo_Errors.DiscardUnknown(m)
}

var xxx_messageInfo_Errors proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("go.micro.srv.test.Errors_ErrorCode", Errors_ErrorCode_name, Errors_ErrorCode_value)
	proto.RegisterType((*Errors)(nil), "go.micro.srv.test.Errors")
}

func init() { proto.RegisterFile("errors.proto", fileDescriptor_24fe73c7f0ddb19c) }

var fileDescriptor_24fe73c7f0ddb19c = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2d, 0x2a, 0xca,
	0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c,
	0x2e, 0xca, 0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x2b, 0x49, 0x2d, 0x2e, 0x51, 0x8a, 0xe6, 0x62, 0x73,
	0x05, 0x2b, 0x51, 0x0a, 0xe4, 0xe2, 0x04, 0xb3, 0x9c, 0xf3, 0x53, 0x52, 0x85, 0xb8, 0xb9, 0xd8,
	0x5d, 0x52, 0xd3, 0x12, 0x4b, 0x73, 0x4a, 0x04, 0x18, 0x84, 0x04, 0xb9, 0xb8, 0x43, 0x8b, 0x53,
	0x8b, 0xfc, 0xf2, 0x4b, 0xdc, 0x32, 0xf3, 0x52, 0x04, 0x1e, 0x2e, 0x17, 0x12, 0x12, 0xe2, 0xe2,
	0x81, 0x0a, 0xf9, 0xe4, 0xa7, 0x67, 0xe6, 0x09, 0x3c, 0x5a, 0x2e, 0x24, 0x24, 0xc0, 0xc5, 0xe5,
	0x99, 0x97, 0x56, 0x94, 0x08, 0x36, 0x45, 0xe0, 0xf1, 0x72, 0x21, 0x27, 0xe1, 0x28, 0x4c, 0x1b,
	0x93, 0xd8, 0xc0, 0x6e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x15, 0x7b, 0xad, 0x90, 0x9b,
	0x00, 0x00, 0x00,
}