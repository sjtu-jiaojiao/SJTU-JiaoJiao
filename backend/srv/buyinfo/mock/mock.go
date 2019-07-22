package mock

import (
	"context"
	"errors"
	buyinfo "jiaojiao/srv/buyinfo/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

func (a *mockSrv) Create(ctx context.Context, req *buyinfo.BuyInfoCreateRequest, opts ...client.CallOption) (*buyinfo.BuyInfoCreateResponse, error) {
	var rsp buyinfo.BuyInfoCreateResponse
	if req.ValidTime == 0 || req.GoodName == "" || req.UserId == 0 {
		rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
		return &rsp, nil
	}
	if req.ContentId == "" && req.ContentToken == "" {
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoId = 1000
	} else if req.ContentId != "" && req.ContentToken != "" {
		if req.ContentId == "error" {
			return nil, errors.New("")
		}
		if req.ContentToken == "invalid_token" {
			rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_TOKEN
			return &rsp, nil
		}
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoId = 1000
	} else {
		rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
	}
	return &rsp, nil
}

func (a *mockSrv) Query(ctx context.Context, req *buyinfo.BuyInfoQueryRequest, opts ...client.CallOption) (*buyinfo.BuyInfoMsg, error) {
	var rsp buyinfo.BuyInfoMsg
	if req.BuyInfoId != 0 {
		if req.BuyInfoId == 1000 {
			rsp.BuyInfoId = 1000
			rsp.GoodName = "good"
			rsp.ValidTime = 1234567890
			rsp.Description = "very good!"
			rsp.ContentId = "123456789abc123456789abc"
			rsp.UserId = 1000
		} else if req.BuyInfoId == 2000 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

func (a *mockSrv) Find(ctx context.Context, req *buyinfo.BuyInfoFindRequest, opts ...client.CallOption) (*buyinfo.BuyInfoFindResponse, error) {
	var rsp buyinfo.BuyInfoFindResponse
	return &rsp, nil
}

func NewBuyInfoService() buyinfo.BuyInfoService {
	return new(mockSrv)
}
