package mock

import (
	"bytes"
	"context"
	"errors"
	content "jiaojiao/srv/content/proto"

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
func (a *mockSrv) Check(ctx context.Context, in *content.ContentCheckRequest, opts ...client.CallOption) (*content.ContentCheckResponse, error) {
	var rsp content.ContentCheckResponse
	// TODO
	return &rsp, nil
}

// Create is content create mock
func (a *mockSrv) Create(ctx context.Context, req *content.ContentCreateRequest, opts ...client.CallOption) (*content.ContentCreateResponse, error) {
	var rsp content.ContentCreateResponse
	if bytes.Equal(req.Content, []byte{0}) || req.Type == 0 {
		rsp.Status = content.ContentCreateResponse_INVALID_PARAM
	} else if req.ContentId == "" && req.ContentToken == "" {
		rsp.Status = content.ContentCreateResponse_SUCCESS
	} else if req.ContentId != "" && req.ContentToken != "" {
		if req.ContentId == "invalid" {
			rsp.Status = content.ContentCreateResponse_INVALID_TYPE
			return &rsp, nil
		}
		if req.ContentId == "error" {
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
	if req.ContentId == "" || req.ContentToken == "" {
		rsp.Status = content.ContentDeleteResponse_INVALID_PARAM
	} else {
		if req.ContentId == "1000" {
			if req.ContentToken == "valid_token" {
				rsp.Status = content.ContentDeleteResponse_SUCCESS
			} else {
				rsp.Status = content.ContentDeleteResponse_INVALID_TOKEN
			}
		} else if req.ContentId == "2000" {
			return nil, errors.New("")
		} else {
			rsp.Status = content.ContentDeleteResponse_INVALID_TOKEN
		}
	}
	return &rsp, nil
}

// NewContentService is content service mock
func NewContentService() content.ContentService {
	return new(mockSrv)
}
