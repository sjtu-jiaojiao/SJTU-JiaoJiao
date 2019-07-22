package main

import (
	"context"
	db "jiaojiao/database"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSrvInfoQuery(t *testing.T) {
	var s srv
	var req buyinfo.BuyInfoQueryRequest

	info := db.BuyInfo{
		ID:          1000,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1000,
		GoodId:      1000,
	}
	good := db.Good{
		ID:          1000,
		GoodName:    "good",
		Description: "Very good!",
		ContentId:   "123456789",
	}
	tf := func(buyId int, goodName string, description string, contentId string, userId int) {
		var rsp buyinfo.BuyInfoMsg
		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.BuyInfoId, ShouldEqual, buyId)
		So(rsp.GoodName, ShouldEqual, goodName)
		So(rsp.Description, ShouldEqual, description)
		So(rsp.ContentId, ShouldEqual, contentId)
		So(rsp.UserId, ShouldEqual, userId)
	}
	Convey("Test BuyInfo Query", t, func() {
		tf(0, "", "", "", 0)

		So(db.Ormer.Create(&good).Error, ShouldBeNil)
		So(db.Ormer.Create(&info).Error, ShouldBeNil)

		req.BuyInfoId = 1000
		tf(1000, "good", "Very good!", "123456789", 1000)

		req.BuyInfoId = 1001
		tf(0, "", "", "", 0)

		So(db.Ormer.Delete(&db.BuyInfo{ID: 1000}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 1000}).Error, ShouldBeNil)
	})

}

func TestSrvInfoFind(t *testing.T) {
	var s srv
	var req buyinfo.BuyInfoFindRequest

	info1 := db.BuyInfo{
		ID:          1000,
		Status:      1,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1000,
		GoodId:      1000,
	}
	info2 := db.BuyInfo{
		ID:          1001,
		Status:      2,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1000,
		GoodId:      1001,
	}
	info3 := db.BuyInfo{
		ID:          1002,
		Status:      3,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1001,
		GoodId:      1002,
	}
	good1 := db.Good{
		ID:          1000,
		GoodName:    "good",
		Description: "Very good!",
		ContentId:   "123456789",
	}
	good2 := db.Good{
		ID:          1001,
		GoodName:    "good",
		Description: "Very good!",
		ContentId:   "123456789",
	}
	good3 := db.Good{
		ID:          1002,
		GoodName:    "good",
		Description: "Very good!",
		ContentId:   "123456789",
	}

	prepare := func() {
		So(db.Ormer.Create(&good1).Error, ShouldBeNil)
		So(db.Ormer.Create(&good2).Error, ShouldBeNil)
		So(db.Ormer.Create(&good3).Error, ShouldBeNil)
		So(db.Ormer.Create(&info1).Error, ShouldBeNil)
		So(db.Ormer.Create(&info2).Error, ShouldBeNil)
		So(db.Ormer.Create(&info3).Error, ShouldBeNil)
	}
	end := func() {
		So(db.Ormer.Delete(&db.BuyInfo{ID: 1000}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.BuyInfo{ID: 1001}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.BuyInfo{ID: 1002}).Error, ShouldBeNil)

		So(db.Ormer.Delete(&db.Good{ID: 1000}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 1001}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 1002}).Error, ShouldBeNil)
	}
	testLen := func(length int) {
		var rsp buyinfo.BuyInfoFindResponse

		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.BuyInfo), ShouldEqual, length)
	}

	Convey("Test BuyInfo Find", t, func() {
		testLen(0)

		prepare()

		testLen(3)

		req.UserId = 1001
		testLen(1)

		req.UserId = 1000
		testLen(2)

		req.Limit = 1
		testLen(1)

		req.Offset = 1
		testLen(1)

		req.Offset = 2
		testLen(0)

		end()
	})
}

func TestSrvContentDelete(t *testing.T) {
	// TODO
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
