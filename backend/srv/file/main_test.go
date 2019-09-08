package main

import (
	"bytes"
	"context"
	db "jiaojiao/database"
	file "jiaojiao/srv/file/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func TestQuery(t *testing.T) {
	tf := func(id string, ctx string, status file.FileQueryResponse_Status) {
		var s srv
		var rsp file.FileQueryResponse
		So(s.Query(context.TODO(), &file.FileRequest{
			FileID: id,
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == file.FileQueryResponse_SUCCESS {
			So(string(rsp.File), ShouldEqual, ctx)
			So(rsp.Size, ShouldEqual, len(ctx))
		}
	}
	Convey("Test file query", t, func() {
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		So(err, ShouldBeNil)
		objID, err := bucket.UploadFromStream("", bytes.NewReader([]byte("valid")))
		So(err, ShouldBeNil)
		defer func() { So(bucket.Delete(objID), ShouldBeNil) }()

		tf("", "", file.FileQueryResponse_INVALID_PARAM)
		tf("invalid", "", file.FileQueryResponse_INVALID_PARAM)
		tf("012345678901234567890123", "", file.FileQueryResponse_NOT_FOUND)
		tf(objID.Hex(), "valid", file.FileQueryResponse_SUCCESS)
	})
}

func TestCreate(t *testing.T) {
	tf := func(f string, status file.FileCreateResponse_Status) string {
		var s srv
		var rsp file.FileCreateResponse
		So(s.Create(context.TODO(), &file.FileCreateRequest{
			File: []byte(f),
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		return rsp.FileID
	}
	Convey("Test file create", t, func() {
		tf("", file.FileCreateResponse_INVALID_PARAM)
		id := tf("valid", file.FileCreateResponse_SUCCESS)
		defer func() {
			bucket, err := gridfs.NewBucket(db.MongoDatabase)
			So(err, ShouldBeNil)
			fid, err := primitive.ObjectIDFromHex(id)
			So(err, ShouldBeNil)
			So(bucket.Delete(fid), ShouldBeNil)
		}()
	})
}

func TestDelete(t *testing.T) {
	tf := func(id string, status file.FileDeleteResponse_Status) {
		var s srv
		var rsp file.FileDeleteResponse
		So(s.Delete(context.TODO(), &file.FileRequest{
			FileID: id,
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
	}
	Convey("Test file delete", t, func() {
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		So(err, ShouldBeNil)
		objID, err := bucket.UploadFromStream("", bytes.NewReader([]byte("valid")))
		So(err, ShouldBeNil)

		tf("", file.FileDeleteResponse_INVALID_PARAM)
		tf("invalid", file.FileDeleteResponse_NOT_FOUND)
		tf("012345678901234567890123", file.FileDeleteResponse_NOT_FOUND)
		tf(objID.Hex(), file.FileDeleteResponse_SUCCESS)
		tf(objID.Hex(), file.FileDeleteResponse_NOT_FOUND)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
