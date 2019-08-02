package main

import (
	"context"
	"errors"
	db "jiaojiao/database"
	content "jiaojiao/srv/content/proto"
	"jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/h2non/filetype"
	"github.com/micro/go-micro/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type srv struct{}

type files struct {
	FileID primitive.ObjectID `bson:"fileID"`
	Type   content.Type       `bson:"type"`
}
type result struct {
	ID    primitive.ObjectID `bson:"_id"`
	Token string             `bson:"token"`
	Files []files            `bson:"files"`
	Tags  []string           `bson:"tags"`
}

/**
 * @api {rpc} /rpc Content.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Content.Create
 * @apiDescription create sell info content
 *
 * @apiParam {string} [contentID] 24 bytes content id, left empty for first upload
 * @apiParam {string} [contentToken] content token, left empty for first upload
 * @apiParam {bytes} content binary bytes, file accept [image](https://github.com/h2non/filetype#image) and [video](https://github.com/h2non/filetype#video)
 * @apiParam {int32} type 1 for picture <br> 2 for video
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token <br> 3 for invalid type
 * @apiSuccess {string} contentID 24 bytes contentID
 * @apiSuccess {string} contentToken random uuid content token
 * @apiSuccess {string} fileID 24 bytes fileID
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *content.ContentCreateRequest, rsp *content.ContentCreateResponse) error {
	// upload file
	upload := func() (primitive.ObjectID, error) {
		srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
			func() interface{} { return mock.NewFileService() }).(file.FileService)
		rsp, err := srv.Create(context.TODO(), &file.FileCreateRequest{
			File: req.Content,
		})
		if err != nil {
			return primitive.ObjectID{}, err
		}
		if rsp.Status != file.FileCreateResponse_SUCCESS {
			_, s := utils.LogContinueS("File create return "+rsp.Status.String(), utils.Error)
			return primitive.ObjectID{}, errors.New(s)
		}

		fid, err := primitive.ObjectIDFromHex(rsp.FileID)
		if err != nil {
			return primitive.ObjectID{}, err
		}

		return fid, nil
	}

	// check param
	if !utils.RequireParam(req.Content, req.Type) {
		rsp.Status = content.ContentCreateResponse_INVALID_PARAM
		return nil
	}

	if !utils.CheckFile(req.Content, filetype.IsImage, filetype.IsVideo) {
		rsp.Status = content.ContentCreateResponse_INVALID_TYPE
		return nil
	}

	if utils.IsEmpty(req.ContentID) && utils.IsEmpty(req.ContentToken) { // create new
		objID, err := upload()
		if utils.LogContinue(err, utils.Error) {
			return err
		}

		token := uuid.NewV4().String()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		collection := db.MongoDatabase.Collection("content")
		res, err := collection.InsertOne(ctx, bson.M{
			"token": token,
			"files": bson.A{
				bson.M{
					"fileID": objID,
					"type":   req.Type,
				}},
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}

		rsp.ContentID = res.InsertedID.(primitive.ObjectID).Hex()
		rsp.ContentToken = token
		rsp.FileID = objID.Hex()
		rsp.Status = content.ContentCreateResponse_SUCCESS
	} else if !utils.IsEmpty(req.ContentID) && !utils.IsEmpty(req.ContentToken) { // add exist one
		if !validCheck(req.ContentID, req.ContentToken) {
			rsp.Status = content.ContentCreateResponse_INVALID_TOKEN
			return nil
		}

		objID, err := upload()
		if utils.LogContinue(err, utils.Error) {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		collection := db.MongoDatabase.Collection("content")
		rid, _ := primitive.ObjectIDFromHex(req.ContentID)
		_, err = collection.UpdateOne(ctx, bson.D{
			{"_id", rid},
			{"token", req.ContentToken},
		},
			bson.D{
				{"$push", bson.D{
					{"files", bson.D{
						{"fileID", objID},
						{"type", req.Type},
					}},
				}}})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		rsp.ContentID = req.ContentID
		rsp.ContentToken = req.ContentToken
		rsp.FileID = objID.Hex()
		rsp.Status = content.ContentCreateResponse_SUCCESS
	} else {
		rsp.Status = content.ContentCreateResponse_INVALID_PARAM
	}
	return nil
}

/**
 * @api {rpc} /rpc Content.CreateTag
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Content.CreateTag
 * @apiDescription create tags
 *
 * @apiParam {string} [contentID] 24 bytes content id, left empty for first upload
 * @apiParam {string} [contentToken] content token, left empty for first upload
 * @apiParam {list} tags {string} tag
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token <br> 3 for invalid type
 * @apiSuccess {string} contentID 24 bytes contentID
 * @apiSuccess {string} contentToken random uuid content token
 * @apiUse DBServerDown
 */
func (a *srv) CreateTag(ctx context.Context, req *content.ContentCreateTagRequest, rsp *content.ContentCreateTagResponse) error {
	// check param
	if !utils.RequireParam(req.Tags) {
		rsp.Status = content.ContentCreateTagResponse_INVALID_PARAM
		return nil
	}

	if utils.IsEmpty(req.ContentID) && utils.IsEmpty(req.ContentToken) { // create new
		token := uuid.NewV4().String()
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		collection := db.MongoDatabase.Collection("content")
		res, err := collection.InsertOne(ctx, bson.M{
			"token": token,
			"tags":  req.Tags,
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}

		rsp.ContentID = res.InsertedID.(primitive.ObjectID).Hex()
		rsp.ContentToken = token
		rsp.Status = content.ContentCreateTagResponse_SUCCESS
	} else if !utils.IsEmpty(req.ContentID) && !utils.IsEmpty(req.ContentToken) { // add exist one
		if !validCheck(req.ContentID, req.ContentToken) {
			rsp.Status = content.ContentCreateTagResponse_INVALID_TOKEN
			return nil
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		collection := db.MongoDatabase.Collection("content")
		rid, err := primitive.ObjectIDFromHex(req.ContentID)
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = content.ContentCreateTagResponse_INVALID_TOKEN
			return nil
		}
		_, err = collection.UpdateOne(ctx, bson.D{
			{"_id", rid},
			{"token", req.ContentToken},
		},
			bson.D{
				{"$set", bson.D{
					{"tags", req.Tags},
				}},
			})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		rsp.ContentID = req.ContentID
		rsp.ContentToken = req.ContentToken
		rsp.Status = content.ContentCreateTagResponse_SUCCESS
	} else {
		rsp.Status = content.ContentCreateTagResponse_INVALID_PARAM
	}
	return nil
}

/**
 * @api {rpc} /rpc Content.Update
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Content.Update
 * @apiDescription update sell info content
 *
 * @apiParam {string} contentID 24 bytes content id
 * @apiParam {string} contentToken content token
 * @apiParam {string} fileID 24 bytes file id
 * @apiParam {bytes} content binary bytes, file accept [image](https://github.com/h2non/filetype#image)
 *                             and [video](https://github.com/h2non/filetype#video)
 * @apiParam {int32} type 1 for picture <br> 2 for video
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token <br> 3 for not found <br> 4 for invalid type
 * @apiSuccess {string} fileID 24 bytes updated file id
 * @apiUse DBServerDown
 */
func (a *srv) Update(ctx context.Context, req *content.ContentUpdateRequest, rsp *content.ContentUpdateResponse) error {
	if !utils.RequireParam(req.ContentID, req.ContentToken, req.FileID, req.Content, req.Type) {
		rsp.Status = content.ContentUpdateResponse_INVALID_PARAM
		return nil
	}

	// check token
	if !validCheck(req.ContentID, req.ContentToken) {
		rsp.Status = content.ContentUpdateResponse_INVALID_TOKEN
		return nil
	}

	// check id
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("content")
	rid, _ := primitive.ObjectIDFromHex(req.ContentID)
	oldFid, err := primitive.ObjectIDFromHex(req.FileID)
	if utils.LogContinue(err, utils.Warning) {
		rsp.Status = content.ContentUpdateResponse_NOT_FOUND
		return nil
	}

	// add new file
	srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
		func() interface{} { return mock.NewFileService() }).(file.FileService)
	microCreateRsp, err := srv.Create(context.TODO(), &file.FileCreateRequest{
		File: req.Content,
	})
	if utils.LogContinue(err, utils.Error) {
		return err
	}
	if microCreateRsp.Status != file.FileCreateResponse_SUCCESS {
		_, s := utils.LogContinueS("File create return "+microCreateRsp.Status.String(), utils.Error)
		return errors.New(s)
	}
	fid, err := primitive.ObjectIDFromHex(microCreateRsp.FileID)
	if utils.LogContinue(err, utils.Error) {
		return err
	}

	_, err = collection.UpdateOne(ctx, bson.D{
		{"_id", rid},
		{"token", req.ContentToken},
		{"files.fileID", oldFid},
	}, bson.D{
		{"$set", bson.D{
			{"files.$.fileID", fid},
			{"files.$.type", req.Type},
		}},
	})
	if utils.LogContinue(err, utils.Error) {
		return err
	}

	// delete old file
	microDeleteRsp, err := srv.Delete(context.TODO(), &file.FileRequest{
		FileID: req.FileID,
	})
	if utils.LogContinue(err, utils.Error) {
		return err
	}
	if microDeleteRsp.Status != file.FileDeleteResponse_SUCCESS {
		_, s := utils.LogContinueS("File delete return "+microDeleteRsp.Status.String(), utils.Error)
		return errors.New(s)
	}

	rsp.FileID = microCreateRsp.FileID
	rsp.Status = content.ContentUpdateResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc Content.Delete
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Content.Delete
 * @apiDescription delete sell info content
 *
 * @apiParam {string} contentID 24 bytes content id
 * @apiParam {string} contentToken content token
 * @apiParam {string} [fileID] 24 bytes file id, if not empty only delete this file
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiUse DBServerDown
 */
func (a *srv) Delete(ctx context.Context, req *content.ContentDeleteRequest, rsp *content.ContentDeleteResponse) error {
	if !utils.RequireParam(req.ContentID, req.ContentToken) {
		rsp.Status = content.ContentDeleteResponse_INVALID_PARAM
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("content")
	rid, err := primitive.ObjectIDFromHex(req.ContentID)
	if utils.LogContinue(err, utils.Warning) {
		rsp.Status = content.ContentDeleteResponse_INVALID_TOKEN
		return nil
	}

	srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
		func() interface{} { return mock.NewFileService() }).(file.FileService)
	if utils.RequireParam(req.FileID) { // delete single file
		fid, err := primitive.ObjectIDFromHex(req.FileID)
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = content.ContentDeleteResponse_NOT_FOUND
			return nil
		}

		res, err := collection.UpdateOne(ctx, bson.D{
			{"_id", rid},
			{"token", req.ContentToken},
		}, bson.D{
			{"$pull", bson.D{
				{"files", bson.D{
					{"fileID", fid},
				}},
			}},
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		if res.ModifiedCount == 0 {
			rsp.Status = content.ContentDeleteResponse_NOT_FOUND
			return nil
		}

		microRsp, err := srv.Delete(context.TODO(), &file.FileRequest{
			FileID: req.FileID,
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		if microRsp.Status != file.FileDeleteResponse_SUCCESS {
			_, s := utils.LogContinueS("File delete return "+microRsp.Status.String(), utils.Error)
			return errors.New(s)
		}
	} else {
		var res result
		err = collection.FindOneAndDelete(ctx, bson.D{
			{"_id", rid},
			{"token", req.ContentToken},
		}).Decode(&res)
		if utils.LogContinue(err, utils.Warning) {
			rsp.Status = content.ContentDeleteResponse_INVALID_TOKEN
			return nil
		}

		srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
			func() interface{} { return mock.NewFileService() }).(file.FileService)
		for _, v := range res.Files {
			microRsp, err := srv.Delete(context.TODO(), &file.FileRequest{
				FileID: v.FileID.Hex(),
			})
			if utils.LogContinue(err, utils.Error) {
				return err
			}
			if microRsp.Status != file.FileDeleteResponse_SUCCESS {
				_, s := utils.LogContinueS("File delete return "+microRsp.Status.String(), utils.Error)
				return errors.New(s)
			}
		}
	}
	rsp.Status = content.ContentDeleteResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc Content.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Content.Query
 * @apiDescription query sell info content
 *
 * @apiParam {string} contentID 24 bytes content id
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for not found
 * @apiSuccess {string} contentToken content token
 * @apiSuccess {list} files {string} fileID : file id <br> {int32} type : file type 1 for picture, 2 for video
 * @apiSuccess {list} tags {string} tag : tag name
 * @apiUse DBServerDown
 */
func (a *srv) Query(ctx context.Context, req *content.ContentQueryRequest, rsp *content.ContentQueryResponse) error {
	if !utils.RequireParam(req.ContentID) {
		rsp.Status = content.ContentQueryResponse_INVALID_PARAM
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("content")
	rid, err := primitive.ObjectIDFromHex(req.ContentID)
	if utils.LogContinue(err, utils.Warning) {
		rsp.Status = content.ContentQueryResponse_INVALID_PARAM
		return nil
	}
	var res result
	err = collection.FindOne(ctx, bson.D{
		{"_id", rid},
	}).Decode(&res)
	if utils.LogContinue(err, utils.Warning) {
		rsp.Status = content.ContentQueryResponse_NOT_FOUND
		return nil
	}

	rsp.ContentToken = res.Token
	for _, v := range res.Files {
		rsp.Files = append(rsp.Files, &content.FileMsg{
			FileID: v.FileID.Hex(),
			Type:   v.Type,
		})
	}
	rsp.Tags = res.Tags
	rsp.Status = content.ContentQueryResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc Content.Check
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Content.Check
 * @apiDescription check sell info content
 *
 * @apiParam {string} contentID 24 bytes content id
 * @apiParam {string} contentToken content token
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for valid <br> 2 for invalid
 * @apiUse DBServerDown
 */
func (a *srv) Check(ctx context.Context, req *content.ContentCheckRequest, rsp *content.ContentCheckResponse) error {
	if !utils.RequireParam(req.ContentID, req.ContentToken) {
		rsp.Status = content.ContentCheckResponse_INVALID_PARAM
		return nil
	}
	if !validCheck(req.ContentID, req.ContentToken) {
		rsp.Status = content.ContentCheckResponse_INVALID
		return nil
	}

	rsp.Status = content.ContentCheckResponse_VALID
	return nil
}

func validCheck(contentID string, contentToken string) bool {
	if !utils.RequireParam(contentID, contentToken) {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MongoDatabase.Collection("content")
	rid, err := primitive.ObjectIDFromHex(contentID)
	if utils.LogContinue(err, utils.Warning) {
		return false
	}

	_, err = collection.FindOne(ctx, bson.D{
		{"_id", rid},
		{"token", contentToken},
	}).DecodeBytes()
	if utils.LogContinue(err, utils.Warning) {
		return false
	}
	return true
}

func main() {
	db.InitMongoDB("contentmongo")
	service := utils.InitMicroService("content")
	utils.LogPanic(content.RegisterContentHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
