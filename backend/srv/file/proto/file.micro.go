// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: file.proto

package file

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

// Client API for File service

type FileService interface {
	Query(ctx context.Context, in *FileQueryRequest, opts ...client.CallOption) (*FileQueryResponse, error)
}

type fileService struct {
	c    client.Client
	name string
}

func NewFileService(name string, c client.Client) FileService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "file"
	}
	return &fileService{
		c:    c,
		name: name,
	}
}

func (c *fileService) Query(ctx context.Context, in *FileQueryRequest, opts ...client.CallOption) (*FileQueryResponse, error) {
	req := c.c.NewRequest(c.name, "File.Query", in)
	out := new(FileQueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for File service

type FileHandler interface {
	Query(context.Context, *FileQueryRequest, *FileQueryResponse) error
}

func RegisterFileHandler(s server.Server, hdlr FileHandler, opts ...server.HandlerOption) error {
	type file interface {
		Query(ctx context.Context, in *FileQueryRequest, out *FileQueryResponse) error
	}
	type File struct {
		file
	}
	h := &fileHandler{hdlr}
	return s.Handle(s.NewHandler(&File{h}, opts...))
}

type fileHandler struct {
	FileHandler
}

func (h *fileHandler) Query(ctx context.Context, in *FileQueryRequest, out *FileQueryResponse) error {
	return h.FileHandler.Query(ctx, in, out)
}
