package main

import (
	"context"
	db "jiaojiao/database"
	mockbuy "jiaojiao/srv/buyinfo/mock"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	mocksell "jiaojiao/srv/sellinfo/mock"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	transaction "jiaojiao/srv/transaction/proto"
	"jiaojiao/utils"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/client"
)

type srv struct{}

/**
 * @api {rpc} /rpc Transaction.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Transaction.Create
 * @apiDescription Create transaction
 *
 * @apiParam {int32} infoID sellInfoID or buyInfoID.
 * @apiParam {int32} category 1 for sell <br> 2 for buy
 * @apiParam {int32} fromUserID userID whose create the transaction
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for info id not found
 * @apiSuccess {int32} transactionID transaction id
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *transaction.TransactionCreateRequest, rsp *transaction.TransactionCreateResponse) error {
	if !utils.RequireParam(req.InfoID, req.FromUserID, req.Category) {
		rsp.Status = transaction.TransactionCreateResponse_INVALID_PARAM
		return nil
	}

	var toUserID int32
	if req.Category == transaction.Category_SELL {
		microSrv := utils.CallMicroService("sellInfo", func(name string, c client.Client) interface{} { return sellinfo.NewSellInfoService(name, c) },
			func() interface{} { return mocksell.NewSellInfoService() }).(sellinfo.SellInfoService)
		srvRsp, err := microSrv.Query(context.TODO(), &sellinfo.SellInfoQueryRequest{
			SellInfoID: req.InfoID,
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		if srvRsp.UserID == 0 {
			rsp.Status = transaction.TransactionCreateResponse_NOT_FOUND
			return nil
		}
		toUserID = srvRsp.UserID
	} else {
		microSrv := utils.CallMicroService("buyInfo", func(name string, c client.Client) interface{} { return buyinfo.NewBuyInfoService(name, c) },
			func() interface{} { return mockbuy.NewBuyInfoService() }).(buyinfo.BuyInfoService)
		srvRsp, err := microSrv.Query(context.TODO(), &buyinfo.BuyInfoQueryRequest{
			BuyInfoID: req.InfoID,
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		if srvRsp.UserID == 0 {
			rsp.Status = transaction.TransactionCreateResponse_NOT_FOUND
			return nil
		}
		toUserID = srvRsp.UserID
	}

	tran := db.Transaction{
		InfoID:     req.InfoID,
		Category:   int32(req.Category),
		FromUserID: req.FromUserID,
		ToUserID:   toUserID,
		CreateTime: time.Now(),
		Status:     int32(transaction.TransStatus_ASKING),
	}
	err := db.Ormer.Create(&tran).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Error) {
		return err
	}

	rsp.TransactionID = tran.ID
	rsp.Status = transaction.TransactionCreateResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc Transaction.Update
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Transaction.Update
 * @apiDescription Update transaction status
 *
 * @apiParam {int32} transactionID transaction ID
 * @apiParam {int32} status 1 for asking <br> 2 for accepted <br> 3 for rejected <br> 4 for closed <br> 5 for pending <br> 6 for done
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success
 * @apiUse DBServerDown
 */
func (a *srv) Update(ctx context.Context, req *transaction.TransactionUpdateRequest, rsp *transaction.TransactionUpdateResponse) error {
	if !utils.RequireParam(req.TransactionID, req.Status) {
		rsp.Status = transaction.TransactionUpdateResponse_INVALID_PARAM
		return nil
	}

	tran := db.Transaction{
		ID: req.TransactionID,
	}
	err := db.Ormer.First(&tran).Error
	if gorm.IsRecordNotFoundError(err) {
		rsp.Status = transaction.TransactionUpdateResponse_NOT_FOUND
		return nil
	} else if utils.LogContinue(err, utils.Error) {
		return err
	}
	tran.Status = int32(req.Status)
	err = db.Ormer.Save(&tran).Error
	if utils.LogContinue(err, utils.Error) {
		return err
	}

	rsp.Status = transaction.TransactionUpdateResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc Transaction.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Transaction.Find
 * @apiDescription Find transactions
 *
 * @apiParam {int32} [infoID] sellInfoID or buyInfoID.
 * @apiParam {int32} [category] 1 for sell <br> 2 for buy
 * @apiParam {int32} [userID] userID who create the transaction
 * @apiParam {int64} [lowCreateTime] low boundary of CreateTime
 * @apiParam {int64} [highCreateTime] high boundary of CreateTime
 * @apiParam {int32} [status] 1 for asking <br> 2 for accepted <br> 3 for rejected <br> 4 for closed <br> 5 for pending <br> 6 for done
 * @apiParam {uint32{0-100}} limit=100 row limit
 * @apiParam {uint32} offset=0 row offset
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiSuccess {array} transactions see below
 * @apiSuccess (Transaction Array) {int32} transactionID transaction ID
 * @apiSuccess (Transaction Array) {int32} infoID sellInfoID or buyInfoID
 * @apiSuccess (Transaction Array) {int32} category 1 for sell <br> 2 for buy
 * @apiSuccess (Transaction Array) {int32} userID userID whose create the transaction
 * @apiSuccess (Transaction Array) {int64} createTime create time
 * @apiSuccess (Transaction Array) {int32} status 1 for order <br> 2 for done
 * @apiUse DBServerDown
 */
func (a *srv) Find(ctx context.Context, req *transaction.TransactionFindRequest, rsp *transaction.TransactionFindResponse) error {
	if req.Limit == 0 {
		req.Limit = 100
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	var res []*db.Transaction
	tx := db.Ormer.Table("transactions")
	if req.InfoID != 0 {
		tx = tx.Where("info_id = ?", req.InfoID)
	}
	if req.Category != transaction.Category_CATEGORY_UNKNOWN {
		tx = tx.Where("category = ?", int32(req.Category))
	}
	if req.UserID != 0 {
		tx = tx.Where("from_user_id = ?", req.UserID)
		tx = tx.Or("to_user_id = ?", req.UserID)
	}
	if req.LowCreateTime != 0 {
		tx = tx.Where("create_time >= ?", time.Unix(req.LowCreateTime, 0))
	}
	if req.HighCreateTime != 0 {
		tx = tx.Where("create_time <= ?", time.Unix(req.HighCreateTime, 0))
	}
	if req.Status != transaction.TransStatus_UNKNOWN {
		tx = tx.Where("status = ?", int32(req.Status))
	}
	ans := tx.Limit(req.Limit).Offset(req.Offset).Find(&res)
	if utils.LogContinue(ans.Error, utils.Error) {
		return ans.Error
	}
	if ans.RowsAffected == 0 {
		rsp.Status = transaction.TransactionFindResponse_NOT_FOUND
		return nil
	}

	for _, v := range res {
		rsp.Transactions = append(rsp.Transactions, &transaction.TransactionMsg{
			TransactionID: v.ID,
			InfoID:        v.InfoID,
			Category:      transaction.Category(utils.EnumConvert(v.Category, transaction.Category_name)),
			FromUserID:    v.FromUserID,
			ToUserID:      v.ToUserID,
			CreateTime:    v.CreateTime.Unix(),
			Status:        transaction.TransStatus(utils.EnumConvert(v.Status, transaction.TransStatus_name)),
		})
	}
	rsp.Status = transaction.TransactionFindResponse_SUCCESS
	return nil
}

func main() {
	db.InitORM("transactiondb", new(db.Transaction))
	defer db.CloseORM()
	service := utils.InitMicroService("transaction")
	utils.LogPanic(transaction.RegisterTransactionHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
