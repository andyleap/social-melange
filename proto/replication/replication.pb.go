// Code generated by protoc-gen-go. DO NOT EDIT.
// source: replication/replication.proto

/*
Package replication is a generated protocol buffer package.

It is generated from these files:
	replication/replication.proto

It has these top-level messages:
	SignerRequest
	SignerUpdate
	FeedRangeRequest
	FeedRangeResponse
	ItemRequest
	ItemResponse
*/
package replication

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import messages "github.com/andyleap/social-melange/proto/messages"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignerRequest struct {
	Signer *messages.Signer  `protobuf:"bytes,1,opt,name=Signer" json:"Signer,omitempty"`
	Feeds  map[string]uint64 `protobuf:"bytes,2,rep,name=Feeds" json:"Feeds,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *SignerRequest) Reset()                    { *m = SignerRequest{} }
func (m *SignerRequest) String() string            { return proto.CompactTextString(m) }
func (*SignerRequest) ProtoMessage()               {}
func (*SignerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignerRequest) GetSigner() *messages.Signer {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *SignerRequest) GetFeeds() map[string]uint64 {
	if m != nil {
		return m.Feeds
	}
	return nil
}

type SignerUpdate struct {
	Feeds map[string]uint64 `protobuf:"bytes,1,rep,name=Feeds" json:"Feeds,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *SignerUpdate) Reset()                    { *m = SignerUpdate{} }
func (m *SignerUpdate) String() string            { return proto.CompactTextString(m) }
func (*SignerUpdate) ProtoMessage()               {}
func (*SignerUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignerUpdate) GetFeeds() map[string]uint64 {
	if m != nil {
		return m.Feeds
	}
	return nil
}

type FeedRangeRequest struct {
	Feed  *messages.Feed `protobuf:"bytes,1,opt,name=Feed" json:"Feed,omitempty"`
	Start uint64         `protobuf:"varint,2,opt,name=Start" json:"Start,omitempty"`
	Stop  uint64         `protobuf:"varint,3,opt,name=Stop" json:"Stop,omitempty"`
}

func (m *FeedRangeRequest) Reset()                    { *m = FeedRangeRequest{} }
func (m *FeedRangeRequest) String() string            { return proto.CompactTextString(m) }
func (*FeedRangeRequest) ProtoMessage()               {}
func (*FeedRangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FeedRangeRequest) GetFeed() *messages.Feed {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *FeedRangeRequest) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *FeedRangeRequest) GetStop() uint64 {
	if m != nil {
		return m.Stop
	}
	return 0
}

