package mock

import (
	"context"
	avatar "jiaojiao/srv/avatar/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

// Create is avatar create mock
func (a *mockSrv) Create(ctx context.Context, in *avatar.AvatarCreateRequest, opts ...client.CallOption) (*avatar.AvatarCreateResponse, error) {
	var rsp avatar.AvatarCreateResponse
	// TODO
	return &rsp, nil
}

// NewAvatarService is avatar service mock
func NewAvatarService() avatar.AvatarService {
	return new(mockSrv)
}
