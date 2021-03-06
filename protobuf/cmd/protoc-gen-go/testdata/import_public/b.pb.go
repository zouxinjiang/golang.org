// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/b.proto

package import_public

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	sub "github.com/golang/protobuf/v2/cmd/protoc-gen-go/testdata/import_public/sub"
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

type Local struct {
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	E                    *sub.E   `protobuf:"varint,2,opt,name=e,enum=goproto.protoc.import_public.sub.E" json:"e,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Local) Reset()         { *m = Local{} }
func (m *Local) String() string { return proto.CompactTextString(m) }
func (*Local) ProtoMessage()    {}
func (*Local) Descriptor() ([]byte, []int) {
	return fileDescriptor_84995586b3d09710, []int{0}
}

func (m *Local) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Local.Unmarshal(m, b)
}
func (m *Local) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Local.Marshal(b, m, deterministic)
}
func (m *Local) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Local.Merge(m, src)
}
func (m *Local) XXX_Size() int {
	return xxx_messageInfo_Local.Size(m)
}
func (m *Local) XXX_DiscardUnknown() {
	xxx_messageInfo_Local.DiscardUnknown(m)
}

var xxx_messageInfo_Local proto.InternalMessageInfo

func (m *Local) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Local) GetE() sub.E {
	if m != nil && m.E != nil {
		return *m.E
	}
	return sub.E_ZERO
}

func init() {
	proto.RegisterType((*Local)(nil), "goproto.protoc.import_public.Local")
}

func init() { proto.RegisterFile("import_public/b.proto", fileDescriptor_84995586b3d09710) }

var fileDescriptor_84995586b3d09710 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xd6, 0x4f, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x49, 0xcf, 0x07, 0x33, 0x20, 0xdc, 0x64, 0x3d, 0x14, 0x55, 0x52, 0x92, 0xa8,
	0x9a, 0x8a, 0x4b, 0x93, 0xf4, 0x13, 0x21, 0x2a, 0x95, 0x72, 0xb9, 0x58, 0x7d, 0xf2, 0x93, 0x13,
	0x73, 0x84, 0x0c, 0xb9, 0x18, 0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x94, 0xf5, 0xf0,
	0x99, 0xa6, 0x57, 0x5c, 0x9a, 0xa4, 0xe7, 0x1b, 0xc4, 0x98, 0x0b, 0xd2, 0x92, 0x2a, 0xc1, 0xa4,
	0xc0, 0xa8, 0xc1, 0x47, 0x8c, 0x16, 0xd7, 0x20, 0xc6, 0x54, 0x27, 0x8f, 0x28, 0xb7, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf4, 0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0x7d,
	0xb0, 0x96, 0xa4, 0xd2, 0x34, 0xfd, 0x32, 0x23, 0xfd, 0xe4, 0xdc, 0x14, 0x08, 0x3f, 0x59, 0x37,
	0x3d, 0x35, 0x4f, 0x37, 0x3d, 0x5f, 0xbf, 0x24, 0xb5, 0xb8, 0x24, 0x25, 0xb1, 0x24, 0x51, 0x1f,
	0xc5, 0x48, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x34, 0xb9, 0xda, 0x09, 0x01, 0x00, 0x00,
}
