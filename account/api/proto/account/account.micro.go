// Code generated by protoc-gen-micro. DO NOT EDIT.
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
import go_api "github.com/micro/go-api/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = go_api.Response{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Account service

type AccountService interface {
	Login(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	Register(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.api.account"
	}
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) Login(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Account.Login", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Register(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Account.Register", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Login(context.Context, *go_api.Request, *go_api.Response) error
	Register(context.Context, *go_api.Request, *go_api.Response) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) {
	type account interface {
		Login(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		Register(ctx context.Context, in *go_api.Request, out *go_api.Response) error
	}
	type Account struct {
		account
	}
	h := &accountHandler{hdlr}
	s.Handle(s.NewHandler(&Account{h}, opts...))
}

type accountHandler struct {
	AccountHandler
}

func (h *accountHandler) Login(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *accountHandler) Register(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.AccountHandler.Register(ctx, in, out)
}
