// Code generated by protoc-gen-go. DO NOT EDIT.
// source: BIT_STRING.proto

package e2ctypes

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

type BIT_STRING struct {
	BitString            []byte   `protobuf:"bytes,1,opt,name=bitString,proto3" json:"bitString,omitempty"`
	Numbits              uint32   `protobuf:"varint,2,opt,name=numbits,proto3" json:"numbits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BIT_STRING) Reset()         { *m = BIT_STRING{} }
func (m *BIT_STRING) String() string { return proto.CompactTextString(m) }
func (*BIT_STRING) ProtoMessage()    {}
func (*BIT_STRING) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b06a613d6c54766, []int{0}
}

func (m *BIT_STRING) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BIT_STRING.Unmarshal(m, b)
}
func (m *BIT_STRING) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BIT_STRING.Marshal(b, m, deterministic)
}
func (m *BIT_STRING) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BIT_STRING.Merge(m, src)
}
func (m *BIT_STRING) XXX_Size() int {
	return xxx_messageInfo_BIT_STRING.Size(m)
}
func (m *BIT_STRING) XXX_DiscardUnknown() {
	xxx_messageInfo_BIT_STRING.DiscardUnknown(m)
}

var xxx_messageInfo_BIT_STRING proto.InternalMessageInfo

func (m *BIT_STRING) GetBitString() []byte {
	if m != nil {
		return m.BitString
	}
	return nil
}

func (m *BIT_STRING) GetNumbits() uint32 {
	if m != nil {
		return m.Numbits
	}
	return 0
}

func init() {
	proto.RegisterType((*BIT_STRING)(nil), "e2ctypes.BIT_STRING")
}

func init() { proto.RegisterFile("BIT_STRING.proto", fileDescriptor_3b06a613d6c54766) }

var fileDescriptor_3b06a613d6c54766 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x70, 0xf2, 0x0c, 0x89,
	0x0f, 0x0e, 0x09, 0xf2, 0xf4, 0x73, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x35,
	0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x72, 0xe1, 0xe2, 0x42, 0xc8, 0x0a, 0xc9, 0x70, 0x71,
	0x26, 0x65, 0x96, 0x04, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04,
	0x21, 0x04, 0x84, 0x24, 0xb8, 0xd8, 0xf3, 0x4a, 0x73, 0x93, 0x32, 0x4b, 0x8a, 0x25, 0x98, 0x14,
	0x18, 0x35, 0x78, 0x83, 0x60, 0xdc, 0x24, 0x36, 0xb0, 0xb1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8e, 0x4b, 0x25, 0x06, 0x6a, 0x00, 0x00, 0x00,
}