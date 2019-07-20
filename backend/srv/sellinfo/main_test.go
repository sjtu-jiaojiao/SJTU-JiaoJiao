package main

import (
	"context"
	db "jiaojiao/database"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSrvInfoQuery(t *testing.T) {
	var s srvInfo
	var req sellinfo.SellInfoQueryRequest

	info := db.SellInfo{
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
	tf := func(sellId int, goodName string, description string, contentId string, userId int) {
		var rsp sellinfo.SellInfoMsg
		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.SellInfoId, ShouldEqual, sellId)
		So(rsp.GoodName, ShouldEqual, goodName)
		So(rsp.Description, ShouldEqual, description)
		So(rsp.ContentId, ShouldEqual, contentId)
		So(rsp.UserId, ShouldEqual, userId)
	}
	Convey("Test SellInfo Query", t, func() {
		tf(0, "", "", "", 0)

		So(db.Ormer.Create(&good).Error, ShouldBeNil)
		So(db.Ormer.Create(&info).Error, ShouldBeNil)

		req.SellInfoId = 1000
		tf(1000, "good", "Very good!", "123456789", 1000)

		req.SellInfoId = 1001
		tf(0, "", "", "", 0)

		So(db.Ormer.Delete(&db.SellInfo{ID: 1000}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 1000}).Error, ShouldBeNil)
	})

}

func TestSrvInfoCreate(t *testing.T) {
	var s srvInfo
	var req sellinfo.SellInfoCreateRequest

	getToken := func() (string, string) {
		var sc srvContent
		var reqc sellinfo.ContentCreateRequest
		var rspc sellinfo.ContentCreateResponse

		reqc.Content = []byte{1, 2, 3, 4, 5, 6}
		reqc.Type = sellinfo.ContentCreateRequest_PICTURE
		So(sc.Create(context.TODO(), &reqc, &rspc), ShouldBeNil)
		So(rspc.Status, ShouldEqual, sellinfo.ContentCreateResponse_SUCCESS)
		So(rspc.ContentId, ShouldNotBeBlank)
		So(rspc.ContentToken, ShouldNotBeBlank)

		return rspc.ContentId, rspc.ContentToken
	}

	tf := func(status sellinfo.SellInfoCreateResponse_Status, success bool) int32 {
		var rsp sellinfo.SellInfoCreateResponse
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if success {
			So(rsp.SellInfoId, ShouldNotEqual, 0)
		} else {
			So(rsp.SellInfoId, ShouldEqual, 0)
		}
		return rsp.SellInfoId
	}

	Convey("Test SellInfo Create", t, func() {
		tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)

		req.GoodName = "good"
		tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)

		req.ValidTime = 1893427200
		tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)

		req.UserId = 1000
		id := tf(sellinfo.SellInfoCreateResponse_SUCCESS, true)

		tmp := db.SellInfo{
			ID: id,
		}
		So(db.Ormer.First(&tmp).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: tmp.GoodId}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&tmp).Error, ShouldBeNil)

		req.ContentId = "123456789abc123456789abc"
		tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)

		req.ContentToken = "jlkfjaoiu2709429-98247ksf"
		tf(sellinfo.SellInfoCreateResponse_INVALID_TOKEN, false)

		req.ContentId = "1234"
		tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)

		req.ContentId, req.ContentToken = getToken()
		id = tf(sellinfo.SellInfoCreateResponse_SUCCESS, true)

		tmp = db.SellInfo{
			ID: id,
		}
		So(db.Ormer.First(&tmp).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: tmp.GoodId}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&tmp).Error, ShouldBeNil)
		var sc srvContent
		var rspc sellinfo.ContentDeleteResponse
		err := sc.Delete(context.TODO(), &sellinfo.ContentDeleteRequest{
			ContentId:    req.ContentId,
			ContentToken: req.ContentToken,
		}, &rspc)
		So(err, ShouldBeNil)
	})
}

func TestSrvInfoFind(t *testing.T) {
	var s srvInfo
	var req sellinfo.SellInfoFindRequest

	info1 := db.SellInfo{
		ID:          1000,
		Status:      1,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1000,
		GoodId:      1000,
	}
	info2 := db.SellInfo{
		ID:          1001,
		Status:      2,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1000,
		GoodId:      1001,
	}
	info3 := db.SellInfo{
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
		So(db.Ormer.Delete(&db.SellInfo{ID: 1000}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.SellInfo{ID: 1001}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.SellInfo{ID: 1002}).Error, ShouldBeNil)

		So(db.Ormer.Delete(&db.Good{ID: 1000}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 1001}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 1002}).Error, ShouldBeNil)
	}
	testLen := func(length int) {
		var rsp sellinfo.SellInfoFindResponse

		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.SellInfo), ShouldEqual, length)
	}

	Convey("Test SellInfo Find", t, func() {
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

func TestSrvContentCreate(t *testing.T) {
	var req sellinfo.ContentCreateRequest

	tf := func(status sellinfo.ContentCreateResponse_Status, success bool) (string, string) {
		var s srvContent
		var rsp sellinfo.ContentCreateResponse
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if success {
			So(rsp.ContentId, ShouldNotBeBlank)
			So(rsp.ContentToken, ShouldNotBeBlank)
		} else {
			So(rsp.ContentId, ShouldBeBlank)
			So(rsp.ContentToken, ShouldBeBlank)
		}
		return rsp.ContentId, rsp.ContentToken
	}

	Convey("Test SellInfo Content Create", t, func() {
		req.Content = []byte{0}
		tf(sellinfo.ContentCreateResponse_INVALID_PARAM, false)

		req.Type = sellinfo.ContentCreateRequest_PICTURE
		tf(sellinfo.ContentCreateResponse_INVALID_PARAM, false)
		req.Type = 0
		req.Content = []byte{1, 2, 3, 4, 5, 6}
		tf(sellinfo.ContentCreateResponse_INVALID_PARAM, false)
		req.Type = sellinfo.ContentCreateRequest_PICTURE
		req.ContentId = "1234"
		tf(sellinfo.ContentCreateResponse_INVALID_PARAM, false)
		req.ContentId = ""
		req.ContentToken = "12463-25897fsfs-5232"
		tf(sellinfo.ContentCreateResponse_INVALID_PARAM, false)

		req.ContentId = "1234"
		tf(sellinfo.ContentCreateResponse_INVALID_TOKEN, false)

		req.ContentId = ""
		req.ContentToken = ""
		id, token := tf(sellinfo.ContentCreateResponse_SUCCESS, true)

		req.ContentId = id
		req.ContentToken = token
		tf(sellinfo.ContentCreateResponse_SUCCESS, true)
		tf(sellinfo.ContentCreateResponse_SUCCESS, true)

		req.ContentToken = "12463-25897fsfs-5232"
		tf(sellinfo.ContentCreateResponse_INVALID_TOKEN, false)

		var sc srvContent
		var rspc sellinfo.ContentDeleteResponse
		err := sc.Delete(context.TODO(), &sellinfo.ContentDeleteRequest{
			ContentId:    id,
			ContentToken: token,
		}, &rspc)
		So(err, ShouldBeNil)
	})
}

func TestSrvContentDelete(t *testing.T) {
	// TODO
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
