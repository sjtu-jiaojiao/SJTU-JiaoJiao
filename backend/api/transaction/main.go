package main

import (
	"context"
	"jiaojiao/srv/transaction/mock"
	transaction "jiaojiao/srv/transaction/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/transaction", findTransaction)
	rg.POST("/transaction", addTransaction)
	rg.PUT("/transaction", updateTransaction)
	return router
}

/**
 * @apiDefine TransactionServiceDown
 * @apiError (Error 500) TransactionServiceDown Transaction service down
 */

/**
 * @api {get} /transaction FindTransaction
 * @apiVersion 1.0.0
 * @apiGroup Transaction
 * @apiPermission self/admin
 * @apiName FindTransaction
 * @apiDescription Find transaction
 *
 * @apiParam {--} Param see [Transaction Service](#api-Service-Transaction_Find) <br> Self permission MUST have userID field
 * @apiSuccess {Response} response see [Transaction Service](#api-Service-Transaction_Find)
 * @apiUse InvalidParam
 * @apiUse TransactionServiceDown
 */
func findTransaction(c *gin.Context) {
	type param struct {
		InfoID         int32  `form:"infoID"`
		Category       int32  `form:"category"`
		UserID         int32  `form:"userID"`
		LowCreateTime  int64  `form:"lowCreateTime"`
		HighCreateTime int64  `form:"highCreateTime"`
		Status         int32  `form:"status"`
		Limit          uint32 `form:"limit"`
		Offset         uint32 `form:"offset"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindQuery(&p), utils.Warning) {
		role := utils.GetRoleID(c, p.UserID)
		if !role.Admin && p.UserID == 0 {
			c.AbortWithStatus(400)
			return
		}
		if !role.Self && !role.Admin {
			c.AbortWithStatus(403)
			return
		}

		srv := utils.CallMicroService("transaction", func(name string, c client.Client) interface{} { return transaction.NewTransactionService(name, c) },
			func() interface{} { return mock.NewTransactionService() }).(transaction.TransactionService)
		rsp, err := srv.Find(context.TODO(), &transaction.TransactionFindRequest{
			InfoID:         p.InfoID,
			Category:       transaction.Category(utils.EnumConvert(p.Category, transaction.Category_name)),
			UserID:         p.UserID,
			LowCreateTime:  p.LowCreateTime,
			HighCreateTime: p.HighCreateTime,
			Status:         transaction.TransStatus(utils.EnumConvert(p.Status, transaction.TransStatus_name)),
			Limit:          p.Limit,
			Offset:         p.Offset,
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
 * @api {post} /transaction AddTransaction
 * @apiVersion 1.0.0
 * @apiGroup Transaction
 * @apiPermission user/admin
 * @apiName AddTransaction
 * @apiDescription Add transaction
 *
 * @apiParam {--} Param see [Transaction Service](#api-Service-Transaction_Create)
 * @apiSuccess {Response} response see [Transaction Service](#api-Service-Transaction_Create)
 * @apiUse InvalidParam
 * @apiUse TransactionServiceDown
 */
func addTransaction(c *gin.Context) {
	type param struct {
		InfoID     int32 `form:"infoID"`
		Category   int32 `form:"category"`
		FromUserID int32 `form:"fromUserID"`
	}
	var p param
	role := utils.GetRole(c)

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if !role.User && !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("transaction", func(name string, c client.Client) interface{} { return transaction.NewTransactionService(name, c) },
			func() interface{} { return mock.NewTransactionService() }).(transaction.TransactionService)
		rsp, err := srv.Create(context.TODO(), &transaction.TransactionCreateRequest{
			InfoID:     p.InfoID,
			Category:   transaction.Category(utils.EnumConvert(p.Category, transaction.Category_name)),
			FromUserID: p.FromUserID,
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
 * @api {put} /transaction UpdateTransaction
 * @apiVersion 1.0.0
 * @apiGroup Transaction
 * @apiPermission admin
 * @apiName UpdateTransaction
 * @apiDescription Update transaction
 *
 * @apiParam {--} Param see [Transaction Service](#api-Service-Transaction_Update)
 * @apiSuccess {Response} response see [Transaction Service](#api-Service-Transaction_Update)
 * @apiUse InvalidParam
 * @apiUse TransactionServiceDown
 */
func updateTransaction(c *gin.Context) {
	type param struct {
		TransactionID int32 `form:"transactionID"`
		Status        int32 `form:"status"`
	}
	var p param
	role := utils.GetRole(c)

	if !utils.LogContinue(c.ShouldBind(&p), utils.Warning) {
		if !role.Admin {
			c.AbortWithStatus(403)
			return
		}
		srv := utils.CallMicroService("transaction", func(name string, c client.Client) interface{} { return transaction.NewTransactionService(name, c) },
			func() interface{} { return mock.NewTransactionService() }).(transaction.TransactionService)
		rsp, err := srv.Update(context.TODO(), &transaction.TransactionUpdateRequest{
			TransactionID: p.TransactionID,
			Status:        transaction.TransStatus(utils.EnumConvert(p.Status, transaction.TransStatus_name)),
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
	utils.RunWebService("transaction", setupRouter())
}
