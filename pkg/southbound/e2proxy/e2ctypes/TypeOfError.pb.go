// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TypeOfError.proto

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

type TypeOfErrorT int32

const (
	TypeOfErrorT_TypeOfError_not_understood TypeOfErrorT = 0
	TypeOfErrorT_TypeOfError_missing        TypeOfErrorT = 1
)

var TypeOfErrorT_name = map[int32]string{
	0: "TypeOfError_not_understood",
	1: "TypeOfError_missing",
}

var TypeOfErrorT_value = map[string]int32{
	"TypeOfError_not_understood": 0,
	"TypeOfError_missing":        1,
}

func (x TypeOfErrorT) String() string {
	return proto.EnumName(TypeOfErrorT_name, int32(x))
}

func (TypeOfErrorT) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6ba9b06385dc0771, []int{0}
}

func init() {
	proto.RegisterEnum("e2ctypes.TypeOfErrorT", TypeOfErrorT_name, TypeOfErrorT_value)
}

func init() { proto.RegisterFile("TypeOfError.proto", fileDescriptor_6ba9b06385dc0771) }

var fileDescriptor_6ba9b06385dc0771 = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0c, 0xa9, 0x2c, 0x48,
	0xf5, 0x4f, 0x73, 0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48,
	0x35, 0x4a, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0xf2, 0xe0, 0xe2, 0x45, 0x92, 0x8e, 0x2f, 0x11,
	0x92, 0xe3, 0x92, 0x42, 0x16, 0xc8, 0xcb, 0x2f, 0x89, 0x2f, 0xcd, 0x4b, 0x49, 0x2d, 0x2a, 0x2e,
	0xc9, 0xcf, 0x4f, 0x11, 0x60, 0x10, 0x12, 0xe7, 0x12, 0x46, 0x96, 0xcf, 0xcd, 0x2c, 0x2e, 0xce,
	0xcc, 0x4b, 0x17, 0x60, 0x4c, 0x62, 0x03, 0x1b, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x8f, 0xea, 0x7e, 0x6f, 0x00, 0x00, 0x00,
}