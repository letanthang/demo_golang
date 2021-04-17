// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hellogrpc/hellogrpc.proto

package hellogrpc

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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a661547b9c73aad7, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a661547b9c73aad7, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "hellogrpc.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "hellogrpc.HelloReply")
}

func init() { proto.RegisterFile("hellogrpc/hellogrpc.proto", fileDescriptor_a661547b9c73aad7) }

var fileDescriptor_a661547b9c73aad7 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x48, 0xcd, 0xc9,
	0xc9, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38,
	0xe1, 0x02, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0x20, 0x4e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89,
	0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98,
	0xad, 0xa4, 0xc6, 0xc5, 0x05, 0x55, 0x53, 0x90, 0x53, 0x29, 0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a,
	0x5c, 0x9c, 0x98, 0x0e, 0x53, 0x04, 0xe3, 0x1a, 0x75, 0x31, 0x72, 0xb1, 0xbb, 0x17, 0xa5, 0xa6,
	0x96, 0xa4, 0x16, 0x09, 0xd9, 0x70, 0x71, 0x04, 0x27, 0x56, 0x82, 0xb5, 0x09, 0x89, 0xeb, 0x21,
	0x1c, 0x80, 0x6c, 0x99, 0x94, 0x28, 0xa6, 0x44, 0x41, 0x4e, 0xa5, 0x12, 0x83, 0x90, 0x23, 0x17,
	0x2f, 0x4c, 0xb7, 0x63, 0x7a, 0x62, 0x66, 0x1e, 0xe9, 0x46, 0x38, 0xd9, 0x72, 0x49, 0x67, 0xe6,
	0xeb, 0x81, 0xc5, 0x53, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0x8b, 0x21, 0x4a, 0xcb, 0xf3, 0x8b,
	0x72, 0x52, 0x9c, 0xf8, 0xc1, 0x8a, 0xc3, 0x41, 0xec, 0x00, 0x50, 0x98, 0x04, 0x30, 0x46, 0x71,
	0xeb, 0x21, 0x02, 0x2a, 0x89, 0x0d, 0x1c, 0x52, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31,
	0x6a, 0x07, 0xb2, 0x46, 0x01, 0x00, 0x00,
}
