package main

import (
	"context"
	"jiaojiao/srv/avatar/mock"
	avatar "jiaojiao/srv/avatar/proto"
	"jiaojiao/utils"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.POST("/avatar", addAvatar)
	return router
}

/**
 * @api {post} /avatar AddAvatar
 * @apiVersion 1.0.0
 * @apiGroup Avatar
 * @apiPermission self/admin
 * @apiName AddAvatar
 * @apiDescription Add avatar
 *
 * @apiParam {--} Param see [Avatar Service](#api-Service-Avatar_Create) <br> Max size is 5M
 * @apiSuccess {Response} response see [Avatar Service](#api-Service-Avatar_Create)
 * @apiUse InvalidParam
 * @apiUse UserServiceDown
 */
func addAvatar(c *gin.Context) {
	type param struct {
		UserID int32 `form:"userID" binding:"required,min=1"`
	}
	var p param

	data, code, _ := utils.GetQueryFile(c, "file", 1024*1024*5) // 5M
	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if code != 200 {
			c.AbortWithStatus(code)
			return
		}

		role := utils.GetRoleID(c, p.UserID)
		if !role.Self && !role.Admin {
			c.AbortWithStatus(403)
			return
		}

		srv := utils.CallMicroService("avatar", func(name string, c client.Client) interface{} { return avatar.NewAvatarService(name, c) },
			func() interface{} { return mock.NewAvatarService() }).(avatar.AvatarService)
		rsp, err := srv.Create(context.TODO(), &avatar.AvatarCreateRequest{
			UserID: p.UserID,
			File:   data,
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
	utils.RunWebService("avatar", setupRouter())
}
