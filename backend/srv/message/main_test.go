package main

import (
	"context"
	db "jiaojiao/database"
	message "jiaojiao/srv/message/proto"
	"jiaojiao/utils"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//
//func TestFind(t *testing.T) {
//	var s srv
//	var req message.MessageFindRequest
//	var rsp message.MessageFindResponse
//	filter := primitive.M{
//		"fromUser": 2000,
//		"toUser":   2001,
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	collection := db.MongoDatabase.Collection("message")
//
//	prepareData := func() {
//		_, err := collection.InsertOne(ctx, bson.M{
//			"fromUser": 2000,
//			"toUser":   2001,
//			"badge":    3,
//			"infos": bson.A{bson.M{
//				"time":    time.Now(),
//				"forward": true,
//				"type":    message.Type_TEXT,
//				"msg":     "你好，我是小明1(⊙﹏⊙)，🔺",
//				"unread":  false,
//			}, bson.M{
//				"time":    time.Now(),
//				"forward": false,
//				"type":    message.Type_TEXT,
//				"msg":     "你好，我是小明2(⊙﹏⊙)，🔺",
//				"unread":  false,
//			}, bson.M{
//				"time":    time.Now(),
//				"forward": true,
//				"type":    message.Type_TEXT,
//				"msg":     "你好，我是小明3(⊙﹏⊙)，🔺",
//				"unread":  true,
//			}, bson.M{
//				"time":    time.Now(),
//				"forward": false,
//				"type":    message.Type_TEXT,
//				"msg":     "你好，我是小明4(⊙﹏⊙)，🔺",
//				"unread":  true,
//			}, bson.M{
//				"time":    time.Now(),
//				"forward": true,
//				"type":    message.Type_PICTURE,
//				"msg":     "012345678901234567890123",
//				"unread":  true,
//			}},
//		})
//		So(err, ShouldBeNil)
//	}
//
//	testBase := func(infoLen int, badge int, status message.MessageFindResponse_Status) {
//		rsp.Reset()
//		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
//		So(rsp.Status, ShouldEqual, status)
//		So(rsp.Badge, ShouldEqual, badge)
//		So(len(rsp.Infos), ShouldEqual, infoLen)
//	}
//	Convey("Test Find Message", t, func() {
//		prepareData()
//		defer func() {
//			_, err := collection.DeleteOne(ctx, filter)
//			So(err, ShouldBeNil)
//		}()
//
//		req.FromUser = 2000
//		req.ToUser = 2001
//		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
//		So(rsp.Status, ShouldEqual, message.MessageFindResponse_INVALID_PARAM)
//
//		req.Way = message.MessageFindRequest_READ_MESSAGE
//		testBase(2, 3, message.MessageFindResponse_SUCCESS)
//		So(rsp.Infos[1].Msg, ShouldEqual, "012345678901234567890123")
//
//		testBase(0, 0, message.MessageFindResponse_SUCCESS)
//
//		req.Way = message.MessageFindRequest_HISTORY
//		testBase(5, 0, message.MessageFindResponse_SUCCESS)
//
//		req.Limit = 3
//		testBase(3, 0, message.MessageFindResponse_SUCCESS)
//		rsp.Infos[0].Msg = "你好，我是小明1(⊙﹏⊙)，🔺"
//
//		req.Offset = 1
//		testBase(3, 0, message.MessageFindResponse_SUCCESS)
//		rsp.Infos[0].Msg = "你好，我是小明2(⊙﹏⊙)，🔺"
//
//		req.FromUser = 2001
//		req.ToUser = 2000
//		req.Way = message.MessageFindRequest_HISTORY
//		testBase(3, 0, message.MessageFindResponse_SUCCESS)
//
//		req.Way = message.MessageFindRequest_READ_MESSAGE
//		testBase(1, 0, message.MessageFindResponse_SUCCESS)
//
//		req.FromUser = 2011
//		req.ToUser = 2010
//		req.Limit = 100
//		req.Way = message.MessageFindRequest_HISTORY
//		testBase(0, 0, message.MessageFindResponse_SUCCESS)
//	})
//
//}
//

func ParseCreate(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp message.MessageCreateResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Create(context.TODO(), &message.MessageCreateRequest{
		FromUser: utils.TestInt(input["fromUser"]),
		ToUser:   utils.TestInt(input["toUser"]),
		Type:     message.Type(utils.EnumConvert(utils.TestInt(input["type"]), message.Type_name)),
		Msg:      utils.TestByte(input["msg"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	return ret
}

func VerifyData(verify utils.StringMap, output utils.StringMap) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")
	filter := primitive.M{
		"fromUser": utils.TestInt(verify["fromUser"]),
		"toUser":   utils.TestInt(verify["toUser"]),
	}
	var chat chatLog

	if exist, ok := verify["_exist"]; !ok || exist.(bool) {
		So(collection.FindOne(ctx, filter).Decode(&chat), ShouldBeNil)
		So(chat.Badge, ShouldEqual, utils.TestInt(verify["badge"]))

		verifyInfo := verify["infos"].([]interface{})
		So(len(chat.Infos), ShouldEqual, len(verifyInfo))
		for k, v := range chat.Infos {
			verifyData := verifyInfo[k].(utils.StringMap)
			So(v.Forward, ShouldEqual, verifyData["forward"])
			So(v.Type, ShouldEqual, verifyData["type"])
			So(v.Msg, ShouldEqual, verifyData["msg"])
			So(v.Unread, ShouldEqual, verifyData["unread"])
		}
	} else {
		So(collection.FindOne(ctx, filter).Decode(&chat), ShouldNotBeNil)
	}
}

func ParseQuery(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp message.MessageQueryResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Query(context.TODO(), &message.MessageQueryRequest{
		UserID: utils.TestInt(input["userID"]),
		OldMsg: utils.TestBool(input["oldMsg"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["news"] = rsp.News
	return ret
}

func CheckQuery(actual utils.StringMap, expect utils.StringMap) {
	So(actual["status"], ShouldEqual, expect["status"])
	rspNews := actual["news"].([]*message.NewMessage)
	outNews := expect["news"].([]interface{})
	So(len(rspNews), ShouldEqual, len(outNews))
	for k, v := range rspNews {
		outData := outNews[k].(utils.StringMap)
		So(v.FromUser, ShouldEqual, outData["fromUser"])
		So(v.ToUser, ShouldEqual, outData["toUser"])
		So(v.Badge, ShouldEqual, outData["badge"])
		So(v.Info.Msg, ShouldEqual, outData["msg"])
	}
}

func InsertData(data utils.StringMap) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")

	var info bson.A
	for _, v := range data["infos"].([]interface{}) {
		t := v.(utils.StringMap)
		info = append(info, bson.M{
			"time":    time.Now(),
			"forward": t["forward"],
			"type":    t["type"],
			"msg":     t["msg"],
			"unread":  t["unread"],
		})
	}
	_, err := collection.InsertOne(ctx,
		bson.M{
			"fromUser": utils.TestInt(data["fromUser"]),
			"toUser":   utils.TestInt(data["toUser"]),
			"badge":    utils.TestInt(data["badge"]),
			"infos":    info,
		})
	So(err, ShouldBeNil)
}

func ParseFind(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp message.MessageFindResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Find(context.TODO(), &message.MessageFindRequest{
		FromUser: utils.TestInt(input["fromUser"]),
		ToUser:   utils.TestInt(input["toUser"]),
		Way:      message.MessageFindRequest_Way(utils.EnumConvert(utils.TestInt(input["way"]), message.MessageFindRequest_Way_name)),
		Limit:    utils.TestUint(input["limit"]),
		Offset:   utils.TestUint(input["offset"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["fromUser"] = rsp.FromUser
	ret["toUser"] = rsp.ToUser
	ret["badge"] = rsp.Badge
	ret["infos"] = rsp.Infos
	return ret
}

func CheckFind(actual utils.StringMap, expect utils.StringMap) {
	So(actual["status"], ShouldEqual, expect["status"])
	So(actual["fromUser"], ShouldEqual, expect["fromUser"])
	So(actual["toUser"], ShouldEqual, expect["toUser"])
	So(actual["badge"], ShouldEqual, expect["badge"])
	rspMsg := actual["infos"].([]*message.MessageInfo)
	outMsg := expect["infos"].([]interface{})
	So(len(rspMsg), ShouldEqual, len(outMsg))
	for k, v := range rspMsg {
		So(v.Msg, ShouldEqual, outMsg[k])
	}
}

func cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("message")
	err := collection.Drop(ctx)
	utils.LogPanic(err)
}

func TestAll(t *testing.T) {
	cleanup()
	utils.Test(t, "test/test_create.json", InsertData, ParseCreate, VerifyData, nil)
	cleanup()
	utils.Test(t, "test/test_query.json", InsertData, ParseQuery, nil, CheckQuery)
	cleanup()
	utils.Test(t, "test/test_find.json", InsertData, ParseFind, VerifyData, CheckFind)
	cleanup()
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