type FeedRangeResponse struct {
	FeedChanges map[uint64]*messages.SignedFeedChange `protobuf:"bytes,1,rep,name=FeedChanges" json:"FeedChanges,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FeedRangeResponse) Reset()                    { *m = FeedRangeResponse{} }
func (m *FeedRangeResponse) String() string            { return proto.CompactTextString(m) }
func (*FeedRangeResponse) ProtoMessage()               {}
func (*FeedRangeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FeedRangeResponse) GetFeedChanges() map[uint64]*messages.SignedFeedChange {
	if m != nil {
		return m.FeedChanges
	}
	return nil
}

type ItemRequest struct {
	Ref *messages.Ref `protobuf:"bytes,1,opt,name=Ref" json:"Ref,omitempty"`
}

func (m *ItemRequest) Reset()                    { *m = ItemRequest{} }
func (m *ItemRequest) String() string            { return proto.CompactTextString(m) }
func (*ItemRequest) ProtoMessage()               {}
func (*ItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ItemRequest) GetRef() *messages.Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

type ItemResponse struct {
	Item *messages.SignedItem `protobuf:"bytes,1,opt,name=Item" json:"Item,omitempty"`
}

func (m *ItemResponse) Reset()                    { *m = ItemResponse{} }
func (m *ItemResponse) String() string            { return proto.CompactTextString(m) }
func (*ItemResponse) ProtoMessage()               {}
func (*ItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ItemResponse) GetItem() *messages.SignedItem {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterType((*SignerRequest)(nil), "replication.SignerRequest")
	proto.RegisterType((*SignerUpdate)(nil), "replication.SignerUpdate")
	proto.RegisterType((*FeedRangeRequest)(nil), "replication.FeedRangeRequest")
	proto.RegisterType((*FeedRangeResponse)(nil), "replication.FeedRangeResponse")
	proto.RegisterType((*ItemRequest)(nil), "replication.ItemRequest")
	proto.RegisterType((*ItemResponse)(nil), "replication.ItemResponse")
}

func init() { proto.RegisterFile("replication/replication.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0x65, 0x9a, 0xb4, 0xe0, 0x4d, 0x2b, 0x75, 0x28, 0x18, 0x02, 0x62, 0x09, 0x0a, 0xd9, 0x98,
	0x48, 0x15, 0x29, 0x75, 0xa7, 0x28, 0xb8, 0x74, 0x8a, 0x9b, 0xae, 0x9c, 0x36, 0xb7, 0x69, 0x30,
	0x7f, 0x66, 0xa6, 0x42, 0xf7, 0x3e, 0x90, 0xcf, 0xe0, 0x93, 0xc9, 0xcc, 0x24, 0x66, 0xda, 0xef,
	0x5b, 0x7e, 0xbb, 0x3b, 0x67, 0xce, 0x39, 0x73, 0xee, 0x21, 0x81, 0x67, 0x2d, 0x36, 0x45, 0x7e,
	0xe0, 0x32, 0xaf, 0xab, 0xc4, 0x9a, 0xe3, 0xa6, 0xad, 0x65, 0x4d, 0x3d, 0x0b, 0x0a, 0x9e, 0x96,
	0x28, 0x04, 0xcf, 0x50, 0x24, 0xfd, 0x60, 0x58, 0xe1, 0x1f, 0x02, 0xb3, 0x6d, 0x9e, 0x55, 0xd8,
	0x32, 0xfc, 0x79, 0x46, 0x21, 0x69, 0x04, 0x13, 0x03, 0xf8, 0x64, 0x49, 0x22, 0x6f, 0x35, 0x8f,
	0xff, 0x4b, 0x3a, 0x62, 0x77, 0x4f, 0xdf, 0xc3, 0xf8, 0x33, 0x62, 0x2a, 0xfc, 0xd1, 0xd2, 0x89,
	0xbc, 0xd5, 0xcb, 0xd8, 0x0e, 0x71, 0x65, 0x1a, 0x6b, 0xde, 0xa7, 0x4a, 0xb6, 0x17, 0x66, 0x34,
	0xc1, 0x1a, 0x60, 0x00, 0xe9, 0x1c, 0x9c, 0x1f, 0x78, 0xd1, 0x2f, 0x3e, 0x62, 0x6a, 0xa4, 0x0b,
	0x18, 0xff, 0xe2, 0xc5, 0x19, 0xfd, 0xd1, 0x92, 0x44, 0x2e, 0x33, 0x87, 0xcd, 0x68, 0x4d, 0xc2,
	0xdf, 0x04, 0xa6, 0xc6, 0xfd, 0x5b, 0x93, 0x72, 0x89, 0x74, 0xd3, 0xe7, 0x20, 0x3a, 0xc7, 0x8b,
	0x7b, 0x72, 0x18, 0xe6, 0x83, 0xc6, 0xf8, 0x0e, 0x73, 0xa5, 0x64, 0xbc, 0xca, 0xb0, 0xef, 0x2e,
	0x04, 0x57, 0x61, 0x5d, 0x73, 0x8f, 0x87, 0xe6, 0x34, 0x53, 0xdf, 0x29, 0xc7, 0xad, 0xe4, 0xad,
	0xec, 0x1d, 0xf5, 0x81, 0x52, 0x70, 0xb7, 0xb2, 0x6e, 0x7c, 0x47, 0x83, 0x7a, 0x0e, 0xff, 0x12,
	0x78, 0x62, 0x3d, 0x21, 0x9a, 0xba, 0x12, 0x48, 0xbf, 0x82, 0xa7, 0xc0, 0x8f, 0x27, 0x85, 0xf6,
	0x3b, 0x27, 0x57, 0x3b, 0xdf, 0x11, 0xc5, 0x96, 0xc2, 0xac, 0x6f, 0x7b, 0x04, 0x3b, 0xb3, 0x8a,
	0x4d, 0xb0, 0xab, 0x70, 0x4d, 0x15, 0xaf, 0xed, 0x2a, 0xbc, 0x55, 0x70, 0xf3, 0x5d, 0xa4, 0x83,
	0x85, 0x5d, 0x53, 0x0c, 0xde, 0x17, 0x89, 0x65, 0xdf, 0xd0, 0x73, 0x70, 0x18, 0x1e, 0xbb, 0x82,
	0x66, 0x83, 0x05, 0xc3, 0x23, 0x53, 0x37, 0xe1, 0x1a, 0xa6, 0x86, 0xdf, 0xad, 0x1b, 0x81, 0xab,
	0xce, 0x9d, 0x62, 0x71, 0xfb, 0xa8, 0xe6, 0x6a, 0xc6, 0x87, 0x77, 0xbb, 0xb7, 0x59, 0x2e, 0x4f,
	0xe7, 0x7d, 0x7c, 0xa8, 0xcb, 0x84, 0x57, 0xe9, 0xa5, 0x40, 0xde, 0x24, 0xa2, 0x3e, 0xe4, 0xbc,
	0x78, 0x55, 0x62, 0xa1, 0xa2, 0x25, 0xfa, 0xb3, 0xb7, 0x7f, 0x97, 0xfd, 0x44, 0x43, 0x6f, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x17, 0x62, 0x8f, 0x27, 0x50, 0x03, 0x00, 0x00,
}