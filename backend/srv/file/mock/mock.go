package mock

import (
	"context"
	file "jiaojiao/srv/file/proto"

	"github.com/micro/go-micro/client"
)

type mockFileSrv struct{}

// Create is file create mock
func (a *mockFileSrv) Create(ctx context.Context, in *file.FileCreateRequest, opts ...client.CallOption) (*file.FileCreateResponse, error) {
	var rsp file.FileCreateResponse
	// TODO
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
