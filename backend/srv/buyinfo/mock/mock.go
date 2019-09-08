package mock

import (
	"context"
	"errors"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

// Find is buyinfo create mock
func (a *mockSrv) Create(ctx context.Context, req *buyinfo.BuyInfoCreateRequest, opts ...client.CallOption) (*buyinfo.BuyInfoCreateResponse, error) {
	var rsp buyinfo.BuyInfoCreateResponse
	if !utils.RequireParam(req.ValidTime, req.GoodName, req.UserID) {
		rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
		return &rsp, nil
	}

	if utils.IsEmpty(req.ContentID) && utils.IsEmpty(req.ContentToken) {
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoID = 1000
	} else if !utils.IsEmpty(req.ContentID) && !utils.IsEmpty(req.ContentToken) {
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
	if utils.RequireParam(req.BuyInfoID) {
		if req.BuyInfoID == 1000 {
			rsp.BuyInfoID = 1000
			rsp.GoodName = "good"
			rsp.ValidTime = 1234567890
			rsp.Description = "very good!"
			rsp.ContentID = "012345678901234567890123"
			rsp.UserID = 1000
		} else if req.BuyInfoID == 3000 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

// Find is buyinfo find mock
func (a *mockSrv) Find(ctx context.Context, req *buyinfo.BuyInfoFindRequest, opts ...client.CallOption) (*buyinfo.BuyInfoFindResponse, error) {
	var rsp buyinfo.BuyInfoFindResponse
	info := buyinfo.BuyInfoMsg{UserID: 1000}
	if req.UserID == 1000 {
		if req.Status == 1 {
			rsp.BuyInfo = append(rsp.BuyInfo, &info)
		} else {
			rsp.BuyInfo = append(rsp.BuyInfo, &info)
			rsp.BuyInfo = append(rsp.BuyInfo, &info)
		}
	} else {
		return nil, errors.New("")
	}
	return &rsp, nil
}

// Update is buyinfo update mock
func (a *mockSrv) Update(ctx context.Context, req *buyinfo.BuyInfoUpdateRequest, opts ...client.CallOption) (*buyinfo.BuyInfoUpdateResponse, error) {
	var rsp buyinfo.BuyInfoUpdateResponse
	return &rsp, nil
}

// NewBuyInfoService is buyinfo service mock
func NewBuyInfoService() buyinfo.BuyInfoService {
	return new(mockSrv)
}
