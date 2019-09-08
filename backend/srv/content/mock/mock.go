package mock

import (
	"context"
	"errors"
	content "jiaojiao/srv/content/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

// Update is content update mock
func (a *mockSrv) Update(ctx context.Context, in *content.ContentUpdateRequest, opts ...client.CallOption) (*content.ContentUpdateResponse, error) {
	var rsp content.ContentUpdateResponse
	// TODO
	return &rsp, nil
}

// Query is content query mock
func (a *mockSrv) Query(ctx context.Context, in *content.ContentQueryRequest, opts ...client.CallOption) (*content.ContentQueryResponse, error) {
	var rsp content.ContentQueryResponse
	// TODO
	return &rsp, nil
}

// Check is content check mock
func (a *mockSrv) Check(ctx context.Context, req *content.ContentCheckRequest, opts ...client.CallOption) (*content.ContentCheckResponse, error) {
	var rsp content.ContentCheckResponse
	if !utils.RequireParam(req.ContentID, req.ContentToken) {
		rsp.Status = content.ContentCheckResponse_INVALID_PARAM
		return &rsp, nil
	}

	if req.ContentID == "012345678901234567890123" {
		if req.ContentToken == "valid_token" {
			rsp.Status = content.ContentCheckResponse_VALID
		} else {
			rsp.Status = content.ContentCheckResponse_INVALID
		}
	} else if req.ContentID == "987654321098765432109876" {
		return nil, errors.New("")
	} else {
		rsp.Status = content.ContentCheckResponse_INVALID
	}
	return &rsp, nil
}

// Create is content create mock
func (a *mockSrv) Create(ctx context.Context, req *content.ContentCreateRequest, opts ...client.CallOption) (*content.ContentCreateResponse, error) {
	var rsp content.ContentCreateResponse
	if !utils.RequireParam(req.Content, req.Type) {
		rsp.Status = content.ContentCreateResponse_INVALID_PARAM
	} else if utils.IsEmpty(req.ContentID) && utils.IsEmpty(req.ContentToken) {
		rsp.Status = content.ContentCreateResponse_SUCCESS
	} else if !utils.IsEmpty(req.ContentID) && !utils.IsEmpty(req.ContentToken) {
		if req.ContentID == "invalid" {
			rsp.Status = content.ContentCreateResponse_INVALID_TYPE
			return &rsp, nil
		}
		if req.ContentID == "error" {
			return nil, errors.New("")
		}
		if req.ContentToken == "invalid_token" {
			rsp.Status = content.ContentCreateResponse_INVALID_TOKEN
		} else {
			rsp.Status = content.ContentCreateResponse_SUCCESS
		}
	} else {
		rsp.Status = content.ContentCreateResponse_INVALID_PARAM
	}
	return &rsp, nil
}

// Delete is content delete mock
func (a *mockSrv) Delete(ctx context.Context, req *content.ContentDeleteRequest, opts ...client.CallOption) (*content.ContentDeleteResponse, error) {
	var rsp content.ContentDeleteResponse
	if !utils.RequireParam(req.ContentID, req.ContentToken) {
		rsp.Status = content.ContentDeleteResponse_INVALID_PARAM
		return &rsp, nil
	}

	if req.ContentID == "012345678901234567890123" {
		if req.ContentToken == "valid_token" {
			rsp.Status = content.ContentDeleteResponse_SUCCESS
		} else {
			rsp.Status = content.ContentDeleteResponse_INVALID_TOKEN
		}
	} else if req.ContentID == "987654321098765432109876" {
		return nil, errors.New("")
	} else {
		rsp.Status = content.ContentDeleteResponse_INVALID_TOKEN
	}
	return &rsp, nil
}

// CreateTag is tag create mock
func (a *mockSrv) CreateTag(ctx context.Context, req *content.ContentCreateTagRequest, opts ...client.CallOption) (*content.ContentCreateTagResponse, error) {
	var rsp content.ContentCreateTagResponse
	if !utils.RequireParam(req.Tags) {
		rsp.Status = content.ContentCreateTagResponse_INVALID_PARAM
		return &rsp, nil
	}

	if utils.IsEmpty(req.ContentID) && utils.IsEmpty(req.ContentToken) { // create new
		rsp.ContentID = "1234567890abcdef12345678"
		rsp.ContentToken = "valid_token"
		rsp.Status = content.ContentCreateTagResponse_SUCCESS
	} else if !utils.IsEmpty(req.ContentID) && !utils.IsEmpty(req.ContentToken) { // add exist one
		if req.ContentToken != "valid_token" {
			rsp.Status = content.ContentCreateTagResponse_INVALID_TOKEN
			return &rsp, nil
		}

		rsp.ContentID = "012345678901234567890123"
		rsp.ContentToken = "valid_token"
		rsp.Status = content.ContentCreateTagResponse_SUCCESS
	} else {
		rsp.Status = content.ContentCreateTagResponse_INVALID_PARAM
	}
	return &rsp, nil
}

// NewContentService is content service mock
func NewContentService() content.ContentService {
	return new(mockSrv)
}
