package main

import (
	"context"
	db "jiaojiao/database"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func TestCreate(t *testing.T) {
	var s srvContent
	var req sellinfo.ContentCreateRequest
	var rsp sellinfo.ContentCreateResponse
	Convey("Test SellInfo Create", t, func() {
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
