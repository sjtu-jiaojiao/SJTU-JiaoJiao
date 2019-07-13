package main

import (
	"bytes"
	"context"
	db "jiaojiao/database"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type srvInfo struct{}

type srvContent struct{}

/**
 * @apiIgnore Not finished Method
 * @api {rpc} /rpc sellinfo.SellInfo.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Query
 * @apiDescription Query sell info
 *
 * @apiParam {int32} sellInfoId sell info id.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for non-exist
 * @apiUse DBServerDown
 */
//func (a *srv) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, rsp *sellinfo.SellInfoQueryResponse) error {
//
//	return nil
//}

/**
 * @apiIgnore Not finished Method
 * @api {rpc} /rpc sellinfo.SellInfo.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Create
 * @apiDescription create sell info
 *
 * @apiParam {int32} sellInfoId sell info id
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for non-exist
 * @apiUse DBServerDown
 */
//func (a *srvInfo) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, rsp *sellinfo.SellInfoCreateResponse) error {
//	return nil
//}

/**
 * @api {rpc} /rpc sellinfo.Content.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.Content.Create
 * @apiDescription create sell info content
 *
 * @apiParam {string} [contentId] content id, left empty for first upload
 * @apiParam {string} [contentToken] content token, left empty for first upload
 * @apiParam {bytes} content binary content
 * @apiParam {int32} type 1 for picture <br> 2 for video
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiUse DBServerDown
 */
func (a *srvContent) Create(ctx context.Context, req *sellinfo.ContentCreateRequest, rsp *sellinfo.ContentCreateResponse) error {
	upload := func() (primitive.ObjectID, error) {
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		if utils.LogContinue(err, utils.Warning) {
			return primitive.ObjectID{}, err
		}

		objId, err := bucket.UploadFromStream("", bytes.NewReader(req.Content))
		if utils.LogContinue(err, utils.Warning) {
			return primitive.ObjectID{}, err
		}
		return objId, nil
	}

	if req.Content == nil || req.Type == 0 {
		rsp.Status = sellinfo.ContentCreateResponse_INVALID_PARAM
	} else if req.ContentId == "" && req.ContentToken == "" {
		objId, err := upload()
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		token := uuid.NewV4().String()
		collection := db.MongoDatabase.Collection("sellinfo")
		res, err := collection.InsertOne(db.MongoContext, bson.M{
			"token": token,
			"files": bson.A{
				bson.M{
					"fileId": objId,
					"type":   req.Type.String(),
				}},
		})
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		rsp.ContentId = res.InsertedID.(primitive.ObjectID).Hex()
		rsp.ContentToken = token
		rsp.Status = sellinfo.ContentCreateResponse_SUCCESS
	} else if req.ContentId != "" && req.ContentToken != "" {
		//check token
		collection := db.MongoDatabase.Collection("sellinfo")
		rid, err := primitive.ObjectIDFromHex(req.ContentId)
		if err != nil {
			rsp.Status = sellinfo.ContentCreateResponse_INVALID_TOKEN
			return nil
		}
		_, err = collection.FindOne(db.MongoContext, bson.D{
			{"_id", rid},
			{"token", req.ContentToken},
		}).DecodeBytes()
		if err != nil {
			rsp.Status = sellinfo.ContentCreateResponse_INVALID_TOKEN
			return nil
		}

		objId, err := upload()
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		_, err = collection.UpdateOne(db.MongoContext, bson.D{
			{"_id", rid},
			{"token", req.ContentToken},
		},
			bson.D{
				{"$push", bson.D{
					{"files", bson.D{
						{"fileId", objId},
						{"type", req.Type.String()},
					}},
				}},
			})
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		rsp.ContentId = req.ContentId
		rsp.ContentToken = req.ContentToken
		rsp.Status = sellinfo.ContentCreateResponse_SUCCESS
	} else {
		rsp.Status = sellinfo.ContentCreateResponse_INVALID_PARAM
	}
	return nil
}

func main() {
	db.InitMongoDB("sellinfomongo")
	service := utils.InitMicroService("sellinfo")
	//utils.LogPanic(sellinfo.RegisterSellInfoHandler(service.Server(), new(srvInfo)))
	utils.LogPanic(sellinfo.RegisterContentHandler(service.Server(), new(srvContent)))
	utils.RunMicroService(service)
}
