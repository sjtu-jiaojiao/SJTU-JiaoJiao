package main

import (
	"bytes"
	"context"
	db "jiaojiao/database"
	files "jiaojiao/srv/file/proto"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/gridfs"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuery(t *testing.T) {
	tf := func(fileID string, file []byte, size int64, status files.FileQueryResponse_Status) {
		var s srvFile
		var rsp files.FileQueryResponse
		So(s.Query(context.TODO(), &files.FileRequest{
			FileID: fileID,
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if status == files.FileQueryResponse_SUCCESS {
			So(rsp.File, ShouldResemble, file)
			So(rsp.Size, ShouldEqual, size)
		}
	}

	Convey("Test File Query", t, func() {
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		So(err, ShouldBeNil)
		objID, err := bucket.UploadFromStream("", bytes.NewReader([]byte{1, 2, 3, 4, 5, 6}))
		So(err, ShouldBeNil)

		tf("", []byte{0}, 0, files.FileQueryResponse_INVALID_PARAM)
		tf("1234567890", []byte{0}, 0, files.FileQueryResponse_INVALID_PARAM)
		tf("1234567890abcdef12345678", []byte{0}, 0, files.FileQueryResponse_NOT_FOUND)
		tf(objID.Hex(), []byte{1, 2, 3, 4, 5, 6}, 6, files.FileQueryResponse_SUCCESS)
	})
}

func TestCreate(t *testing.T) {
	tf := func(file []byte, status files.FileCreateResponse_Status) string {
		var s srvFile
		var rsp files.FileCreateResponse
		So(s.Create(context.TODO(), &files.FileCreateRequest{
			File: file,
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		return rsp.FileID
	}

	Convey("Test File Create", t, func() {
		fileID1 := tf([]byte{0}, files.FileCreateResponse_INVALID_PARAM)
		So(fileID1, ShouldBeEmpty)

		fileID2 := tf([]byte{0, 1, 2, 3, 4, 5}, files.FileCreateResponse_SUCCESS)
		So(len(fileID2), ShouldEqual, 24)
	})
}

func TestDelete(t *testing.T) {
	// TODO
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
