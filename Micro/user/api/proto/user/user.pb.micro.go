// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package com_github_hollson_api_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/api/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Info(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	GetUserByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "com.github.hollson.api.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Info(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "User.Info", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.GetUserByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Info(context.Context, *proto1.Request, *proto1.Response) error
	GetUserByID(context.Context, *Request, *Response) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Info(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		GetUserByID(ctx context.Context, in *Request, out *Response) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Info(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.UserHandler.Info(ctx, in, out)
}

func (h *userHandler) GetUserByID(ctx context.Context, in *Request, out *Response) error {
	return h.UserHandler.GetUserByID(ctx, in, out)
}
