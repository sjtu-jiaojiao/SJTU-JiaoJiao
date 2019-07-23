package mock

import (
	"context"
	"errors"
	buyinfo "jiaojiao/srv/buyinfo/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

// Find is buyinfo create mock
func (a *mockSrv) Create(ctx context.Context, req *buyinfo.BuyInfoCreateRequest, opts ...client.CallOption) (*buyinfo.BuyInfoCreateResponse, error) {
	var rsp buyinfo.BuyInfoCreateResponse
	if req.ValidTime == 0 || req.GoodName == "" || req.UserID == 0 {
		rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
		return &rsp, nil
	}
	if req.ContentID == "" && req.ContentToken == "" {
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoID = 1000
	} else if req.ContentID != "" && req.ContentToken != "" {
		if req.ContentID == "error" {
			return nil, errors.New("")
		}
		if req.ContentToken == "invalid_token" {
			rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_TOKEN
			return &rsp, nil
		}
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoID = 1000
	} else {
		rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
	}
	return &rsp, nil
}

// Find is buyinfo query mock
func (a *mockSrv) Query(ctx context.Context, req *buyinfo.BuyInfoQueryRequest, opts ...client.CallOption) (*buyinfo.BuyInfoMsg, error) {
	var rsp buyinfo.BuyInfoMsg
	if req.BuyInfoID != 0 {
		if req.BuyInfoID == 1000 {
			rsp.BuyInfoID = 1000
			rsp.GoodName = "good"
			rsp.ValidTime = 1234567890
			rsp.Description = "very good!"
			rsp.ContentID = "123456789abc123456789abc"
			rsp.UserID = 1000
		} else if req.BuyInfoID == 2000 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

// Find is buyinfo find mock
func (a *mockSrv) Find(ctx context.Context, req *buyinfo.BuyInfoFindRequest, opts ...client.CallOption) (*buyinfo.BuyInfoFindResponse, error) {
	var rsp buyinfo.BuyInfoFindResponse
	// TODO
	return &rsp, nil
}

// NewBuyInfoService is buyinfo service mock
func NewBuyInfoService() buyinfo.BuyInfoService {
	return new(mockSrv)
}
