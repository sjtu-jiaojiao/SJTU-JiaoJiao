package main

import (
	"context"
	"jiaojiao/srv/content/mock"
	content "jiaojiao/srv/content/proto"
	"jiaojiao/utils"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.POST("/content", addContent)
	rg.DELETE("/content", deleteContent)
	rg.PUT("/content", updateContent)
	return router
}

/**
 * @apiDefine ContentServiceDown
 * @apiError (Error 500) ContentServiceDown Content service down
 */

/**
 * @api {post} /content AddContent
 * @apiVersion 1.0.0
 * @apiGroup Content
 * @apiPermission user/admin
 * @apiName AddContent
 * @apiDescription Add sell info content
 *
 * @apiParam {--} Param see [Content Service](#api-Service-Content_Create)
 * @apiSuccess (Success 200) {Response} response see [Content Service](#api-Service-Content_Create)
 * @apiUse InvalidParam
 * @apiUse ContentServiceDown
 */
func addContent(c *gin.Context) {
	type param struct {
		ContentId    string `form:"contentId"`
		ContentToken string `form:"contentToken"`
		Content      []byte `form:"content" binding:"required"`
		Type         int32  `form:"type" binding:"required"`
	}
	var p param
	role := utils.GetRole(c)

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if (p.ContentId == "" && p.ContentToken != "") || (p.ContentId != "" && p.ContentToken == "") {
			c.AbortWithStatus(400)
			return
		}

		if !role.User && !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("content", func(name string, c client.Client) interface{} { return content.NewContentService(name, c) },
			func() interface{} { return mock.NewContentService() }).(content.ContentService)
		rsp, err := srv.Create(context.TODO(), &content.ContentCreateRequest{
			ContentId:    p.ContentId,
			ContentToken: p.ContentToken,
			Content:      p.Content,
			Type:         content.ContentCreateRequest_Type(p.Type),
		})
		if utils.LogContinue(err, utils.Warning, "Content service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

/**
 * @api {delete} /content DeleteContent
 * @apiVersion 1.0.0
 * @apiGroup Content
 * @apiPermission user/admin
 * @apiName DeleteContent
 * @apiDescription Delete sell info content
 *
 * @apiParam {--} Param see [Content Service](#api-Service-Content_Delete)
 * @apiSuccess (Success 200) {Response} response see [Content Service](#api-Service-Content_Delete)
 * @apiUse InvalidParam
 * @apiUse ContentServiceDown
 */
func deleteContent(c *gin.Context) {
	type param struct {
		ContentId    string `form:"contentId" binding:"required"`
		ContentToken string `form:"contentToken" binding:"required"`
	}
	var q param
	role := utils.GetRole(c)

	if !utils.LogContinue(c.ShouldBindQuery(&q), utils.Warning) {
		if !role.User && !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("content", func(name string, c client.Client) interface{} { return content.NewContentService(name, c) },
			func() interface{} { return mock.NewContentService() }).(content.ContentService)
		rsp, err := srv.Delete(context.TODO(), &content.ContentDeleteRequest{
			ContentId:    q.ContentId,
			ContentToken: q.ContentToken,
		})
		if utils.LogContinue(err, utils.Warning, "Content service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

/**
 * @api {update} /content UpdateContent
 * @apiVersion 1.0.0
 * @apiGroup Content
 * @apiPermission user/admin
 * @apiName UpdateContent
 * @apiDescription Update sell info content
 *
 * @apiParam {--} Param see [Content Service](#api-Service-Content_Update)
 * @apiSuccess (Success 200) {Response} response see [Content Service](#api-Service-Content_Update)
 * @apiUse InvalidParam
 * @apiUse ContentServiceDown
 */
func updateContent(c *gin.Context) {
	type param struct {
		ContentId    string `form:"contentId" binding:"required"`
		ContentToken string `form:"contentToken" binding:"required"`
		FileId       string `form:"fileId" binding:"required"`
		Content      []byte `form:"content" binding:"required"`
		Type         int32  `form:"type" binding:"required"`
	}
	var q param
	role := utils.GetRole(c)

	if !utils.LogContinue(c.ShouldBindQuery(&q), utils.Warning) {
		if !role.User && !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("content", func(name string, c client.Client) interface{} { return content.NewContentService(name, c) },
			func() interface{} { return mock.NewContentService() }).(content.ContentService)
		rsp, err := srv.Update(context.TODO(), &content.ContentUpdateRequest{
			ContentId:    q.ContentId,
			ContentToken: q.ContentToken,
			FileId:       q.FileId,
			Content:      q.Content,
			Type:         content.ContentUpdateRequest_Type(q.Type),
		})
		if utils.LogContinue(err, utils.Warning, "Content service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("content", setupRouter())
}
