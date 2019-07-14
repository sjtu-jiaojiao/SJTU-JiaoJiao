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
	rg.PUT("/user", addUser)
	return router
}

type userInfo struct {
	UserId int32 `uri:"userId" binding:"required,min=1"`
}

/**
 * @api {get} /user/:userId GetUserInfo
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission none/self
 * @apiName GetUserInfo
 * @apiDescription Get user info
 *
 * @apiParam {--} Param see [User Service](#api-Service-user_User_Query)
 * @apiSuccess (None - Success 200) {Response} response see [User Service](#api-Service-user_User_Query) <br>
 * 											   studentId: hidden <br> studentName: hidden
 * @apiSuccess (Self - Success 200) {Response} response see [User Service](#api-Service-user_User_Query)
 * @apiError (Error 500) UserServiceDown User service down
 */
func getUserInfo(c *gin.Context) {
	var info userInfo

	if !utils.LogContinue(c.ShouldBindUri(&info), utils.Warning) {
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Query(context.TODO(), &user.UserQueryRequest{
			UserId: info.UserId,
		})
		if utils.LogContinue(err, utils.Warning, "User service error: %v", err) {
			c.JSON(500, err)
			return
		}
		if !utils.CheckUserId(c, info.UserId) {
			rsp.StudentId = ""
			rsp.StudentName = ""
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

type findCond struct {
	UserName string `form:"userName"`
	Limit    uint32 `form:"limit"`
	Offset   uint32 `form:"offset"`
}

/**
 * @api {get} /user FindUser
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission none/admin
 * @apiName FindUser
 * @apiDescription Find user
 *
 * @apiParam {--} Param see [User Service](#api-Service-user_User_Find) <br> No param need admin permission!
 * @apiSuccess {Response} response see [User Service](#api-Service-user_User_Find) <br>
 * 											   None - studentId: hidden <br> None - studentName: hidden
 * @apiError (Error 500) UserServiceDown User service down
 */
func findUser(c *gin.Context) {
	var cond findCond
	p := utils.CheckAdmin(c)

	if !utils.LogContinue(c.ShouldBindQuery(&cond), utils.Warning) {
		if cond.UserName == "" && !p {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Find(context.TODO(), &user.UserFindRequest{
			UserName: cond.UserName,
			Limit:    cond.Limit,
			Offset:   cond.Offset,
		})
		if utils.LogContinue(err, utils.Warning, "User service error: %v", err) {
			c.JSON(500, err)
			return
		}
		if !p {
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

type addInfo struct {
	StudentId   string `form:"studentId"`
	StudentName string `form:"studentName"`
}

/**
 * @api {put} /user AddUser
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission admin
 * @apiName AddUser
 * @apiDescription Add user, use default value.
 *
 * @apiParam {--} Param see [User Service](#api-Service-user_User_Create)
 * @apiSuccess {Response} response see [User Service](#api-Service-user_User_Create) <br>
 * @apiError (Error 500) UserServiceDown User service down
 */
func addUser(c *gin.Context) {
	var info addInfo

	if !utils.LogContinue(c.ShouldBindUri(&info), utils.Warning) {
		if !utils.CheckAdmin(c) {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Create(context.TODO(), &user.UserCreateRequest{
			StudentId:   info.StudentId,
			StudentName: info.StudentName,
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
