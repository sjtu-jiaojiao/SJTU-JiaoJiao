package main

import (
	"bytes"
	"context"
	db "jiaojiao/database"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"
	"time"

	"github.com/astaxie/beego/orm"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type srvInfo struct{}
type srvContent struct{}

var o orm.Ormer

/**
 * @api {rpc} /rpc sellinfo.SellInfo.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Query
 * @apiDescription Query sell info
 *
 * @apiParam {int32} sellInfoId sellInfo id.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for non-exist
 * @apiSuccess {int32} sellInfoId sellInfoId
 * @apiSuccess {int64} validTime sellInfo validate time
 * @apiSuccess {string} goodName good name
 * @apiSuccess {string} description good description
 * @apiSuccess {string} contentId multimedia data
 * @apiSuccess {array} tag good tags (un-finished)
 * @apiUse DBServerDown
 */
func (a *srvInfo) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, rsp *sellinfo.SellInfoQueryResponse) error {
	if req.SellInfoId == 0 {
		rsp.Status = sellinfo.SellInfoQueryResponse_EMPTY_PARAM
		return nil
	}
	info := db.SellInfo{
		Id: int(req.SellInfoId),
	}
	err := o.Read(&info)
	if err == orm.ErrNoRows {
		rsp.Status = sellinfo.SellInfoQueryResponse_NOT_EXIST
		return nil
	} else if err != nil {
		return err
	}
	good := db.Good{
		Id: info.Good.Id,
	}
	err = o.Read(&good)
	if err == orm.ErrNoRows {
		rsp.Status = sellinfo.SellInfoQueryResponse_NOT_EXIST
		return nil
	} else if err != nil {
		return err
	}

	rsp.SellInfoId = int32(info.Id)
	rsp.ValidTime = info.ValidDate.Unix()
	rsp.GoodName = good.GoodName
	rsp.Description = good.Description
	rsp.ContentId = good.ContentId
	rsp.Status = sellinfo.SellInfoQueryResponse_SUCCESS
	return nil
}

/**
 * @api {rpc} /rpc sellinfo.SellInfo.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Create
 * @apiDescription create sell info
 *
 * @apiParam {int64} validTime valid timestamp
 * @apiParam {string} goodName good name
 * @apiParam {string} description description for good
 * @apiParam {string} [contentId] content id of good
 * @apiParam {array} [tag] tags for good(un-finished)
 * @apiParam {string} [contentToken] content token
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiSuccess {int32} sellInfoId created sellInfoId
 * @apiUse DBServerDown
 */
func (a *srvInfo) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, rsp *sellinfo.SellInfoCreateResponse) error {
	good := db.Good{
		GoodName:    req.GoodName,
		Description: req.Description,
	}
	info := db.SellInfo{
		Status:      1,
		ReleaseTime: time.Now(),
		ValidDate:   time.Unix(req.ValidTime, 0),
		Good:        &good,
	}

	insert := func() (int32, error) {
		err := o.Begin()
		_, err1 := o.Insert(&good)
		id, err2 := o.Insert(&info)
		if err != nil || err1 != nil || err2 != nil {
			err = o.Rollback()
			if err != nil {
				utils.LogContinue(err, utils.Warning)
			}
			return 0, err
		}

		err = o.Commit()
		if err != nil {
			utils.LogContinue(err, utils.Warning)
			return 0, err
		}
		return int32(id), nil
	}

	if req.ValidTime == 0 || req.GoodName == "" {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
	} else if req.ContentId == "" && req.ContentToken == "" {
		id, err := insert()
		if err != nil || id == 0 {
			return nil
		}
		rsp.Status = sellinfo.SellInfoCreateResponse_SUCCESS
		rsp.SellInfoId = int32(id)
	} else if req.ContentId != "" && req.ContentToken != "" {
		collection := db.MongoDatabase.Collection("sellinfo")
		rid, err := primitive.ObjectIDFromHex(req.ContentId)
		if err != nil {
			rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
			return nil
		}
		_, err = collection.FindOne(db.MongoContext, bson.D{
			{"_id", rid},
			{"token", req.ContentToken},
		}).DecodeBytes()
		if err != nil {
			rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_TOKEN
			return nil
		}

		good.ContentId = req.ContentId
		id, err := insert()
		if err != nil || id == 0 {
			return nil
		}
		rsp.Status = sellinfo.SellInfoCreateResponse_SUCCESS
		rsp.SellInfoId = int32(id)
	} else {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
	}
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
	o = db.InitORM(new(db.SellInfo), new(db.Good))
	service := utils.InitMicroService("sellinfo")
	utils.LogPanic(sellinfo.RegisterSellInfoHandler(service.Server(), new(srvInfo)))
	utils.LogPanic(sellinfo.RegisterContentHandler(service.Server(), new(srvContent)))
	utils.RunMicroService(service)
}
