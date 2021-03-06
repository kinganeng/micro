// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/api/proto/account/account.proto

/*
Package go_micro_api_account is a generated protocol buffer package.

It is generated from these files:
	account/api/proto/account/account.proto

It has these top-level messages:
	ReqLogin
	ReqRegister
	Rsp
*/
package go_micro_api_account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/micro/go-api/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ReqLogin struct {
	Nickname string `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
	Pwd      string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *ReqLogin) Reset()                    { *m = ReqLogin{} }
func (m *ReqLogin) String() string            { return proto.CompactTextString(m) }
func (*ReqLogin) ProtoMessage()               {}
func (*ReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqLogin) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *ReqLogin) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type ReqRegister struct {
	Nickname string `protobuf:"bytes,1,opt,name=nickname" json:"nickname,omitempty"`
	Pwd      string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	PwdRe    string `protobuf:"bytes,3,opt,name=pwdRe" json:"pwdRe,omitempty"`
}

func (m *ReqRegister) Reset()                    { *m = ReqRegister{} }
func (m *ReqRegister) String() string            { return proto.CompactTextString(m) }
func (*ReqRegister) ProtoMessage()               {}
func (*ReqRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReqRegister) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *ReqRegister) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *ReqRegister) GetPwdRe() string {
	if m != nil {
		return m.PwdRe
	}
	return ""
}

type Rsp struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *Rsp) Reset()                    { *m = Rsp{} }
func (m *Rsp) String() string            { return proto.CompactTextString(m) }
func (*Rsp) ProtoMessage()               {}
func (*Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Rsp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Rsp) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Rsp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqLogin)(nil), "go.micro.api.account.ReqLogin")
	proto.RegisterType((*ReqRegister)(nil), "go.micro.api.account.ReqRegister")
	proto.RegisterType((*Rsp)(nil), "go.micro.api.account.Rsp")
}

func init() { proto.RegisterFile("account/api/proto/account/account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x69, 0xa2, 0x40, 0x39, 0x24, 0xa8, 0xac, 0x0c, 0x51, 0x26, 0x94, 0xa5, 0x0c, 0x60,
	0x4b, 0xb0, 0xb0, 0x32, 0xb1, 0xb0, 0xe0, 0x7f, 0x90, 0x26, 0x27, 0xf7, 0x54, 0xc5, 0xe7, 0xc4,
	0x8e, 0xfa, 0xf7, 0x51, 0xed, 0x52, 0x3e, 0x26, 0x98, 0xec, 0x7b, 0x4e, 0xef, 0x23, 0xdd, 0x0b,
	0xeb, 0xb6, 0xeb, 0x78, 0xb6, 0x41, 0xb5, 0x8e, 0x94, 0x9b, 0x38, 0xb0, 0x3a, 0x91, 0xf4, 0xca,
	0x48, 0x45, 0x69, 0x58, 0x0e, 0xd4, 0x4d, 0x2c, 0x5b, 0x47, 0xf2, 0xb8, 0xab, 0xd7, 0x86, 0xc2,
	0x76, 0xde, 0xc8, 0x8e, 0x07, 0x15, 0xb7, 0xca, 0xf0, 0xc3, 0x37, 0x95, 0xa3, 0x14, 0x6f, 0x9e,
	0x61, 0xa9, 0x71, 0x7c, 0x63, 0x43, 0x56, 0xd4, 0xb0, 0xb4, 0xd4, 0xed, 0x6c, 0x3b, 0x60, 0xb5,
	0xb8, 0x5d, 0xdc, 0x5d, 0xea, 0xd3, 0x2c, 0x56, 0x90, 0xbb, 0x7d, 0x5f, 0x65, 0x11, 0x1f, 0xbe,
	0xcd, 0x3b, 0x5c, 0x69, 0x1c, 0x35, 0x1a, 0xf2, 0x01, 0xa7, 0xff, 0x85, 0x45, 0x09, 0x85, 0xdb,
	0xf7, 0x1a, 0xab, 0x3c, 0xb2, 0x34, 0x34, 0xaf, 0x90, 0x6b, 0xef, 0xc4, 0x35, 0x64, 0xd4, 0x47,
	0x49, 0xae, 0x33, 0xea, 0x7f, 0xa8, 0xb3, 0x5f, 0xea, 0x12, 0x8a, 0xc0, 0x3b, 0xb4, 0x9f, 0xa2,
	0x38, 0x3c, 0x6e, 0xe1, 0xe2, 0x25, 0x35, 0x21, 0xee, 0xa1, 0x48, 0xd7, 0xdd, 0x48, 0x93, 0x3a,
	0xd2, 0x38, 0xce, 0xe8, 0x43, 0xbd, 0xfa, 0x02, 0xde, 0xb1, 0xf5, 0xd8, 0x9c, 0x09, 0x75, 0xa8,
	0xe3, 0x78, 0xd1, 0x5f, 0x02, 0x9b, 0xf3, 0x58, 0xe3, 0xd3, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5c, 0x21, 0x0c, 0x7d, 0xb0, 0x01, 0x00, 0x00,
}
