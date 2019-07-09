package main

import (
	"bytes"
	"context"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"jiaojiao/database"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"
)

type srvInfo struct{}

type srvContent struct{}

/**
 * @api {rpc} /rpc sellinfo.SellInfo.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Query
 * @apiDescription Query sell info
 *
 * @apiParam {int32} sellInfoId sell info id.
 * @apiSuccess {number} status -1 for empty param <br> 1 for success <br> 2 for non-exist
 * @apiUse DBServerDown
 */
//func (a *srv) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, rsp *sellinfo.SellInfoQueryResponse) error {
//
//	return nil
//}

/**
 * @api {rpc} /rpc sellinfo.SellInfo.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Create
 * @apiDescription create sell info
 *
 * @apiParam {int32} sellInfoId sell info id
 * @apiSuccess {number} status -1 for empty param <br> 1 for success <br> 2 for non-exist
 * @apiUse DBServerDown
 */
func (a *srvInfo) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, rsp *sellinfo.SellInfoCreateResponse) error {
	return nil
}

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
 * @apiSuccess {number} status -1 for empty param <br> 1 for success <br> 2 for invalid token
 * @apiUse DBServerDown
 */
func (a *srvContent) Create(ctx context.Context, req *sellinfo.ContentCreateRequest, rsp *sellinfo.ContentCreateResponse) error {
	if req.Content == nil || req.Type == 0 {
		rsp.Status = sellinfo.ContentCreateResponse_INVALID_PARAM
	} else if req.ContentId == "" && req.ContentToken == "" {
		//create new info
		bucket, err := gridfs.NewBucket(db.MongoDatabase)
		if err != nil {
			utils.Error("Failed to create bucket: %v", err)
		}

		objId, err := bucket.UploadFromStream("filename", bytes.NewReader(req.Content))
		if err != nil {
			utils.Error("Failed to upload stream: %v", err)
		}

		token := uuid.NewV4().String()
		collection := db.MongoDatabase.Collection("release")
		res, err := collection.InsertOne(db.MongoContext, bson.M{
			"token": token,
			"files": bson.A{
				bson.M{
					"fileId": objId.String(),
					"type":   req.Type.String(),
				}},
		})
		if err != nil {
			utils.Error("Failed to create collection: %v", err)
		}

		rsp.ContentId = res.InsertedID.(primitive.ObjectID).String()
		rsp.ContentToken = token
		rsp.Status = sellinfo.ContentCreateResponse_SUCCESS
	} else if req.ContentId != "" && req.ContentToken != "" {
		//check token
		collection := db.MongoDatabase.Collection("release")
		cur, err := collection.Find(db.MongoContext, bson.D{
			{"_id", req.ContentId},
			{"token", req.ContentToken},
		})
		if err != nil {
			utils.Error("Failed to find: %v", err)
		}

		if !cur.Next(db.MongoContext) {
			rsp.Status = sellinfo.ContentCreateResponse_INVALID_TOKEN
		} else {
			bucket, err := gridfs.NewBucket(db.MongoDatabase)
			if err != nil {
				utils.Error("Failed to create bucket: %v", err)
			}

			objId, err := bucket.UploadFromStream("filename", bytes.NewReader(req.Content))
			if err != nil {
				utils.Error("Failed to upload stream: %v", err)
			}

			res, err := collection.UpdateOne(db.MongoContext, bson.D{
				{"_id", req.ContentId},
				{"token", req.ContentToken},
			},
				bson.D{
					{"$push", bson.D{
						{"files", bson.D{
							{"fileId", objId.String()},
							{"type", req.Type.String()},
						}},
					}},
				})
			if err != nil {
				utils.Error("Failed to update: %v", err)
			}
			rsp.ContentId = res.UpsertedID.(string)
			rsp.ContentToken = req.ContentToken
			rsp.Status = sellinfo.ContentCreateResponse_SUCCESS
		}
	} else {
		rsp.Status = sellinfo.ContentCreateResponse_INVALID_PARAM
	}
	return nil
}

func main() {
	service := utils.InitMicroService("sellinfo")
	utils.LogPanic(sellinfo.RegisterSellInfoHandler(service.Server(), new(srvInfo)))
	utils.LogPanic(sellinfo.RegisterContentHandler(service.Server(), new(srvContent)))
	utils.RunMicroService(service)
}
