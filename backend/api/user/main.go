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
	rg.POST("/user", updateUser)

	rg.GET("/admin/:studentId", findAdmin)
	rg.PUT("/admin", addAdmin)
	rg.DELETE("/admin", deleteAdmin)
	return router
}

type userInfo struct {
	UserId int32 `uri:"userId" binding:"required,min=1"`
}

/**
 * @apiDefine UserServiceDown
 * @apiError (Error 500) UserServiceDown User service down
 */

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
 * @apiUse UserServiceDown
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
 * @apiUse UserServiceDown
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
 * @apiSuccess {Response} response see [User Service](#api-Service-user_User_Create)
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

type updateInfo struct {
	UserId      int32  `form:"userId" binding:"required,min=1"`
	UserName    string `form:"userName"`
	AvatarId    string `form:"avatarId"`
	Telephone   string `form:"telephone"`
	StudentId   string `form:"studentId"`
	StudentName string `form:"studentName"`
	Status      int32  `form:"status"`
	ClearEmpty  bool   `form:"clearEmpty"`
}

/**
 * @api {post} /user UpdateUser
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission self/admin
 * @apiName UpdateUser
 * @apiDescription Update user
 *
 * @apiParam {--} Param see [User Service](#api-Service-user_User_Update)
 * @apiSuccess {Response} response see [User Service](#api-Service-user_User_Update)
 * @apiUse UserServiceDown
 */
func updateUser(c *gin.Context) {
	var usrInfo updateInfo

	if !utils.LogContinue(c.ShouldBind(&usrInfo), utils.Warning) {
		if !utils.CheckAdmin(c) && !utils.CheckUserId(c, usrInfo.UserId) {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
			func() interface{} { return mock.NewUserService() }).(user.UserService)
		rsp, err := srv.Update(context.TODO(), &user.UserInfo{
			UserId:      usrInfo.UserId,
			UserName:    usrInfo.UserName,
			AvatarId:    usrInfo.AvatarId,
			Telephone:   usrInfo.Telephone,
			StudentId:   usrInfo.StudentId,
			StudentName: usrInfo.StudentName,
			Status:      usrInfo.Status,
			ClearEmpty:  usrInfo.ClearEmpty,
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

type addAdminInfo struct {
	StudentId string `form:"studentId"`
}

/**
 * @api {put} /admin AddAdmin
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission admin
 * @apiName AddAdmin
 * @apiDescription Add admin
 *
 * @apiParam {--} Param see [User Service](#api-Service-user_AdminUser_Create)
 * @apiSuccess {Response} response see [User Service](#api-Service-user_AdminUser_Create) <br>
 * @apiError (Error 500) UserServiceDown User service down
 */
func addAdmin(c *gin.Context) {
	var info addAdminInfo

	if !utils.LogContinue(c.ShouldBindUri(&info), utils.Warning) {
		if !utils.CheckAdmin(c) {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewAdminUserService(name, c) },
			func() interface{} { return mock.NewAdminUserService() }).(user.AdminUserService)
		rsp, err := srv.Create(context.TODO(), &user.AdminUserRequest{
			StudentId: info.StudentId,
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

type findAdminInfo struct {
	StudentId string `uri:"studentId" binding:"required,min=1"`
}

/**
 * @api {get} /admin/:studentId FindAdmin
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission admin
 * @apiName FindAdmin
 * @apiDescription Find admin
 *
 * @apiParam {--} Param see [User Service](#api-Service-user_AdminUser_Find)
 * @apiSuccess {Response} response see [User Service](#api-Service-user_AdminUser_Find) <br>
 * @apiError (Error 500) UserServiceDown User service down
 */
func findAdmin(c *gin.Context) {
	var info findAdminInfo

	if !utils.LogContinue(c.ShouldBindUri(&info), utils.Warning) {
		if !utils.CheckAdmin(c) {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewAdminUserService(name, c) },
			func() interface{} { return mock.NewAdminUserService() }).(user.AdminUserService)
		rsp, err := srv.Find(context.TODO(), &user.AdminUserRequest{
			StudentId: info.StudentId,
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
 * @api {delete} /admin DeleteAdmin
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission admin
 * @apiName DeleteAdmin
 * @apiDescription Delete admin
 *
 * @apiParam {--} Param see [User Service](#api-Service-user_AdminUser_Delete)
 * @apiSuccess {Response} response see [User Service](#api-Service-user_AdminUser_Delete) <br>
 * @apiError (Error 500) UserServiceDown User service down
 */
func deleteAdmin(c *gin.Context) {
	var info addAdminInfo

	if !utils.LogContinue(c.ShouldBindUri(&info), utils.Warning) {
		if !utils.CheckAdmin(c) {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewAdminUserService(name, c) },
			func() interface{} { return mock.NewAdminUserService() }).(user.AdminUserService)
		rsp, err := srv.Delete(context.TODO(), &user.AdminUserRequest{
			StudentId: info.StudentId,
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
