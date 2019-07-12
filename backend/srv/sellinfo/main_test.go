package main

import (
	"context"
	db "jiaojiao/database"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func TestSrvInfo_Query(t *testing.T) {
	var s srvInfo
	var req sellinfo.SellInfoQueryRequest

	var good = db.Good{
		Id:          1000,
		GoodName:    "good",
		Description: "Very good!",
		ContentId:   "123456789",
	}
	var info = db.SellInfo{
		Id:          1000,
		Status:      1,
		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
		ValidDate:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
		Good:        &good,
	}
	tf := func(sellId int, validTime int64, goodName string,
		description string, contentId string) {
		var rsp sellinfo.SellInfoQueryResponse
		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.SellInfoId, ShouldEqual, sellId)
		So(rsp.ValidTime, ShouldEqual, validTime)
		So(rsp.GoodName, ShouldEqual, goodName)
		So(rsp.Description, ShouldEqual, description)
		So(rsp.ContentId, ShouldEqual, contentId)
	}
	Convey("Test SellInfo Query", t, func() {
		tf(0, 0, "", "", "")

		_, err := o.Insert(&good)
		So(err, ShouldBeNil)
		_, err = o.Insert(&info)
		So(err, ShouldBeNil)

		req.SellInfoId = 1000
		tf(1000, time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local).Unix(), "good",
			"Very good!", "123456789")

		req.SellInfoId = 1001
		tf(0, 0, "", "", "")

		_, err = o.Delete(&db.SellInfo{
			Id: 1000,
		})
		So(err, ShouldBeNil)
	})

}

func TestSrvInfo_Create(t *testing.T) {
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
		So(rsp.Status, ShouldEqual, sellinfo.SellInfoCreateResponse_SUCCESS)
		So(rsp.SellInfoId, ShouldNotEqual, 0)

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
	})
}

func TestSrvContent_Create(t *testing.T) {
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

		// clean up
		collection := db.MongoDatabase.Collection("sellinfo")
		rid, err := primitive.ObjectIDFromHex(req.ContentId)
		So(err, ShouldBeNil)
		type files struct {
			FileId primitive.ObjectID `bson:"fileId"`
			Type   string             `bson:"type"`
		}
		type result struct {
			Id    primitive.ObjectID `bson:"_id"`
			Files []files            `bson:"files"`
		}
		var res result
		err = collection.FindOneAndDelete(db.MongoContext, bson.D{
			{"_id", rid},
		}).Decode(&res)
		So(err, ShouldBeNil)
		for _, v := range res.Files {
			bucket, err := gridfs.NewBucket(db.MongoDatabase)
			So(err, ShouldBeNil)
			err = bucket.Delete(v.FileId)
			So(err, ShouldBeNil)
		}
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
