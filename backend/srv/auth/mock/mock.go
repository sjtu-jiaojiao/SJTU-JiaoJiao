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
		ret.Status = auth.AuthResponse_INVALID_PARAM
	} else {
		if req.Code == "valid" {
			ret.Status = auth.AuthResponse_SUCCESS
			ret.Token = "test_token"
			ret.StudentId = "1234"
			ret.StudentName = "test"
		} else if req.Code == "down" {
			return &ret, errors.New("")

		} else if req.Code == "userdown" {
			ret.Status = auth.AuthResponse_SUCCESS
			ret.Token = "test_token"
			ret.StudentId = "2345"
			ret.StudentName = "down"
		} else {
			ret.Status = auth.AuthResponse_INVALID_CODE
		}
	}
	return &ret, nil
}

func NewAuthService() auth.AuthService {
	return new(mockSrv)
}
