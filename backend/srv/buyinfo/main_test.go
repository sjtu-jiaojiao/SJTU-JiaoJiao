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
		ID:          1100,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserID:      1000,
		GoodID:      2110,
	}
	good := db.Good{
		ID:          2110,
		GoodName:    "good",
		Description: "Very good!",
		ContentID:   "012345678901234567890123",
	}
	tf := func(sellID int, contentID string, userID int) {
		var rsp buyinfo.BuyInfoMsg
		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.BuyInfoID, ShouldEqual, sellID)
		So(rsp.ContentID, ShouldEqual, contentID)
		So(rsp.UserID, ShouldEqual, userID)
	}
	Convey("Test BuyInfo Query", t, func() {
		tf(0, "", 0)

		So(db.Ormer.Create(&good).Error, ShouldBeNil)
		So(db.Ormer.Create(&info).Error, ShouldBeNil)
		defer func() {
			So(db.Ormer.Delete(&db.Good{ID: 2110}).Error, ShouldBeNil)
			So(db.Ormer.Delete(&db.BuyInfo{ID: 1100}).Error, ShouldBeNil)
		}()

		req.BuyInfoID = 1100
		tf(1100, "012345678901234567890123", 1000)

		req.BuyInfoID = 1101
		tf(0, "", 0)
	})
}

func TestSrvInfoCreate(t *testing.T) {
	// TODO
}

func TestSrvInfoFind(t *testing.T) {
	var s srv
	var req buyinfo.BuyInfoFindRequest

	info1 := db.BuyInfo{
		ID:          1000,
		Status:      1,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserID:      1000,
		GoodID:      2010,
	}
	info2 := db.BuyInfo{
		ID:          1001,
		Status:      2,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserID:      1000,
		GoodID:      2011,
	}
	info3 := db.BuyInfo{
		ID:          1002,
		Status:      3,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserID:      1001,
		GoodID:      2012,
	}
	good1 := db.Good{
		ID:          2010,
		GoodName:    "good",
		Description: "Very good!",
		ContentID:   "012345678901234567890123",
	}
	good2 := db.Good{
		ID:          2011,
		GoodName:    "good",
		Description: "Very good!",
		ContentID:   "012345678901234567890123",
	}
	good3 := db.Good{
		ID:          2012,
		GoodName:    "good",
		Description: "Very good!",
		ContentID:   "012345678901234567890123",
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

		So(db.Ormer.Delete(&db.Good{ID: 2010}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 2011}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 2012}).Error, ShouldBeNil)
	}
	testLen := func(length int) {
		var rsp buyinfo.BuyInfoFindResponse

		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.BuyInfo), ShouldEqual, length)
	}

	Convey("Test BuyInfo Find", t, func() {
		testLen(0)

		prepare()
		defer end()

		testLen(3)

		req.UserID = 1001
		testLen(1)

		req.UserID = 1000
		testLen(2)

		req.Limit = 1
		testLen(1)

		req.Offset = 1
		testLen(1)

		req.Offset = 2
		testLen(0)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
