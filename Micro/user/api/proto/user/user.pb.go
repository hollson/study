// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package com_github_hollson_api_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/micro/go-micro/api/proto"
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
	proto.RegisterType((*Request)(nil), "com.github.hollson.api.user.Request")
	proto.RegisterType((*UserInfo)(nil), "com.github.hollson.api.user.UserInfo")
	proto.RegisterType((*Error)(nil), "com.github.hollson.api.user.Error")
	proto.RegisterType((*Response)(nil), "com.github.hollson.api.user.Response")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0x4d, 0xb5, 0x4e, 0x11, 0xcb, 0x82, 0x50, 0xda, 0x8b, 0x04, 0x0b, 0x82, 0x74,
	0x03, 0xf5, 0xa2, 0xd7, 0xa2, 0x48, 0xaf, 0x0b, 0xde, 0xbc, 0xa4, 0xe9, 0x34, 0x5d, 0xc8, 0x66,
	0xe2, 0xee, 0xe6, 0xe0, 0xcf, 0x10, 0xff, 0xb0, 0xec, 0xa4, 0xc1, 0x83, 0x10, 0x2f, 0xe1, 0x65,
	0xe6, 0xbd, 0x6f, 0x26, 0x13, 0xb8, 0xae, 0x2d, 0x79, 0x4a, 0x1b, 0x87, 0x96, 0x1f, 0x92, 0xdf,
	0xc5, 0x3c, 0x27, 0x23, 0x0b, 0xed, 0x0f, 0xcd, 0x56, 0x1e, 0xa8, 0x2c, 0x1d, 0x55, 0x32, 0xab,
	0xb5, 0x0c, 0x96, 0xd9, 0xf2, 0xd8, 0xc8, 0xc9, 0xa4, 0x46, 0xe7, 0x96, 0xd2, 0x82, 0x96, 0xad,
	0xc8, 0x6a, 0x9d, 0xb6, 0xc4, 0xe0, 0x66, 0x95, 0xcc, 0xe1, 0x5c, 0xe1, 0x47, 0x83, 0xce, 0x8b,
	0x09, 0x0c, 0x1a, 0xbd, 0x9b, 0x46, 0x37, 0xd1, 0xdd, 0x85, 0x0a, 0x32, 0x59, 0xc3, 0xe8, 0xcd,
	0xa1, 0xdd, 0x54, 0x7b, 0xfa, 0xdb, 0x15, 0x02, 0xe2, 0x2a, 0x33, 0x38, 0x3d, 0xe5, 0x12, 0xeb,
	0xe0, 0xf2, 0x58, 0x4e, 0x07, 0xad, 0xcb, 0x63, 0x99, 0x2c, 0x61, 0xf8, 0x62, 0x2d, 0xd9, 0xd0,
	0x32, 0xae, 0xe8, 0x00, 0xc6, 0x15, 0x01, 0xa0, 0xab, 0x3d, 0x75, 0x80, 0xa0, 0x93, 0xef, 0x08,
	0x46, 0x0a, 0x5d, 0x4d, 0x95, 0xc3, 0x60, 0xc8, 0x69, 0x87, 0x9c, 0xb9, 0x54, 0xac, 0xc5, 0x23,
	0x0c, 0x31, 0xf0, 0x38, 0x35, 0x5e, 0x25, 0xb2, 0xe7, 0x18, 0x92, 0x27, 0xab, 0x36, 0x20, 0x9e,
	0x20, 0xde, 0x65, 0x3e, 0xe3, 0xe5, 0xc6, 0xab, 0x45, 0x6f, 0xb0, 0xfb, 0x6c, 0xc5, 0x91, 0xd5,
	0x57, 0x04, 0x71, 0x28, 0x89, 0x7b, 0x88, 0xf9, 0x1a, 0x57, 0xb2, 0x20, 0x76, 0x1f, 0x8f, 0x37,
	0x9b, 0xfc, 0x16, 0xda, 0xe5, 0x93, 0x13, 0xf1, 0x0e, 0xe3, 0x57, 0xf4, 0x21, 0xb7, 0xfe, 0xdc,
	0x3c, 0x8b, 0xdb, 0xde, 0x89, 0x1d, 0x68, 0xf1, 0x8f, 0xab, 0xa3, 0x6f, 0xcf, 0xf8, 0x07, 0x3e,
	0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x12, 0x7d, 0xab, 0x9f, 0x25, 0x02, 0x00, 0x00,
}
