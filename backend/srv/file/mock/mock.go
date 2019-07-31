package mock

import (
	"bytes"
	"context"
	file "jiaojiao/srv/file/proto"

	"github.com/micro/go-micro/client"
)

type mockFileSrv struct{}

// Create is file create mock
func (a *mockFileSrv) Create(ctx context.Context, req *file.FileCreateRequest, opts ...client.CallOption) (*file.FileCreateResponse, error) {
	var rsp file.FileCreateResponse
	if bytes.Equal(req.File, []byte{0, 1, 2, 3, 4, 5}) {
		rsp.Status = file.FileCreateResponse_SUCCESS
		rsp.FileID = "1234567890abcdef12345678"
	}
	return &rsp, nil
}

// Query is file query mock
func (a *mockFileSrv) Query(ctx context.Context, in *file.FileRequest, opts ...client.CallOption) (*file.FileQueryResponse, error) {
	var rsp file.FileQueryResponse
	// TODO
	return &rsp, nil
}

// Delete is file delete mock
func (a *mockFileSrv) Delete(ctx context.Context, in *file.FileRequest, opts ...client.CallOption) (*file.FileDeleteResponse, error) {
	var rsp file.FileDeleteResponse
	// TODO
	return &rsp, nil
}

// NewFileService is file service mock
func NewFileService() file.FileService {
	return new(mockFileSrv)
}
