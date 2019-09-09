package main

import (
	"context"
	mockbuy "jiaojiao/srv/buyinfo/mock"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	mockmsg "jiaojiao/srv/message/mock"
	message "jiaojiao/srv/message/proto"
	"jiaojiao/srv/sellinfo/mock"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"
	"strconv"

	"github.com/micro/go-micro/client"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/sellInfo/:sellInfoID", getSellInfo)
	rg.PUT("/sellInfo/:sellInfoID", updateSellInfo)
	rg.GET("/sellInfo", findSellInfo)
	rg.POST("/sellInfo", addSellInfo)
	return router
}

/**
 * @apiDefine SellInfoServiceDown
 * @apiError (Error 500) SellInfoServiceDown SellInfo service down
 */

/**
 * @api {get} /sellInfo/:sellInfoID GetSellInfo
 * @apiVersion 1.0.0
 * @apiGroup SellInfo
 * @apiPermission none
 * @apiName GetSellInfo
 * @apiDescription Get sell info
 *
 * @apiParam {--} Param see [SellInfo Service](#api-Service-SellInfo_Query)
 * @apiSuccess (Success 200) {Response} response see [SellInfo Service](#api-Service-SellInfo_Query)
 * @apiUse InvalidParam
 * @apiUse SellInfoServiceDown
 */
