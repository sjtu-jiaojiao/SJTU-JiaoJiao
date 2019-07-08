package mock

import (
	"context"
	"errors"
	user "jiaojiao/srv/user/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

func (a *mockSrv) Create(ctx context.Context, req *user.UserCreateRequest, opts ...client.CallOption) (*user.UserCreateResponse, error) {
	var ret user.UserCreateResponse
	if req.StudentId == "" || req.StudentName == "" {
		ret.Status = user.UserCreateResponse_EMPTY_PARAM
	} else {
		if req.StudentId == "1234" {
			ret.Status = user.UserCreateResponse_SUCCESS
			ret.UserId = 1
		} else if req.StudentId == "2345" {
			return &ret, errors.New("")
		} else {
			ret.Status = user.UserCreateResponse_USER_EXIST
			ret.UserId = 1
		}
	}
	return &ret, nil
}

func NewUserService() user.UserService {
	return new(mockSrv)
}
