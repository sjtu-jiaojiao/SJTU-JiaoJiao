package main

import (
	"context"
	"jiaojiao/srv/content/mock"
	content "jiaojiao/srv/content/proto"
	"jiaojiao/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.POST("/content", addContent)
	rg.DELETE("/content", deleteContent)
	rg.PUT("/content", updateContent)
	rg.GET("/content/:contentID", getContent)
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
 * @apiParam {--} Param see [Content Service](#api-Service-Content_Create) <br> Max size is 50M
 * @apiSuccess (Success 200) {Response} response see [Content Service](#api-Service-Content_Create)
 * @apiUse InvalidParam
 * @apiUse ContentServiceDown
 */
func addContent(c *gin.Context) {
	type param struct {
		ContentID    string `form:"contentID"`
		ContentToken string `form:"contentToken"`
		Type         int32  `form:"type" binding:"required"`
	}
	var p param
	role := utils.GetRole(c)
	data, code, _ := utils.GetQueryFile(c, "content", 1024*1024*50) // 50M

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		// check param&role
		if code != 200 {
			c.AbortWithStatus(code)
			return
		}

		if (p.ContentID == "" && p.ContentToken != "") || (p.ContentID != "" && p.ContentToken == "") {
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
			ContentID:    p.ContentID,
			ContentToken: p.ContentToken,
			Content:      data,
			Type:         content.Type(utils.EnumConvert(p.Type, content.Type_name)),
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
		ContentID    string `form:"contentID" binding:"required"`
		ContentToken string `form:"contentToken" binding:"required"`
		FileID       string `form:"fileID"`
	}
	var p param
	role := utils.GetRole(c)

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if !role.User && !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("content", func(name string, c client.Client) interface{} { return content.NewContentService(name, c) },
			func() interface{} { return mock.NewContentService() }).(content.ContentService)
		rsp, err := srv.Delete(context.TODO(), &content.ContentDeleteRequest{
			ContentID:    p.ContentID,
			ContentToken: p.ContentToken,
			FileID:       p.FileID,
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
 * @api {put} /content UpdateContent
 * @apiVersion 1.0.0
 * @apiGroup Content
 * @apiPermission user/admin
 * @apiName UpdateContent
 * @apiDescription Update sell info content
 *
 * @apiParam {--} Param see [Content Service](#api-Service-Content_Update) <br> Max size is 50M
 * @apiSuccess (Success 200) {Response} response see [Content Service](#api-Service-Content_Update)
 * @apiUse InvalidParam
 * @apiUse ContentServiceDown
 */
func updateContent(c *gin.Context) {
	type param struct {
		ContentID    string `form:"contentID" binding:"required"`
		ContentToken string `form:"contentToken" binding:"required"`
		FileID       string `form:"fileID" binding:"required"`
		Type         int32  `form:"type"`
	}
	var p param
	role := utils.GetRole(c)
	data, code, err := utils.GetQueryFile(c, "content", 1024*1024*50) // 50M
	if err == http.ErrMissingFile {                                   // allow no file
		code = 200
		err = nil
	}

	if err == nil && !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if code != 200 {
			c.AbortWithStatus(code)
			return
		}

		if !role.User && !role.Admin {
			c.AbortWithStatus(403)
			return
		}

		srv := utils.CallMicroService("content", func(name string, c client.Client) interface{} { return content.NewContentService(name, c) },
			func() interface{} { return mock.NewContentService() }).(content.ContentService)
		rsp, err := srv.Update(context.TODO(), &content.ContentUpdateRequest{
			ContentID:    p.ContentID,
			ContentToken: p.ContentToken,
			FileID:       p.FileID,
			Content:      data,
			Type:         content.Type(utils.EnumConvert(p.Type, content.Type_name)),
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
 * @api {get} /content/:contentID GetContent
 * @apiVersion 1.0.0
 * @apiGroup Content
 * @apiPermission none/self/admin
 * @apiName GetContent
 * @apiDescription Get sell info content
 *
 * @apiParam {int32} [userID] user id, left empty for guest
 * @apiParam {--} Param see [Content Service](#api-Service-Content_Query)
 * @apiSuccess (None - Success 200) {Response} response see [Content Service](#api-Service-Content_Query) <br>
 *														contentToken: hidden
 * @apiSuccess (Self/Admin - Success 200) {Response} response see [Content Service](#api-Service-Content_Query)
 * @apiUse InvalidParam
 * @apiUse ContentServiceDown
 */
func getContent(c *gin.Context) {
	type param struct {
		ContentID string `uri:"contentID" binding:"required,min=1"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindUri(&p), utils.Warning) {
		tmp := c.DefaultQuery("userID", "0")
		userID, err := strconv.Atoi(tmp)
		if utils.LogContinue(err, utils.Warning) {
			c.AbortWithStatus(400)
			return
		}
		role := utils.GetRoleID(c, int32(userID))

		srv := utils.CallMicroService("content", func(name string, c client.Client) interface{} { return content.NewContentService(name, c) },
			func() interface{} { return mock.NewContentService() }).(content.ContentService)
		rsp, err := srv.Query(context.TODO(), &content.ContentQueryRequest{
			ContentID: p.ContentID,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}
		if !role.User && !role.Admin {
			rsp.ContentToken = ""
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("content", setupRouter())
}
