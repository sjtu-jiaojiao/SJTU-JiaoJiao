package main

import (
	"context"
	db "jiaojiao/database"
	avatar "jiaojiao/srv/avatar/proto"
	"jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"

	"github.com/h2non/filetype"
	"github.com/jinzhu/gorm"
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
	if !utils.RequreParam(req.File, req.UserID) {
		rsp.Status = avatar.AvatarCreateResponse_INVALID_PARAM
		return nil
	}

	if !utils.CheckInTest() && !filetype.IsImage(req.File) {
		rsp.Status = avatar.AvatarCreateResponse_INVALID_TYPE
		return nil
	}

	usr := db.User{
		ID: req.UserID,
	}
	err := db.Ormer.First(&usr).Error
	if gorm.IsRecordNotFoundError(err) {
		rsp.Status = avatar.AvatarCreateResponse_NOT_FOUND
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
		return err
	}

	srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
		func() interface{} { return mock.NewFileService() }).(file.FileService)
	microRsp, err := srv.Create(context.TODO(), &file.FileCreateRequest{
		File: req.File,
	})
	if utils.LogContinue(err, utils.Warning, "File service error: %v", err) || microRsp.Status != file.FileCreateResponse_SUCCESS {
		return err
	}

	usr.AvatarID = microRsp.FileID
	err = db.Ormer.Save(&usr).Error
	if utils.LogContinue(err, utils.Warning) {
		return err
	}

	rsp.AvatarID = microRsp.FileID
	rsp.Status = avatar.AvatarCreateResponse_SUCCESS
	return nil
}

func main() {
	db.InitORM("userdb", new(db.User))
	defer db.CloseORM()
	service := utils.InitMicroService("avatar")
	utils.LogPanic(avatar.RegisterAvatarHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
