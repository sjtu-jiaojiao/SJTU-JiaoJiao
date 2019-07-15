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

/**
 * @api {rpc} /rpc sellinfo.SellInfo.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Query
 * @apiDescription Query sell info
 *
 * @apiParam {int32} sellInfoId sellInfo id.
 * @apiSuccess {int32} sellInfoId sellInfoId
 * @apiSuccess {int32} sellInfoState 1 for selling <br> 2 for reserved <br> 3 for done <br> 4 for expired
 * @apiSuccess {int64} releaseTime sellInfo release time
 * @apiSuccess {int64} validTime sellInfo validate time
 * @apiSuccess {string} goodName good name
 * @apiSuccess {double} price good price
 * @apiSuccess {string} description good description
 * @apiSuccess {string} contentId multimedia data
 * @apiSuccess {int32} userId userId
 * @apiUse DBServerDown
 */
func (a *srvInfo) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, rsp *sellinfo.SellInfoMsg) error {
	if req.SellInfoId == 0 {
		return nil
	}
	info := db.SellInfo{
		Id: req.SellInfoId,
	}
	err := db.Ormer.Read(&info)
	if err == orm.ErrNoRows {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
		return err
	}
	good := db.Good{
		Id: info.GoodId,
	}
	err = db.Ormer.Read(&good)
	if err == orm.ErrNoRows {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
		return err
	}

	rsp.SellInfoId = info.Id
	rsp.Status = info.Status
	rsp.ReleaseTime = info.ReleaseTime.Unix()
	rsp.ValidTime = info.ValidTime.Unix()
	rsp.GoodName = good.GoodName
	rsp.Price = good.Price
	rsp.Description = good.Description
	rsp.ContentId = good.ContentId
	rsp.UserId = info.UserId
	return nil
}

/**
 * @api {rpc} /rpc sellinfo.SellInfo.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Create
 * @apiDescription create sell info
 *
 * @apiParam {int32} userId sellinfo userid
 * @apiParam {int64} validTime valid timestamp
 * @apiParam {string} goodName good name
 * @apiParam {string} [description] description for good
 * @apiParam {double} [price] good price
 * @apiParam {string} [contentId] content id of good
 * @apiParam {string} [contentToken] content token
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiSuccess {int32} sellInfoId created sellInfoId
 * @apiUse DBServerDown
 */
func (a *srvInfo) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, rsp *sellinfo.SellInfoCreateResponse) error {
	if req.ValidTime == 0 || req.GoodName == "" || req.UserId == 0 {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
		return nil
	}

	good := db.Good{
		GoodName:    req.GoodName,
		Price:       req.Price,
		Description: req.Description,
	}
	info := db.SellInfo{
		Status:    1,
		ValidTime: time.Unix(req.ValidTime, 0),
		UserId:    req.UserId,
	}

	insert := func() (int32, error) {
		err := db.Ormer.Begin()
		if utils.LogContinue(err, utils.Warning) {
			return 0, err
		}
		id, err := db.Ormer.Insert(&good)
		if id == 0 || utils.LogContinue(err, utils.Warning) {
			utils.LogContinue(db.Ormer.Rollback(), utils.Warning)
			return 0, err
		}
		info.GoodId = int32(id)
		id, err = db.Ormer.Insert(&info)
		if id == 0 || utils.LogContinue(err, utils.Warning) {
			utils.LogContinue(db.Ormer.Rollback(), utils.Warning)
			return 0, err
		}

		err = db.Ormer.Commit()
		if utils.LogContinue(err, utils.Warning) {
			utils.LogContinue(db.Ormer.Rollback(), utils.Warning)
			return 0, err
		}
		return int32(id), nil
	}

	if req.ContentId == "" && req.ContentToken == "" {
		id, err := insert()
		if err != nil || id == 0 {
			return nil
		}
		rsp.Status = sellinfo.SellInfoCreateResponse_SUCCESS
		rsp.SellInfoId = id
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
		rsp.SellInfoId = id
	} else {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
	}
	return nil
}

/**
 * @api {rpc} /rpc sellinfo.SellInfo.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.SellInfo.Find
 * @apiDescription Find SellInfo.
 *
 * @apiParam {int32} [userId] userId
 * @apiParam {uint32} limit=100 row limit
 * @apiParam {uint32} offset=0 row offset
 * @apiSuccess {list} sellInfo see [SellInfo Service](#api-Service-sellinfo_SellInfo_Query)
 * @apiUse DBServerDown
 */
func (a *srvInfo) Find(ctx context.Context, req *sellinfo.SellInfoFindRequest, rsp *sellinfo.SellInfoFindResponse) error {
	if req.Limit == 0 {
		req.Limit = 100
	}

	var res []*db.SellInfo
	tb := db.Ormer.QueryTable(&db.SellInfo{})
	if req.UserId != 0 {
		tb = tb.Filter("userId", req.UserId)
	}
	_, err := tb.Limit(req.Limit, req.Offset).All(&res)
	if utils.LogContinue(err, utils.Warning) {
		return err
	}
	for i, v := range res {
		rsp.SellInfo = append(rsp.SellInfo, new(sellinfo.SellInfoMsg))
		good := db.Good{
			Id: v.GoodId,
		}
		err = db.Ormer.Read(&good)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		parseSellInfo(v, &good, rsp.SellInfo[i])
	}
	return nil
}

func parseSellInfo(s *db.SellInfo, g *db.Good, d *sellinfo.SellInfoMsg) {
	d.SellInfoId = s.Id
	d.Status = s.Status
	d.ReleaseTime = s.ReleaseTime.Unix()
	d.ValidTime = s.ValidTime.Unix()
	d.GoodName = g.GoodName
	d.Price = g.Price
	d.Description = g.Description
	d.ContentId = g.ContentId
	d.UserId = s.UserId
}

/**
 * @api {rpc} /rpc sellinfo.Content.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName sellinfo.Content.Create
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
	db.InitORM("sellinfodb", new(db.SellInfo), new(db.Good))
	db.InitMongoDB("sellinfomongo")
	service := utils.InitMicroService("sellInfo")
	utils.LogPanic(sellinfo.RegisterSellInfoHandler(service.Server(), new(srvInfo)))
	utils.LogPanic(sellinfo.RegisterContentHandler(service.Server(), new(srvContent)))
	utils.RunMicroService(service)
}
