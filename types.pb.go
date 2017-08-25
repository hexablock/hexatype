// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

/*
Package hexatype is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	Entry
	KeylogIndex
	ReqResp
	RequestOptions
	KeyValuePair
*/
package hexatype

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hexaring "github.com/hexablock/hexaring"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Hexalog entry
type Entry struct {
	Previous  []byte `protobuf:"bytes,1,opt,name=Previous,json=previous,proto3" json:"Previous,omitempty"`
	Timestamp uint64 `protobuf:"varint,2,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
	Height    uint32 `protobuf:"varint,3,opt,name=Height,json=height" json:"Height,omitempty"`
	Key       []byte `protobuf:"bytes,4,opt,name=Key,json=key,proto3" json:"Key,omitempty"`
	Data      []byte `protobuf:"bytes,5,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

func (m *Entry) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Entry) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Entry) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type KeylogIndex struct {
	Key     []byte   `protobuf:"bytes,1,opt,name=Key,json=key,proto3" json:"Key,omitempty"`
	Height  uint32   `protobuf:"varint,2,opt,name=Height,json=height" json:"Height,omitempty"`
	Marker  []byte   `protobuf:"bytes,3,opt,name=Marker,json=marker,proto3" json:"Marker,omitempty"`
	Entries [][]byte `protobuf:"bytes,4,rep,name=Entries,json=entries,proto3" json:"Entries,omitempty"`
}

func (m *KeylogIndex) Reset()                    { *m = KeylogIndex{} }
func (m *KeylogIndex) String() string            { return proto.CompactTextString(m) }
func (*KeylogIndex) ProtoMessage()               {}
func (*KeylogIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeylogIndex) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeylogIndex) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *KeylogIndex) GetMarker() []byte {
	if m != nil {
		return m.Marker
	}
	return nil
}

