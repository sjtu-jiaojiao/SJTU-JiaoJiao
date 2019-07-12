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
	return router
}

type sellInfo struct {
	SellInfoId int32 `uri:"sellInfoId" binding:"required,min=1"`
}

/**
 * @api {get} /sellInfo/:sellInfoId GetSellInfo
 * @apiVersion 1.0.0
 * @apiGroup SellInfo
 * @apiPermission none/self
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
		if !utils.CheckSellInfo(c, info.SellInfoId) {
			rsp.SellInfoId = 0
			rsp.GoodName = ""
			rsp.ValidTime = 0
			rsp.ContentId = ""
			rsp.Description = ""
		}
		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("sellInfo", setupRouter())
}
