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
	rg.GET("/message/:userID", getMessage)
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
 * @apiPermission self
 * @apiName AddMessage
 * @apiDescription Add chat message
 *
 * @apiParam {--} Param see [Message Service](#api-Service-Message_Create) <br> Self userID should in fromUser
 * @apiSuccess (Success 200) {Response} response see [Message Service](#api-Service-Message_Create)
 * @apiUse InvalidParam
 * @apiUse MessageServiceDown
 */
func addMessage(c *gin.Context) {
	type param struct {
		FromUser int32 `form:"fromUser" binding:"required"`
		ToUser   int32 `form:"toUser" binding:"required"`
		Type     int32 `form:"type" binding:"required"`
	}
	var p param
	var msg []byte

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if p.Type == int32(message.Type_PICTURE) || p.Type == int32(message.Type_VIDEO) {
			var code int
			msg, code, _ = utils.GetQueryFile(c, "msg", int64(utils.If(p.Type == int32(message.Type_PICTURE), 1024*1024*5, 1024*1024*50).(int)))
			if code != 200 {
				c.AbortWithStatus(code)
				return
			}
		} else {
			s := c.PostForm("msg")
			if s != "" {
				msg = []byte(s)
			} else {
				c.AbortWithStatus(400)
				return
			}
		}
		role := utils.GetRoleID(c, int32(p.FromUser))
		if !role.Self {
			c.Status(403)
			return
		}

		srv := utils.CallMicroService("message", func(name string, c client.Client) interface{} { return message.NewMessageService(name, c) },
			func() interface{} { return mock.NewMessageService() }).(message.MessageService)
		rsp, err := srv.Create(context.TODO(), &message.MessageCreateRequest{
			FromUser: p.FromUser,
			ToUser:   p.ToUser,
			Type:     message.Type(utils.EnumConvert(p.Type, message.Type_name)),
			Msg:      msg,
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
		FromUser int32  `form:"fromUser" binding:"required"`
		ToUser   int32  `form:"toUser" binding:"required"`
		Way      int32  `form:"way" binding:"required"`
		Limit    uint32 `form:"limit"`
		Offset   uint32 `form:"offset"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindQuery(&p), utils.Warning) {
		role1 := utils.GetRoleID(c, int32(p.FromUser))
		role2 := utils.GetRoleID(c, int32(p.ToUser))
		if !(role1.Self || role1.Admin || role2.Self || role2.Admin) {
			c.Status(403)
			return
		}

		srv := utils.CallMicroService("message", func(name string, c client.Client) interface{} { return message.NewMessageService(name, c) },
			func() interface{} { return mock.NewMessageService() }).(message.MessageService)
		rsp, err := srv.Find(context.TODO(), &message.MessageFindRequest{
			FromUser: p.FromUser,
			ToUser:   p.ToUser,
			Way:      message.MessageFindRequest_Way(utils.EnumConvert(p.Way, message.MessageFindRequest_Way_name)),
			Limit:    p.Limit,
			Offset:   p.Offset,
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
 * @api {get} /message/:userID GetMessage
 * @apiVersion 1.0.0
 * @apiGroup Message
 * @apiPermission self/admin
 * @apiName GetMessage
 * @apiDescription Get all (new) message about user
 *
 * @apiParam {bool} oldMsg=0 true to get all message, not only the new
 * @apiSuccess (Success 200) {Response} response see [Message Service](#api-Service-Message_Query)
 * @apiUse InvalidParam
 * @apiUse MessageServiceDown
 */
func getMessage(c *gin.Context) {
	type param struct {
		UserID int32 `uri:"userID" binding:"required,min=1"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindUri(&p), utils.Warning) {
		role := utils.GetRoleID(c, int32(p.UserID))
		if !role.Self && !role.Admin {
			c.Status(403)
			return
		}

		srv := utils.CallMicroService("message", func(name string, c client.Client) interface{} { return message.NewMessageService(name, c) },
			func() interface{} { return mock.NewMessageService() }).(message.MessageService)
		rsp, err := srv.Query(context.TODO(), &message.MessageQueryRequest{
			UserID: p.UserID,
			OldMsg: c.DefaultQuery("oldMsg", "0") == "1",
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
	utils.RunWebService("message", setupRouter())
}
