// Code generated by protoc-gen-go. DO NOT EDIT.
// source: any/any.proto

/*
Package any is a generated protocol buffer package.

It is generated from these files:
	any/any.proto

It has these top-level messages:
	Any
*/
package any

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Any struct {
	Type  string `protobuf:"bytes,1,opt,name=Type" json:"Type,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *Any) Reset()                    { *m = Any{} }
func (m *Any) String() string            { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()               {}
func (*Any) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Any) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Any) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Any)(nil), "any.Any")
}

func init() { proto.RegisterFile("any/any.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xcc, 0xab, 0xd4,
	0x4f, 0xcc, 0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xcc, 0xab, 0x54, 0xd2,
	0xe7, 0x62, 0x76, 0xcc, 0xab, 0x14, 0x12, 0xe2, 0x62, 0x09, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0xc3, 0x12, 0x73, 0x4a, 0x53, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0x20, 0x1c, 0x27, 0xbd, 0x28, 0x9d, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xc4, 0xbc, 0x94, 0xca, 0x9c, 0xd4, 0xc4, 0x02, 0xfd,
	0xe2, 0xfc, 0xe4, 0xcc, 0xc4, 0x1c, 0xdd, 0xdc, 0xd4, 0x9c, 0xc4, 0xbc, 0xf4, 0x54, 0x7d, 0xb0,
	0x0d, 0x20, 0xbb, 0x92, 0xd8, 0xc0, 0x4c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x96,
	0x19, 0xba, 0x7d, 0x00, 0x00, 0x00,
}
