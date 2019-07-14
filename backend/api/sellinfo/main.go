package main

import (
	"context"
	"jiaojiao/srv/sellinfo/mock"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/sellInfo/:sellInfoId", getSellInfo)
	rg.PUT("/content", addContent)
	rg.PUT("/sellInfo", addSellInfo)
	return router
}

type sellInfo struct {
	SellInfoId int32 `uri:"sellInfoId" binding:"required,min=1"`
}

/**
 * @api {get} /sellInfo/:sellInfoId GetSellInfo
 * @apiVersion 1.0.0
 * @apiGroup SellInfo
 * @apiPermission none
 * @apiName GetSellInfo
 * @apiDescription Get sell info
 *
 * @apiParam {--} Param see [SellInfo Service](#api-Service-sellinfo_SellInfo_Query)
 * @apiSuccess (Success 200) {Response} response see [SellInfo Service](#api-Service-sellinfo_SellInfo_Query) <br>
 * @apiError (Error 500) SellInfoServiceDown SellInfo service down
 */
func getSellInfo(c *gin.Context) {
	var info sellInfo
	if !utils.LogContinue(c.ShouldBindUri(&info), utils.Warning) {
		srv := utils.CallMicroService("sellInfo", func(name string, c client.Client) interface{} { return sellinfo.NewSellInfoService(name, c) },
			func() interface{} { return mock.NewSellInfoService() }).(sellinfo.SellInfoService)
		rsp, err := srv.Query(context.TODO(), &sellinfo.SellInfoQueryRequest{
			SellInfoId: info.SellInfoId,
		})
		if utils.LogContinue(err, utils.Warning, "SellInfo service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

type content struct {
	ContentId    string `form:"contentId"`
	ContentToken string `form:"contentToken"`
	Content      []byte `form:"content"`
	Type         int32  `form:"type"`
}

/**
 * @api {put} /content AddContent
 * @apiVersion 1.0.0
 * @apiGroup SellInfo
 * @apiPermission user/admin
 * @apiName AddContent
 * @apiDescription Add sell info content
 *
 * @apiParam {--} Param see [SellInfo Service](#api-Service-sellinfo_Content_Create)
 * @apiSuccess (Success 200) {Response} response see [SellInfo Service](#api-Service-sellinfo_Content_Create) <br>
 * @apiError (Error 500) SellInfoServiceDown SellInfo service down
 */
func addContent(c *gin.Context) {
	var cont content
	if !utils.LogContinue(c.ShouldBindQuery(&cont), utils.Warning) {
		if !utils.CheckUser(c) {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("sellInfo", func(name string, c client.Client) interface{} { return sellinfo.NewContentService(name, c) },
			func() interface{} { return mock.NewContentService() }).(sellinfo.ContentService)
		rsp, err := srv.Create(context.TODO(), &sellinfo.ContentCreateRequest{
			ContentId:    cont.ContentId,
			ContentToken: cont.ContentToken,
			Content:      cont.Content,
			Type:         sellinfo.ContentCreateRequest_Type(cont.Type),
		})
		if utils.LogContinue(err, utils.Warning, "SellInfo service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

type createSellInfo struct {
	ValidTime    int64   `form:"validTime"`
	GoodName     string  `form:"goodName"`
	Price        float64 `form:"price"`
	Description  string  `form:"description"`
	ContentId    string  `form:"contentId"`
	ContentToken string  `form:"contentToken"`
	UserId       int32   `form:"userId"`
}

/**
 * @api {put} /sellInfo AddSellInfo
 * @apiVersion 1.0.0
 * @apiGroup SellInfo
 * @apiPermission self/admin
 * @apiName AddSellInfo
 * @apiDescription Add sell info
 *
 * @apiParam {--} Param see [SellInfo Service](#api-Service-sellinfo_SellInfo_Create)
 * @apiSuccess (Success 200) {Response} response see [SellInfo Service](#api-Service-sellinfo_SellInfo_Create) <br>
 * @apiError (Error 500) SellInfoServiceDown SellInfo service down
 */
func addSellInfo(c *gin.Context) {
	var info createSellInfo
	if !utils.LogContinue(c.ShouldBindQuery(&info), utils.Warning) {
		if !utils.CheckUserId(c, info.UserId) && !utils.CheckAdmin(c) {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("sellInfo", func(name string, c client.Client) interface{} { return sellinfo.NewSellInfoService(name, c) },
			func() interface{} { return mock.NewSellInfoService() }).(sellinfo.SellInfoService)
		rsp, err := srv.Create(context.TODO(), &sellinfo.SellInfoCreateRequest{
			ValidTime:    info.ValidTime,
			GoodName:     info.GoodName,
			Price:        info.Price,
			Description:  info.Description,
			ContentId:    info.ContentId,
			ContentToken: info.ContentToken,
			UserId:       info.UserId,
		})
		if utils.LogContinue(err, utils.Warning, "SellInfo service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("sellInfo", setupRouter())
}
