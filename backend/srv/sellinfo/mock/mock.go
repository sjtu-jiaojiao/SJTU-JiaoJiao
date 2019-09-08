package mock

import (
	"context"
	"errors"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

// Create is sellinfo create mock
func (a *mockSrv) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, opts ...client.CallOption) (*sellinfo.SellInfoCreateResponse, error) {
	var rsp sellinfo.SellInfoCreateResponse
	if !utils.RequireParam(req.ValidTime, req.GoodName, req.UserID) {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
		return &rsp, nil
	}

	if utils.IsEmpty(req.ContentID) && utils.IsEmpty(req.ContentToken) {
		rsp.Status = sellinfo.SellInfoCreateResponse_SUCCESS
		rsp.SellInfoID = 1000
	} else if !utils.IsEmpty(req.ContentID) && !utils.IsEmpty(req.ContentToken) {
		if req.ContentID == "error" {
			return nil, errors.New("")
		}
		if req.ContentToken == "invalid_token" {
			rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_TOKEN
			return &rsp, nil
		}
		rsp.Status = sellinfo.SellInfoCreateResponse_SUCCESS
		rsp.SellInfoID = 1000
	} else {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
	}
	return &rsp, nil
}

// Query is sellinfo query mock
func (a *mockSrv) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, opts ...client.CallOption) (*sellinfo.SellInfoMsg, error) {
	var rsp sellinfo.SellInfoMsg
	if utils.RequireParam(req.SellInfoID) {
		if req.SellInfoID == 1000 {
			rsp.SellInfoID = 1000
			rsp.GoodName = "good"
			rsp.ValidTime = 1234567890
			rsp.Description = "very good!"
			rsp.ContentID = "012345678901234567890123"
			rsp.UserID = 1000
		} else if req.SellInfoID == 3000 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

// Find is sellinfo find mock
func (a *mockSrv) Find(ctx context.Context, req *sellinfo.SellInfoFindRequest, opts ...client.CallOption) (*sellinfo.SellInfoFindResponse, error) {
	var rsp sellinfo.SellInfoFindResponse
	info := sellinfo.SellInfoMsg{UserID: 1000}
	if req.UserID == 1000 {
		if req.Status == 1 {
			rsp.SellInfo = append(rsp.SellInfo, &info)
		} else {
			rsp.SellInfo = append(rsp.SellInfo, &info)
			rsp.SellInfo = append(rsp.SellInfo, &info)
		}
	} else {
		return nil, errors.New("")
	}
	return &rsp, nil
}

// Update is sellinfo update mock
func (a *mockSrv) Update(ctx context.Context, req *sellinfo.SellInfoUpdateRequest, opts ...client.CallOption) (*sellinfo.SellInfoUpdateResponse, error) {
	var rsp sellinfo.SellInfoUpdateResponse
	return &rsp, nil
}

// NewSellInfoService is sellinfo service mock
func NewSellInfoService() sellinfo.SellInfoService {
	return new(mockSrv)
}
