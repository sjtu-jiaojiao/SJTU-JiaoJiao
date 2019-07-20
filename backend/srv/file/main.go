package main

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	db "jiaojiao/database"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type srvFile struct{}

/**
 * @api {rpc} /rpc file.File.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName file.File.Query
 * @apiDescription Query file stream
 *
 * @apiParam {string} fileId file id.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiSuccess {bytes} file file stream
 * @apiSuccess {int64} size file size
 * @apiUse DBServerDown
 */
func (a *srvFile) Query(ctx context.Context, req *file.FileRequest, rsp *file.FileQueryResponse) error {
	if req.FileId == "" {
		rsp.Status = file.FileQueryResponse_INVALID_PARAM
	} else {
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		fid, err := primitive.ObjectIDFromHex(req.FileId)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		var buf bytes.Buffer
		size, err := bucket.DownloadToStream(fid, &buf)

		if err != nil {
			rsp.Status = file.FileQueryResponse_NOT_FOUND
			return nil
		} else {
			rsp.File = buf.Bytes()
			rsp.Size = size
			rsp.Status = file.FileQueryResponse_SUCCESS
		}
	}
	return nil
}

/**
 * @api {rpc} /rpc file.File.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName file.File.Create
 * @apiDescription Create file
 *
 * @apiParam {bytes} file file stream bytes
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success
 * @apiSuccess {string} fileId file id
 * @apiUse DBServerDown
 */
func (a *srvFile) Create(ctx context.Context, req *file.FileCreateRequest, rsp *file.FileCreateResponse) error {
	if bytes.Equal(req.File, []byte{0}) {
		rsp.Status = file.FileCreateResponse_INVALID_PARAM
	} else {
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		objId, err := bucket.UploadFromStream("", bytes.NewReader(req.File))
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		rsp.FileId = objId.Hex()
		rsp.Status = file.FileCreateResponse_SUCCESS
	}
	return nil
}

/**
 * @api {rpc} /rpc file.File.Delete
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName file.File.Delete
 * @apiDescription Delete file
 *
 * @apiParam {string} fileId file id.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiUse DBServerDown
 */
func (a *srvFile) Delete(ctx context.Context, req *file.FileRequest, rsp *file.FileDeleteResponse) error {
	// TODO
	var s file.FileCreateResponse
	f, _ := os.Open("/home/imwxz/Documents/code/SJTU-JiaoJiao/backend/srv/file/1.png")
	z, _ := ioutil.ReadAll(f)
	_ = a.Create(context.TODO(), &file.FileCreateRequest{
		File: z,
	}, &s)
	return errors.New(s.FileId)
}

func main() {
	db.InitMongoDB("filemongo")
	service := utils.InitMicroService("file")
	utils.LogPanic(file.RegisterFileHandler(service.Server(), new(srvFile)))
	utils.RunMicroService(service)
}
