// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc/rpc.proto

It has these top-level messages:
	Request
	Response
	Packet
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/andyleap/social-melange/proto/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	ID      uint64   `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Request *any.Any `protobuf:"bytes,2,opt,name=Request" json:"Request,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Request) GetRequest() *any.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

type Response struct {
	ID       uint64   `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Response *any.Any `protobuf:"bytes,2,opt,name=Response" json:"Response,omitempty"`
	End      bool     `protobuf:"varint,3,opt,name=End" json:"End,omitempty"`
	Error    string   `protobuf:"bytes,4,opt,name=Error" json:"Error,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Response) GetResponse() *any.Any {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Response) GetEnd() bool {
	if m != nil {
		return m.End
	}
	return false
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Packet struct {
	// Types that are valid to be assigned to Packet:
	//	*Packet_Request
	//	*Packet_Response
	Packet isPacket_Packet `protobuf_oneof:"packet"`
}

func (m *Packet) Reset()                    { *m = Packet{} }
func (m *Packet) String() string            { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()               {}
func (*Packet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isPacket_Packet interface {
	isPacket_Packet()
}

type Packet_Request struct {
	Request *Request `protobuf:"bytes,1,opt,name=Request,oneof"`
}
type Packet_Response struct {
	Response *Response `protobuf:"bytes,2,opt,name=Response,oneof"`
}

func (*Packet_Request) isPacket_Packet()  {}
func (*Packet_Response) isPacket_Packet() {}

func (m *Packet) GetPacket() isPacket_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *Packet) GetRequest() *Request {
	if x, ok := m.GetPacket().(*Packet_Request); ok {
		return x.Request
	}
	return nil
}

func (m *Packet) GetResponse() *Response {
	if x, ok := m.GetPacket().(*Packet_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Packet) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Packet_OneofMarshaler, _Packet_OneofUnmarshaler, _Packet_OneofSizer, []interface{}{
		(*Packet_Request)(nil),
		(*Packet_Response)(nil),
	}
}

func _Packet_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Packet)
	// packet
	switch x := m.Packet.(type) {
	case *Packet_Request:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Request); err != nil {
			return err
		}
	case *Packet_Response:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Response); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Packet.Packet has unexpected type %T", x)
	}
	return nil
}

func _Packet_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Packet)
	switch tag {
	case 1: // packet.Request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Request)
		err := b.DecodeMessage(msg)
		m.Packet = &Packet_Request{msg}
		return true, err
	case 2: // packet.Response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Response)
		err := b.DecodeMessage(msg)
		m.Packet = &Packet_Response{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Packet_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Packet)
	// packet
	switch x := m.Packet.(type) {
	case *Packet_Request:
		s := proto.Size(x.Request)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_Response:
		s := proto.Size(x.Response)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Request)(nil), "rpc.Request")
	proto.RegisterType((*Response)(nil), "rpc.Response")
	proto.RegisterType((*Packet)(nil), "rpc.Packet")
}

func init() { proto.RegisterFile("rpc/rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0xa4, 0x84, 0x70, 0x10, 0x84, 0x2c, 0x86, 0x88, 0x29, 0x8a, 0x18, 0x2c, 0x01,
	0xb6, 0x04, 0x33, 0x03, 0x55, 0x2b, 0xb5, 0x1b, 0xf2, 0xc8, 0xe6, 0xba, 0xa7, 0x36, 0x22, 0xb5,
	0x8d, 0x93, 0x0e, 0xf9, 0xf7, 0x28, 0x49, 0x9b, 0x0c, 0xd9, 0xce, 0xef, 0xf9, 0x9d, 0x3f, 0x3f,
	0x48, 0xbc, 0xd3, 0xc2, 0x3b, 0xcd, 0x9d, 0xb7, 0xb5, 0xa5, 0xa1, 0x77, 0xfa, 0x29, 0x51, 0xa6,
	0x11, 0xca, 0x34, 0xbd, 0x96, 0x7f, 0xc2, 0xb5, 0xc4, 0xbf, 0x13, 0x56, 0x35, 0xbd, 0x87, 0x60,
	0xb3, 0x4c, 0x49, 0x46, 0xd8, 0x5c, 0x06, 0x9b, 0x25, 0xcd, 0x07, 0x2b, 0x0d, 0x32, 0xc2, 0x6e,
	0xdf, 0x63, 0xde, 0xe6, 0xbe, 0x4c, 0x23, 0x2f, 0x46, 0x7e, 0x80, 0x58, 0x62, 0xe5, 0xac, 0xa9,
	0x70, 0x92, 0x7f, 0x1e, 0xbd, 0xc9, 0x82, 0x31, 0xf5, 0x00, 0xe1, 0xca, 0xec, 0xd2, 0x30, 0x23,
	0x2c, 0x96, 0xed, 0x48, 0x1f, 0xe1, 0x6a, 0xe5, 0xbd, 0xf5, 0xe9, 0x3c, 0x23, 0xec, 0x46, 0xf6,
	0x87, 0xbc, 0x80, 0xe8, 0x5b, 0xe9, 0x5f, 0xac, 0x29, 0x1b, 0xb9, 0x48, 0xb7, 0xf6, 0x8e, 0xb7,
	0x7f, 0x3c, 0x6b, 0xeb, 0xd9, 0x40, 0x47, 0x5f, 0x26, 0x04, 0xc9, 0xf9, 0x6a, 0x2f, 0xae, 0x67,
	0x23, 0xc8, 0x22, 0x86, 0xc8, 0x75, 0x0f, 0x2c, 0xf8, 0xcf, 0xeb, 0xbe, 0xa8, 0x0f, 0xa7, 0x2d,
	0xd7, 0xf6, 0x28, 0x94, 0xd9, 0x35, 0x25, 0x2a, 0x27, 0x2a, 0xab, 0x0b, 0x55, 0xbe, 0x1d, 0xb1,
	0x54, 0x66, 0x8f, 0xa2, 0xeb, 0xaf, 0x6d, 0x77, 0x1b, 0x75, 0xe3, 0xc7, 0x7f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x11, 0xe1, 0x3b, 0xd0, 0x6f, 0x01, 0x00, 0x00,
}