// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3/enum.proto

package proto3

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Enum int32

const (
	Enum_ZERO Enum = 0
	Enum_ONE  Enum = 1
	Enum_TWO  Enum = 2
)

var Enum_name = map[int32]string{
	0: "ZERO",
	1: "ONE",
	2: "TWO",
}

var Enum_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
	"TWO":  2,
}

func (x Enum) String() string {
	return proto.EnumName(Enum_name, int32(x))
}

func (Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4b9b1e8d161a9a6, []int{0}
}

func init() {
	proto.RegisterEnum("goproto.protoc.proto3.Enum", Enum_name, Enum_value)
}

func init() { proto.RegisterFile("proto3/enum.proto", fileDescriptor_b4b9b1e8d161a9a6) }

var fileDescriptor_b4b9b1e8d161a9a6 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd6, 0x4f, 0xcd, 0x2b, 0xcd, 0xd5, 0x03, 0xb3, 0x85, 0x44, 0xd3, 0xf3, 0xc1, 0x0c,
	0x08, 0x37, 0x19, 0x42, 0x19, 0x6b, 0x29, 0x71, 0xb1, 0xb8, 0xe6, 0x95, 0xe6, 0x0a, 0x71, 0x70,
	0xb1, 0x44, 0xb9, 0x06, 0xf9, 0x0b, 0x30, 0x08, 0xb1, 0x73, 0x31, 0xfb, 0xfb, 0xb9, 0x0a, 0x30,
	0x82, 0x18, 0x21, 0xe1, 0xfe, 0x02, 0x4c, 0x4e, 0x8e, 0x51, 0xf6, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0xfa, 0x60, 0xfd, 0x49,
	0xa5, 0x69, 0xfa, 0x65, 0x46, 0xfa, 0xc9, 0xb9, 0x29, 0x10, 0x7e, 0xb2, 0x6e, 0x7a, 0x6a, 0x9e,
	0x6e, 0x7a, 0xbe, 0x7e, 0x49, 0x6a, 0x71, 0x49, 0x4a, 0x62, 0x49, 0x22, 0x44, 0xd8, 0x38, 0x89,
	0x0d, 0x42, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x96, 0x4a, 0x7d, 0xc7, 0x99, 0x00, 0x00, 0x00,
}
