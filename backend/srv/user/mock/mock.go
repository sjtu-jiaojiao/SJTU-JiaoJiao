package mock

import (
	"context"
	"errors"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"
)

type mockUserSrv struct{}

// Create is user create mock
func (a *mockUserSrv) Create(ctx context.Context, req *user.UserCreateRequest, opts ...client.CallOption) (*user.UserCreateResponse, error) {
	var rsp user.UserCreateResponse
	if !utils.RequireParam(req.StudentID, req.StudentName) {
		rsp.Status = user.UserCreateResponse_INVALID_PARAM
		return &rsp, nil
	}

	if req.StudentID == "10000" {
		rsp.Status = user.UserCreateResponse_SUCCESS
		rsp.User = new(user.UserInfo)
		rsp.User.UserID = 1000
		rsp.User.Status = user.UserInfo_NORMAL
		rsp.User.Role = user.UserInfo_USER
	} else if req.StudentID == "10001" {
		rsp.Status = user.UserCreateResponse_SUCCESS
		rsp.User = new(user.UserInfo)
		rsp.User.UserID = 1001
		rsp.User.Status = user.UserInfo_NORMAL
		rsp.User.Role = user.UserInfo_ADMIN
	} else if req.StudentID == "20000" {
		rsp.Status = user.UserCreateResponse_SUCCESS
		rsp.User = new(user.UserInfo)
		rsp.User.UserID = 2000
		rsp.User.Status = user.UserInfo_FROZEN
		rsp.User.Role = user.UserInfo_USER
	} else if req.StudentID == "30000" {
		return nil, errors.New("")
	} else {
		rsp.Status = user.UserCreateResponse_USER_EXIST
		rsp.User = new(user.UserInfo)
		rsp.User.UserID = 1000
	}
	return &rsp, nil
}

// Query is user query mock
func (a *mockUserSrv) Query(ctx context.Context, req *user.UserQueryRequest, opts ...client.CallOption) (*user.UserInfo, error) {
	var rsp user.UserInfo
	if utils.RequireParam(req.UserID) {
		if req.UserID == 1000 {
			rsp.UserID = 1000
			rsp.UserName = "test"
			rsp.AvatarID = "012345678901234567890123"
			rsp.Telephone = "12345678901"
			rsp.StudentID = "10000"
			rsp.StudentName = "jiang"
			rsp.Status = user.UserInfo_NORMAL
		} else if req.UserID == 3000 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

// Find is user find mock
func (a *mockUserSrv) Find(ctx context.Context, req *user.UserFindRequest, opts ...client.CallOption) (*user.UserFindResponse, error) {
	user1 := user.UserInfo{
		UserID:      1000,
		UserName:    "test1",
		AvatarID:    "012345678901234567890123",
		Telephone:   "12345678901",
		StudentID:   "10000",
		StudentName: "user1",
		Status:      user.UserInfo_NORMAL,
	}
	user2 := user.UserInfo{
		UserID:      1001,
		UserName:    "test2",
		AvatarID:    "012345678901234567890123",
		Telephone:   "12345678901",
		StudentID:   "10001",
		StudentName: "user2",
		Status:      user.UserInfo_NORMAL,
	}
	user3 := user.UserInfo{
		UserID:      1002,
		UserName:    "test3",
		AvatarID:    "012345678901234567890123",
		Telephone:   "12345678901",
		StudentID:   "10002",
		StudentName: "user3",
		Status:      user.UserInfo_NORMAL,
	}

	var rsp user.UserFindResponse
	if req.UserName == "test1" {
		rsp.User = append(rsp.User, &user1)
	} else if req.UserName == "test2" {
		rsp.User = append(rsp.User, &user2)
		rsp.User = append(rsp.User, &user3)
	} else if req.UserName == "" {
		if req.Limit == 0 {
			rsp.User = append(rsp.User, &user1)
			rsp.User = append(rsp.User, &user2)
			rsp.User = append(rsp.User, &user3)
		} else if req.Limit == 2 {
			if req.Offset == 0 {
				rsp.User = append(rsp.User, &user1)
				rsp.User = append(rsp.User, &user2)
			} else if req.Offset == 1 {
				rsp.User = append(rsp.User, &user2)
				rsp.User = append(rsp.User, &user3)
			} else if req.Offset == 2 {
				rsp.User = append(rsp.User, &user3)
			}
		}
	} else if req.UserName == "error" {
		return nil, errors.New("")
	}
	return &rsp, nil
}

// Update is user update mock
func (a *mockUserSrv) Update(ctx context.Context, req *user.UserInfo, opts ...client.CallOption) (*user.UserUpdateResponse, error) {
	var rsp user.UserUpdateResponse
	if !utils.RequireParam(req.UserID) {
		rsp.Status = user.UserUpdateResponse_INVALID_PARAM
		return &rsp, nil
	}

	if req.UserID == 1000 {
		rsp.Status = user.UserUpdateResponse_SUCCESS
	} else if req.UserID == 3000 {
		return &rsp, errors.New("")
	} else {
		rsp.Status = user.UserUpdateResponse_NOT_FOUND
	}
	return &rsp, nil
}

// NewUserService is user service mock
func NewUserService() user.UserService {
	return new(mockUserSrv)
}
