package main

import (
	"context"
	"errors"
	avatar "jiaojiao/srv/avatar/proto"
	mockFile "jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	mockUser "jiaojiao/srv/user/mock"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"

	"github.com/h2non/filetype"
	"github.com/micro/go-micro/client"
)

type srv struct{}

/**
 * @api {rpc} /rpc Avatar.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Avatar.Create
 * @apiDescription Create avatar and return avatarID.
 *
 * @apiParam {int32} userID user id
 * @apiParam {bytes} file file bytes, file accept [file type](https://github.com/h2non/filetype#image)
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found <br> 3 for invalid file type
 * @apiSuccess {int32} avatarID new avatar id
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *avatar.AvatarCreateRequest, rsp *avatar.AvatarCreateResponse) error {
	if !utils.RequireParam(req.File, req.UserID) {
		rsp.Status = avatar.AvatarCreateResponse_INVALID_PARAM
		return nil
	}

	if !utils.CheckFile(req.File, filetype.IsImage) {
		rsp.Status = avatar.AvatarCreateResponse_INVALID_TYPE
		return nil
	}

	// upload file
	srvFile := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
		func() interface{} { return mockFile.NewFileService() }).(file.FileService)
	fileRsp, err := srvFile.Create(context.TODO(), &file.FileCreateRequest{
		File: req.File,
	})
	if utils.LogContinue(err, utils.Error) {
		return err
	}
	if fileRsp.Status != file.FileCreateResponse_SUCCESS {
		_, s := utils.LogContinueS("File create return "+fileRsp.Status.String(), utils.Error)
		return errors.New(s)
	}

	// update user
	srvUser := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
		func() interface{} { return mockUser.NewUserService() }).(user.UserService)
	userRsp, err := srvUser.Update(context.TODO(), &user.UserInfo{
		UserID:   req.UserID,
		AvatarID: fileRsp.FileID,
	})
	if utils.LogContinue(err, utils.Error) {
		return err
	}
	if userRsp.Status != user.UserUpdateResponse_SUCCESS {
		_, s := utils.LogContinueS("User update return "+userRsp.Status.String(), utils.Error)
		return errors.New(s)
	}

	rsp.AvatarID = fileRsp.FileID
	rsp.Status = avatar.AvatarCreateResponse_SUCCESS
	return nil
}

func main() {
	service := utils.InitMicroService("avatar")
	utils.LogPanic(avatar.RegisterAvatarHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
