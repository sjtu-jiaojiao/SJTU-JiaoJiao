package main

import (
	"context"
	db "jiaojiao/database"
	transaction "jiaojiao/srv/transaction/proto"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	tf := func(infoID int32, category int32, fromUserID int32, status transaction.TransactionCreateResponse_Status) {
		var s srv
		var rsp transaction.TransactionCreateResponse
		So(s.Create(context.TODO(), &transaction.TransactionCreateRequest{
			InfoID:     infoID,
			Category:   transaction.Category(category),
			FromUserID: fromUserID,
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if status == transaction.TransactionCreateResponse_SUCCESS {
			So(db.Ormer.Delete(&db.Transaction{ID: rsp.TransactionID}).RowsAffected, ShouldEqual, 1)
		}
	}

	Convey("Test Transaction Create", t, func() {
		tf(0, 1, 0, transaction.TransactionCreateResponse_INVALID_PARAM)
		tf(3000, 1, 1001, transaction.TransactionCreateResponse_NOT_FOUND)
		tf(3000, 2, 1001, transaction.TransactionCreateResponse_NOT_FOUND)
		tf(1000, 1, 1001, transaction.TransactionCreateResponse_SUCCESS)
		tf(1000, 2, 1001, transaction.TransactionCreateResponse_SUCCESS)
	})
}

func TestUpdate(t *testing.T) {
	tf := func(tranID int32, status transaction.TransStatus, rspStatus transaction.TransactionUpdateResponse_Status) {
		var s srv
		var rsp transaction.TransactionUpdateResponse
		So(s.Update(context.TODO(), &transaction.TransactionUpdateRequest{
			TransactionID: tranID,
			Status:        status,
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, rspStatus)
		if rspStatus == transaction.TransactionUpdateResponse_SUCCESS {
			tran := db.Transaction{ID: tranID}
			db.Ormer.First(&tran)
			So(tran.Status, ShouldEqual, status)
		}
	}

	Convey("Test Transaction Update", t, func() {
		So(db.Ormer.Create(&db.Transaction{
			ID:         100,
			InfoID:     1000,
			Category:   1,
			FromUserID: 1001,
			ToUserID:   2001,
			CreateTime: time.Now(),
			Status:     int32(transaction.TransStatus_ACCEPTED),
		}).Error, ShouldBeNil)
		defer func() {
			So(db.Ormer.Delete(&db.Transaction{ID: 100}).Error, ShouldBeNil)
		}()

		tf(100, transaction.TransStatus_UNKNOWN, transaction.TransactionUpdateResponse_INVALID_PARAM)
		tf(200, transaction.TransStatus_CLOSED, transaction.TransactionUpdateResponse_NOT_FOUND)
		tf(100, transaction.TransStatus_CLOSED, transaction.TransactionUpdateResponse_SUCCESS)
	})
}

func TestFind(t *testing.T) {
	tf := func(infoID int32, category transaction.Category, userID int32, lowCreateTime int64, highCreateTime int64, status transaction.TransStatus, limit uint32, offset uint32,
		arrayLen int32, rspStatus transaction.TransactionFindResponse_Status) {
		var s srv
		var rsp transaction.TransactionFindResponse
		So(s.Find(context.TODO(), &transaction.TransactionFindRequest{
			InfoID:         infoID,
			Category:       category,
			UserID:         userID,
			LowCreateTime:  lowCreateTime,
			HighCreateTime: highCreateTime,
			Status:         status,
			Limit:          limit,
			Offset:         offset,
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, rspStatus)
		So(len(rsp.Transactions), ShouldEqual, arrayLen)
	}

	Convey("Test Transaction Update", t, func() {
		So(db.Ormer.Create(&db.Transaction{
			ID:         100,
			InfoID:     1000,
			Category:   1,
			FromUserID: 1001,
			ToUserID:   2002,
			CreateTime: time.Unix(100000000, 0),
			Status:     int32(transaction.TransStatus_ACCEPTED),
		}).Error, ShouldBeNil)
		So(db.Ormer.Create(&db.Transaction{
			ID:         101,
			InfoID:     1000,
			Category:   2,
			FromUserID: 1001,
			ToUserID:   2001,
			CreateTime: time.Unix(200000000, 0),
			Status:     int32(transaction.TransStatus_ACCEPTED),
		}).Error, ShouldBeNil)
		So(db.Ormer.Create(&db.Transaction{
			ID:         102,
			InfoID:     1002,
			Category:   1,
			FromUserID: 2001,
			ToUserID:   1002,
			CreateTime: time.Unix(300000000, 0),
			Status:     int32(transaction.TransStatus_ASKING),
		}).Error, ShouldBeNil)
		So(db.Ormer.Create(&db.Transaction{
			ID:         103,
			InfoID:     1003,
			Category:   1,
			FromUserID: 1001,
			ToUserID:   2001,
			CreateTime: time.Unix(400000000, 0),
			Status:     int32(transaction.TransStatus_ASKING),
		}).Error, ShouldBeNil)
		So(db.Ormer.Create(&db.Transaction{
			ID:         104,
			InfoID:     1004,
			Category:   2,
			FromUserID: 1002,
			ToUserID:   2002,
			CreateTime: time.Unix(500000000, 0),
			Status:     int32(transaction.TransStatus_ACCEPTED),
		}).Error, ShouldBeNil)
		defer func() {
			So(db.Ormer.Delete(&db.Transaction{ID: 100}).Error, ShouldBeNil)
			So(db.Ormer.Delete(&db.Transaction{ID: 101}).Error, ShouldBeNil)
			So(db.Ormer.Delete(&db.Transaction{ID: 102}).Error, ShouldBeNil)
			So(db.Ormer.Delete(&db.Transaction{ID: 103}).Error, ShouldBeNil)
			So(db.Ormer.Delete(&db.Transaction{ID: 104}).Error, ShouldBeNil)
		}()

		tf(1005, 0, 0, 0, 0, 0, 0, 0, 0, transaction.TransactionFindResponse_NOT_FOUND)
		tf(1000, 0, 0, 0, 0, 0, 0, 0, 2, transaction.TransactionFindResponse_SUCCESS)
		tf(0, 1, 0, 0, 0, 0, 0, 0, 3, transaction.TransactionFindResponse_SUCCESS)
		tf(0, 1, 0, 200000000, 0, 0, 0, 0, 2, transaction.TransactionFindResponse_SUCCESS)
		tf(0, 0, 0, 0, 0, transaction.TransStatus_ASKING, 0, 0, 2, transaction.TransactionFindResponse_SUCCESS)
		tf(0, 0, 1001, 0, 0, 0, 0, 0, 3, transaction.TransactionFindResponse_SUCCESS)
		tf(0, 0, 2001, 0, 0, 0, 0, 0, 3, transaction.TransactionFindResponse_SUCCESS)
		tf(0, 0, 0, 0, 0, 0, 0, 0, 5, transaction.TransactionFindResponse_SUCCESS)
		tf(0, 0, 0, 0, 0, 0, 2, 2, 2, transaction.TransactionFindResponse_SUCCESS)
		tf(0, 0, 0, 0, 0, 0, 3, 3, 2, transaction.TransactionFindResponse_SUCCESS)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
