package main

import (
	"context"
	db "jiaojiao/database"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func ParseCreate(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp sellinfo.SellInfoCreateResponse
	ret := make(utils.StringMap)
	var tags []string
	for _, v := range input["tags"].([]interface{}) {
		tags = append(tags, utils.TestString(v))
	}

	ret["_error"] = s.Create(context.TODO(), &sellinfo.SellInfoCreateRequest{
		UserID:       utils.TestInt(input["userID"]),
		ValidTime:    utils.TestInt64(input["validTime"]),
		GoodName:     utils.TestString(input["goodName"]),
		Description:  utils.TestString(input["description"]),
		Price:        utils.TestFloat64(input["price"]),
		ContentID:    utils.TestString(input["contentID"]),
		ContentToken: utils.TestString(input["contentToken"]),
		Tags:         tags,
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["sellInfoID"] = float64(rsp.SellInfoID)
	return ret
}

func ParseQuery(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp sellinfo.SellInfoMsg
	ret := make(utils.StringMap)

	ret["_error"] = s.Query(context.TODO(), &sellinfo.SellInfoQueryRequest{
		SellInfoID: utils.TestInt(input["sellInfoID"]),
	}, &rsp)
	ret["status"] = int32(rsp.Status)
	ret["sellInfoID"] = rsp.SellInfoID
	ret["releaseTime"] = rsp.ReleaseTime
	ret["validTime"] = rsp.ValidTime
	ret["goodName"] = rsp.GoodName
	ret["price"] = rsp.Price
	ret["description"] = rsp.Description
	ret["contentID"] = rsp.ContentID
	ret["userID"] = rsp.UserID
	return ret
}

func ParseFind(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp sellinfo.SellInfoFindResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Find(context.TODO(), &sellinfo.SellInfoFindRequest{
		UserID:    utils.TestInt(input["userID"]),
		Status:    sellinfo.SellStatus(utils.EnumConvert(utils.TestInt(input["status"]), sellinfo.SellStatus_name)),
		GoodName:  utils.TestString(input["goodName"]),
		LowPrice:  utils.TestFloat64(input["lowPrice"]),
		HighPrice: utils.TestFloat64(input["highPrice"]),
		Limit:     utils.TestUint(input["limit"]),
		Offset:    utils.TestUint(input["offset"]),
	}, &rsp)
	ret["sellInfo"] = rsp.SellInfo
	return ret
}

func ParseUpdate(input utils.StringMap) utils.StringMap {
	var s srv
	var rsp sellinfo.SellInfoUpdateResponse
	ret := make(utils.StringMap)

	ret["_error"] = s.Update(context.TODO(), &sellinfo.SellInfoUpdateRequest{
		SellInfoID: utils.TestInt(input["sellInfoID"]),
		Status:     sellinfo.SellStatus(utils.EnumConvert(utils.TestInt(input["status"]), sellinfo.SellStatus_name)),
	}, &rsp)
	ret["status"] = rsp.Status
	return ret
}

func CheckFind(actual utils.StringMap, expect utils.StringMap) {
	rsp := actual["sellInfo"].([]*sellinfo.SellInfoMsg)
	data := expect["sellInfo"].([]interface{})
	So(len(rsp), ShouldEqual, len(data))
	for k, v := range rsp {
		t := data[k].(utils.StringMap)
		utils.TestCheck(v.SellInfoID, t["sellInfoID"])
		utils.TestCheck(v.SellInfoID, t["sellInfoID"])
		utils.TestCheck(v.Status, t["status"])
		utils.TestCheck(v.ReleaseTime, t["releaseTime"])
		utils.TestCheck(v.ValidTime, t["validTime"])
		utils.TestCheck(v.GoodName, t["goodName"])
		utils.TestCheck(v.Price, t["price"])
		utils.TestCheck(v.Description, t["description"])
		utils.TestCheck(v.ContentID, t["contentID"])
		utils.TestCheck(v.UserID, t["userID"])
	}
}

func VerifyData(verify utils.StringMap, output utils.StringMap) {
	info := db.SellInfo{ID: int32(utils.TransVar("sellInfoID", verify, output).(float64))}
	if exist, ok := verify["_exist"]; !ok || exist.(bool) {
		So(info.ReleaseTime, ShouldNotBeEmpty)
		So(db.Ormer.Find(&info).Error, ShouldBeNil)
		So(info.Status, ShouldEqual, verify["status"])
		So(info.UserID, ShouldEqual, verify["userID"])
		So(info.ValidTime.Unix(), ShouldEqual, verify["validTime"])
		So(info.GoodID, ShouldNotBeNil)

		good := db.Good{ID: info.GoodID}
		So(db.Ormer.Find(&good).Error, ShouldBeNil)
		So(good.GoodName, ShouldEqual, verify["goodName"])
		So(good.Description, ShouldEqual, verify["description"])
		So(good.Price, ShouldEqual, verify["price"])
		utils.TestCheck(good.ContentID, verify["contentID"])
	} else {
		So(db.Ormer.Find(&info).Error, ShouldNotBeNil)
	}
}

func InsertData(data utils.StringMap) {
	info := db.SellInfo{
		ID:          utils.TestInt(data["sellInfoID"]),
		Status:      utils.TestInt(data["status"]),
		ReleaseTime: time.Now(),
		ValidTime:   time.Unix(utils.TestInt64(data["validTime"]), 0),
		UserID:      utils.TestInt(data["userID"]),
		GoodID:      utils.TestInt(data["goodID"]),
	}
	good := db.Good{
		ID:          utils.TestInt(data["goodID"]),
		GoodName:    utils.TestString(data["goodName"]),
		Price:       data["price"].(float64),
		Description: utils.TestString(data["description"]),
		ContentID:   utils.TestString(data["contentID"]),
	}

	So(db.Ormer.Create(&good).Error, ShouldBeNil)
	So(db.Ormer.Create(&info).Error, ShouldBeNil)
}

func cleanup() {
	var infos []db.SellInfo
	db.Ormer.Find(&infos)
	for _, v := range infos {
		db.Ormer.Delete(&db.Good{ID: v.GoodID})
	}
	db.Ormer.Delete(db.SellInfo{})
}

func TestAll(t *testing.T) {
	cleanup()
	utils.Test(t, "test/test_create.json", nil, ParseCreate, VerifyData, nil)
	cleanup()
	utils.Test(t, "test/test_query.json", InsertData, ParseQuery, nil, nil)
	cleanup()
	utils.Test(t, "test/test_find.json", InsertData, ParseFind, nil, CheckFind)
	cleanup()
	utils.Test(t, "test/test_update.json", InsertData, ParseUpdate, VerifyData, nil)
	cleanup()
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
