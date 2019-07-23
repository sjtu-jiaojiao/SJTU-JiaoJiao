package mock

import (
	"context"
	"errors"
	auth "jiaojiao/srv/auth/proto"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

// Auth is auth auth mock
func (a *mockSrv) Auth(ctx context.Context, req *auth.AuthRequest, opts ...client.CallOption) (*auth.AuthResponse, error) {
	var ret auth.AuthResponse
	if req.Code == "" {
		ret.Status = auth.AuthResponse_INVALID_PARAM
	} else {
		if req.Code == "valid_user" {
			ret.Status = auth.AuthResponse_SUCCESS
			ret.Token = "test_token"
			ret.StudentId = "1000"
			ret.StudentName = "test"
		} else if req.Code == "valid_admin" {
			ret.Status = auth.AuthResponse_SUCCESS
			ret.Token = "test_token"
			ret.StudentId = "1001"
			ret.StudentName = "test"
		} else if req.Code == "frozen_user" {
			ret.Status = auth.AuthResponse_SUCCESS
			ret.Token = "test_token"
			ret.StudentId = "3000"
			ret.StudentName = "test"
		} else if req.Code == "down" {
			return &ret, errors.New("")
		} else if req.Code == "userdown" {
			ret.Status = auth.AuthResponse_SUCCESS
			ret.Token = "test_token"
			ret.StudentId = "2000"
			ret.StudentName = "down"
		} else if req.Code == "frozen" {
			ret.Status = auth.AuthResponse_SUCCESS
			ret.Token = "test_token"
			ret.StudentId = "3000"
			ret.StudentName = "frozen"
		} else {
			ret.Status = auth.AuthResponse_INVALID_CODE
		}
	}
	return &ret, nil
}

// NewAuthService is auth service mock
func NewAuthService() auth.AuthService {
	return new(mockSrv)
}
