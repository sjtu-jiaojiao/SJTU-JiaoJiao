package main

import (
	"context"
	db "jiaojiao/database"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	"jiaojiao/srv/content/mock"
	content "jiaojiao/srv/content/proto"
	"jiaojiao/utils"
	"time"

	"github.com/micro/go-micro/client"

	"github.com/jinzhu/gorm"
)

type srv struct{}

/**
 * @api {rpc} /rpc BuyInfo.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName BuyInfo.Query
 * @apiDescription Query buy info
 *
 * @apiParam {int32} buyInfoID buyInfo id.
 * @apiSuccess {int32} buyInfoID buyInfoID
 * @apiSuccess {int32} status 1 for buying <br> 2 for reserved <br> 3 for done <br> 4 for expired <br> 5 for closed
 * @apiSuccess {int64} releaseTime buyInfo release time
 * @apiSuccess {int64} validTime buyInfo validate time
 * @apiSuccess {string} goodName good name
 * @apiSuccess {double} price good price
 * @apiSuccess {string} description good description
 * @apiSuccess {string} contentID multimedia data
 * @apiSuccess {int32} userID userID
 * @apiUse DBServerDown
 */
func (a *srv) Query(ctx context.Context, req *buyinfo.BuyInfoQueryRequest, rsp *buyinfo.BuyInfoMsg) error {
	if !utils.RequireParam(req.BuyInfoID) {
		return nil
	}
	info := db.BuyInfo{
		ID: req.BuyInfoID,
	}
	err := db.Ormer.First(&info).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Error) {
		return err
	}
	good := db.Good{
		ID: info.GoodID,
	}
	err = db.Ormer.First(&good).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Error) {
		return err
	}

	rsp.BuyInfoID = info.ID
	rsp.Status = buyinfo.BuyStatus(utils.EnumConvert(info.Status, buyinfo.BuyStatus_name))
	rsp.ReleaseTime = info.ReleaseTime.Unix()
	rsp.ValidTime = info.ValidTime.Unix()
	rsp.GoodName = good.GoodName
	rsp.Price = good.Price
	rsp.Description = good.Description
	rsp.ContentID = good.ContentID
	rsp.UserID = info.UserID
	return nil
}

/**
 * @api {rpc} /rpc BuyInfo.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName BuyInfo.Create
 * @apiDescription Create buy info
 *
 * @apiParam {int32} userID user id
 * @apiParam {int64} validTime valid timestamp
 * @apiParam {string} goodName good name
 * @apiParam {string} [description] description for good
 * @apiParam {double} [price] good price
 * @apiParam {string} [contentID] content id of good
 * @apiParam {string} [contentToken] content token
 * @apiParam {list} [tags] {string} tag
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiSuccess {int32} buyInfoID created buyInfoID
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *buyinfo.BuyInfoCreateRequest, rsp *buyinfo.BuyInfoCreateResponse) error {
	if !utils.RequireParam(req.ValidTime, req.GoodName, req.UserID) {
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
		UserID:      req.UserID,
	}

	insert := func() (int32, error) {
		tx := db.Ormer.Begin()
		if tx.Error != nil {
			return 0, tx.Error
		}
		err := tx.Create(&good).Error
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		info.GoodID = good.ID
		err = tx.Create(&info).Error
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		return info.ID, nil
	}

	srv := utils.CallMicroService("content", func(name string, c client.Client) interface{} { return content.NewContentService(name, c) },
		func() interface{} { return mock.NewContentService() }).(content.ContentService)
	if utils.RequireParam(req.Tags) {
		microRsp, err := srv.CreateTag(context.TODO(), &content.ContentCreateTagRequest{
			ContentID:    req.ContentID,
			ContentToken: req.ContentToken,
			Tags:         req.Tags,
		})
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		if microRsp.Status == content.ContentCreateTagResponse_INVALID_PARAM {
			rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
			return nil
		}
		if microRsp.Status != content.ContentCreateTagResponse_SUCCESS {
			rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_TOKEN
			return nil
		}
		req.ContentID = microRsp.ContentID
		req.ContentToken = microRsp.ContentToken
	}

	if utils.IsEmpty(req.ContentID) && utils.IsEmpty(req.ContentToken) {
		id, err := insert()
		if utils.LogContinue(err, utils.Error) || id == 0 {
			return err
		}
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoID = id
	} else if !utils.IsEmpty(req.ContentID) && !utils.IsEmpty(req.ContentToken) {
		microRsp, err := srv.Check(context.TODO(), &content.ContentCheckRequest{
			ContentID:    req.ContentID,
			ContentToken: req.ContentToken,
		})
		if err != nil || microRsp.Status != content.ContentCheckResponse_VALID {
			rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_TOKEN
			return nil
		}

		good.ContentID = req.ContentID
		id, err := insert()
		if utils.LogContinue(err, utils.Error) || id == 0 {
			return err
		}
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoID = id
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
 * @apiParam {int32} [userID] userID
 * @apiParam {int32} [status] status 1 for waiting <br> 2 for reserved <br> 3 for done <br> 4 for expired <br> 5 for closed
 * @apiParam {string} [goodName] good name(fuzzy)
 * @apiParam {double} lowPrice=0 low bound of price, included
 * @apiParam {double} highPrice=inf high bound of price, included
 * @apiParam {uint32{0-100}} limit=100 row limit
 * @apiParam {uint32} offset=0 row offset
 * @apiSuccess {list} buyInfo see [BuyInfo Service](#api-Service-BuyInfo_Query)
 * @apiUse DBServerDown
 */
