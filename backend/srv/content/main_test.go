package main

import (
	"context"
	db "jiaojiao/database"
	content "jiaojiao/srv/content/proto"
	"jiaojiao/utils"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	. "github.com/smartystreets/goconvey/convey"
)

func ParseCreate(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp content.ContentCreateResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Create(context.TODO(), &content.ContentCreateRequest{
		ContentID:    utils.TestString(input["contentID"]),
		ContentToken: utils.TestString(input["contentToken"]),
		Content:      utils.TestByte(input["content"]),
		Type:         content.Type(utils.EnumConvert(utils.TestInt(input["type"]), content.Type_name)),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["contentID"] = rsp.ContentID
	ret["contentToken"] = rsp.ContentToken
	ret["fileID"] = rsp.FileID
	return ret
}

func ParseCreateTag(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp content.ContentCreateTagResponse
	ret := make(utils.StringMap)
	var tags []string
	for _, v := range input["tags"].([]interface{}) {
		tags = append(tags, utils.TestString(v))
	}

	ret["_error"] = s.CreateTag(context.TODO(), &content.ContentCreateTagRequest{
		ContentID:    utils.TestString(input["contentID"]),
		ContentToken: utils.TestString(input["contentToken"]),
		Tags:         tags,
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["contentID"] = rsp.ContentID
	ret["contentToken"] = rsp.ContentToken
	return ret
}

func ParseQuery(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp content.ContentQueryResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Query(context.TODO(), &content.ContentQueryRequest{
		ContentID: utils.TestString(input["contentID"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["contentToken"] = rsp.ContentToken
	ret["files"] = rsp.Files
	ret["tags"] = rsp.Tags
	return ret
}

func ParseCheck(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp content.ContentCheckResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Check(context.TODO(), &content.ContentCheckRequest{
		ContentID:    utils.TestString(input["contentID"]),
		ContentToken: utils.TestString(input["contentToken"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	return ret
}

func ParseDelete(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp content.ContentDeleteResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Delete(context.TODO(), &content.ContentDeleteRequest{
		ContentID:    utils.TestString(input["contentID"]),
		ContentToken: utils.TestString(input["contentToken"]),
		FileID:       utils.TestString(input["fileID"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	return ret
}
func ParseUpdate(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp content.ContentUpdateResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Update(context.TODO(), &content.ContentUpdateRequest{
		ContentID:    utils.TestString(input["contentID"]),
		ContentToken: utils.TestString(input["contentToken"]),
		FileID:       utils.TestString(input["fileID"]),
		Content:      utils.TestByte(input["content"]),
		Type:         content.Type(utils.EnumConvert(utils.TestInt(input["type"]), content.Type_name)),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["fileID"] = rsp.FileID
	return ret
}

func CheckQuery(actual utils.StringMap, expect utils.StringMap) {
	utils.TestCheck(actual["status"], expect["status"])
	utils.TestCheck(actual["contentToken"], expect["contentToken"])
	rspTags := actual["tags"].([]string)
	outTags := expect["tags"].([]interface{})
	utils.TestCheck(len(rspTags), len(outTags))
	for k, v := range rspTags {
		So(v, ShouldEqual, outTags[k])
	}
	rspFiles := actual["files"].([]*content.FileMsg)
	outFiles := expect["files"].([]interface{})
	utils.TestCheck(len(rspFiles), len(outFiles))
	for k, v := range rspFiles {
		outData := outFiles[k].(utils.StringMap)
		utils.TestCheck(v.FileID, outData["fileID"])
		utils.TestCheck(v.Type, outData["type"])
	}
}

func InsertData(data utils.StringMap) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("content")
	id, err := primitive.ObjectIDFromHex(utils.TestString(data["contentID"]))
	So(err, ShouldBeNil)

	var files bson.A
	for _, v := range data["files"].([]interface{}) {
		t := v.(utils.StringMap)
		fid, err := primitive.ObjectIDFromHex(utils.TestString(t["fileID"]))
		So(err, ShouldBeNil)
		files = append(files, bson.M{
			"fileID": fid,
			"type":   t["type"],
		})
	}
	_, err = collection.InsertOne(ctx,
		bson.M{
			"_id":   id,
			"token": data["contentToken"],
			"tags":  data["tags"],
			"files": files,
		})
	So(err, ShouldBeNil)
}

func VerifyData(verify utils.StringMap, output utils.StringMap) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("content")
	id, err := primitive.ObjectIDFromHex(utils.TransVarString("contentID", verify, output))
	So(err, ShouldBeNil)
	filter := primitive.M{
		"_id":   id,
		"token": utils.TransVarString("contentToken", verify, output),
	}
	var res result

	if exist, ok := verify["_exist"]; !ok || exist.(bool) {
		So(collection.FindOne(ctx, filter).Decode(&res), ShouldBeNil)

		verifyFile := verify["files"].([]interface{})
		So(len(res.Files), ShouldEqual, len(verifyFile))
		for k, v := range res.Files {
			verifyData := verifyFile[k].(utils.StringMap)
			So(v.FileID.Hex(), ShouldEqual, utils.TransVarString("fileID", verifyData, output))
			So(v.Type, ShouldEqual, verifyData["type"])
		}
		verifyTags := verify["tags"].([]interface{})
		So(len(res.Tags), ShouldEqual, len(verifyTags))
		for k, v := range res.Tags {
			So(v, ShouldEqual, verifyTags[k])
		}
	} else {
		So(collection.FindOne(ctx, filter).Decode(&res), ShouldNotBeNil)
	}
}

func cleanup() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("content")
	err := collection.Drop(ctx)
	utils.LogPanic(err)
}

func TestAll(t *testing.T) {
	cleanup()
	utils.Test(t, "test/test_create.json", InsertData, ParseCreate, VerifyData, nil)
	cleanup()
	utils.Test(t, "test/test_createtag.json", InsertData, ParseCreateTag, VerifyData, nil)
	cleanup()
	utils.Test(t, "test/test_query.json", InsertData, ParseQuery, nil, CheckQuery)
	cleanup()
	utils.Test(t, "test/test_check.json", InsertData, ParseCheck, nil, nil)
	cleanup()
	utils.Test(t, "test/test_delete.json", InsertData, ParseDelete, VerifyData, nil)
	cleanup()
	utils.Test(t, "test/test_update.json", InsertData, ParseUpdate, VerifyData, nil)
	cleanup()
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
