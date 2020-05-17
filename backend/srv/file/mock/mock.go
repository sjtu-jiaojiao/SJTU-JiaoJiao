package mock

import (
	"context"
	"errors"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"
)

type mockFileSrv struct{}

// Create is file create mock
func (a *mockFileSrv) Create(ctx context.Context, req *file.FileCreateRequest, opts ...client.CallOption) (*file.FileCreateResponse, error) {
	var rsp file.FileCreateResponse
	if !utils.RequireParam(req.File) {
		rsp.Status = file.FileCreateResponse_INVALID_PARAM
		return &rsp, nil
	}

	if string(req.File) == "valid_file_error" {
		return nil, errors.New("")
	}

	if string(req.File) == "valid_file_unknown" {
		rsp.Status = file.FileCreateResponse_UNKNOWN
		rsp.FileID = ""
	} else {
		rsp.Status = file.FileCreateResponse_SUCCESS
		if string(req.File) == "valid_file_2" {
			rsp.FileID = "000000000000000000000002"
		} else {
			rsp.FileID = "000000000000000000000001"
		}
	}
	return &rsp, nil
}

// Query is file query mock
func (a *mockFileSrv) Query(ctx context.Context, req *file.FileRequest, opts ...client.CallOption) (*file.FileQueryResponse, error) {
	var rsp file.FileQueryResponse
	if !utils.RequireParam(req.FileID) {
		rsp.Status = file.FileQueryResponse_INVALID_PARAM
		return &rsp, nil
	}

	switch req.FileID {
	case "valid":
		rsp.Status = file.FileQueryResponse_SUCCESS
		rsp.Size = 5
		rsp.File = []byte("valid_file")
	case "invalid":
		rsp.Status = file.FileQueryResponse_NOT_FOUND
	case "invalid_type":
		rsp.Status = file.FileQueryResponse_SUCCESS
		rsp.Size = 5
		rsp.File = []byte("invalid_file")
	case "error":
		return nil, errors.New("")
	}

	return &rsp, nil
}

// Delete is file delete mock
func (a *mockFileSrv) Delete(ctx context.Context, req *file.FileRequest, opts ...client.CallOption) (*file.FileDeleteResponse, error) {
	var rsp file.FileDeleteResponse
	if !utils.RequireParam(req.FileID) {
		rsp.Status = file.FileDeleteResponse_INVALID_PARAM
		return &rsp, nil
	}

	if req.FileID == "000000000000000000000001" {
		rsp.Status = file.FileDeleteResponse_SUCCESS
	} else if req.FileID == "100000000000000000000000" {
		return nil, errors.New("")
	} else {
		rsp.Status = file.FileDeleteResponse_NOT_FOUND
	}
	return &rsp, nil
}

// NewFileService is file service mock
func NewFileService() file.FileService {
	return new(mockFileSrv)
}