func (a *srv) Find(ctx context.Context, req *buyinfo.BuyInfoFindRequest, rsp *buyinfo.BuyInfoFindResponse) error {
	type result struct {
		BuyInfoID   int32
		Status      int32
		ReleaseTime time.Time
		ValidTime   time.Time
		GoodName    string
		Price       float64
		Description string
		ContentID   string
		UserID      int32
	}

	if req.Limit == 0 {
		req.Limit = 100
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.LowPrice < 0 {
		req.LowPrice = 0
	}
	if req.HighPrice < 0 {
		req.HighPrice = 0
	}

	var res []*result
	tb := db.Ormer.Table("buy_infos, goods").Select("buy_infos.id as buy_info_id, status, release_time, " +
		"valid_time, good_name, price, description, content_id, user_id").Where("buy_infos.good_id = goods.id")
	if req.UserID != 0 {
		tb = tb.Where("user_id = ?", req.UserID)
	}
	if req.Status != 0 {
		tb = tb.Where("status = ?", req.Status)
	}
	if req.GoodName != "" {
		tb = tb.Where("good_name LIKE ?", "%"+req.GoodName+"%")
	}
	if req.LowPrice != 0 {
		tb = tb.Where("price >= ?", req.LowPrice)
	}
	if req.HighPrice != 0 {
		tb = tb.Where("price <= ?", req.HighPrice)
	}

	err := tb.Limit(req.Limit).Offset(req.Offset).Find(&res).Error
	if utils.LogContinue(err, utils.Error) {
		return err
	}
	for _, v := range res {
		rsp.BuyInfo = append(rsp.BuyInfo, &buyinfo.BuyInfoMsg{
			BuyInfoID:   v.BuyInfoID,
			Status:      buyinfo.BuyStatus(utils.EnumConvert(v.Status, buyinfo.BuyStatus_name)),
			ReleaseTime: v.ReleaseTime.Unix(),
			ValidTime:   v.ValidTime.Unix(),
			GoodName:    v.GoodName,
			Price:       v.Price,
			Description: v.Description,
			ContentID:   v.ContentID,
			UserID:      v.UserID,
		})
	}
	return nil
}

/**
 * @api {rpc} /rpc BuyInfo.Update
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName BuyInfo.Update
 * @apiDescription Update buy info.
 *
 * @apiParam {int32} buyInfoID buyInfo id.
 * @apiParam {int32} status 1 for selling <br> 2 for reserved <br> 3 for done <br> 4 for expired <br> 5 for closed
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for buy info not found
 * @apiUse DBServerDown
 */
func (a *srv) Update(ctx context.Context, req *buyinfo.BuyInfoUpdateRequest, rsp *buyinfo.BuyInfoUpdateResponse) error {
	if !utils.RequireParam(req.BuyInfoID, req.Status) {
		rsp.Status = buyinfo.BuyInfoUpdateResponse_INVALID_PARAM
		return nil
	}

	info := db.BuyInfo{
		ID: req.BuyInfoID,
	}
	err := db.Ormer.First(&info).Error
	if err == nil {
		info.Status = int32(req.Status)
		err := db.Ormer.Save(&info).Error
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		rsp.Status = buyinfo.BuyInfoUpdateResponse_SUCCESS
	} else if gorm.IsRecordNotFoundError(err) {
		rsp.Status = buyinfo.BuyInfoUpdateResponse_NOT_FOUND
		return nil
	} else {
		utils.Error(err)
		return err
	}
	return nil
}

func main() {
	db.InitORM("buyinfodb", new(db.BuyInfo), new(db.Good))
	defer db.CloseORM()
	service := utils.InitMicroService("buyInfo")
	utils.LogPanic(buyinfo.RegisterBuyInfoHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
