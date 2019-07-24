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

func (a *srv) Find(context.Context, *transaction.TransactionFindRequest, *transaction.TransactionFindResponse) error {
	panic("implement me")
}

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
	if req.InfoID == 0 || req.UserID == 0 || req.Category == transaction.TransactionCreateRequest_UNKNOWN {
		rsp.Status = transaction.TransactionCreateResponse_INVALID_PARAM
		return nil
	}
	tran := db.Transaction{
		InfoID:     req.InfoID,
		Category:   int32(req.Category),
		UserID:     req.UserID,
		CreateTime: time.Now(),
		Status:     int32(transaction.TransactionMsg_ORDER),
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
 * @apiParam {int32} status 1 for order <br> 2 for done
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success
 * @apiUse DBServerDown
 */
func (a *srv) Update(ctx context.Context, req *transaction.TransactionUpdateRequest, rsp *transaction.TransactionUpdateResponse) error {
	if req.TransactionID == 0 || req.Status == transaction.TransactionUpdateRequest_STATUS_UNKNOWN {
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

func main() {
	db.InitORM("transactiondb", new(db.Transaction))
	defer db.CloseORM()
	service := utils.InitMicroService("transaction")
	utils.LogPanic(transaction.RegisterTransactionHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
