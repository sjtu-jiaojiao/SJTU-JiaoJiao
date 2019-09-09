package main

import (
	"context"
	"jiaojiao/srv/buyinfo/mock"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	mockmsg "jiaojiao/srv/message/mock"
	message "jiaojiao/srv/message/proto"
	mocksell "jiaojiao/srv/sellinfo/mock"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"
	"strconv"

	"github.com/micro/go-micro/client"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/buyInfo/:buyInfoID", getBuyInfo)
	rg.PUT("/buyInfo/:buyInfoID", updateBuyInfo)
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
 * @api {put} /buyInfo/:buyInfoID UpdateBuyInfo
 * @apiVersion 1.0.0
 * @apiGroup BuyInfo
 * @apiPermission self/admin
 * @apiName UpdateBuyInfo
 * @apiDescription Update buy info
 *
 * @apiParam {int32} status 1 for buying <br> 2 for reserved <br> 3 for done <br> 4 for expired <br> 5 for closed
 * @apiSuccess (Success 200) {Response} response see [BuyInfo Service](#api-Service-BuyInfo_Update)
 * @apiUse InvalidParam
 * @apiUse BuyInfoServiceDown
 */
func updateBuyInfo(c *gin.Context) {
	type paramUri struct {
		BuyInfoID int32 `uri:"buyInfoID" binding:"required,min=1"`
	}
	var pu paramUri

	type param struct {
		Status int32 `form:"status" binding:"required"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindUri(&pu), utils.Warning) && !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		srv := utils.CallMicroService("buyInfo", func(name string, c client.Client) interface{} { return buyinfo.NewBuyInfoService(name, c) },
			func() interface{} { return mock.NewBuyInfoService() }).(buyinfo.BuyInfoService)
		rsp2, err := srv.Query(context.TODO(), &buyinfo.BuyInfoQueryRequest{
			BuyInfoID: pu.BuyInfoID,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}
		if rsp2.BuyInfoID == 0 {
			c.JSON(200, &buyinfo.BuyInfoUpdateResponse{
				Status: buyinfo.BuyInfoUpdateResponse_NOT_FOUND,
			})
			return
		}

		role := utils.GetRoleID(c, rsp2.UserID)

		if !role.Self && !role.Admin {
			c.AbortWithStatus(403)
			return
		}

		rsp, err := srv.Update(context.TODO(), &buyinfo.BuyInfoUpdateRequest{
			BuyInfoID: pu.BuyInfoID,
			Status:    buyinfo.BuyStatus(utils.EnumConvert(p.Status, buyinfo.BuyStatus_name)),
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
			Tags:         p.Tags,
		})
		if utils.LogContinue(err, utils.Error) {
			c.JSON(500, err)
			return
		}

		// Just for check
		if len(p.Tags) >= 1 && len(p.Tags[0]) > 6 {
			srv2 := utils.CallMicroService("sellInfo", func(name string, c client.Client) interface{} { return sellinfo.NewSellInfoService(name, c) },
				func() interface{} { return mocksell.NewSellInfoService() }).(sellinfo.SellInfoService)
			rsp2, err := srv2.Find(context.TODO(), &sellinfo.SellInfoFindRequest{
				GoodName: p.Tags[0][2:5],
				Limit:    1,
			})
			if utils.LogContinue(err, utils.Error) {
				c.JSON(500, err)
				return
			}
			if rsp2.SellInfo != nil {
				srv3 := utils.CallMicroService("message", func(name string, c client.Client) interface{} { return message.NewMessageService(name, c) },
					func() interface{} { return mockmsg.NewMessageService() }).(message.MessageService)
				_, err := srv3.Create(context.TODO(), &message.MessageCreateRequest{
					FromUser: 1,
					ToUser:   p.UserID,
					Type:     1,
					Msg:      []byte(strconv.Itoa(int(rsp2.SellInfo[0].SellInfoID))),
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
			Status:    buyinfo.BuyStatus(utils.EnumConvert(p.Status, buyinfo.BuyStatus_name)),
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
	utils.RunWebService("buyInfo", setupRouter())
}
