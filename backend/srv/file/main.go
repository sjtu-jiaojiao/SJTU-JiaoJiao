package main

import (
	"bytes"
	"context"
	db "jiaojiao/database"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type srvFile struct{}

/**
 * @api {rpc} /rpc File.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName File.Query
 * @apiDescription Query file stream
 *
 * @apiParam {string} fileID file id.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiSuccess {bytes} file file stream
 * @apiSuccess {int64} size file size
 * @apiUse DBServerDown
 */
func (a *srvFile) Query(ctx context.Context, req *file.FileRequest, rsp *file.FileQueryResponse) error {
	if !utils.RequireParam(req.FileID) {
		rsp.Status = file.FileQueryResponse_INVALID_PARAM
		return nil
	}

	bucket, err := gridfs.NewBucket(db.MongoDatabase)
	if utils.LogContinue(err, utils.Error) {
		return err
	}

	fid, err := primitive.ObjectIDFromHex(req.FileID)
	if utils.LogContinue(err, utils.Error) {
		return err
	}

	var buf bytes.Buffer
	size, err := bucket.DownloadToStream(fid, &buf)

	if err != nil {
		rsp.Status = file.FileQueryResponse_NOT_FOUND
		return nil
	}
	rsp.File = buf.Bytes()
	rsp.Size = size
	rsp.Status = file.FileQueryResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc File.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName File.Create
 * @apiDescription Create file
 *
 * @apiParam {bytes} file file stream bytes
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success
 * @apiSuccess {string} fileID file id
 * @apiUse DBServerDown
 */
func (a *srvFile) Create(ctx context.Context, req *file.FileCreateRequest, rsp *file.FileCreateResponse) error {
	if !utils.RequireParam(req.File) {
		rsp.Status = file.FileCreateResponse_INVALID_PARAM
		return nil
	}

	bucket, err := gridfs.NewBucket(db.MongoDatabase)
	if utils.LogContinue(err, utils.Error) {
		return err
	}

	objID, err := bucket.UploadFromStream("", bytes.NewReader(req.File))
	if utils.LogContinue(err, utils.Error) {
		return err
	}
	rsp.FileID = objID.Hex()
	rsp.Status = file.FileCreateResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc File.Delete
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName File.Delete
 * @apiDescription Delete file
 *
 * @apiParam {string} fileID file id.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiUse DBServerDown
 */
func (a *srvFile) Delete(ctx context.Context, req *file.FileRequest, rsp *file.FileDeleteResponse) error {
	// TODO
	rsp.Status = file.FileDeleteResponse_SUCCESS
	return nil
}

func main() {
	db.InitMongoDB("filemongo")
	service := utils.InitMicroService("file")
	utils.LogPanic(file.RegisterFileHandler(service.Server(), new(srvFile)))
	utils.RunMicroService(service)
}
