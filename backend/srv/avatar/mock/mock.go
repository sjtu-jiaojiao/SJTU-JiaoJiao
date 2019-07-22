package mock

import (
	"context"
	avatar "jiaojiao/srv/avatar/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

func (a *mockSrv) Create(ctx context.Context, in *avatar.AvatarCreateRequest, opts ...client.CallOption) (*avatar.AvatarCreateResponse, error) {
	var rsp avatar.AvatarCreateResponse
	// TODO
	return &rsp, nil
}

func NewAvatarService() avatar.AvatarService {
	return new(mockSrv)
}
