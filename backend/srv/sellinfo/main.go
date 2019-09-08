package main

import (
	"context"
	db "jiaojiao/database"
	"jiaojiao/srv/content/mock"
	content "jiaojiao/srv/content/proto"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"
	"time"

	"github.com/micro/go-micro/client"

	"github.com/jinzhu/gorm"
)

type srv struct{}

/**
 * @api {rpc} /rpc SellInfo.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName SellInfo.Query
 * @apiDescription Query sell info
 *
 * @apiParam {int32} sellInfoID sellInfo id.
 * @apiSuccess {int32} sellInfoID sellInfoID
 * @apiSuccess {int32} status 1 for selling <br> 2 for reserved <br> 3 for done <br> 4 for expired <br> 5 for closed
 * @apiSuccess {int64} releaseTime sellInfo release time
 * @apiSuccess {int64} validTime sellInfo validate time
 * @apiSuccess {string} goodName good name
 * @apiSuccess {double} price good price
 * @apiSuccess {string} description good description
 * @apiSuccess {string} contentID multimedia data
 * @apiSuccess {int32} userID userID
 * @apiUse DBServerDown
 */
func (a *srv) Query(ctx context.Context, req *sellinfo.SellInfoQueryRequest, rsp *sellinfo.SellInfoMsg) error {
	if !utils.RequireParam(req.SellInfoID) {
		return nil
	}

	info := db.SellInfo{
		ID: req.SellInfoID,
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

	rsp.SellInfoID = info.ID
	rsp.Status = sellinfo.SellStatus(utils.EnumConvert(info.Status, sellinfo.SellStatus_name))
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
 * @api {rpc} /rpc SellInfo.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName SellInfo.Create
 * @apiDescription create sell info
 *
 * @apiParam {int32} userID sellinfo userid
 * @apiParam {int64} validTime valid timestamp
 * @apiParam {string} goodName good name
 * @apiParam {string} [description] description for good
 * @apiParam {double} [price] good price
 * @apiParam {string} [contentID] content id of good
 * @apiParam {string} [contentToken] content token
 * @apiParam {list} [tags] {string} tag
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiSuccess {int32} sellInfoID created sellInfoID
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *sellinfo.SellInfoCreateRequest, rsp *sellinfo.SellInfoCreateResponse) error {
	if !utils.RequireParam(req.ValidTime, req.GoodName, req.UserID) {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
		return nil
	}

	good := db.Good{
		GoodName:    req.GoodName,
		Price:       req.Price,
		Description: req.Description,
	}
	info := db.SellInfo{
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
			rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
			return nil
		}
		if microRsp.Status != content.ContentCreateTagResponse_SUCCESS {
			rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_TOKEN
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
		rsp.Status = sellinfo.SellInfoCreateResponse_SUCCESS
		rsp.SellInfoID = id
	} else if !utils.IsEmpty(req.ContentID) && !utils.IsEmpty(req.ContentToken) {
		microRsp, err := srv.Check(context.TODO(), &content.ContentCheckRequest{
			ContentID:    req.ContentID,
			ContentToken: req.ContentToken,
		})
		if err != nil || microRsp.Status != content.ContentCheckResponse_VALID {
			rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_TOKEN
			return nil
		}

		good.ContentID = req.ContentID
		id, err := insert()
		if utils.LogContinue(err, utils.Error) || id == 0 {
			return err
		}
		rsp.Status = sellinfo.SellInfoCreateResponse_SUCCESS
		rsp.SellInfoID = id
	} else {
		rsp.Status = sellinfo.SellInfoCreateResponse_INVALID_PARAM
	}
	return nil
}

/**
 * @api {rpc} /rpc SellInfo.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName SellInfo.Find
 * @apiDescription Find SellInfo.
 *
 * @apiParam {int32} [userID] userID
 * @apiParam {int32} [status] status 1 for selling <br> 2 for reserved <br> 3 for done <br> 4 for expired <br> 5 for closed
 * @apiParam {string} [goodName] good name(fuzzy)
 * @apiParam {double} lowPrice=0 low bound of price
 * @apiParam {double} highPrice=inf high bound of price
 * @apiParam {uint32{0-100}} limit=100 row limit
 * @apiParam {uint32} offset=0 row offset
 * @apiSuccess {list} sellInfo see [SellInfo Service](#api-Service-sellinfo_SellInfo_Query)
 * @apiUse DBServerDown
 */
func (a *srv) Find(ctx context.Context, req *sellinfo.SellInfoFindRequest, rsp *sellinfo.SellInfoFindResponse) error {
	type result struct {
		SellInfoID  int32
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
	tb := db.Ormer.Table("sell_infos, goods").Select("sell_infos.id as sell_info_id, status, release_time, " +
		"valid_time, good_name, price, description, content_id, user_id").Where("sell_infos.good_id = goods.id")
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
		rsp.SellInfo = append(rsp.SellInfo, &sellinfo.SellInfoMsg{
			SellInfoID:  v.SellInfoID,
			Status:      sellinfo.SellStatus(utils.EnumConvert(v.Status, sellinfo.SellStatus_name)),
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
 * @api {rpc} /rpc SellInfo.Update
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName SellInfo.Update
 * @apiDescription Update sell info.
 *
 * @apiParam {int32} sellInfoID sellInfo id.
 * @apiParam {int32} status 1 for selling <br> 2 for reserved <br> 3 for done <br> 4 for expired <br> 5 for closed
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for sell info not found
 * @apiUse DBServerDown
 */
func (a *srv) Update(ctx context.Context, req *sellinfo.SellInfoUpdateRequest, rsp *sellinfo.SellInfoUpdateResponse) error {
	if !utils.RequireParam(req.SellInfoID, req.Status) {
		rsp.Status = sellinfo.SellInfoUpdateResponse_INVALID_PARAM
		return nil
	}

	info := db.SellInfo{
		ID: req.SellInfoID,
	}
	err := db.Ormer.First(&info).Error
	if err == nil {
		info.Status = int32(req.Status)
		err := db.Ormer.Save(&info).Error
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		rsp.Status = sellinfo.SellInfoUpdateResponse_SUCCESS
	} else if gorm.IsRecordNotFoundError(err) {
		rsp.Status = sellinfo.SellInfoUpdateResponse_NOT_FOUND
		return nil
	} else {
		utils.Error(err)
		return err
	}
	return nil
}

func main() {
	db.InitORM("sellinfodb", new(db.SellInfo), new(db.Good))
	defer db.CloseORM()
	service := utils.InitMicroService("sellInfo")
	utils.LogPanic(sellinfo.RegisterSellInfoHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
