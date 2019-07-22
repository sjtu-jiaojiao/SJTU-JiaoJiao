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
	rg.GET("/user/:userId", getUserInfo)
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
 * @api {get} /user/:userId GetUserInfo
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission none/self/admin
 * @apiName GetUserInfo
 * @apiDescription Get user info
 *
 * @apiParam {--} Param see [User Service](#api-Service-User_Query)
 * @apiSuccess (None - Success 200) {Response} response see [User Service](#api-Service-User_Query) <br>
 * 											   studentId: hidden <br> studentName: hidden
 * @apiSuccess (Self/Admin - Success 200) {Response} response see [User Service](#api-Service-User_Query)
 * @apiUse InvalidParam
 * @apiUse UserServiceDown
 */
func getUserInfo(c *gin.Context) {
	type param struct {
		UserId int32 `uri:"userId" binding:"required,min=1"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindUri(&p), utils.Warning) {
		role := utils.GetRoleID(c, p.UserId)

		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Query(context.TODO(), &user.UserQueryRequest{
			UserId: p.UserId,
		})
		if utils.LogContinue(err, utils.Warning, "User service error: %v", err) {
			c.JSON(500, err)
			return
		}
		if !role.Self && !role.Admin {
			rsp.StudentId = ""
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
 * @apiPermission none/admin
 * @apiName FindUser
 * @apiDescription Find user
 *
 * @apiParam {--} Param see [User Service](#api-Service-User_Find) <br> No param need admin permission!
 * @apiSuccess {Response} response see [User Service](#api-Service-User_Find) <br>
 * 											   None - studentId: hidden <br> None - studentName: hidden
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
		if utils.LogContinue(err, utils.Warning, "User service error: %v", err) {
			c.JSON(500, err)
			return
		}
		if !role.Admin {
			for _, v := range rsp.User {
				v.StudentId = ""
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
		StudentId   string `form:"studentId" binding:"required"`
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
			StudentId:   p.StudentId,
			StudentName: p.StudentName,
		})
		if utils.LogContinue(err, utils.Warning, "User service error: %v", err) {
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
 * @apiParam {--} Param see [User Service](#api-Service-User_Update) <br> self not allow edit StudentId,StudentName,Status,Role
 * @apiSuccess {Response} response see [User Service](#api-Service-User_Update)
 * @apiUse InvalidParam
 * @apiUse UserServiceDown
 */
func updateUser(c *gin.Context) {
	type param struct {
		UserId      int32  `form:"userId" binding:"required,min=1"`
		UserName    string `form:"userName"`
		AvatarId    string `form:"avatarId"`
		Telephone   string `form:"telephone"`
		StudentId   string `form:"studentId"`
		StudentName string `form:"studentName"`
		Status      int32  `form:"status"`
		Role        int32  `form:"role"`
		ClearEmpty  bool   `form:"clearEmpty"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		role := utils.GetRoleID(c, p.UserId)

		if !role.Self && !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		if !role.Admin { // only admin allow change these fields
			p.StudentId = ""
			p.StudentName = ""
			p.Status = 0
			p.Role = 0
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Update(context.TODO(), &user.UserInfo{
			UserId:      p.UserId,
			UserName:    p.UserName,
			AvatarId:    p.AvatarId,
			Telephone:   p.Telephone,
			StudentId:   p.StudentId,
			StudentName: p.StudentName,
			Status:      user.UserInfo_Status(p.Status),
			Role:        user.UserInfo_Role(p.Role),
			ClearEmpty:  p.ClearEmpty,
		})
		if utils.LogContinue(err, utils.Warning, "User service error: %v", err) {
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
