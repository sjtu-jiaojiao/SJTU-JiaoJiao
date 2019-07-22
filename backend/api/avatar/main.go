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
 * @apiParam {--} Param see [Avatar Service](#api-Service-Avatar_Create) <br>
 *						file accept [file type](https://github.com/h2non/filetype#image) <br> Max size is 5M
 * @apiSuccess {Response} response see [Avatar Service](#api-Service-Avatar_Create)
 * @apiUse InvalidParam
 * @apiUse UserServiceDown
 */
func addAvatar(c *gin.Context) {
	type param struct {
		UserId int32 `form:"userId" binding:"required,min=1"`
	}
	var p param

	file, err := c.FormFile("file")
	if err == nil && !utils.LogContinue(c.ShouldBindQuery(&p), utils.Warning) {
		if file.Size > 1024*1024*5 { // 5M
			c.AbortWithStatus(413)
			return
		}

		role := utils.GetRoleID(c, p.UserId)

		if !role.Self && !role.Admin {
			c.AbortWithStatus(403)
			return
		}

		f, err := file.Open()
		if utils.LogContinue(err, utils.Warning) {
			c.JSON(500, err)
			return
		}
		defer f.Close()
		data := make([]byte, file.Size)
		_, err = f.Read(data)
		if utils.LogContinue(err, utils.Warning) {
			c.JSON(500, err)
			return
		}

		srv := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return avatar.NewAvatarService(name, c) },
			func() interface{} { return mock.NewAvatarService() }).(avatar.AvatarService)
		rsp, err := srv.Create(context.TODO(), &avatar.AvatarCreateRequest{
			UserId: p.UserId,
			File:   data,
		})
		if utils.LogContinue(err, utils.Warning, "Avatar service error: %v", err) {
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
