package mock

import (
	"context"
	"errors"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	user "jiaojiao/srv/user/proto"

	"github.com/micro/go-micro/client"
)

type mockSellInfoSrv struct{}
type mockContentSrv struct{}

func (a *mockSellInfoSrv) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, opts ...client.CallOption) (*sellinfo.SellInfoCreateResponse, error) {
	var rsp sellinfo.SellInfoCreateResponse
	return &rsp, nil
}

func (a *mockSellInfoSrv) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, opts ...client.CallOption) (*sellinfo.SellInfoQueryResponse, error) {
	var rsp sellinfo.SellInfoQueryResponse
	if req.SellInfoId != 0 {
		if req.SellInfoId == 1000 {
			rsp.SellInfoId = 1000
			rsp.GoodName = "good"
			rsp.ValidTime = 1234567890
			rsp.Description = "very good!"
			rsp.ContentId = "123456789abc123456789abc"
			rsp.ContentToken = "fajoijrw78029347-fsay"
		} else if req.SellInfoId == 2000 {
			return nil, errors.New("")
		}
	}
	return &rsp, nil
}

func (a *mockSellInfoSrv) Find(ctx context.Context, req *user.UserFindRequest, opts ...client.CallOption) (*user.UserFindResponse, error) {
	user1 := user.UserInfo{
		UserId:      1000,
		UserName:    "test1",
		AvatarId:    "5d23ea2c32311335f935cd14",
		Telephone:   "12345224232",
		StudentId:   "517397299873",
		StudentName: "Xiao Ming",
	}
	user2 := user.UserInfo{
		UserId:      1001,
		UserName:    "test2",
		AvatarId:    "jksfa0980923jkjoifu92323",
		Telephone:   "67307269876",
		StudentId:   "517234731342",
		StudentName: "Xiao Huang",
	}
	user3 := user.UserInfo{
		UserId:      1002,
		UserName:    "test2",
		AvatarId:    "yuwry981hkjbgmxnlaud9u34352",
		Telephone:   "16539896792",
		StudentId:   "517357253234",
		StudentName: "Xiao Bai",
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

func (a *mockContentSrv) Create(ctx context.Context, req *sellinfo.ContentCreateRequest, opts ...client.CallOption) (*sellinfo.ContentCreateResponse, error) {
	var rsp sellinfo.ContentCreateResponse
	return &rsp, nil
}

func (a *mockContentSrv) Find(ctx context.Context, req *user.AdminUserRequest, opts ...client.CallOption) (*user.AdminUserResponse, error) {
	var rsp user.AdminUserResponse
	if req.StudentId == "" {
		rsp.Status = user.AdminUserResponse_INVALID_PARAM
	} else {
		if req.StudentId == "1001" {
			rsp.Status = user.AdminUserResponse_SUCCESS
			rsp.AdminId = 1
		} else if req.StudentId == "2001" {
			return &rsp, errors.New("")
		} else {
			rsp.Status = user.AdminUserResponse_NOT_FOUND
		}
	}
	return &rsp, nil
}

func NewSellInfoService() sellinfo.SellInfoService {
	return new(mockSellInfoSrv)
}

func NewContentService() sellinfo.ContentService {
	return new(mockContentSrv)
}
