package main

import (
	"context"
	db "jiaojiao/database"
	message "jiaojiao/srv/message/proto"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	var s srv
	var req message.MessageCreateRequest
	var rsp message.MessageCreateResponse
	filter := primitive.M{
		"fromUser": 1001,
		"toUser":   2001,
	}
	Convey("Test Create Message", t, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := db.MongoDatabase.Collection("message")

		_, err := collection.DeleteOne(ctx, filter)
		So(err, ShouldBeNil)

		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_INVALID_PARAM)

		req.FromUser = 1001
		req.ToUser = 2001
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_INVALID_PARAM)

		req.Type = message.MessageCreateRequest_TEXT
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_INVALID_PARAM)

		req.Text = "你好，我是小明(⊙﹏⊙)，🔺"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_SUCCESS)

		req.Type = message.MessageCreateRequest_PICTURE
		req.File = []byte{0, 1, 2, 3, 4, 5}
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_SUCCESS)

		req.FromUser = 2001
		req.ToUser = 1001
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_SUCCESS)

		count, err := collection.CountDocuments(ctx, filter)
		So(count, ShouldEqual, 1)
		So(err, ShouldBeNil)

		var chat ChatLog
		So(collection.FindOne(ctx, filter).Decode(&chat), ShouldBeNil)
		So(chat.Badge, ShouldEqual, 3)
		So(len(chat.Infos), ShouldEqual, 3)
		So(chat.Infos[2].Forward, ShouldEqual, false)
		So(chat.Infos[1].Text, ShouldEqual, "1234567890abcdef12345678")

		defer func() {
			_, err = collection.DeleteOne(ctx, filter)
			So(err, ShouldBeNil)
		}()
	})
}

func TestFind(t *testing.T) {
	var s srv
	var req message.MessageFindRequest
	Convey("Test Find Message", t, func() {
		err := s.Find(context.TODO(), &req, &message.MessageFindResponse{})
		ShouldBeNil(err)
	})

}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
