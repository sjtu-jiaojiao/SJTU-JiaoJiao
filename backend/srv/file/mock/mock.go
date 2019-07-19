package mock

import (
	"context"
	file "jiaojiao/srv/file/proto"

	"github.com/micro/go-micro/client"
)

type mockFileSrv struct{}

func (a *mockFileSrv) Create(ctx context.Context, in *file.FileCreateRequest, opts ...client.CallOption) (*file.FileCreateResponse, error) {
	panic("implement me")
}

func (a *mockFileSrv) Query(ctx context.Context, in *file.FileQueryRequest, opts ...client.CallOption) (*file.FileQueryResponse, error) {
	panic("implement me")
}

func NewFileService() file.FileService {
	return new(mockFileSrv)
}
