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
		"fromUser": 1000,
		"toUser":   1001,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")

	checkResult := func() {
		count, err := collection.CountDocuments(ctx, filter)
		So(count, ShouldEqual, 1)
		So(err, ShouldBeNil)

		var chat chatLog
		So(collection.FindOne(ctx, filter).Decode(&chat), ShouldBeNil)
		So(chat.Badge, ShouldEqual, 3)
		So(len(chat.Infos), ShouldEqual, 3)
		So(chat.Infos[0].Forward, ShouldEqual, false)
		So(chat.Infos[1].Msg, ShouldEqual, "012345678901234567891234")
	}

	Convey("Test Create Message", t, func() {
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_INVALID_PARAM)

		req.FromUser = 1000
		req.ToUser = 1001
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_INVALID_PARAM)

		req.Type = message.Type_TEXT
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_INVALID_PARAM)

		req.Msg = []byte("你好，我是小明(⊙﹏⊙)，🔺")
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_SUCCESS)
		defer func() {
			_, err := collection.DeleteOne(ctx, filter)
			So(err, ShouldBeNil)
		}()

		req.FromUser = 1000
		req.ToUser = 1001
		req.Type = message.Type_PICTURE
		req.Msg = []byte("valid_file")
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_SUCCESS)

		req.Msg = []byte("invalid_file")
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_INVALID_TYPE)

		req.FromUser = 1001
		req.ToUser = 1000
		req.Msg = []byte("valid_file")
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageCreateResponse_SUCCESS)

		checkResult()
	})
}

