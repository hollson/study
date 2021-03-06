// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package com_github_hollson_srv_user

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

// 请求
type Request struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type UserInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tel                  string   `protobuf:"bytes,3,opt,name=tel,proto3" json:"tel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

type Error struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Error) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

// 响应
type Response struct {
	Code                 uint32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                *Error    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data                 *UserInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetData() *UserInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "com.github.hollson.srv.user.Request")
	proto.RegisterType((*UserInfo)(nil), "com.github.hollson.srv.user.UserInfo")
	proto.RegisterType((*Error)(nil), "com.github.hollson.srv.user.Error")
	proto.RegisterType((*Response)(nil), "com.github.hollson.srv.user.Response")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x5d, 0xb7, 0xab, 0xeb, 0x14, 0x41, 0x02, 0x42, 0x71, 0x2f, 0x12, 0x5c, 0xf0, 0x62,
	0x84, 0x7a, 0xd1, 0xeb, 0xa2, 0xc8, 0x5e, 0x03, 0xde, 0xbc, 0x74, 0xb7, 0xb3, 0xdd, 0x42, 0x9b,
	0x59, 0x93, 0x54, 0xf0, 0x77, 0xf8, 0x87, 0x65, 0xa6, 0xf6, 0x24, 0xd4, 0x4b, 0x78, 0x99, 0x79,
	0xef, 0x4b, 0x26, 0x81, 0xcb, 0x83, 0xa7, 0x48, 0xf7, 0x5d, 0x40, 0x2f, 0x8b, 0x91, 0xbd, 0x5a,
	0x6c, 0xa9, 0x35, 0x55, 0x1d, 0xf7, 0xdd, 0xc6, 0xec, 0xa9, 0x69, 0x02, 0x39, 0x13, 0xfc, 0xa7,
	0x61, 0x8b, 0x5e, 0xc0, 0xa9, 0xc5, 0x8f, 0x0e, 0x43, 0x54, 0x17, 0x30, 0xed, 0xea, 0x32, 0x9b,
	0x5c, 0x4f, 0x6e, 0xcf, 0x2c, 0x4b, 0xbd, 0x82, 0xf9, 0x5b, 0x40, 0xbf, 0x76, 0x3b, 0xfa, 0xdb,
	0x55, 0x0a, 0x12, 0x57, 0xb4, 0x98, 0x1d, 0x4b, 0x49, 0x34, 0xbb, 0x22, 0x36, 0xd9, 0xb4, 0x77,
	0x45, 0x6c, 0xf4, 0x1d, 0xcc, 0x5e, 0xbc, 0x27, 0xcf, 0xad, 0x36, 0x54, 0x03, 0xa0, 0x0d, 0x15,
	0x03, 0x6a, 0xb7, 0xa3, 0x01, 0xc0, 0x5a, 0x7f, 0x4f, 0x60, 0x6e, 0x31, 0x1c, 0xc8, 0x05, 0x64,
	0xc3, 0x96, 0x4a, 0x94, 0xcc, 0xb9, 0x15, 0xad, 0x1e, 0x61, 0x86, 0xcc, 0x93, 0x54, 0x9a, 0x6b,
	0x33, 0x32, 0x9d, 0x91, 0x93, 0x6d, 0x1f, 0x50, 0x4f, 0x90, 0x94, 0x45, 0x2c, 0xe4, 0x72, 0x69,
	0xbe, 0x1c, 0x0d, 0x0e, 0x63, 0x5b, 0x89, 0xe4, 0x25, 0x24, 0x5c, 0x51, 0xef, 0x90, 0xbe, 0x62,
	0x64, 0xb9, 0xfa, 0x5a, 0x3f, 0xab, 0x9b, 0x51, 0xc6, 0xef, 0xbb, 0x5e, 0x2d, 0xff, 0x71, 0xf5,
	0xc3, 0xea, 0xa3, 0xcd, 0x89, 0xfc, 0xd7, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0xd9,
	0x20, 0xce, 0xc8, 0x01, 0x00, 0x00,
}
