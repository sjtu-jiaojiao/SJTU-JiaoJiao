package mock

import (
	"context"
	"errors"
	auth "jiaojiao/srv/auth/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"
)

type mockSrv struct{}

// Auth is auth auth mock
func (a *mockSrv) Auth(ctx context.Context, req *auth.AuthRequest, opts ...client.CallOption) (*auth.AuthResponse, error) {
	var ret auth.AuthResponse
	if !utils.RequireParam(req.Code) {
		ret.Status = auth.AuthResponse_INVALID_PARAM
		return &ret, nil
	}
	if req.Code == "valid_user" {
		ret.Status = auth.AuthResponse_SUCCESS
		ret.Token = "valid_token"
		ret.StudentID = "10000"
		ret.StudentName = "user"
	} else if req.Code == "valid_admin" {
		ret.Status = auth.AuthResponse_SUCCESS
		ret.Token = "valid_token"
		ret.StudentID = "10001"
		ret.StudentName = "admin"
	} else if req.Code == "frozen_user" {
		ret.Status = auth.AuthResponse_SUCCESS
		ret.Token = "valid_token"
		ret.StudentID = "20000"
		ret.StudentName = "frozen"
	} else if req.Code == "error" {
		return &ret, errors.New("")
	} else if req.Code == "user_error" {
		ret.Status = auth.AuthResponse_SUCCESS
		ret.Token = "valid_token"
		ret.StudentID = "30000"
		ret.StudentName = "error"
	} else {
		ret.Status = auth.AuthResponse_INVALID_CODE
	}
	return &ret, nil
}

// NewAuthService is auth service mock
func NewAuthService() auth.AuthService {
	return new(mockSrv)
}
