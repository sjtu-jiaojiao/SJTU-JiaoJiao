package mock

import (
	"context"
	"errors"
	sellinfo "jiaojiao/srv/sellinfo/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

// Create is sellinfo create mock
func (a *mockSrv) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, opts ...client.CallOption) (*sellinfo.SellInfoCreateResponse, error) {
	var rsp sellinfo.SellInfoCreateResponse
	if req.ValidTime == 0 || req.GoodName == "" || req.UserId == 0 {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
		return &rsp, nil
	}
	if req.ContentId == "" && req.ContentToken == "" {
		rsp.Status = sellinfo.SellInfoCreateResponse_SUCCESS
		rsp.SellInfoId = 1000
	} else if req.ContentId != "" && req.ContentToken != "" {
		if req.ContentId == "error" {
			return nil, errors.New("")
		}
		if req.ContentToken == "invalid_token" {
			rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_TOKEN
			return &rsp, nil
		}
		rsp.Status = sellinfo.SellInfoCreateResponse_SUCCESS
		rsp.SellInfoId = 1000
	} else {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
	}
	return &rsp, nil
}

// Query is sellinfo query mock
func (a *mockSrv) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, opts ...client.CallOption) (*sellinfo.SellInfoMsg, error) {
	var rsp sellinfo.SellInfoMsg
	if req.SellInfoId != 0 {
		if req.SellInfoId == 1000 {
			rsp.SellInfoId = 1000
			rsp.GoodName = "good"
			rsp.ValidTime = 1234567890
			rsp.Description = "very good!"
			rsp.ContentId = "123456789abc123456789abc"
			rsp.UserId = 1000
		} else if req.SellInfoId == 2000 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

// Find is sellinfo find mock
func (a *mockSrv) Find(ctx context.Context, req *sellinfo.SellInfoFindRequest, opts ...client.CallOption) (*sellinfo.SellInfoFindResponse, error) {
	var rsp sellinfo.SellInfoFindResponse
	return &rsp, nil
}

// NewSellInfoService is sellinfo service mock
func NewSellInfoService() sellinfo.SellInfoService {
	return new(mockSrv)
}
