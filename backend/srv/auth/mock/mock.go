package mock

import (
	"context"
	"errors"
	auth "jiaojiao/srv/auth/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

func (a *mockSrv) Auth(ctx context.Context, req *auth.AuthRequest, opts ...client.CallOption) (*auth.AuthResponse, error) {
	var ret auth.AuthResponse
	if req.Code == "" {
		ret.Status = -1
	} else {
		if req.Code == "valid" {
			ret.Status = 1
			ret.Token = "test_token"
			ret.StudentId = 1234
			ret.StudentName = "test"
		} else if req.Code == "down" {
			return &ret, errors.New("")
		} else {
			ret.Status = 2
		}
	}
	return &ret, nil
}

func NewAuthService() auth.AuthService {
	return new(mockSrv)
}
