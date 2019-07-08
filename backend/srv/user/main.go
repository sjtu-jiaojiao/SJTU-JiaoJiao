package main

import (
	"context"
	db "jiaojiao/database"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"

	"github.com/astaxie/beego/orm"
)

type srv struct{}

/**
 * @apiDefine DBServerDown
 * @apiError (Error 500) DBServerDown can't connect to database server
 */

/**
 * @api {rpc} /rpc user.User.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.User.Create
 * @apiDescription Create new user.
 *
 * @apiParam {string} studentId student id.
 * @apiParam {string} studentName student name.
 * @apiSuccess {number} status -1 for empty param <br> 1 for success <br> 2 for exist user
 * @apiSuccess {int32} userId created or existed userid
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *user.UserCreateRequest, rsp *user.UserCreateResponse) error {
	if req.StudentId == "" || req.StudentName == "" {
		rsp.Status = user.UserCreateResponse_EMPTY_PARAM
	} else {
		o := orm.NewOrm()
		usr := db.User{
			UserName:    req.StudentName,
			StudentId:   req.StudentId,
			StudentName: req.StudentName,
		}
		created, id, err := o.ReadOrCreate(&usr, "StudentId")
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		if created {
			rsp.Status = user.UserCreateResponse_SUCCESS
		} else {
			rsp.Status = user.UserCreateResponse_USER_EXIST
		}
		rsp.UserId = int32(id)
	}
	return nil
}

func main() {
	service := utils.InitMicroService("user")
	utils.LogPanic(user.RegisterUserHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
