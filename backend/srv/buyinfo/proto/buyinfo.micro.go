// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: buyinfo.proto

package buyinfo

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

// Client API for BuyInfo service

type BuyInfoService interface {
	Query(ctx context.Context, in *BuyInfoQueryRequest, opts ...client.CallOption) (*BuyInfoMsg, error)
	Create(ctx context.Context, in *BuyInfoCreateRequest, opts ...client.CallOption) (*BuyInfoCreateResponse, error)
	Find(ctx context.Context, in *BuyInfoFindRequest, opts ...client.CallOption) (*BuyInfoFindResponse, error)
}

type buyInfoService struct {
	c    client.Client
	name string
}

func NewBuyInfoService(name string, c client.Client) BuyInfoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "buyinfo"
	}
	return &buyInfoService{
		c:    c,
		name: name,
	}
}

func (c *buyInfoService) Query(ctx context.Context, in *BuyInfoQueryRequest, opts ...client.CallOption) (*BuyInfoMsg, error) {
	req := c.c.NewRequest(c.name, "BuyInfo.Query", in)
	out := new(BuyInfoMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buyInfoService) Create(ctx context.Context, in *BuyInfoCreateRequest, opts ...client.CallOption) (*BuyInfoCreateResponse, error) {
	req := c.c.NewRequest(c.name, "BuyInfo.Create", in)
	out := new(BuyInfoCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buyInfoService) Find(ctx context.Context, in *BuyInfoFindRequest, opts ...client.CallOption) (*BuyInfoFindResponse, error) {
	req := c.c.NewRequest(c.name, "BuyInfo.Find", in)
	out := new(BuyInfoFindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BuyInfo service

type BuyInfoHandler interface {
	Query(context.Context, *BuyInfoQueryRequest, *BuyInfoMsg) error
	Create(context.Context, *BuyInfoCreateRequest, *BuyInfoCreateResponse) error
	Find(context.Context, *BuyInfoFindRequest, *BuyInfoFindResponse) error
}

func RegisterBuyInfoHandler(s server.Server, hdlr BuyInfoHandler, opts ...server.HandlerOption) error {
	type buyInfo interface {
		Query(ctx context.Context, in *BuyInfoQueryRequest, out *BuyInfoMsg) error
		Create(ctx context.Context, in *BuyInfoCreateRequest, out *BuyInfoCreateResponse) error
		Find(ctx context.Context, in *BuyInfoFindRequest, out *BuyInfoFindResponse) error
	}
	type BuyInfo struct {
		buyInfo
	}
	h := &buyInfoHandler{hdlr}
	return s.Handle(s.NewHandler(&BuyInfo{h}, opts...))
}

type buyInfoHandler struct {
	BuyInfoHandler
}

func (h *buyInfoHandler) Query(ctx context.Context, in *BuyInfoQueryRequest, out *BuyInfoMsg) error {
	return h.BuyInfoHandler.Query(ctx, in, out)
}

func (h *buyInfoHandler) Create(ctx context.Context, in *BuyInfoCreateRequest, out *BuyInfoCreateResponse) error {
	return h.BuyInfoHandler.Create(ctx, in, out)
}

func (h *buyInfoHandler) Find(ctx context.Context, in *BuyInfoFindRequest, out *BuyInfoFindResponse) error {
	return h.BuyInfoHandler.Find(ctx, in, out)
}