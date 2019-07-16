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
		Id:          1000,
		Status:      1,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1000,
		GoodId:      1000,
	}
	good := db.Good{
		Id:          1000,
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

		_, err := db.Ormer.Insert(&good)
		So(err, ShouldBeNil)
		_, err = db.Ormer.Insert(&info)
		So(err, ShouldBeNil)

		req.SellInfoId = 1000
		tf(1000, "good", "Very good!", "123456789", 1000)

		req.SellInfoId = 1001
		tf(0, "", "", "", 0)

		_, err = db.Ormer.Delete(&db.SellInfo{
			Id: 1000,
		})
		So(err, ShouldBeNil)

		_, err = db.Ormer.Delete(&db.Good{
			Id: 1000,
		})
		So(err, ShouldBeNil)
	})

}

func TestSrvInfoCreate(t *testing.T) {
	var s srvInfo
	var req sellinfo.SellInfoCreateRequest
	var rsp sellinfo.SellInfoCreateResponse

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

	Convey("Test SellInfo Create", t, func() {
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.SellInfoCreateResponse_INVALID_PARAM)

		req.GoodName = "good"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.SellInfoCreateResponse_INVALID_PARAM)

		req.ValidTime = 19087982694
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.SellInfoCreateResponse_INVALID_PARAM)

		req.UserId = 1000
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.SellInfoCreateResponse_SUCCESS)
		So(rsp.SellInfoId, ShouldNotEqual, 0)

		tmp := db.SellInfo{
			Id: rsp.SellInfoId,
		}
		err := db.Ormer.Read(&tmp)
		So(err, ShouldBeNil)
		_, err = db.Ormer.Delete(&db.Good{
			Id: tmp.GoodId,
		})
		So(err, ShouldBeNil)
		_, err = db.Ormer.Delete(&tmp)
		So(err, ShouldBeNil)

		req.ContentId = "123456789abc123456789abc"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.SellInfoCreateResponse_INVALID_PARAM)

		req.ContentToken = "jlkfjaoiu2709429-98247ksf"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.SellInfoCreateResponse_INVALID_TOKEN)

		req.ContentId = "1234"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.SellInfoCreateResponse_INVALID_PARAM)

		req.ContentId, req.ContentToken = getToken()
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.SellInfoCreateResponse_SUCCESS)
		So(rsp.SellInfoId, ShouldNotEqual, 0)

		tmp = db.SellInfo{
			Id: rsp.SellInfoId,
		}
		err = db.Ormer.Read(&tmp)
		So(err, ShouldBeNil)
		_, err = db.Ormer.Delete(&db.Good{
			Id: tmp.GoodId,
		})
		So(err, ShouldBeNil)
		_, err = db.Ormer.Delete(&tmp)
		So(err, ShouldBeNil)
		var sc srvContent
		var rspc sellinfo.ContentDeleteResponse
		err = sc.Delete(context.TODO(), &sellinfo.ContentDeleteRequest{
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
		Id:          1000,
		Status:      1,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1000,
		GoodId:      1000,
	}
	info2 := db.SellInfo{
		Id:          1001,
		Status:      2,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1000,
		GoodId:      1001,
	}
	info3 := db.SellInfo{
		Id:          1002,
		Status:      3,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserId:      1001,
		GoodId:      1002,
	}
	good1 := db.Good{
		Id:          1000,
		GoodName:    "good",
		Description: "Very good!",
		ContentId:   "123456789",
	}
	good2 := db.Good{
		Id:          1001,
		GoodName:    "good",
		Description: "Very good!",
		ContentId:   "123456789",
	}
	good3 := db.Good{
		Id:          1002,
		GoodName:    "good",
		Description: "Very good!",
		ContentId:   "123456789",
	}

	prepare := func() {
		_, err := db.Ormer.Insert(&good1)
		So(err, ShouldBeNil)
		_, err = db.Ormer.Insert(&good2)
		So(err, ShouldBeNil)
		_, err = db.Ormer.Insert(&good3)
		So(err, ShouldBeNil)
		_, err = db.Ormer.Insert(&info1)
		So(err, ShouldBeNil)
		_, err = db.Ormer.Insert(&info2)
		So(err, ShouldBeNil)
		_, err = db.Ormer.Insert(&info3)
		So(err, ShouldBeNil)
	}
	end := func() {
		_, err := db.Ormer.Delete(&db.SellInfo{Id: 1000})
		So(err, ShouldBeNil)
		_, err = db.Ormer.Delete(&db.SellInfo{Id: 1001})
		So(err, ShouldBeNil)
		_, err = db.Ormer.Delete(&db.SellInfo{Id: 1002})
		So(err, ShouldBeNil)

		_, err = db.Ormer.Delete(&db.Good{Id: 1000})
		So(err, ShouldBeNil)
		_, err = db.Ormer.Delete(&db.Good{Id: 1001})
		So(err, ShouldBeNil)
		_, err = db.Ormer.Delete(&db.Good{Id: 1002})
		So(err, ShouldBeNil)
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
	var s srvContent
	var req sellinfo.ContentCreateRequest
	var rsp sellinfo.ContentCreateResponse
	Convey("Test SellInfo Content Create", t, func() {
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_PARAM)

		req.Content = []byte{1, 2, 3, 4, 5, 6}
		req.Type = sellinfo.ContentCreateRequest_PICTURE
		req.ContentId = "1234"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_PARAM)

		req.ContentId = ""
		req.ContentToken = "12463-25897fsfs-5232"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_PARAM)

		req.ContentId = "1234"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_TOKEN)

		req.ContentId = ""
		req.ContentToken = ""
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_SUCCESS)
		So(rsp.ContentId, ShouldNotBeBlank)
		So(rsp.ContentToken, ShouldNotBeBlank)

		req.ContentId = rsp.ContentId
		req.ContentToken = rsp.ContentToken
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_SUCCESS)
		So(rsp.ContentId, ShouldNotBeBlank)
		So(rsp.ContentToken, ShouldNotBeBlank)

		req.ContentToken = "12463-25897fsfs-5232"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_TOKEN)

		var sc srvContent
		var rspc sellinfo.ContentDeleteResponse
		err := sc.Delete(context.TODO(), &sellinfo.ContentDeleteRequest{
			ContentId:    rsp.ContentId,
			ContentToken: rsp.ContentToken,
		}, &rspc)
		So(err, ShouldBeNil)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
