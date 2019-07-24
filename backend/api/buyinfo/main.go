package main

import (
	"context"
	"jiaojiao/srv/buyinfo/mock"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/buyInfo/:buyInfoID", getBuyInfo)
	rg.GET("/buyInfo", findBuyInfo)
	rg.POST("/buyInfo", addBuyInfo)
	return router
}

/**
 * @apiDefine BuyInfoServiceDown
 * @apiError (Error 500) BuyInfoServiceDown BuyInfo service down
 */

/**
 * @api {get} /buyInfo/:buyInfoID GetBuyInfo
 * @apiVersion 1.0.0
 * @apiGroup BuyInfo
 * @apiPermission none
 * @apiName GetBuyInfo
 * @apiDescription Get buy info
 *
 * @apiParam {--} Param see [BuyInfo Service](#api-Service-BuyInfo_Query)
 * @apiSuccess (Success 200) {Response} response see [BuyInfo Service](#api-Service-BuyInfo_Query)
 * @apiUse InvalidParam
 * @apiUse BuyInfoServiceDown
 */
func getBuyInfo(c *gin.Context) {
	type param struct {
		BuyInfoID int32 `uri:"buyInfoID" binding:"required,min=1"`
	}
	var p param
	if !utils.LogContinue(c.ShouldBindUri(&p), utils.Warning) {
		srv := utils.CallMicroService("buyInfo", func(name string, c client.Client) interface{} { return buyinfo.NewBuyInfoService(name, c) },
			func() interface{} { return mock.NewBuyInfoService() }).(buyinfo.BuyInfoService)
		rsp, err := srv.Query(context.TODO(), &buyinfo.BuyInfoQueryRequest{
			BuyInfoID: p.BuyInfoID,
		})
		if utils.LogContinue(err, utils.Warning, "BuyInfo service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

/**
 * @api {post} /buyInfo AddBuyInfo
 * @apiVersion 1.0.0
 * @apiGroup BuyInfo
 * @apiPermission self/admin
 * @apiName AddBuyInfo
 * @apiDescription Add buy info
 *
 * @apiParam {--} Param see [BuyInfo Service](#api-Service-BuyInfo_Create)
 * @apiSuccess (Success 200) {Response} response see [BuyInfo Service](#api-Service-BuyInfo_Create)
 * @apiUse InvalidParam
 * @apiUse BuyInfoServiceDown
 */
func addBuyInfo(c *gin.Context) {
	type param struct {
		ValidTime    int64   `form:"validTime" binding:"required"`
		GoodName     string  `form:"goodName" binding:"required"`
		Price        float64 `form:"price"`
		Description  string  `form:"description"`
		ContentID    string  `form:"contentID"`
		ContentToken string  `form:"contentToken"`
		UserID       int32   `form:"userID" binding:"required"`
	}
	var p param
	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if (p.ContentID == "" && p.ContentToken != "") || (p.ContentID != "" && p.ContentToken == "") {
			c.AbortWithStatus(400)
			return
		}

		role := utils.GetRoleID(c, p.UserID)

		if !role.Self && !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("buyInfo", func(name string, c client.Client) interface{} { return buyinfo.NewBuyInfoService(name, c) },
			func() interface{} { return mock.NewBuyInfoService() }).(buyinfo.BuyInfoService)
		rsp, err := srv.Create(context.TODO(), &buyinfo.BuyInfoCreateRequest{
			ValidTime:    p.ValidTime,
			GoodName:     p.GoodName,
			Price:        p.Price,
			Description:  p.Description,
			ContentID:    p.ContentID,
			ContentToken: p.ContentToken,
			UserID:       p.UserID,
		})
		if utils.LogContinue(err, utils.Warning, "BuyInfo service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

/**
 * @api {get} /buyInfo FindBuyInfo
 * @apiVersion 1.0.0
 * @apiGroup BuyInfo
 * @apiPermission none
 * @apiName FindBuyInfo
 * @apiDescription Find buy info
 *
 * @apiParam {--} Param see [BuyInfo Service](#api-Service-BuyInfo_Find)
 * @apiSuccess {Response} response see [BuyInfo Service](#api-Service-BuyInfo_Find)
 * @apiUse InvalidParam
 * @apiUse BuyInfoServiceDown
 */
func findBuyInfo(c *gin.Context) {
	type param struct {
		UserID    int32   `form:"userID"`
		Status    int32   `form:"status"`
		GoodName  string  `form:"goodName"`
		LowPrice  float64 `form:"lowPrice"`
		HighPrice float64 `form:"highPrice"`
		Limit     uint32  `form:"limit"`
		Offset    uint32  `form:"offset"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindQuery(&p), utils.Warning) {
		srv := utils.CallMicroService("buyInfo", func(name string, c client.Client) interface{} { return buyinfo.NewBuyInfoService(name, c) },
			func() interface{} { return mock.NewBuyInfoService() }).(buyinfo.BuyInfoService)
		rsp, err := srv.Find(context.TODO(), &buyinfo.BuyInfoFindRequest{
			UserID:    p.UserID,
			Status:    p.Status,
			GoodName:  p.GoodName,
			LowPrice:  p.LowPrice,
			HighPrice: p.HighPrice,
			Limit:     p.Limit,
			Offset:    p.Offset,
		})
		if utils.LogContinue(err, utils.Warning, "BuyInfo service error: %v", err) {
			c.JSON(500, err)
			return
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("buyInfo", setupRouter())
}
