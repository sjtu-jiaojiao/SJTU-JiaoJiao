package main

import (
	"context"
	db "jiaojiao/database"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"

	"github.com/astaxie/beego/orm"
)

type srvUser struct{}
type srvAdmin struct{}

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
func (a *srvUser) Create(ctx context.Context, req *user.UserCreateRequest, rsp *user.UserCreateResponse) error {
	if req.StudentId == "" || req.StudentName == "" {
		rsp.Status = user.UserCreateResponse_INVALID_PARAM
	} else {
		usr := db.User{
			UserName:    req.StudentName,
			AvatarId:    utils.GetStringConfig("srv_config", "default_avatar"),
			StudentId:   req.StudentId,
			StudentName: req.StudentName,
			Status:      1,
		}
		created, id, err := db.Ormer.ReadOrCreate(&usr, "StudentId")
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
 * @apiSuccess {int32} status user status, 1 for normal <br> 2 for frozen
 * @apiUse DBServerDown
 */
func (a *srvUser) Query(ctx context.Context, req *user.UserQueryRequest, rsp *user.UserInfo) error {
	if req.UserId == 0 {
		return nil
	}
	usr := db.User{
		Id: int(req.UserId),
	}
	err := db.Ormer.Read(&usr)
	if err == orm.ErrNoRows {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
		return err
	}
	parseUser(&usr, rsp)
	return nil
}

/**
 * @api {rpc} /rpc user.User.Update
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.User.Update
 * @apiDescription Update user info, only update provided field. If clearEmpty=1 and param support allow clear, clear the field when not provided.
 *
 * @apiParam {int32} userId user id
 * @apiParam {string} [userName] user name
 * @apiParam {string} [avatarId] user avatar id
 * @apiParam {string} [telephone] user telephone, allow clear
 * @apiParam {string} [studentId] student id
 * @apiParam {string} [studentName] student name
 * @apiParam {int32} status user status, 1 for normal <br> 2 for frozen
 * @apiParam {bool} clearEmpty=0 clear the empty field
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for user not found
 * @apiUse DBServerDown
 */
func (a *srvUser) Update(ctx context.Context, req *user.UserInfo, rsp *user.UserUpdateResponse) error {
	if req.UserId == 0 {
		rsp.Status = user.UserUpdateResponse_INVALID_PARAM
		return nil
	}

	usr := db.User{
		Id: int(req.UserId),
	}
	if err := db.Ormer.Read(&usr); err == nil {
		utils.AssignNotEmpty(&req.UserName, &usr.UserName)
		utils.AssignNotEmpty(&req.AvatarId, &usr.AvatarId)
		if req.ClearEmpty {
			usr.Telephone = req.Telephone
		} else {
			utils.AssignNotEmpty(&req.Telephone, &usr.Telephone)
		}
		utils.AssignNotEmpty(&req.StudentId, &usr.StudentId)
		utils.AssignNotEmpty(&req.StudentName, &usr.StudentName)
		utils.AssignNotZero(&req.Status, &usr.Status)
		_, err := db.Ormer.Update(&usr)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		rsp.Status = user.UserUpdateResponse_SUCCESS
	} else if err == orm.ErrNoRows {
		rsp.Status = user.UserUpdateResponse_NOT_FOUND
		return nil
	} else {
		utils.Warning(err)
		return err
	}
	return nil
}

/**
 * @api {rpc} /rpc user.User.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.User.Find
 * @apiDescription Find user(fuzzy).
 *
 * @apiParam {string} [userName] username
 * @apiParam {uint32} limit=100 row limit
 * @apiParam {uint32} offset=0 row offset
 * @apiSuccess {list} user see [User Service](#api-Service-user_User_Query)
 * @apiUse DBServerDown
 */
func (a *srvUser) Find(ctx context.Context, req *user.UserFindRequest, rsp *user.UserFindResponse) error {
	if req.Limit == 0 {
		req.Limit = 100
	}

	var res []*db.User
	_, err := db.Ormer.QueryTable(&db.User{}).Filter("UserName__icontains", req.UserName).Limit(req.Limit, req.Offset).All(&res)
	if utils.LogContinue(err, utils.Warning) {
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
	d.Status = int32(s.Status)
}

/**
 * @apiDefine DBServerDown
 * @apiError (Error 500) DBServerDown can't connect to database server
 */

/**
 * @api {rpc} /rpc user.AdminUser.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.AdminUser.Create
 * @apiDescription Create new admin user.
 *
 * @apiParam {string} studentId student id.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for exist user
 * @apiSuccess {int32} adminId created or existed adminId
 * @apiUse DBServerDown
 */
func (a *srvAdmin) Create(ctx context.Context, req *user.AdminUserRequest, rsp *user.AdminUserResponse) error {
	if req.StudentId == "" {
		rsp.Status = user.AdminUserResponse_INVALID_PARAM
	} else {
		usr := db.AdminUser{
			StudentId: req.StudentId,
		}
		created, id, err := db.Ormer.ReadOrCreate(&usr, "StudentId")
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		if created {
			rsp.Status = user.AdminUserResponse_SUCCESS
		} else {
			rsp.Status = user.AdminUserResponse_USER_EXIST
		}
		rsp.AdminId = int32(id)
	}
	return nil
}

/**
 * @api {rpc} /rpc user.AdminUser.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName user.AdminUser.Find
 * @apiDescription Find admin user.
 *
 * @apiParam {string} studentId student id.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiSuccess {int32} adminId
 * @apiUse DBServerDown
 */
func (a *srvAdmin) Find(ctx context.Context, req *user.AdminUserRequest, rsp *user.AdminUserResponse) error {
	if req.StudentId == "" {
		rsp.Status = user.AdminUserResponse_INVALID_PARAM
	} else {
		usr := db.AdminUser{
			StudentId: req.StudentId,
		}
		err := db.Ormer.Read(&usr, "StudentId")
		if err == orm.ErrNoRows {
			rsp.Status = user.AdminUserResponse_NOT_FOUND
			return nil
		} else if utils.LogContinue(err, utils.Warning) {
			return err
		}
		rsp.Status = user.AdminUserResponse_SUCCESS
		rsp.AdminId = int32(usr.Id)
	}
	return nil
}

func main() {
	db.InitORM("userdb", new(db.User), new(db.AdminUser))
	service := utils.InitMicroService("user")
	utils.LogPanic(user.RegisterUserHandler(service.Server(), new(srvUser)))
	utils.LogPanic(user.RegisterAdminUserHandler(service.Server(), new(srvAdmin)))
	utils.RunMicroService(service)
}
