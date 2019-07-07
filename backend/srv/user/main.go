package main

import (
	"context"
	db "jiaojiao/database"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"

	"github.com/astaxie/beego/orm"
	"github.com/micro/go-micro/server"
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
 * @apiParam {uint64} studentId student id.
 * @apiParam {string} studentName student name.
 * @apiSuccess {number} status -1 for empty param <br> 1 for success <br> 2 for exist user
 * @apiSuccess {int32} userId created or existed userid
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *user.UserCreateRequest, rsp *user.UserCreateResponse) error {
	if req.StudentId == 0 || req.StudentName == "" {
		rsp.Status = -1
	} else {
		o := orm.NewOrm()
		usr := db.User{StudentId: req.StudentId, StudentName: req.StudentName}
		created, id, err := o.ReadOrCreate(&usr, "StudentId")
		if err != nil {
			return err
		}
		if created {
			rsp.Status = 1
		} else {
			rsp.Status = 2
		}
		rsp.UserId = int32(id)
	}
	return nil
}

func main() {
	utils.RunMicroService("user", func(s server.Server, hdlr interface{},
		opts ...server.HandlerOption) error {
		return user.RegisterUserHandler(s, hdlr.(user.UserHandler), opts...)
	}, new(srv))
}
