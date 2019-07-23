package mock

import (
	"context"
	"errors"
	user "jiaojiao/srv/user/proto"

	"github.com/micro/go-micro/client"
)

type mockUserSrv struct{}

// Create is user create mock
func (a *mockUserSrv) Create(ctx context.Context, req *user.UserCreateRequest, opts ...client.CallOption) (*user.UserCreateResponse, error) {
	var rsp user.UserCreateResponse
	if req.StudentID == "" || req.StudentName == "" {
		rsp.Status = user.UserCreateResponse_INVALID_PARAM
	} else {
		if req.StudentID == "1000" {
			rsp.Status = user.UserCreateResponse_SUCCESS
			rsp.User = new(user.UserInfo)
			rsp.User.UserID = 1
			rsp.User.Status = user.UserInfo_NORMAL
			rsp.User.Role = user.UserInfo_USER
		} else if req.StudentID == "1001" {
			rsp.Status = user.UserCreateResponse_SUCCESS
			rsp.User = new(user.UserInfo)
			rsp.User.UserID = 2
			rsp.User.Status = user.UserInfo_NORMAL
			rsp.User.Role = user.UserInfo_ADMIN
		} else if req.StudentID == "2000" {
			return nil, errors.New("")
		} else if req.StudentID == "3000" {
			rsp.Status = user.UserCreateResponse_SUCCESS
			rsp.User = new(user.UserInfo)
			rsp.User.UserID = 3
			rsp.User.Status = user.UserInfo_FROZEN
			rsp.User.Role = user.UserInfo_USER
		} else {
			rsp.Status = user.UserCreateResponse_USER_EXIST
			rsp.User = new(user.UserInfo)
			rsp.User.UserID = 1
		}
	}
	return &rsp, nil
}

// Query is user query mock
func (a *mockUserSrv) Query(ctx context.Context, req *user.UserQueryRequest, opts ...client.CallOption) (*user.UserInfo, error) {
	var rsp user.UserInfo
	if req.UserID != 0 {
		if req.UserID == 1000 {
			rsp.UserID = 1000
			rsp.UserName = "test"
			rsp.AvatarID = "5d23ea2c32311335f935cd14"
			rsp.Telephone = "12345678901"
			rsp.StudentID = "1000"
			rsp.StudentName = "jiang"
			rsp.Status = 1
		} else if req.UserID == 2000 {
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
		AvatarID:    "5d23ea2c32311335f935cd14",
		Telephone:   "12345224232",
		StudentID:   "1",
		StudentName: "Xiao Ming",
		Status:      1,
	}
	user2 := user.UserInfo{
		UserID:      1001,
		UserName:    "test2",
		AvatarID:    "jksfa0980923jkjoifu92323",
		Telephone:   "67307269876",
		StudentID:   "2",
		StudentName: "Xiao Huang",
		Status:      1,
	}
	user3 := user.UserInfo{
		UserID:      1002,
		UserName:    "test2",
		AvatarID:    "yuwry981hkjbgmxnlaud9u34352",
		Telephone:   "16539896792",
		StudentID:   "3",
		StudentName: "Xiao Bai",
		Status:      1,
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
	} else if req.UserName == "down" {
		return nil, errors.New("")
	}
	return &rsp, nil
}

// Update is user update mock
func (a *mockUserSrv) Update(ctx context.Context, req *user.UserInfo, opts ...client.CallOption) (*user.UserUpdateResponse, error) {
	var rsp user.UserUpdateResponse
	if req.UserID == 0 {
		rsp.Status = user.UserUpdateResponse_INVALID_PARAM
	} else {
		if req.UserID == 1000 {
			rsp.Status = user.UserUpdateResponse_SUCCESS
		} else if req.UserID == 2000 {
			return &rsp, errors.New("")
		} else {
			rsp.Status = user.UserUpdateResponse_NOT_FOUND
		}
	}
	return &rsp, nil
}

// NewUserService is user service mock
func NewUserService() user.UserService {
	return new(mockUserSrv)
}
