package main

import (
	"context"
	db "jiaojiao/database"
	content "jiaojiao/srv/content/proto"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
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
		req.Content = []byte("invalid_file")
		tf(content.ContentCreateResponse_INVALID_TYPE, false)
		req.Content = []byte("valid_file")
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
		req.ContentID = "123"
		tf(content.ContentCreateTagResponse_INVALID_PARAM)
		req.ContentToken = "456"
		tf(content.ContentCreateTagResponse_INVALID_TOKEN)
		req.ContentID = ""
		req.ContentToken = ""
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
	tf := func(status content.ContentCheckResponse_Status, id string, token string) {
		var s srv
		var rsp content.ContentCheckResponse
		So(s.Check(context.TODO(), &content.ContentCheckRequest{
			ContentID:    id,
			ContentToken: token,
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
	}

	Convey("Test content check", t, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		collection := db.MongoDatabase.Collection("content")
		token := uuid.NewV4().String()
		res, err := collection.InsertOne(ctx, bson.M{
			"token": token,
			"files": bson.A{
				bson.M{
					"fileID": "012345678901234567891234",
					"type":   1,
				}},
		})
		So(err, ShouldBeNil)
		id := res.InsertedID.(primitive.ObjectID).Hex()
		defer func() {
			err = collection.FindOneAndDelete(ctx, bson.D{
				{"_id", res.InsertedID.(primitive.ObjectID)},
				{"token", token},
			}).Decode(&res)
			So(err, ShouldBeNil)
		}()

		tf(content.ContentCheckResponse_INVALID_PARAM, "", "")
		tf(content.ContentCheckResponse_INVALID_PARAM, id, "")
		tf(content.ContentCheckResponse_INVALID_PARAM, "", token)
		tf(content.ContentCheckResponse_INVALID, id, "123123")
		tf(content.ContentCheckResponse_VALID, id, token)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
