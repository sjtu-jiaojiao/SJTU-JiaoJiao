// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sellinfo.proto

package sellinfo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Client API for SellInfo service

type SellInfoService interface {
}

type sellInfoService struct {
	c    client.Client
	name string
}

func NewSellInfoService(name string, c client.Client) SellInfoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "sellinfo"
	}
	return &sellInfoService{
		c:    c,
		name: name,
	}
}

// Server API for SellInfo service

type SellInfoHandler interface {
}

func RegisterSellInfoHandler(s server.Server, hdlr SellInfoHandler, opts ...server.HandlerOption) error {
	type sellInfo interface {
	}
	type SellInfo struct {
		sellInfo
	}
	h := &sellInfoHandler{hdlr}
	return s.Handle(s.NewHandler(&SellInfo{h}, opts...))
}

type sellInfoHandler struct {
	SellInfoHandler
}

// Client API for Content service

type ContentService interface {
	Create(ctx context.Context, in *ContentCreateRequest, opts ...client.CallOption) (*ContentCreateResponse, error)
}

type contentService struct {
	c    client.Client
	name string
}

func NewContentService(name string, c client.Client) ContentService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "content"
	}
	return &contentService{
		c:    c,
		name: name,
	}
}

func (c *contentService) Create(ctx context.Context, in *ContentCreateRequest, opts ...client.CallOption) (*ContentCreateResponse, error) {
	req := c.c.NewRequest(c.name, "Content.Create", in)
	out := new(ContentCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Content service

type ContentHandler interface {
	Create(context.Context, *ContentCreateRequest, *ContentCreateResponse) error
}

func RegisterContentHandler(s server.Server, hdlr ContentHandler, opts ...server.HandlerOption) error {
	type content interface {
		Create(ctx context.Context, in *ContentCreateRequest, out *ContentCreateResponse) error
	}
	type Content struct {
		content
	}
	h := &contentHandler{hdlr}
	return s.Handle(s.NewHandler(&Content{h}, opts...))
}

type contentHandler struct {
	ContentHandler
}

func (h *contentHandler) Create(ctx context.Context, in *ContentCreateRequest, out *ContentCreateResponse) error {
	return h.ContentHandler.Create(ctx, in, out)
}
