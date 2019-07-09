package mock

import (
	"context"
	"errors"
	user "jiaojiao/srv/user/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

func (a *mockSrv) Create(ctx context.Context, req *user.UserCreateRequest, opts ...client.CallOption) (*user.UserCreateResponse, error) {
	var rsp user.UserCreateResponse
	if req.StudentId == "" || req.StudentName == "" {
		rsp.Status = user.UserCreateResponse_INVALID_PARAM
	} else {
		if req.StudentId == "1234" {
			rsp.Status = user.UserCreateResponse_SUCCESS
			rsp.UserId = 1
		} else if req.StudentId == "2345" {
			return &rsp, errors.New("")
		} else {
			rsp.Status = user.UserCreateResponse_USER_EXIST
			rsp.UserId = 1
		}
	}
	return &rsp, nil
}

func (a *mockSrv) Query(ctx context.Context, req *user.UserQueryRequest, opts ...client.CallOption) (*user.UserInfo, error) {
	var rsp user.UserInfo
	if req.UserId != 0 {
		if req.UserId == 1234 {
			rsp.UserId = 1234
			rsp.UserName = "test"
			rsp.AvatarId = "5d23ea2c32311335f935cd14"
			rsp.Telephone = "12345678901"
			rsp.StudentId = "1234"
			rsp.StudentName = "jiang"
		} else if req.UserId != 2345 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

func (a *mockSrv) Find(ctx context.Context, req *user.UserFindRequest, opts ...client.CallOption) (*user.UserFindResponse, error) {
	var rsp user.UserFindResponse
	return &rsp, nil
}

func NewUserService() user.UserService {
	return new(mockSrv)
}
