// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/slot/slot.proto

package slot

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

type Slot struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Slot) Reset()         { *m = Slot{} }
func (m *Slot) String() string { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()    {}
func (*Slot) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc39ccfcc247c8d, []int{0}
}

func (m *Slot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slot.Unmarshal(m, b)
}
func (m *Slot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slot.Marshal(b, m, deterministic)
}
func (m *Slot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slot.Merge(m, src)
}
func (m *Slot) XXX_Size() int {
	return xxx_messageInfo_Slot.Size(m)
}
func (m *Slot) XXX_DiscardUnknown() {
	xxx_messageInfo_Slot.DiscardUnknown(m)
}

var xxx_messageInfo_Slot proto.InternalMessageInfo

func (m *Slot) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Slot)(nil), "slot.Slot")
}

func init() { proto.RegisterFile("proto/slot/slot.proto", fileDescriptor_7dc39ccfcc247c8d) }

var fileDescriptor_7dc39ccfcc247c8d = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0xc9, 0x2f, 0x01, 0x13, 0x7a, 0x60, 0xbe, 0x10, 0x0b, 0x88, 0xad, 0x24,
	0xc6, 0xc5, 0x12, 0x9c, 0x93, 0x5f, 0x22, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x1c, 0xc4, 0x94, 0x99, 0xe2, 0x64, 0x1e, 0x65, 0x9a, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9a, 0x9e, 0x5f, 0x64, 0x68, 0x6c, 0x62, 0xa2, 0x9f, 0x94,
	0x98, 0x97, 0x97, 0x5a, 0xa4, 0x5f, 0x94, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x17, 0x0f, 0xe5,
	0x23, 0x6c, 0x48, 0x62, 0x03, 0xb3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x54, 0xa6,
	0x80, 0x76, 0x00, 0x00, 0x00,
}