func TestFind(t *testing.T) {
	var s srv
	var req message.MessageFindRequest
	var rsp message.MessageFindResponse
	filter := primitive.M{
		"fromUser": 2000,
		"toUser":   2001,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")

	prepareData := func() {
		_, err := collection.InsertOne(ctx, bson.M{
			"fromUser": 2000,
			"toUser":   2001,
			"badge":    3,
			"infos": bson.A{bson.M{
				"time":    time.Now(),
				"forward": true,
				"type":    message.Type_TEXT,
				"msg":     "你好，我是小明1(⊙﹏⊙)，🔺",
				"unread":  false,
			}, bson.M{
				"time":    time.Now(),
				"forward": false,
				"type":    message.Type_TEXT,
				"msg":     "你好，我是小明2(⊙﹏⊙)，🔺",
				"unread":  false,
			}, bson.M{
				"time":    time.Now(),
				"forward": true,
				"type":    message.Type_TEXT,
				"msg":     "你好，我是小明3(⊙﹏⊙)，🔺",
				"unread":  true,
			}, bson.M{
				"time":    time.Now(),
				"forward": false,
				"type":    message.Type_TEXT,
				"msg":     "你好，我是小明4(⊙﹏⊙)，🔺",
				"unread":  true,
			}, bson.M{
				"time":    time.Now(),
				"forward": true,
				"type":    message.Type_PICTURE,
				"msg":     "012345678901234567890123",
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
		defer func() {
			_, err := collection.DeleteOne(ctx, filter)
			So(err, ShouldBeNil)
		}()

		req.FromUser = 2000
		req.ToUser = 2001
		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageFindResponse_INVALID_PARAM)

		req.Way = message.MessageFindRequest_READ_MESSAGE
		testBase(2, 3, message.MessageFindResponse_SUCCESS)
		So(rsp.Infos[1].Msg, ShouldEqual, "012345678901234567890123")

		testBase(0, 0, message.MessageFindResponse_SUCCESS)

		req.Way = message.MessageFindRequest_HISTORY
		testBase(5, 0, message.MessageFindResponse_SUCCESS)

		req.Limit = 3
		testBase(3, 0, message.MessageFindResponse_SUCCESS)
		rsp.Infos[0].Msg = "你好，我是小明1(⊙﹏⊙)，🔺"

		req.Offset = 1
		testBase(3, 0, message.MessageFindResponse_SUCCESS)
		rsp.Infos[0].Msg = "你好，我是小明2(⊙﹏⊙)，🔺"

		req.FromUser = 2001
		req.ToUser = 2000
		req.Way = message.MessageFindRequest_HISTORY
		testBase(3, 0, message.MessageFindResponse_SUCCESS)

		req.Way = message.MessageFindRequest_READ_MESSAGE
		testBase(1, 0, message.MessageFindResponse_SUCCESS)

		req.FromUser = 2011
		req.ToUser = 2010
		req.Limit = 100
		req.Way = message.MessageFindRequest_HISTORY
		testBase(0, 0, message.MessageFindResponse_SUCCESS)
	})

}

func TestQuery(t *testing.T) {
	var s srv
	var req message.MessageQueryRequest
	var rsp message.MessageQueryResponse
	filter1 := primitive.M{
		"fromUser": 3000,
		"toUser":   3001,
	}
	filter2 := primitive.M{
		"fromUser": 3000,
		"toUser":   3002,
	}
	filter3 := primitive.M{
		"fromUser": 3003,
		"toUser":   3000,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")

	prepareData := func() {
		_, err := collection.InsertMany(ctx, bson.A{
			bson.M{
				"fromUser": 3000,
				"toUser":   3001,
				"badge":    1,
				"infos": bson.A{
					bson.M{"time": time.Now(), "forward": false, "type": message.Type_TEXT, "msg": "你好，我是小明1(⊙﹏⊙)，🔺", "unread": true},
					bson.M{"time": time.Now(), "forward": false, "type": message.Type_TEXT, "msg": "你好，我是小明2(⊙﹏⊙)，🔺", "unread": false},
				},
			},
			bson.M{
				"fromUser": 3000,
				"toUser":   3002,
				"badge":    1,
				"infos": bson.A{
					bson.M{"time": time.Now(), "forward": true, "type": message.Type_TEXT, "msg": "你好，我是小明3(⊙﹏⊙)，🔺", "unread": true},
					bson.M{"time": time.Now(), "forward": false, "type": message.Type_TEXT, "msg": "你好，我是小明4(⊙﹏⊙)，🔺", "unread": false},
				},
			},
			bson.M{
				"fromUser": 3003,
				"toUser":   3000,
				"badge":    2,
				"infos": bson.A{
					bson.M{"time": time.Now(), "forward": true, "type": message.Type_TEXT, "msg": "你好，我是小明5(⊙﹏⊙)，🔺", "unread": true},
					bson.M{"time": time.Now(), "forward": true, "type": message.Type_TEXT, "msg": "你好，我是小明6(⊙﹏⊙)，🔺", "unread": true},
				},
			},
		})
		So(err, ShouldBeNil)
	}

	testBase := func(newsLen int, status message.MessageQueryResponse_Status) {
		rsp.Reset()
		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		So(len(rsp.News), ShouldEqual, newsLen)
	}
	Convey("Test Find Message", t, func() {
		prepareData()
		defer func() {
			_, err := collection.DeleteOne(ctx, filter1)
			So(err, ShouldBeNil)
			_, err = collection.DeleteOne(ctx, filter2)
			So(err, ShouldBeNil)
			_, err = collection.DeleteOne(ctx, filter3)
			So(err, ShouldBeNil)
		}()

		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, message.MessageQueryResponse_INVALID_PARAM)

		req.UserID = 3000
		testBase(2, message.MessageQueryResponse_SUCCESS)
		So(rsp.News[0].Badge, ShouldEqual, 1)
		So(rsp.News[1].Badge, ShouldEqual, 2)
		So(rsp.News[1].Info.Msg, ShouldEqual, "你好，我是小明5(⊙﹏⊙)，🔺")

		req.UserID = 3001
		testBase(0, message.MessageQueryResponse_SUCCESS)

		req.UserID = 3002
		testBase(1, message.MessageQueryResponse_SUCCESS)
		So(rsp.News[0].Info.Msg, ShouldEqual, "你好，我是小明3(⊙﹏⊙)，🔺")

		req.UserID = 3003
		testBase(0, message.MessageQueryResponse_SUCCESS)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