func (m *KeylogIndex) GetEntries() [][]byte {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Request and response shared structure for hexalog
type ReqResp struct {
	// ID is based on the request/response
	ID      []byte          `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Entry   *Entry          `protobuf:"bytes,2,opt,name=Entry,json=entry" json:"Entry,omitempty"`
	Options *RequestOptions `protobuf:"bytes,3,opt,name=Options,json=options" json:"Options,omitempty"`
}

func (m *ReqResp) Reset()                    { *m = ReqResp{} }
func (m *ReqResp) String() string            { return proto.CompactTextString(m) }
func (*ReqResp) ProtoMessage()               {}
func (*ReqResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReqResp) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *ReqResp) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *ReqResp) GetOptions() *RequestOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// Hexalog request options
type RequestOptions struct {
	// Index of the source in the PeerSet.  This is set internally by the log
	SourceIndex int32 `protobuf:"varint,1,opt,name=SourceIndex,json=sourceIndex" json:"SourceIndex,omitempty"`
	// Set of peers for the request.
	PeerSet []*hexaring.Location `protobuf:"bytes,2,rep,name=PeerSet,json=peerSet" json:"PeerSet,omitempty"`
	// Number of times to retry for a given request
	Retries int32 `protobuf:"varint,3,opt,name=Retries,json=retries" json:"Retries,omitempty"`
}

func (m *RequestOptions) Reset()                    { *m = RequestOptions{} }
func (m *RequestOptions) String() string            { return proto.CompactTextString(m) }
func (*RequestOptions) ProtoMessage()               {}
func (*RequestOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestOptions) GetSourceIndex() int32 {
	if m != nil {
		return m.SourceIndex
	}
	return 0
}

func (m *RequestOptions) GetPeerSet() []*hexaring.Location {
	if m != nil {
		return m.PeerSet
	}
	return nil
}

func (m *RequestOptions) GetRetries() int32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

type KeyValuePair struct {
	Key   []byte `protobuf:"bytes,1,opt,name=Key,json=key,proto3" json:"Key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
	Entry *Entry `protobuf:"bytes,3,opt,name=Entry,json=entry" json:"Entry,omitempty"`
}

func (m *KeyValuePair) Reset()                    { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string            { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()               {}
func (*KeyValuePair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *KeyValuePair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValuePair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyValuePair) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func init() {
	proto.RegisterType((*Entry)(nil), "hexatype.Entry")
	proto.RegisterType((*KeylogIndex)(nil), "hexatype.KeylogIndex")
	proto.RegisterType((*ReqResp)(nil), "hexatype.ReqResp")
	proto.RegisterType((*RequestOptions)(nil), "hexatype.RequestOptions")
	proto.RegisterType((*KeyValuePair)(nil), "hexatype.KeyValuePair")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0xf1, 0x87, 0xd2, 0x6b, 0xaf, 0x1b, 0x62, 0x14, 0x51, 0xf6, 0x60, 0x0c, 0x03, 0xc3,
	0x86, 0x0b, 0xd9, 0x5f, 0xc8, 0x60, 0x25, 0x1b, 0x0b, 0xea, 0xd8, 0xdb, 0x1e, 0x14, 0xf7, 0x92,
	0x88, 0x7c, 0xc8, 0x95, 0xe4, 0x52, 0xef, 0x65, 0x7f, 0x7d, 0xf8, 0xda, 0xc1, 0x04, 0xc6, 0xde,
	0x74, 0xce, 0xd5, 0x3d, 0xc7, 0xe7, 0x58, 0x90, 0xfa, 0xae, 0x41, 0x57, 0x35, 0xd6, 0x78, 0xc3,
	0xe7, 0x3b, 0x7c, 0x51, 0x3d, 0x71, 0xfb, 0x61, 0xab, 0xfd, 0xae, 0xdd, 0x54, 0xb5, 0x39, 0xde,
	0xf5, 0xe4, 0xe6, 0x60, 0xea, 0x3d, 0x9d, 0xac, 0x3e, 0x6d, 0xef, 0x9c, 0xb7, 0x6d, 0xed, 0xc7,
	0xb5, 0xe2, 0x0f, 0xc4, 0x9f, 0x4f, 0xde, 0x76, 0xfc, 0x16, 0xe6, 0x6b, 0x8b, 0xcf, 0xda, 0xb4,
	0x4e, 0x04, 0x79, 0x50, 0x66, 0x72, 0xde, 0x8c, 0x98, 0xbf, 0x83, 0xab, 0x1f, 0xfa, 0x88, 0xce,
	0xab, 0x63, 0x23, 0x66, 0x79, 0x50, 0x46, 0xf2, 0xca, 0x9f, 0x09, 0x7e, 0x03, 0xc9, 0x17, 0xd4,
	0xdb, 0x9d, 0x17, 0x61, 0x1e, 0x94, 0xaf, 0x64, 0xb2, 0x23, 0xc4, 0xdf, 0x40, 0xb8, 0xc2, 0x4e,
	0x44, 0x24, 0x16, 0xee, 0xb1, 0xe3, 0x1c, 0xa2, 0xa5, 0xf2, 0x4a, 0xc4, 0x44, 0x45, 0x8f, 0xca,
	0xab, 0x42, 0x43, 0xba, 0xc2, 0xee, 0x60, 0xb6, 0xf7, 0xa7, 0x47, 0x7c, 0x39, 0x2f, 0x05, 0xd3,
	0xd2, 0x24, 0x3f, 0xbb, 0x90, 0xbf, 0x81, 0xe4, 0x9b, 0xb2, 0x7b, 0xb4, 0x64, 0x9b, 0xc9, 0xe4,
	0x48, 0x88, 0x0b, 0x60, 0x7d, 0x22, 0x8d, 0x4e, 0x44, 0x79, 0x58, 0x66, 0x92, 0xe1, 0x00, 0x0b,
	0x0f, 0x4c, 0xe2, 0x93, 0x44, 0xd7, 0xf0, 0x6b, 0x98, 0xdd, 0x2f, 0x47, 0x97, 0x99, 0x5e, 0xf2,
	0xf7, 0x63, 0x0d, 0xe4, 0x91, 0x2e, 0x5e, 0x57, 0xe7, 0x36, 0x2b, 0xa2, 0x65, 0x8c, 0x54, 0xd2,
	0x02, 0xd8, 0xf7, 0xc6, 0x6b, 0x73, 0x72, 0x64, 0x9a, 0x2e, 0xc4, 0x74, 0x51, 0xe2, 0x53, 0x8b,
	0xce, 0x8f, 0x73, 0xc9, 0xcc, 0x70, 0x28, 0x7e, 0xc3, 0xf5, 0xe5, 0x88, 0xe7, 0x90, 0x3e, 0x98,
	0xd6, 0xd6, 0x48, 0x91, 0xe9, 0x2b, 0x62, 0x99, 0xba, 0x89, 0xe2, 0x1f, 0x81, 0xad, 0x11, 0xed,
	0x03, 0xf6, 0xa1, 0xc3, 0x32, 0x5d, 0xf0, 0xea, 0xfc, 0xff, 0xaa, 0xaf, 0xa6, 0x56, 0xbd, 0x8e,
	0x64, 0xcd, 0x70, 0xa5, 0x4f, 0x2c, 0x71, 0x48, 0x1c, 0x92, 0x16, 0xb3, 0x03, 0x2c, 0x7e, 0x41,
	0xb6, 0xc2, 0xee, 0xa7, 0x3a, 0xb4, 0xb8, 0x56, 0xda, 0xfe, 0xa3, 0xdd, 0xb7, 0x10, 0xd3, 0x98,
	0x82, 0x67, 0x32, 0x7e, 0xee, 0xc1, 0x54, 0x47, 0xf8, 0xbf, 0x3a, 0x36, 0x09, 0xbd, 0xa1, 0x4f,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x38, 0x94, 0x24, 0xcc, 0x89, 0x02, 0x00, 0x00,
}
