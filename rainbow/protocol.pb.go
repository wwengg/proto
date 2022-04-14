// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rainbow/protocol.proto

package rainbow

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

type HttpReply struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpReply) Reset()         { *m = HttpReply{} }
func (m *HttpReply) String() string { return proto.CompactTextString(m) }
func (*HttpReply) ProtoMessage()    {}
func (*HttpReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e043f27103e262a, []int{0}
}

func (m *HttpReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpReply.Unmarshal(m, b)
}
func (m *HttpReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpReply.Marshal(b, m, deterministic)
}
func (m *HttpReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpReply.Merge(m, src)
}
func (m *HttpReply) XXX_Size() int {
	return xxx_messageInfo_HttpReply.Size(m)
}
func (m *HttpReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpReply.DiscardUnknown(m)
}

var xxx_messageInfo_HttpReply proto.InternalMessageInfo

func (m *HttpReply) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HttpReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HttpReply) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type HttpArgs struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	UId                  uint64   `protobuf:"varint,4,opt,name=uId,proto3" json:"uId,omitempty"`
	Sign                 string   `protobuf:"bytes,5,opt,name=sign,proto3" json:"sign,omitempty"`
	Data                 []byte   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpArgs) Reset()         { *m = HttpArgs{} }
func (m *HttpArgs) String() string { return proto.CompactTextString(m) }
func (*HttpArgs) ProtoMessage()    {}
func (*HttpArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e043f27103e262a, []int{1}
}

func (m *HttpArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpArgs.Unmarshal(m, b)
}
func (m *HttpArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpArgs.Marshal(b, m, deterministic)
}
func (m *HttpArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpArgs.Merge(m, src)
}
func (m *HttpArgs) XXX_Size() int {
	return xxx_messageInfo_HttpArgs.Size(m)
}
func (m *HttpArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpArgs.DiscardUnknown(m)
}

var xxx_messageInfo_HttpArgs proto.InternalMessageInfo

func (m *HttpArgs) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *HttpArgs) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpArgs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *HttpArgs) GetUId() uint64 {
	if m != nil {
		return m.UId
	}
	return 0
}

func (m *HttpArgs) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *HttpArgs) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpReply)(nil), "proto.protocol.HttpReply")
	proto.RegisterType((*HttpArgs)(nil), "proto.protocol.HttpArgs")
}

func init() {
	proto.RegisterFile("rainbow/protocol.proto", fileDescriptor_3e043f27103e262a)
}

var fileDescriptor_3e043f27103e262a = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x3d, 0x4e, 0xc4, 0x30,
	0x10, 0x85, 0x65, 0x36, 0x1b, 0xc8, 0x08, 0x10, 0xb2, 0xd0, 0xca, 0x1d, 0x61, 0xab, 0x54, 0x49,
	0xc1, 0x09, 0xa0, 0x82, 0x82, 0xc6, 0x25, 0x5d, 0x7e, 0x46, 0xde, 0x88, 0x8d, 0x27, 0xb2, 0x1d,
	0x22, 0xee, 0xc0, 0xa1, 0x91, 0x27, 0x09, 0x54, 0xfe, 0x3e, 0x6b, 0xfc, 0x3c, 0x7a, 0x70, 0x70,
	0x75, 0x6f, 0x1b, 0x9a, 0xab, 0xd1, 0x51, 0xa0, 0x96, 0xce, 0x25, 0x83, 0xbc, 0xe5, 0xa3, 0xdc,
	0x6e, 0x8f, 0xef, 0x90, 0xbd, 0x86, 0x30, 0x6a, 0x1c, 0xcf, 0xdf, 0x52, 0x42, 0xd2, 0x52, 0x87,
	0x4a, 0xe4, 0xa2, 0xb8, 0xd1, 0xcc, 0x52, 0xc1, 0xe5, 0x80, 0xde, 0xd7, 0x06, 0xd5, 0x45, 0x2e,
	0x8a, 0x4c, 0x6f, 0x1a, 0xa7, 0xbb, 0x3a, 0xd4, 0x6a, 0x97, 0x8b, 0xe2, 0x5a, 0x33, 0x1f, 0x7f,
	0x04, 0x5c, 0xc5, 0xbc, 0x67, 0x67, 0x7c, 0x7c, 0xea, 0xd1, 0x7d, 0xf5, 0xed, 0x92, 0x98, 0xe9,
	0x4d, 0xe5, 0x01, 0xd2, 0x01, 0xc3, 0x89, 0xba, 0x35, 0x73, 0x35, 0x79, 0x0f, 0xfb, 0x40, 0x9f,
	0x68, 0x39, 0x33, 0xd3, 0x8b, 0xc8, 0x3b, 0xd8, 0x4d, 0x6f, 0x9d, 0x4a, 0x72, 0x51, 0x24, 0x3a,
	0x62, 0xfc, 0xda, 0xf7, 0xc6, 0xaa, 0x3d, 0x8f, 0x31, 0xff, 0xad, 0x93, 0xfe, 0xaf, 0xf3, 0xf2,
	0xf8, 0xf1, 0x60, 0xfa, 0x70, 0x9a, 0x9a, 0xb2, 0xa5, 0xa1, 0x9a, 0x67, 0xb4, 0xc6, 0x2c, 0x8d,
	0x54, 0x6b, 0x3f, 0x4d, 0xca, 0xfa, 0xf4, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x87, 0x92, 0x0d,
	0x31, 0x01, 0x00, 0x00,
}