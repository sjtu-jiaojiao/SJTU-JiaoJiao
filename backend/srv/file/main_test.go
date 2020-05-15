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

func ParseCreate(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp file.FileCreateResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Create(context.TODO(), &file.FileCreateRequest{
		File: utils.TestByte(input["file"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["fileID"] = rsp.FileID
	return ret
}

func VerifyData(verify utils.StringMap, output utils.StringMap) {

	bucket, err := gridfs.NewBucket(db.MongoDatabase)
	So(err, ShouldBeNil)

	fid, err := primitive.ObjectIDFromHex(utils.TransVarString("fileID", verify, output))
	So(err, ShouldBeNil)

	var buf bytes.Buffer
	_, err = bucket.DownloadToStream(fid, &buf)
	if exist, ok := verify["_exist"]; !ok || exist.(bool) {
		So(err, ShouldBeNil)
		So(buf.String(), ShouldEqual, utils.TestString(verify["file"]))
	} else {
		So(err, ShouldNotBeNil)
	}
}

func ParseQuery(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp file.FileQueryResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Query(context.TODO(), &file.FileRequest{
		FileID: utils.TestString(input["fileID"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["file"] = string(rsp.File)
	ret["size"] = rsp.Size
	return ret
}

func InsertData(data utils.StringMap) {
	bucket, err := gridfs.NewBucket(db.MongoDatabase)
	So(err, ShouldBeNil)
	fid, err := primitive.ObjectIDFromHex(utils.TestString(data["fileID"]))
	So(err, ShouldBeNil)
	err = bucket.UploadFromStreamWithID(fid, "", bytes.NewReader(utils.TestByte(data["file"])))
	So(err, ShouldBeNil)
}

func ParseDelete(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp file.FileDeleteResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Delete(context.TODO(), &file.FileRequest{
		FileID: utils.TestString(input["fileID"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
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
	utils.Test(t, "test/test_create.json", nil, ParseCreate, VerifyData, nil)
	cleanup()
	utils.Test(t, "test/test_query.json", InsertData, ParseQuery, nil, nil)
	cleanup()
	utils.Test(t, "test/test_delete.json", InsertData, ParseDelete, VerifyData, nil)
	cleanup()
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
