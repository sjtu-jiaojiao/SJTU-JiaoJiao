package main

import (
	"context"
	db "jiaojiao/database"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"

	"github.com/astaxie/beego/orm"
)

type srv struct{}

var o orm.Ormer

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
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for exist user
 * @apiSuccess {int32} userId created or existed userid
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *user.UserCreateRequest, rsp *user.UserCreateResponse) error {
	if req.StudentId == "" || req.StudentName == "" {
		rsp.Status = user.UserCreateResponse_INVALID_PARAM
	} else {
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

/**
 * @api {rpc} /rpc user.User.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.User.Query
 * @apiDescription Query user info.
 *
 * @apiParam {int32} userId user id
 * @apiSuccess {int32} userId user id
 * @apiSuccess {string} userName user name
 * @apiSuccess {string} avatarId user avatar id
 * @apiSuccess {string} telephone user telephone
 * @apiSuccess {string} studentId student id
 * @apiSuccess {string} studentName student name
 * @apiUse DBServerDown
 */
func (a *srv) Query(ctx context.Context, req *user.UserQueryRequest, rsp *user.UserInfo) error {
	if req.UserId == 0 {
		return nil
	}
	usr := db.User{
		Id: int(req.UserId),
	}
	err := o.Read(&usr)
	if err == orm.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	parseUser(&usr, rsp)
	return nil
}

/**
 * @api {rpc} /rpc user.User.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.User.Find
 * @apiDescription Find user(fuzzy).
 *
 * @apiParam {string} userName username
 * @apiParam {uint32} limit=100 row limit
 * @apiParam {uint32} offset row offset
 * @apiSuccess {list} user see [User Service](#api-Service-user_User_Query)
 * @apiUse DBServerDown
 */
func (a *srv) Find(ctx context.Context, req *user.UserFindRequest, rsp *user.UserFindResponse) error {
	if req.Limit == 0 {
		req.Limit = 100
	}

	var res []*db.User
	_, err := o.QueryTable(&db.User{}).Filter("UserName__icontains", req.UserName).Limit(req.Limit, req.Offset).All(&res)
	if err != nil {
		return err
	}
	for i, v := range res {
		rsp.User = append(rsp.User, new(user.UserInfo))
		parseUser(v, rsp.User[i])
	}
	return nil
}

func parseUser(s *db.User, d *user.UserInfo) {
	d.UserId = int32(s.Id)
	d.UserName = s.UserName
	d.AvatarId = s.AvatarId
	d.Telephone = s.Telephone
	d.StudentId = s.StudentId
	d.StudentName = s.StudentName
}

func main() {
	o = db.InitORM(new(db.User))
	service := utils.InitMicroService("user")
	utils.LogPanic(user.RegisterUserHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
