package main

import (
	"context"
	db "jiaojiao/database"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"

	"github.com/jinzhu/gorm"
)

type srv struct{}

/**
 * @apiDefine DBServerDown
 * @apiError (Error 500) DBServerDown can't connect to database server
 */

/**
 * @api {rpc} /rpc User.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName User.Create
 * @apiDescription Create new user.
 *
 * @apiParam {string} studentID student id.
 * @apiParam {string} studentName student name.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for exist user
 * @apiSuccess {Response} user see [User Service](#api-Service-User_Query)
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *user.UserCreateRequest, rsp *user.UserCreateResponse) error {
	if !utils.RequireParam(req.StudentID, req.StudentName) {
		rsp.Status = user.UserCreateResponse_INVALID_PARAM
		return nil
	}

	var usr db.User
	err := db.Ormer.Where("student_id = ?", req.StudentID).First(&usr).Error
	if gorm.IsRecordNotFoundError(err) {
		usr = db.User{
			UserName:    req.StudentName,
			AvatarID:    utils.GetStringConfig("srv_config", "default_avatar"),
			StudentID:   req.StudentID,
			StudentName: req.StudentName,
		}
		utils.LogContinue(db.Ormer.Create(&usr).Error, utils.Warning)
		rsp.Status = user.UserCreateResponse_SUCCESS
	} else if utils.LogContinue(err, utils.Error) {
		return err
	} else {
		rsp.Status = user.UserCreateResponse_USER_EXIST
	}
	rsp.User = new(user.UserInfo)
	parseUser(&usr, rsp.User)
	return nil
}

/**
 * @api {rpc} /rpc User.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName User.Query
 * @apiDescription Query user info.
 *
 * @apiParam {int32} userID user id
 * @apiSuccess {int32} userID user id
 * @apiSuccess {string} userName user name
 * @apiSuccess {string} avatarID user avatar id
 * @apiSuccess {string} telephone user telephone
 * @apiSuccess {string} studentID student id
 * @apiSuccess {string} studentName student name
 * @apiSuccess {int32} status user status, 1 for normal <br> 2 for frozen
 * @apiSuccess {int32} role user role, 1 for user <br> 10 for admin
 * @apiUse DBServerDown
 */
func (a *srv) Query(ctx context.Context, req *user.UserQueryRequest, rsp *user.UserInfo) error {
	if !utils.RequireParam(req.UserID) {
		return nil
	}

	usr := db.User{
		ID: req.UserID,
	}
	err := db.Ormer.First(&usr).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Error) {
		return err
	}
	parseUser(&usr, rsp)
	return nil
}

/**
 * @api {rpc} /rpc User.Update
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName User.Update
 * @apiDescription Update user info, only update provided field. If clearEmpty=1 and param support allow clear, clear the field when not provided.
 *
 * @apiParam {int32} userID user id
 * @apiParam {string} [userName] user name
 * @apiParam {string} [avatarID] user avatar id
 * @apiParam {string} [telephone] user telephone, allow clear
 * @apiParam {string} [studentID] student id
 * @apiParam {string} [studentName] student name
 * @apiParam {int32} [status] user status, 1 for normal <br> 2 for frozen
 * @apiParam {int32} [role] user role, 1 for user <br> 10 for admin
 * @apiParam {bool} clearEmpty=0 clear the empty field
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for user not found
 * @apiUse DBServerDown
 */
func (a *srv) Update(ctx context.Context, req *user.UserInfo, rsp *user.UserUpdateResponse) error {
	if !utils.RequireParam(req.UserID) {
		rsp.Status = user.UserUpdateResponse_INVALID_PARAM
		return nil
	}

	usr := db.User{
		ID: req.UserID,
	}
	err := db.Ormer.First(&usr).Error
	if err == nil {
		utils.AssignNotEmpty(&req.UserName, &usr.UserName)
		utils.AssignNotEmpty(&req.AvatarID, &usr.AvatarID)
		if req.ClearEmpty {
			usr.Telephone = req.Telephone
		} else {
			utils.AssignNotEmpty(&req.Telephone, &usr.Telephone)
		}
		utils.AssignNotEmpty(&req.StudentID, &usr.StudentID)
		utils.AssignNotEmpty(&req.StudentName, &usr.StudentName)
		utils.AssignNotZero((*int32)(&req.Status), &usr.Status)
		utils.AssignNotZero((*int32)(&req.Role), &usr.Role)
		err := db.Ormer.Save(&usr).Error
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		rsp.Status = user.UserUpdateResponse_SUCCESS
	} else if gorm.IsRecordNotFoundError(err) {
		rsp.Status = user.UserUpdateResponse_NOT_FOUND
		return nil
	} else {
		utils.Error(err)
		return err
	}
	return nil
}

/**
 * @api {rpc} /rpc User.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName User.Find
 * @apiDescription Find user(fuzzy).
 *
 * @apiParam {string} [userName] username
 * @apiParam {uint32{0-100}} limit=100 row limit
 * @apiParam {uint32} offset=0 row offset
 * @apiSuccess {array} user see [User Service](#api-Service-User_Query)
 * @apiUse DBServerDown
 */
func (a *srv) Find(ctx context.Context, req *user.UserFindRequest, rsp *user.UserFindResponse) error {
	if req.Limit == 0 {
		req.Limit = 100
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	var res []*db.User
	err := db.Ormer.Where("user_name LIKE ?", "%"+req.UserName+"%").Limit(req.Limit).Offset(req.Offset).Find(&res).Error
	if utils.LogContinue(err, utils.Error) {
		return err
	}
	for i, v := range res {
		rsp.User = append(rsp.User, new(user.UserInfo))
		parseUser(v, rsp.User[i])
	}
	return nil
}

func parseUser(s *db.User, d *user.UserInfo) {
	d.UserID = int32(s.ID)
	d.UserName = s.UserName
	d.AvatarID = s.AvatarID
	d.Telephone = s.Telephone
	d.StudentID = s.StudentID
	d.StudentName = s.StudentName
	d.Status = user.UserInfo_Status(utils.EnumConvert(s.Status, user.UserInfo_Status_name))
	d.Role = user.UserInfo_Role(utils.EnumConvert(s.Role, user.UserInfo_Role_name))
}

func main() {
	db.InitORM("userdb", new(db.User))
	defer db.CloseORM()
	service := utils.InitMicroService("user")
	utils.LogPanic(user.RegisterUserHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
