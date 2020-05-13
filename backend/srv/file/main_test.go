package main

import (
	"bytes"
	"context"
	db "jiaojiao/database"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func ParseCreate(param interface{}) map[string]interface{} {
	input := param.(map[string]interface{})

	var s srv
	var rsp file.FileCreateResponse
	So(s.Create(context.TODO(), &file.FileCreateRequest{
		File: []byte(input["file"].(string)),
	}, &rsp), ShouldBeNil)

	ret := make(map[string]interface{})
	ret["status"] = rsp.Status.String()
	ret["fileID"] = rsp.FileID
	return ret
}

func VerifyData(param interface{}, output interface{}) {
	verify := param.(map[string]interface{})
	variable := output.(map[string]interface{})

	bucket, err := gridfs.NewBucket(db.MongoDatabase)
	So(err, ShouldBeNil)

	fid, err := primitive.ObjectIDFromHex(utils.TransVar("fileID", verify, variable).(string))
	So(err, ShouldBeNil)

	var buf bytes.Buffer
	_, err = bucket.DownloadToStream(fid, &buf)
	if exist, ok := verify["_exist"]; !ok || exist.(bool) {
		So(err, ShouldBeNil)
		So(buf.String(), ShouldEqual, verify["file"].(string))
	} else {
		So(err, ShouldNotBeNil)
	}
}

func ParseQuery(param interface{}) map[string]interface{} {
	input := param.(map[string]interface{})

	var s srv
	var rsp file.FileQueryResponse
	So(s.Query(context.TODO(), &file.FileRequest{
		FileID: input["fileID"].(string),
	}, &rsp), ShouldBeNil)

	ret := make(map[string]interface{})
	ret["status"] = rsp.Status.String()
	ret["file"] = string(rsp.File)
	ret["size"] = rsp.Size
	return ret
}

func InsertData(param interface{}) {
	data := param.(map[string]interface{})
	bucket, err := gridfs.NewBucket(db.MongoDatabase)
	So(err, ShouldBeNil)
	fid, err := primitive.ObjectIDFromHex(data["fileID"].(string))
	So(err, ShouldBeNil)
	err = bucket.UploadFromStreamWithID(fid, "", bytes.NewReader([]byte(data["file"].(string))))
	So(err, ShouldBeNil)
}

func ParseDelete(param interface{}) map[string]interface{} {
	input := param.(map[string]interface{})

	var s srv
	var rsp file.FileDeleteResponse
	So(s.Delete(context.TODO(), &file.FileRequest{
		FileID: input["fileID"].(string),
	}, &rsp), ShouldBeNil)

	ret := make(map[string]interface{})
	ret["status"] = rsp.Status.String()
	return ret
}

func cleanup() {
	bucket, err := gridfs.NewBucket(db.MongoDatabase)
	utils.LogPanic(err)
	err = bucket.Drop()
	utils.LogPanic(err)
}

func TestAll(t *testing.T) {
	cleanup()
	utils.Test(t, "test/test_create.json", nil, ParseCreate, VerifyData)
	cleanup()
	utils.Test(t, "test/test_query.json", InsertData, ParseQuery, nil)
	cleanup()
	utils.Test(t, "test/test_delete.json", InsertData, ParseDelete, VerifyData)
	cleanup()
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
