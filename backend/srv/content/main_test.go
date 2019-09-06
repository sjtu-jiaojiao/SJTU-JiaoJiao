package main

import (
	"context"
	db "jiaojiao/database"
	content "jiaojiao/srv/content/proto"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestContentCreate(t *testing.T) {
	var req content.ContentCreateRequest

	tf := func(status content.ContentCreateResponse_Status, success bool) (string, string) {
		var s srv
		var rsp content.ContentCreateResponse
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if success {
			So(rsp.ContentID, ShouldNotBeBlank)
			So(rsp.ContentToken, ShouldNotBeBlank)
		} else {
			So(rsp.ContentID, ShouldBeBlank)
			So(rsp.ContentToken, ShouldBeBlank)
		}
		return rsp.ContentID, rsp.ContentToken
	}

	Convey("Test content create", t, func() {
		req.Content = []byte{0}
		tf(content.ContentCreateResponse_INVALID_PARAM, false)

		req.Type = content.Type_PICTURE
		tf(content.ContentCreateResponse_INVALID_PARAM, false)
		req.Type = 0
		req.Content = []byte("valid_file")
		tf(content.ContentCreateResponse_INVALID_PARAM, false)
		req.Type = content.Type_PICTURE
		req.ContentID = "1234"
		tf(content.ContentCreateResponse_INVALID_PARAM, false)
		req.ContentID = ""
		req.ContentToken = "12463-25897fsfs-5232"
		tf(content.ContentCreateResponse_INVALID_PARAM, false)

		req.ContentID = "1234"
		tf(content.ContentCreateResponse_INVALID_TOKEN, false)

		req.ContentID = ""
		req.ContentToken = ""
		id, token := tf(content.ContentCreateResponse_SUCCESS, true)

		req.ContentID = id
		req.ContentToken = token
		defer func() {
			var sc srv
			var rspc content.ContentDeleteResponse
			err := sc.Delete(context.TODO(), &content.ContentDeleteRequest{
				ContentID:    id,
				ContentToken: token,
			}, &rspc)
			So(err, ShouldBeNil)
		}()
		tf(content.ContentCreateResponse_SUCCESS, true)
		tf(content.ContentCreateResponse_SUCCESS, true)

		req.ContentToken = "12463-25897fsfs-5232"
		tf(content.ContentCreateResponse_INVALID_TOKEN, false)
	})
}

func TestCreateTag(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("content")

	var req content.ContentCreateTagRequest

	tf := func(status content.ContentCreateTagResponse_Status) (string, string) {
		var s srv
		var rsp content.ContentCreateTagResponse
		So(s.CreateTag(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == content.ContentCreateTagResponse_SUCCESS {
			So(rsp.ContentID, ShouldNotBeBlank)
			So(rsp.ContentToken, ShouldNotBeBlank)

			cid, err := primitive.ObjectIDFromHex(rsp.ContentID)
			So(err, ShouldBeNil)
			var res result
			err = collection.FindOne(ctx, bson.D{
				{"_id", cid},
			}).Decode(&res)
			So(err, ShouldBeNil)
			So(res.Tags, ShouldResemble, req.Tags)
		} else {
			So(rsp.ContentID, ShouldBeBlank)
			So(rsp.ContentToken, ShouldBeBlank)
		}
		return rsp.ContentID, rsp.ContentToken
	}

	Convey("Test content create tag", t, func() {
		tf(content.ContentCreateTagResponse_INVALID_PARAM)

		req.Tags = []string{"123", "456"}
		req.ContentID, req.ContentToken = tf(content.ContentCreateTagResponse_SUCCESS)
		defer func() {
			var res result
			cid, err := primitive.ObjectIDFromHex(req.ContentID)
			So(err, ShouldBeNil)
			err = collection.FindOneAndDelete(ctx, bson.D{
				{"_id", cid},
				{"token", req.ContentToken},
			}).Decode(&res)
			So(err, ShouldBeNil)
		}()

		req.Tags = []string{"789"}
		tf(content.ContentCreateTagResponse_SUCCESS)
	})
}

func TestContentDelete(t *testing.T) {
	// TODO
}

func TestContentQuery(t *testing.T) {
	// TODO
}

func TestContentCheck(t *testing.T) {
	// TODO
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
