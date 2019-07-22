package main

import (
	"context"
	db "jiaojiao/database"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	"jiaojiao/utils"
	"time"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type srv struct{}

/**
 * @api {rpc} /rpc BuyInfo.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName BuyInfo.Query
 * @apiDescription Query buy info
 *
 * @apiParam {int32} buyInfoId buyInfo id.
 * @apiSuccess {int32} buyInfoId buyInfoId
 * @apiSuccess {int32} status 1 for buying <br> 2 for reserved <br> 3 for done <br> 4 for expired
 * @apiSuccess {int64} releaseTime buyInfo release time
 * @apiSuccess {int64} validTime buyInfo validate time
 * @apiSuccess {string} goodName good name
 * @apiSuccess {double} price good price
 * @apiSuccess {string} description good description
 * @apiSuccess {string} contentId multimedia data
 * @apiSuccess {int32} userId userId
 * @apiUse DBServerDown
 */
func (a *srv) Query(ctx context.Context, req *buyinfo.BuyInfoQueryRequest, rsp *buyinfo.BuyInfoMsg) error {
	if req.BuyInfoId == 0 {
		return nil
	}
	info := db.BuyInfo{
		ID: req.BuyInfoId,
	}
	err := db.Ormer.First(&info).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
		return err
	}
	good := db.Good{
		ID: info.GoodId,
	}
	err = db.Ormer.First(&good).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
		return err
	}

	rsp.BuyInfoId = info.ID
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
 * @api {rpc} /rpc BuyInfo.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName BuyInfo.Create
 * @apiDescription create buy info
 *
 * @apiParam {int32} userId buyinfo userid
 * @apiParam {int64} validTime valid timestamp
 * @apiParam {string} goodName good name
 * @apiParam {string} [description] description for good
 * @apiParam {double} [price] good price
 * @apiParam {string} [contentId] content id of good
 * @apiParam {string} [contentToken] content token
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiSuccess {int32} buyInfoId created buyInfoId
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *buyinfo.BuyInfoCreateRequest, rsp *buyinfo.BuyInfoCreateResponse) error {
	if req.ValidTime == 0 || req.GoodName == "" || req.UserId == 0 {
		rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
		return nil
	}

	good := db.Good{
		GoodName:    req.GoodName,
		Price:       req.Price,
		Description: req.Description,
	}
	info := db.BuyInfo{
		ReleaseTime: time.Now(),
		ValidTime:   time.Unix(req.ValidTime, 0),
		UserId:      req.UserId,
	}

	insert := func() (int32, error) {
		tx := db.Ormer.Begin()
		if utils.LogContinue(tx.Error, utils.Warning) {
			return 0, tx.Error
		}
		err := tx.Create(&good).Error
		if utils.LogContinue(err, utils.Warning) {
			tx.Rollback()
			return 0, err
		}
		info.GoodId = good.ID
		err = tx.Create(&info).Error
		if utils.LogContinue(err, utils.Warning) {
			tx.Rollback()
			return 0, err
		}

		err = tx.Commit().Error
		if utils.LogContinue(err, utils.Warning) {
			tx.Rollback()
			return 0, err
		}
		return info.ID, nil
	}

	if req.ContentId == "" && req.ContentToken == "" {
		id, err := insert()
		if err != nil || id == 0 {
			return nil
		}
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoId = id
	} else if req.ContentId != "" && req.ContentToken != "" {
		collection := db.MongoDatabase.Collection("buyinfo")
		rid, err := primitive.ObjectIDFromHex(req.ContentId)
		if err != nil {
			rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
			return nil
		}
		_, err = collection.FindOne(db.MongoContext, bson.D{
			{"_id", rid},
			{"token", req.ContentToken},
		}).DecodeBytes()
		if err != nil {
			rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_TOKEN
			return nil
		}

		good.ContentId = req.ContentId
		id, err := insert()
		if err != nil || id == 0 {
			return nil
		}
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoId = id
	} else {
		rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
	}
	return nil
}

/**
 * @api {rpc} /rpc BuyInfo.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName BuyInfo.Find
 * @apiDescription Find BuyInfo.
 *
 * @apiParam {int32} [userId] userId
 * @apiParam {uint32} limit=100 row limit
 * @apiParam {uint32} offset=0 row offset
 * @apiSuccess {list} buyInfo see [BuyInfo Service](#api-Service-buyinfo_BuyInfo_Query)
 * @apiUse DBServerDown
 */
func (a *srv) Find(ctx context.Context, req *buyinfo.BuyInfoFindRequest, rsp *buyinfo.BuyInfoFindResponse) error {
	if req.Limit == 0 {
		req.Limit = 100
	}

	var res []*db.BuyInfo
	tb := db.Ormer
	if req.UserId != 0 {
		tb = tb.Where("user_id = ?", req.UserId)
	}
	err := tb.Limit(req.Limit).Offset(req.Offset).Find(&res).Error
	if utils.LogContinue(err, utils.Warning) {
		return err
	}
	for i, v := range res {
		rsp.BuyInfo = append(rsp.BuyInfo, new(buyinfo.BuyInfoMsg))
		good := db.Good{
			ID: v.GoodId,
		}
		err = db.Ormer.First(&good).Error
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		parseBuyInfo(v, &good, rsp.BuyInfo[i])
	}
	return nil
}

func parseBuyInfo(s *db.BuyInfo, g *db.Good, d *buyinfo.BuyInfoMsg) {
	d.BuyInfoId = s.ID
	d.Status = s.Status
	d.ReleaseTime = s.ReleaseTime.Unix()
	d.ValidTime = s.ValidTime.Unix()
	d.GoodName = g.GoodName
	d.Price = g.Price
	d.Description = g.Description
	d.ContentId = g.ContentId
	d.UserId = s.UserId
}

func main() {
	db.InitORM("buyinfodb", new(db.BuyInfo), new(db.Good))
	defer db.CloseORM()
	service := utils.InitMicroService("buyInfo")
	utils.LogPanic(buyinfo.RegisterBuyInfoHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
