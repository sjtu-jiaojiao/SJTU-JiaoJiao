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
	utils.TestCheck(actual["status"], expect["status"])
	rspNews := actual["news"].([]*message.NewMessage)
	outNews := expect["news"].([]interface{})
	So(len(rspNews), ShouldEqual, len(outNews))
	for k, v := range rspNews {
		outData := outNews[k].(utils.StringMap)
		utils.TestCheck(v.FromUser, outData["fromUser"])
		utils.TestCheck(v.ToUser, outData["toUser"])
		utils.TestCheck(v.Badge, outData["badge"])
		utils.TestCheck(v.Info.Msg, outData["msg"])
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
	utils.TestCheck(actual["status"], expect["status"])
	utils.TestCheck(actual["fromUser"], expect["fromUser"])
	utils.TestCheck(actual["toUser"], expect["toUser"])
	utils.TestCheck(actual["badge"], expect["badge"])
	rspMsg := actual["infos"].([]*message.MessageInfo)
	outMsg := expect["infos"].([]interface{})
	So(len(rspMsg), ShouldEqual, len(outMsg))
	for k, v := range rspMsg {
		utils.TestCheck(v.Msg, outMsg[k])
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