func getSellInfo(c *gin.Context) {
	type param struct {
		SellInfoID int32 `uri:"sellInfoID" binding:"required,min=1"`
	}
	var p param
	if !utils.LogContinue(c.ShouldBindUri(&p), utils.Warning) {
		srv := utils.CallMicroService("sellInfo", func(name string, c client.Client) interface{} { return sellinfo.NewSellInfoService(name, c) },
			func() interface{} { return mock.NewSellInfoService() }).(sellinfo.SellInfoService)
		rsp, err := srv.Query(context.TODO(), &sellinfo.SellInfoQueryRequest{
			SellInfoID: p.SellInfoID,
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
 * @api {put} /sellInfo/:sellInfoID UpdateSellInfo
 * @apiVersion 1.0.0
 * @apiGroup SellInfo
 * @apiPermission self/admin
 * @apiName UpdateSellInfo
 * @apiDescription Update sell info
 *
 * @apiParam {int32} status 1 for selling <br> 2 for reserved <br> 3 for done <br> 4 for expired <br> 5 for closed
 * @apiSuccess (Success 200) {Response} response see [SellInfo Service](#api-Service-SellInfo_Update)
 * @apiUse InvalidParam
 * @apiUse SellInfoServiceDown
 */
func updateSellInfo(c *gin.Context) {
	type paramUri struct {
		SellInfoID int32 `uri:"sellInfoID" binding:"required,min=1"`
	}
	var pu paramUri

	type param struct {
		Status int32 `form:"status" binding:"required"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindUri(&pu), utils.Warning) && !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		srv := utils.CallMicroService("sellInfo", func(name string, c client.Client) interface{} { return sellinfo.NewSellInfoService(name, c) },
			func() interface{} { return mock.NewSellInfoService() }).(sellinfo.SellInfoService)
		rsp2, err := srv.Query(context.TODO(), &sellinfo.SellInfoQueryRequest{
			SellInfoID: pu.SellInfoID,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}
		if rsp2.SellInfoID == 0 {
			c.JSON(200, &sellinfo.SellInfoUpdateResponse{
				Status: sellinfo.SellInfoUpdateResponse_NOT_FOUND,
			})
			return
		}

		role := utils.GetRoleID(c, rsp2.UserID)

		if !role.Self && !role.Admin {
			c.AbortWithStatus(403)
			return
		}

		rsp, err := srv.Update(context.TODO(), &sellinfo.SellInfoUpdateRequest{
			SellInfoID: pu.SellInfoID,
			Status:     sellinfo.SellStatus(utils.EnumConvert(p.Status, sellinfo.SellStatus_name)),
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
 * @api {post} /sellInfo AddSellInfo
 * @apiVersion 1.0.0
 * @apiGroup SellInfo
 * @apiPermission self/admin
 * @apiName AddSellInfo
 * @apiDescription Add sell info
 *
 * @apiParam {--} Param see [SellInfo Service](#api-Service-SellInfo_Create)
 * @apiSuccess (Success 200) {Response} response see [SellInfo Service](#api-Service-SellInfo_Create)
 * @apiUse InvalidParam
 * @apiUse SellInfoServiceDown
 */
func addSellInfo(c *gin.Context) {
	type param struct {
		ValidTime    int64    `form:"validTime" binding:"required"`
		GoodName     string   `form:"goodName" binding:"required"`
		Price        float64  `form:"price"`
		Description  string   `form:"description"`
		ContentID    string   `form:"contentID"`
		ContentToken string   `form:"contentToken"`
		UserID       int32    `form:"userID" binding:"required"`
		Tags         []string `form:"tags"`
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
		srv := utils.CallMicroService("sellInfo", func(name string, c client.Client) interface{} { return sellinfo.NewSellInfoService(name, c) },
			func() interface{} { return mock.NewSellInfoService() }).(sellinfo.SellInfoService)
		rsp, err := srv.Create(context.TODO(), &sellinfo.SellInfoCreateRequest{
			ValidTime:    p.ValidTime,
			GoodName:     p.GoodName,
			Price:        p.Price,
			Description:  p.Description,
			ContentID:    p.ContentID,
			ContentToken: p.ContentToken,
			UserID:       p.UserID,
			Tags:         p.Tags,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}

		// Just for check
		if len(p.Tags) >= 1 && len(p.Tags[0]) > 6 {
			srv2 := utils.CallMicroService("buyInfo", func(name string, c client.Client) interface{} { return buyinfo.NewBuyInfoService(name, c) },
				func() interface{} { return mockbuy.NewBuyInfoService() }).(buyinfo.BuyInfoService)
			rsp2, err := srv2.Find(context.TODO(), &buyinfo.BuyInfoFindRequest{
				GoodName: p.Tags[0][2:5],
				Limit:    1,
			})
			if utils.LogContinue(err, utils.Error) {
				c.JSON(500, err)
				return
			}
			if rsp2.BuyInfo != nil {
				srv3 := utils.CallMicroService("message", func(name string, c client.Client) interface{} { return message.NewMessageService(name, c) },
					func() interface{} { return mockmsg.NewMessageService() }).(message.MessageService)
				_, err := srv3.Create(context.TODO(), &message.MessageCreateRequest{
					FromUser: 1,
					ToUser:   p.UserID,
					Type:     1,
					Msg:      []byte(strconv.Itoa(int(rsp2.BuyInfo[0].BuyInfoID))),
				})
				if utils.LogContinue(err, utils.Error) {
					c.JSON(500, err)
					return
				}
			}
		}

		c.JSON(200, rsp)
	} else {
		c.AbortWithStatus(400)
	}
}

/**
 * @api {get} /sellInfo FindSellInfo
 * @apiVersion 1.0.0
 * @apiGroup SellInfo
 * @apiPermission none
 * @apiName FindSellInfo
 * @apiDescription Find sell info
 *
 * @apiParam {--} Param see [SellInfo Service](#api-Service-SellInfo_Find)
 * @apiSuccess {Response} response see [SellInfo Service](#api-Service-SellInfo_Find)
 * @apiUse InvalidParam
 * @apiUse SellInfoServiceDown
 */
func findSellInfo(c *gin.Context) {
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
		srv := utils.CallMicroService("sellInfo", func(name string, c client.Client) interface{} { return sellinfo.NewSellInfoService(name, c) },
			func() interface{} { return mock.NewSellInfoService() }).(sellinfo.SellInfoService)
		rsp, err := srv.Find(context.TODO(), &sellinfo.SellInfoFindRequest{
			UserID:    p.UserID,
			Status:    sellinfo.SellStatus(utils.EnumConvert(p.Status, sellinfo.SellStatus_name)),
			GoodName:  p.GoodName,
			LowPrice:  p.LowPrice,
			HighPrice: p.HighPrice,
			Limit:     p.Limit,
			Offset:    p.Offset,
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
	utils.RunWebService("sellInfo", setupRouter())
}
