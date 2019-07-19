package main

import (
	"bytes"
	"context"
	db "jiaojiao/database"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"
	"time"

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
 * @apiSuccess {bytes} stream file stream
 * @apiUse DBServerDown
 */
func (a *srvFile) Query(ctx context.Context, req *file.FileQueryRequest, rsp *file.FileQueryResponse) error {
	if req.FileId == "" {
		rsp.Status = file.FileQueryResponse_INVALID_PARAM
	} else {
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		stream, err := bucket.OpenDownloadStream(req.FileId)

		if err != nil {
			rsp.Status = file.FileQueryResponse_NOT_FOUND
			return nil
		} else {
			err = stream.SetReadDeadline(time.Now().Add(time.Second * 30))
			if utils.LogContinue(err, utils.Warning) {
				return err
			}

			downloadBuffer := make([]byte, 0, 1024*1024)
			_, err := stream.Read(downloadBuffer)
			if utils.LogContinue(err, utils.Warning) {
				return err
			}

			rsp.Stream = downloadBuffer
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
 * @apiParam {bytes} stream file stream bytes
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success
 * @apiSuccess {string} fileId file id
 * @apiUse DBServerDown
 */
func (a *srvFile) Create(ctx context.Context, req *file.FileCreateRequest, rsp *file.FileCreateResponse) error {
	if bytes.Equal(req.Stream, []byte{0}) {
		rsp.Status = file.FileCreateResponse_INVALID_PARAM
	} else {
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		objId, err := bucket.UploadFromStream("", bytes.NewReader(req.Stream))
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		rsp.FileId = objId.Hex()
		rsp.Status = file.FileCreateResponse_SUCCESS
	}
	return nil
}

func main() {
	db.InitMongoDB("filemongo")

	service := utils.InitMicroService("file")
	utils.LogPanic(file.RegisterFileHandler(service.Server(), new(srvFile)))
	utils.RunMicroService(service)
}
