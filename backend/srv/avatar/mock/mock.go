package mock

import (
	"context"
	"errors"
	avatar "jiaojiao/srv/avatar/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

// Create is avatar create mock
func (a *mockSrv) Create(ctx context.Context, req *avatar.AvatarCreateRequest, opts ...client.CallOption) (*avatar.AvatarCreateResponse, error) {
	var rsp avatar.AvatarCreateResponse
	if !utils.RequireParam(req.File, req.UserID) {
		rsp.Status = avatar.AvatarCreateResponse_INVALID_PARAM
		return &rsp, nil
	}

	switch string(req.File) {
	case "invalid":
		rsp.Status = avatar.AvatarCreateResponse_INVALID_TYPE
	case "valid":
		rsp.Status = avatar.AvatarCreateResponse_SUCCESS
		rsp.AvatarID = "012345678901234567890123"
	case "error":
		return nil, errors.New("")
	}
	return &rsp, nil
}

// NewAvatarService is avatar service mock
func NewAvatarService() avatar.AvatarService {
	return new(mockSrv)
}
