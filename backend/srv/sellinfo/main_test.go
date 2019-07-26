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
	var s srv
	var req sellinfo.SellInfoQueryRequest

	info := db.SellInfo{
		ID:          1100,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserID:      1000,
		GoodID:      1101,
	}
	good := db.Good{
		ID:          1101,
		GoodName:    "good",
		Description: "Very good!",
		ContentID:   "123456789",
	}
	tf := func(sellID int, contentID string, userID int) {
		var rsp sellinfo.SellInfoMsg
		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.SellInfoID, ShouldEqual, sellID)
		So(rsp.ContentID, ShouldEqual, contentID)
		So(rsp.UserID, ShouldEqual, userID)
	}
	Convey("Test SellInfo Query", t, func() {
		tf(0, "", 0)

		So(db.Ormer.Create(&good).Error, ShouldBeNil)
		So(db.Ormer.Create(&info).Error, ShouldBeNil)
		defer func() {
			So(db.Ormer.Delete(&db.Good{ID: 1101}).Error, ShouldBeNil)
			So(db.Ormer.Delete(&db.SellInfo{ID: 1100}).Error, ShouldBeNil)
		}()

		req.SellInfoID = 1100
		tf(1100, "123456789", 1000)

		req.SellInfoID = 1101
		tf(0, "", 0)
	})
}

func TestSrvInfoCreate(t *testing.T) {
	// TODO
	//var s srv
	//var req sellinfo.SellInfoCreateRequest
	//srv := utils.CallMicroService("content", func(name string, c client.Client) interface{} { return content.NewContentService(name, c) },
	//	func() interface{} { return mock.NewContentService() }).(content.ContentService)
	//
	//getToken := func() (string, string) {
	//	rsp, err := srv.Create(context.TODO(), &content.ContentCreateRequest{
	//		Content: []byte{1, 2, 3, 4, 5, 6},
	//		Type:    content.ContentCreateRequest_PICTURE,
	//	})
	//	So(err, ShouldBeNil)
	//	So(rsp.Status, ShouldEqual, content.ContentCreateResponse_SUCCESS)
	//	So(rsp.ContentID, ShouldNotBeBlank)
	//	So(rsp.ContentToken, ShouldNotBeBlank)
	//
	//	return rsp.ContentID, rsp.ContentToken
	//}
	//
	//tf := func(status sellinfo.SellInfoCreateResponse_Status, success bool) int32 {
	//	var rsp sellinfo.SellInfoCreateResponse
	//	So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
	//	So(rsp.Status, ShouldEqual, status)
	//	if success {
	//		So(rsp.SellInfoID, ShouldNotEqual, 0)
	//	} else {
	//		So(rsp.SellInfoID, ShouldEqual, 0)
	//	}
	//	return rsp.SellInfoID
	//}
	//
	//Convey("Test SellInfo Create", t, func() {
	//	tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)
	//
	//	req.GoodName = "good"
	//	tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)
	//
	//	req.ValidTime = 1893427200
	//	tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)
	//
	//	req.UserID = 1000
	//	id := tf(sellinfo.SellInfoCreateResponse_SUCCESS, true)
	//
	//	tmp := db.SellInfo{
	//		ID: id,
	//	}
	//	So(db.Ormer.First(&tmp).Error, ShouldBeNil)
	//	So(db.Ormer.Delete(&db.Good{ID: tmp.GoodID}).Error, ShouldBeNil)
	//	So(db.Ormer.Delete(&tmp).Error, ShouldBeNil)
	//
	//	req.ContentID = "123456789abc123456789abc"
	//	tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)
	//
	//	req.ContentToken = "jlkfjaoiu2709429-98247ksf"
	//	tf(sellinfo.SellInfoCreateResponse_INVALID_TOKEN, false)
	//
	//	req.ContentID = "1234"
	//	tf(sellinfo.SellInfoCreateResponse_INVALID_PARAM, false)
	//
	//	req.ContentID, req.ContentToken = getToken()
	//	id = tf(sellinfo.SellInfoCreateResponse_SUCCESS, true)
	//
	//	tmp = db.SellInfo{
	//		ID: id,
	//	}
	//	So(db.Ormer.First(&tmp).Error, ShouldBeNil)
	//	So(db.Ormer.Delete(&db.Good{ID: tmp.GoodID}).Error, ShouldBeNil)
	//	So(db.Ormer.Delete(&tmp).Error, ShouldBeNil)
	//
	//	_, err := srv.Delete(context.TODO(), &content.ContentDeleteRequest{
	//		ContentID:    req.ContentID,
	//		ContentToken: req.ContentToken,
	//	})
	//	So(err, ShouldBeNil)
	//})
}

func TestSrvInfoFind(t *testing.T) {
	var s srv
	var req sellinfo.SellInfoFindRequest

	info1 := db.SellInfo{
		ID:          1000,
		Status:      1,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserID:      1000,
		GoodID:      1010,
	}
	info2 := db.SellInfo{
		ID:          1001,
		Status:      2,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserID:      1000,
		GoodID:      1011,
	}
	info3 := db.SellInfo{
		ID:          1002,
		Status:      3,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		UserID:      1001,
		GoodID:      1012,
	}
	good1 := db.Good{
		ID:          1010,
		GoodName:    "good",
		Description: "Very good!",
		ContentID:   "123456789",
	}
	good2 := db.Good{
		ID:          1011,
		GoodName:    "good",
		Description: "Very good!",
		ContentID:   "123456789",
	}
	good3 := db.Good{
		ID:          1012,
		GoodName:    "good",
		Description: "Very good!",
		ContentID:   "123456789",
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

		So(db.Ormer.Delete(&db.Good{ID: 1010}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 1011}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.Good{ID: 1012}).Error, ShouldBeNil)
	}
	testLen := func(length int) {
		var rsp sellinfo.SellInfoFindResponse

		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.SellInfo), ShouldEqual, length)
	}

	Convey("Test SellInfo Find", t, func() {
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
