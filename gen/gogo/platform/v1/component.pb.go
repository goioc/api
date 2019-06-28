// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: platform/v1/component.proto

package platformv1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Specifies one component of a multi-part email body.
type EmailMime struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	BodyRawRef           string   `protobuf:"bytes,2,opt,name=body_raw_ref,json=bodyRawRef,proto3" json:"body_raw_ref,omitempty"`
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	ContentDisposition   string   `protobuf:"bytes,4,opt,name=content_disposition,json=contentDisposition,proto3" json:"content_disposition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailMime) Reset()         { *m = EmailMime{} }
func (m *EmailMime) String() string { return proto.CompactTextString(m) }
func (*EmailMime) ProtoMessage()    {}
func (*EmailMime) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b0822eb19db8309, []int{0}
}
func (m *EmailMime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailMime.Unmarshal(m, b)
}
func (m *EmailMime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailMime.Marshal(b, m, deterministic)
}
func (m *EmailMime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailMime.Merge(m, src)
}
func (m *EmailMime) XXX_Size() int {
	return xxx_messageInfo_EmailMime.Size(m)
}
func (m *EmailMime) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailMime.DiscardUnknown(m)
}

var xxx_messageInfo_EmailMime proto.InternalMessageInfo

func (m *EmailMime) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *EmailMime) GetBodyRawRef() string {
	if m != nil {
		return m.BodyRawRef
	}
	return ""
}

func (m *EmailMime) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *EmailMime) GetContentDisposition() string {
	if m != nil {
		return m.ContentDisposition
	}
	return ""
}

func init() {
	proto.RegisterType((*EmailMime)(nil), "platform.v1.EmailMime")
}

func init() { proto.RegisterFile("platform/v1/component.proto", fileDescriptor_8b0822eb19db8309) }

var fileDescriptor_8b0822eb19db8309 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0xc8, 0x49, 0x2c,
	0x49, 0xcb, 0x2f, 0xca, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0xce, 0xcf, 0x2d, 0xc8, 0xcf, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x49, 0xea, 0x95, 0x19, 0x2a, 0x4d,
	0x65, 0xe4, 0xe2, 0x74, 0xcd, 0x4d, 0xcc, 0xcc, 0xf1, 0xcd, 0xcc, 0x4d, 0x15, 0x12, 0xe2, 0x62,
	0x49, 0xca, 0x4f, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x14, 0xb8,
	0x78, 0x40, 0x74, 0x7c, 0x51, 0x62, 0x79, 0x7c, 0x51, 0x6a, 0x9a, 0x04, 0x13, 0x58, 0x8e, 0x0b,
	0x24, 0x16, 0x94, 0x58, 0x1e, 0x94, 0x9a, 0x26, 0xa4, 0xc8, 0xc5, 0x93, 0x9c, 0x9f, 0x57, 0x92,
	0x9a, 0x57, 0x12, 0x5f, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0x0c, 0x56, 0xc1, 0x0d, 0x15, 0x0b, 0xa9,
	0x2c, 0x48, 0x15, 0xd2, 0xe7, 0x12, 0x86, 0x29, 0x49, 0xc9, 0x2c, 0x2e, 0xc8, 0x2f, 0xce, 0x2c,
	0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x01, 0xab, 0x14, 0x82, 0x4a, 0xb9, 0x20, 0x64, 0x9c, 0xb2, 0xb8,
	0x34, 0xf2, 0x8b, 0xd2, 0xf5, 0xaa, 0x52, 0xf3, 0x32, 0x4b, 0x32, 0x12, 0x8b, 0xf4, 0x32, 0xf3,
	0x93, 0xe1, 0xce, 0x4e, 0x2c, 0xc8, 0xd4, 0x43, 0xf2, 0x83, 0x13, 0x9f, 0x33, 0xcc, 0x87, 0x01,
	0x20, 0x0f, 0x06, 0x30, 0x46, 0x71, 0xc1, 0xa4, 0xcb, 0x0c, 0x17, 0x31, 0x31, 0x07, 0x44, 0x44,
	0xac, 0x62, 0xe2, 0x0e, 0x80, 0xe9, 0x08, 0x33, 0x3c, 0x85, 0xe0, 0xc5, 0x84, 0x19, 0x26, 0xb1,
	0x81, 0xc3, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x08, 0x19, 0xb2, 0x36, 0x01, 0x00,
	0x00,
}
