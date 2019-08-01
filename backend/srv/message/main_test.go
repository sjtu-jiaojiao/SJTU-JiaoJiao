package main

import (
	"context"
	db "jiaojiao/database"
	message "jiaojiao/srv/message/proto"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")

	checkResult := func() {
		count, err := collection.CountDocuments(ctx, filter)
		So(count, ShouldEqual, 1)
		So(err, ShouldBeNil)

		var chat ChatLog
		So(collection.FindOne(ctx, filter).Decode(&chat), ShouldBeNil)
		So(chat.Badge, ShouldEqual, 3)
		So(len(chat.Infos), ShouldEqual, 3)
		So(chat.Infos[0].Forward, ShouldEqual, false)
		So(chat.Infos[1].Text, ShouldEqual, "1234567890abcdef12345678")
	}

	Convey("Test Create Message", t, func() {
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

		checkResult()

		defer func() {
			_, err = collection.DeleteOne(ctx, filter)
			So(err, ShouldBeNil)
		}()
	})
}

func TestFind(t *testing.T) {
	var s srv
	var req message.MessageFindRequest
	var rsp message.MessageFindResponse
	filter := primitive.M{
		"fromUser": 1001,
		"toUser":   2001,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")

	prepareData := func() {
		_, err := collection.DeleteOne(ctx, filter)
		So(err, ShouldBeNil)

		_, err = collection.InsertOne(ctx, bson.M{
			"fromUser": 1001,
			"toUser":   2001,
			"badge":    3,
			"infos": bson.A{bson.M{
				"time":    time.Now(),
				"forward": true,
				"type":    message.MessageInfo_TEXT,
				"text":    "你好，我是小明1(⊙﹏⊙)，🔺",
				"unread":  false,
			}, bson.M{
				"time":    time.Now(),
				"forward": false,
				"type":    message.MessageInfo_TEXT,
				"text":    "你好，我是小明2(⊙﹏⊙)，🔺",
				"unread":  false,
			}, bson.M{
				"time":    time.Now(),
				"forward": true,
				"type":    message.MessageInfo_TEXT,
				"text":    "你好，我是小明3(⊙﹏⊙)，🔺",
				"unread":  true,
			}, bson.M{
				"time":    time.Now(),
				"forward": false,
				"type":    message.MessageInfo_TEXT,
				"text":    "你好，我是小明4(⊙﹏⊙)，🔺",
				"unread":  true,
			}, bson.M{
				"time":    time.Now(),
				"forward": true,
				"type":    message.MessageCreateRequest_PICTURE,
				"text":    "1234567890abcdef12345678",
				"unread":  true,
			}},
		})
		So(err, ShouldBeNil)
	}

	testBase := func(infoLen int, badge int, status message.MessageFindResponse_Status) {
		rsp.Reset()
		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		So(rsp.Badge, ShouldEqual, badge)
		So(len(rsp.Infos), ShouldEqual, infoLen)
	}
	Convey("Test Find Message", t, func() {
		prepareData()

		req.FromUser = 1001
		req.ToUser = 2001
		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageFindResponse_INVALID_PARAM)

		req.Way = message.MessageFindRequest_ONLY_PULL
		testBase(2, 3, message.MessageFindResponse_SUCCESS)
		So(rsp.Infos[1].Text, ShouldEqual, "1234567890abcdef12345678")

		req.Way = message.MessageFindRequest_READ_MESSAGE
		testBase(2, 3, message.MessageFindResponse_SUCCESS)
		So(rsp.Infos[1].Text, ShouldEqual, "1234567890abcdef12345678")

		testBase(0, 0, message.MessageFindResponse_SUCCESS)

		req.Way = message.MessageFindRequest_HISTORY
		testBase(5, 0, message.MessageFindResponse_SUCCESS)

		req.Limit = 3
		testBase(3, 0, message.MessageFindResponse_SUCCESS)
		rsp.Infos[0].Text = "你好，我是小明1(⊙﹏⊙)，🔺"

		req.Offset = 1
		testBase(3, 0, message.MessageFindResponse_SUCCESS)
		rsp.Infos[0].Text = "你好，我是小明2(⊙﹏⊙)，🔺"

		req.FromUser = 2001
		req.Way = message.MessageFindRequest_ONLY_PULL
		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageFindResponse_NOT_FOUND)

		req.ToUser = 1001
		testBase(1, 0, message.MessageFindResponse_SUCCESS)
		So(rsp.Infos[0].Text, ShouldEqual, "你好，我是小明4(⊙﹏⊙)，🔺")

		req.Way = message.MessageFindRequest_HISTORY
		testBase(3, 0, message.MessageFindResponse_SUCCESS)

		req.Way = message.MessageFindRequest_READ_MESSAGE
		testBase(1, 0, message.MessageFindResponse_SUCCESS)

		defer func() {
			_, err := collection.DeleteOne(ctx, filter)
			So(err, ShouldBeNil)
		}()
	})

}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
