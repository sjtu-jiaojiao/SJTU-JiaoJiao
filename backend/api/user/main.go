package main

import (
	"context"
	"jiaojiao/srv/user/mock"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/user/:userID", getUserInfo)
	rg.GET("/user", findUser)
	rg.POST("/user", addUser)
	rg.PUT("/user", updateUser)
	return router
}

/**
 * @apiDefine UserServiceDown
 * @apiError (Error 500) UserServiceDown User service down
 */

/**
 * @api {get} /user/:userID GetUserInfo
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission none/self/admin
 * @apiName GetUserInfo
 * @apiDescription Get user info
 *
 * @apiParam {--} Param see [User Service](#api-Service-User_Query)
 * @apiSuccess (None - Success 200) {Response} response see [User Service](#api-Service-User_Query) <br>
 * 											   studentID: hidden <br> studentName: hidden
 * @apiSuccess (Self/Admin - Success 200) {Response} response see [User Service](#api-Service-User_Query)
 * @apiUse InvalidParam
 * @apiUse UserServiceDown
 */
func getUserInfo(c *gin.Context) {
	type param struct {
		UserID int32 `uri:"userID" binding:"required,min=1"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindUri(&p), utils.Warning) {
		role := utils.GetRoleID(c, p.UserID)

		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Query(context.TODO(), &user.UserQueryRequest{
			UserID: p.UserID,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}
		if !role.Self && !role.Admin {
			rsp.StudentID = ""
			rsp.StudentName = ""
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

/**
 * @api {get} /user FindUser
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission none/self/admin
 * @apiName FindUser
 * @apiDescription Find user
 *
 * @apiParam {--} Param see [User Service](#api-Service-User_Find) <br> No param need admin permission!
 * @apiSuccess (None - Success 200) {Response} response see [User Service](#api-Service-User_Find) <br>
 * 											   studentID: hidden <br> studentName: hidden
 * @apiSuccess (Self/Admin - Success 200) {Response} response see [User Service](#api-Service-User_Find)
 * @apiUse InvalidParam
 * @apiUse UserServiceDown
 */
func findUser(c *gin.Context) {
	type param struct {
		UserName string `form:"userName"`
		Limit    uint32 `form:"limit"`
		Offset   uint32 `form:"offset"`
	}
	var p param
	role := utils.GetRole(c)

	if !utils.LogContinue(c.ShouldBindQuery(&p), utils.Warning) {
		if p.UserName == "" && !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Find(context.TODO(), &user.UserFindRequest{
			UserName: p.UserName,
			Limit:    p.Limit,
			Offset:   p.Offset,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}
		if !role.Admin {
			for _, v := range rsp.User {
				v.StudentID = ""
				v.StudentName = ""
			}
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

/**
 * @api {post} /user AddUser
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission admin
 * @apiName AddUser
 * @apiDescription Add user, use default value.
 *
 * @apiParam {--} Param see [User Service](#api-Service-User_Create)
 * @apiSuccess {Response} response see [User Service](#api-Service-User_Create)
 * @apiUse InvalidParam
 * @apiUse UserServiceDown
 */
func addUser(c *gin.Context) {
	type param struct {
		StudentID   string `form:"studentID" binding:"required"`
		StudentName string `form:"studentName" binding:"required"`
	}
	var p param
	role := utils.GetRole(c)

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Create(context.TODO(), &user.UserCreateRequest{
			StudentID:   p.StudentID,
			StudentName: p.StudentName,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

/**
 * @api {put} /user UpdateUser
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission self/admin
 * @apiName UpdateUser
 * @apiDescription Update user
 *
 * @apiParam {--} Param see [User Service](#api-Service-User_Update) <br> self not allow edit StudentID,StudentName,Status,Role
 * @apiSuccess {Response} response see [User Service](#api-Service-User_Update)
 * @apiUse InvalidParam
 * @apiUse UserServiceDown
 */
func updateUser(c *gin.Context) {
	type param struct {
		UserID      int32  `form:"userID" binding:"required,min=1"`
		UserName    string `form:"userName"`
		AvatarID    string `form:"avatarID"`
		Telephone   string `form:"telephone"`
		StudentID   string `form:"studentID"`
		StudentName string `form:"studentName"`
		Status      int32  `form:"status"`
		Role        int32  `form:"role"`
		ClearEmpty  bool   `form:"clearEmpty"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		role := utils.GetRoleID(c, p.UserID)

		if !role.Self && !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		if !role.Admin { // only admin allow change these fields
			p.StudentID = ""
			p.StudentName = ""
			p.Status = 0
			p.Role = 0
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Update(context.TODO(), &user.UserInfo{
			UserID:      p.UserID,
			UserName:    p.UserName,
			AvatarID:    p.AvatarID,
			Telephone:   p.Telephone,
			StudentID:   p.StudentID,
			StudentName: p.StudentName,
			Status:      user.UserInfo_Status(utils.EnumConvert(p.Status, user.UserInfo_Status_name)),
			Role:        user.UserInfo_Role(utils.EnumConvert(p.Role, user.UserInfo_Role_name)),
			ClearEmpty:  p.ClearEmpty,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("user", setupRouter())
}
