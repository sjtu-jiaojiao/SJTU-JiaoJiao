package main

import (
	"context"
	db "jiaojiao/database"
	transaction "jiaojiao/srv/transaction/proto"
	"jiaojiao/utils"
	"time"

	"github.com/jinzhu/gorm"
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
 * @apiParam {int32} userID userID whose create the transaction
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success
 * @apiSuccess {int32} transactionID transaction id
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *transaction.TransactionCreateRequest, rsp *transaction.TransactionCreateResponse) error {
	if !utils.RequreParam(req.InfoID, req.UserID, req.Category) {
		rsp.Status = transaction.TransactionCreateResponse_INVALID_PARAM
		return nil
	}
	tran := db.Transaction{
		InfoID:     req.InfoID,
		Category:   int32(req.Category),
		UserID:     req.UserID,
		CreateTime: time.Now(),
		Status:     int32(transaction.TransStatus_RESERVED),
	}
	err := db.Ormer.Create(&tran).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
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
	if utils.RequreParam(req.TransactionID, req.Status) {
		rsp.Status = transaction.TransactionUpdateResponse_INVALID_PARAM
		return nil
	}

	tran := db.Transaction{
		ID: req.TransactionID,
	}
	err := db.Ormer.First(&tran).Error
	if utils.LogContinue(err, utils.Warning) {
		rsp.Status = transaction.TransactionUpdateResponse_NOT_FOUND
		return nil
	}
	tran.Status = int32(req.Status)
	err = db.Ormer.Update(&tran).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
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
 * @apiParam {uint32} limit=100 row limit
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

	var res []*db.Transaction
	tx := db.Ormer.Table("transactions")
	if req.InfoID != 0 {
		tx.Where("info_id = ?", req.InfoID)
	}
	if req.Category != transaction.TransactionFindRequest_CATEGORY_UNKNOWN {
		tx.Where("category = ?", int32(req.Category))
	}
	if req.UserID != 0 {
		tx.Where("user_id = ?", req.UserID)
	}
	if req.LowCreateTime != 0 {
		tx.Where("create_time > ?", time.Unix(req.LowCreateTime, 0))
	}
	if req.HighCreateTime != 0 {
		tx.Where("create_time < ?", time.Unix(req.HighCreateTime, 0))
	}
	if req.Status != transaction.TransStatus_UNKNOWN {
		tx.Where("status = ?", int32(req.Status))
	}
	err := tx.Limit(req.Limit).Offset(req.Offset).Find(&res).Error
	if utils.LogContinue(err, utils.Warning) {
		rsp.Status = transaction.TransactionFindResponse_NOT_FOUND
		return nil
	}
	for _, v := range res {
		rsp.Transactions = append(rsp.Transactions, &transaction.TransactionMsg{
			TransactionID: v.ID,
			InfoID:        v.InfoID,
			Category:      transaction.TransactionMsg_Category(v.Category),
			UserID:        v.UserID,
			CreateTime:    v.CreateTime.Unix(),
			Status:        transaction.TransStatus(v.Status),
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
