package main

import (
	"context"
	"jiaojiao/srv/message/mock"
	message "jiaojiao/srv/message/proto"
	"jiaojiao/utils"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/message", findMessage)
	rg.POST("/message", addMessage)
	return router
}

/**
 * @apiDefine MessageServiceDown
 * @apiError (Error 500) MessageServiceDown Message service down
 */

/**
 * @api {post} /message AddMessage
 * @apiVersion 1.0.0
 * @apiGroup Message
 * @apiPermission self/admin
 * @apiName AddMessage
 * @apiDescription Add chat message
 *
 * @apiParam {--} Param see [Message Service](#api-Service-Message_Create)
 * @apiSuccess (Success 200) {Response} response see [Message Service](#api-Service-Message_Create)
 * @apiUse InvalidParam
 * @apiUse MessageServiceDown
 */
func addMessage(c *gin.Context) {
	type param struct {
		FromUser int32                             `form:"fromUser" binding:"required"`
		ToUser   int32                             `form:"toUser" binding:"required"`
		Type     message.MessageCreateRequest_Type `form:"type" binding:"required"`
		Text     string                            `form:"text"`
	}
	var p param
	var data []byte

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if p.Type == message.MessageCreateRequest_PICTURE || p.Type == message.MessageCreateRequest_VIDEO {
			var code int
			var err error
			data, code, err = utils.GetQueryFile(c, "file", 1024*1024*50) // 50M
			if err != nil {
				c.AbortWithStatus(400)
				return
			}
			if code != 200 {
				c.AbortWithStatus(code)
				return
			}
		}
		role1 := utils.GetRoleID(c, int32(p.FromUser))
		role2 := utils.GetRoleID(c, int32(p.FromUser))
		if !(role1.Self || role1.Admin || role2.Self || role2.Admin) {
			c.Status(403)
			return
		}

		srv := utils.CallMicroService("message", func(name string, c client.Client) interface{} { return message.NewMessageService(name, c) },
			func() interface{} { return mock.NewMessageService() }).(message.MessageService)
		rsp, err := srv.Create(context.TODO(), &message.MessageCreateRequest{
			FromUser: p.FromUser,
			ToUser:   p.ToUser,
			Type:     p.Type,
			Text:     p.Text,
			File:     data,
		})
		if utils.LogContinue(err, utils.Warning, "Message service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

/**
 * @api {get} /message FindMessage
 * @apiVersion 1.0.0
 * @apiGroup Message
 * @apiPermission self/admin
 * @apiName FindMessage
 * @apiDescription Find chat message
 *
 * @apiParam {--} Param see [Message Service](#api-Service-Message_Find)
 * @apiSuccess (Success 200) {Response} response see [Message Service](#api-Service-Message_Find)
 * @apiUse InvalidParam
 * @apiUse MessageServiceDown
 */
func findMessage(c *gin.Context) {
	type param struct {
		FromUser int32                          `form:"fromUser" binding:"required"`
		ToUser   int32                          `form:"toUser" binding:"required"`
		Way      message.MessageFindRequest_Way `form:"way" binding:"required"`
		Limit    int32                          `form:"limit"`
		Offset   int32                          `form:"offset"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		role1 := utils.GetRoleID(c, int32(p.FromUser))
		role2 := utils.GetRoleID(c, int32(p.FromUser))
		if !(role1.Self || role1.Admin || role2.Self || role2.Admin) {
			c.Status(403)
			return
		}

		srv := utils.CallMicroService("message", func(name string, c client.Client) interface{} { return message.NewMessageService(name, c) },
			func() interface{} { return mock.NewMessageService() }).(message.MessageService)
		rsp, err := srv.Find(context.TODO(), &message.MessageFindRequest{
			FromUser: p.FromUser,
			ToUser:   p.ToUser,
			Way:      p.Way,
		})
		if utils.LogContinue(err, utils.Warning, "Message service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("message", setupRouter())
}
