package main

import (
	"bytes"
	"context"
	db "jiaojiao/database"
	content "jiaojiao/srv/content/proto"
	"jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"

	"github.com/micro/go-micro/client"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type srv struct{}

/**
 * @api {rpc} /rpc Content.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Content.Create
 * @apiDescription create sell info content
 *
 * @apiParam {string} [contentId] 24 bytes content id, left empty for first upload
 * @apiParam {string} [contentToken] content token, left empty for first upload
 * @apiParam {bytes} content binary content
 * @apiParam {int32} type 1 for picture <br> 2 for video
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiSuccess {string} contentId 24 bytes contentId
 * @apiSuccess {string} contentToken random uuid content token
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *content.ContentCreateRequest, rsp *content.ContentCreateResponse) error {
	upload := func() (primitive.ObjectID, error) {
		srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
			func() interface{} { return mock.NewFileService() }).(file.FileService)
		rsp, err := srv.Create(context.TODO(), &file.FileCreateRequest{
			File: req.Content,
		})
		if utils.LogContinue(err, utils.Warning, "File service error: %v", err) || rsp.Status != file.FileCreateResponse_SUCCESS {
			return primitive.ObjectID{}, err
		}

		fid, err := primitive.ObjectIDFromHex(rsp.FileId)
		if utils.LogContinue(err, utils.Warning, "File service error: %v", err) {
			return primitive.ObjectID{}, err
		}

		return fid, nil
	}

	if bytes.Equal(req.Content, []byte{0}) || req.Type == 0 {
		rsp.Status = content.ContentCreateResponse_INVALID_PARAM
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
		rsp.Status = content.ContentCreateResponse_SUCCESS
	} else if req.ContentId != "" && req.ContentToken != "" {
		//check token
		collection := db.MongoDatabase.Collection("sellinfo")
		rid, err := primitive.ObjectIDFromHex(req.ContentId)
		if err != nil {
			rsp.Status = content.ContentCreateResponse_INVALID_TOKEN
			return nil
		}
		_, err = collection.FindOne(db.MongoContext, bson.D{
			{"_id", rid},
			{"token", req.ContentToken},
		}).DecodeBytes()
		if err != nil {
			rsp.Status = content.ContentCreateResponse_INVALID_TOKEN
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
		rsp.Status = content.ContentCreateResponse_SUCCESS
	} else {
		rsp.Status = content.ContentCreateResponse_INVALID_PARAM
	}
	return nil
}

/**
 * @api {rpc} /rpc Content.Delete
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Content.Delete
 * @apiDescription delete sell info content
 *
 * @apiParam {string} contentId 24 bytes content id
 * @apiParam {string} contentToken content token
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiUse DBServerDown
 */
func (a *srv) Delete(ctx context.Context, req *content.ContentDeleteRequest, rsp *content.ContentDeleteResponse) error {
	if req.ContentId == "" || req.ContentToken == "" {
		rsp.Status = content.ContentDeleteResponse_INVALID_PARAM
		return nil
	}
	type files struct {
		FileId primitive.ObjectID                `bson:"fileId"`
		Type   content.ContentCreateRequest_Type `bson:"type"`
	}
	type result struct {
		Id    primitive.ObjectID `bson:"_id"`
		Files []files            `bson:"files"`
	}

	collection := db.MongoDatabase.Collection("sellinfo")
	rid, err := primitive.ObjectIDFromHex(req.ContentId)
	if err != nil {
		rsp.Status = content.ContentDeleteResponse_INVALID_TOKEN
		return nil
	}
	var res result
	err = collection.FindOneAndDelete(db.MongoContext, bson.D{
		{"_id", rid},
		{"token", req.ContentToken},
	}).Decode(&res)
	if err != nil {
		rsp.Status = content.ContentDeleteResponse_INVALID_TOKEN
		return nil
	}

	srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
		func() interface{} { return mock.NewFileService() }).(file.FileService)
	for _, v := range res.Files {
		microRsp, err := srv.Delete(context.TODO(), &file.FileRequest{
			FileId: v.FileId.Hex(),
		})
		if utils.LogContinue(err, utils.Warning, "File service error: %v", err) || microRsp.Status != file.FileDeleteResponse_SUCCESS {
			return err
		}
	}
	rsp.Status = content.ContentDeleteResponse_SUCCESS
	return nil
}

func main() {
	db.InitMongoDB("sellinfomongo")
	service := utils.InitMicroService("content")
	utils.LogPanic(content.RegisterContentHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
