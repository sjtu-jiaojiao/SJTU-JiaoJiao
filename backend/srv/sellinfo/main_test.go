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

//
//func TestSrvInfoQuery(t *testing.T) {
//	var s srv
//	var req sellinfo.SellInfoQueryRequest
//
//	info := db.SellInfo{
//		ID:          1100,
//		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
//		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
//		UserID:      1000,
//		GoodID:      1110,
//	}
//	good := db.Good{
//		ID:          1110,
//		GoodName:    "good",
//		Description: "Very good!",
//		ContentID:   "012345678901234567890123",
//	}
//	tf := func(sellID int, contentID string, userID int) {
//		var rsp sellinfo.SellInfoMsg
//		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
//		So(rsp.SellInfoID, ShouldEqual, sellID)
//		So(rsp.ContentID, ShouldEqual, contentID)
//		So(rsp.UserID, ShouldEqual, userID)
//	}
//	Convey("Test SellInfo Query", t, func() {
//		tf(0, "", 0)
//
//		So(db.Ormer.Create(&good).Error, ShouldBeNil)
//		So(db.Ormer.Create(&info).Error, ShouldBeNil)
//		defer func() {
//			So(db.Ormer.Delete(&db.Good{ID: 1110}).Error, ShouldBeNil)
//			So(db.Ormer.Delete(&db.SellInfo{ID: 1100}).Error, ShouldBeNil)
//		}()
//
//		req.SellInfoID = 1100
//		tf(1100, "012345678901234567890123", 1000)
//
//		req.SellInfoID = 1101
//		tf(0, "", 0)
//	})
//}
//
//func TestSrvInfoCreate(t *testing.T) {
//	tf := func(userID int32, validTime int64, goodName string, contentID string, contentToken string, tags []string, status sellinfo.SellInfoCreateResponse_Status) {
//		var s srv
//		var rsp sellinfo.SellInfoCreateResponse
//		So(s.Create(context.TODO(), &sellinfo.SellInfoCreateRequest{
//			ValidTime:    validTime,
//			GoodName:     "123459",
//			Price:        0,
//			Description:  "",
//			ContentID:    contentID,
//			ContentToken: contentToken,
//			UserID:       userID,
//			Tags:         tags,
//		}, &rsp), ShouldBeNil)
//		So(rsp.Status, ShouldEqual, status)
//
//		if rsp.Status == sellinfo.SellInfoCreateResponse_SUCCESS {
//			info := db.SellInfo{ID: rsp.SellInfoID}
//			So(db.Ormer.Find(&info).Error, ShouldBeNil)
//			So(info.UserID, ShouldEqual, userID)
//			So(info.ValidTime.Unix(), ShouldEqual, validTime)
//			So(info.GoodID, ShouldNotBeNil)
//
//			good := db.Good{ID: info.GoodID}
//			So(db.Ormer.Find(&good).Error, ShouldBeNil)
//			if len(tags) == 0 {
//				So(good.ContentID, ShouldEqual, contentID)
//			} else {
//				So(len(good.ContentID), ShouldEqual, 24)
//			}
//
//			So(db.Ormer.Delete(&good).RowsAffected, ShouldEqual, 1)
//			So(db.Ormer.Delete(&info).RowsAffected, ShouldEqual, 1)
//		}
//	}
//
//	Convey("Test SellInfo Create", t, func() {
//		tf(1001, 0, "", "1234567890abcdef12345678", "abc", []string{"tag1", "tag2", "tag3"}, sellinfo.SellInfoCreateResponse_INVALID_PARAM)
//		tf(1001, 100000000, "test good", "012345678901234567890123", "invalid_token", []string{"tag1", "tag2", "tag3"}, sellinfo.SellInfoCreateResponse_INVALID_TOKEN)
//		tf(1001, 100000000, "test good", "", "", []string{"tag1", "tag2", "tag3"}, sellinfo.SellInfoCreateResponse_INVALID_TOKEN)
//		tf(1001, 100000000, "test good", "", "", []string{}, sellinfo.SellInfoCreateResponse_SUCCESS)
//		tf(1001, 100000000, "test good", "", "valid_token", []string{"tag1", "tag2", "tag3"}, sellinfo.SellInfoCreateResponse_INVALID_PARAM)
//		tf(1001, 100000000, "test good", "012345678901234567890123", "", []string{}, sellinfo.SellInfoCreateResponse_INVALID_PARAM)
//		tf(1001, 100000000, "test good", "012345678901234567890123", "valid_token", []string{"tag1", "tag2", "tag3"}, sellinfo.SellInfoCreateResponse_SUCCESS)
//	})
//}
//
//func TestSrvInfoFind(t *testing.T) {
//	var s srv
//	var req sellinfo.SellInfoFindRequest
//
//	info1 := db.SellInfo{
//		ID:          1000,
//		Status:      1,
//		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
//		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
//		UserID:      1000,
//		GoodID:      1010,
//	}
//	info2 := db.SellInfo{
//		ID:          1001,
//		Status:      2,
//		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
//		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
//		UserID:      1000,
//		GoodID:      1011,
//	}
//	info3 := db.SellInfo{
//		ID:          1002,
//		Status:      3,
//		ReleaseTime: time.Date(2019, 9, 9, 9, 9, 9, 0, time.Local),
//		ValidTime:   time.Date(2020, 9, 9, 9, 9, 9, 0, time.Local),
//		UserID:      1001,
//		GoodID:      1012,
//	}
//	good1 := db.Good{
//		ID:          1010,
//		GoodName:    "good",
//		Description: "Very good!",
//		ContentID:   "012345678901234567890123",
//	}
//	good2 := db.Good{
//		ID:          1011,
//		GoodName:    "good",
//		Description: "Very good!",
//		ContentID:   "012345678901234567890123",
//	}
//	good3 := db.Good{
//		ID:          1012,
//		GoodName:    "good",
//		Description: "Very good!",
//		ContentID:   "012345678901234567890123",
//		Price:       100,
//	}
//
//	prepare := func() {
//		So(db.Ormer.Create(&good1).Error, ShouldBeNil)
//		So(db.Ormer.Create(&good2).Error, ShouldBeNil)
//		So(db.Ormer.Create(&good3).Error, ShouldBeNil)
//		So(db.Ormer.Create(&info1).Error, ShouldBeNil)
//		So(db.Ormer.Create(&info2).Error, ShouldBeNil)
//		So(db.Ormer.Create(&info3).Error, ShouldBeNil)
//	}
//	end := func() {
//		So(db.Ormer.Delete(&db.SellInfo{ID: 1000}).Error, ShouldBeNil)
//		So(db.Ormer.Delete(&db.SellInfo{ID: 1001}).Error, ShouldBeNil)
//		So(db.Ormer.Delete(&db.SellInfo{ID: 1002}).Error, ShouldBeNil)
//
//		So(db.Ormer.Delete(&db.Good{ID: 1010}).Error, ShouldBeNil)
//		So(db.Ormer.Delete(&db.Good{ID: 1011}).Error, ShouldBeNil)
//		So(db.Ormer.Delete(&db.Good{ID: 1012}).Error, ShouldBeNil)
//	}
//	testLen := func(length int) {
//		var rsp sellinfo.SellInfoFindResponse
//
//		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
//		So(len(rsp.SellInfo), ShouldEqual, length)
//	}
//
//	Convey("Test SellInfo Find", t, func() {
//		testLen(0)
//
//		prepare()
//		defer end()
//
//		req.LowPrice = -1
//		req.HighPrice = -1
//		testLen(3)
//
//		req.UserID = 1001
//		req.Limit = 101
//		req.Status = 3
//		req.GoodName = "good"
//		req.LowPrice = 99
//		req.HighPrice = 101
//		testLen(1)
//		req.GoodName = ""
//		req.LowPrice = 0
//		req.HighPrice = 0
//		req.Status = 0
//
//		req.UserID = 1000
//		testLen(2)
//
//		req.Limit = 1
//		testLen(1)
//
//		req.Offset = 1
//		testLen(1)
//
//		req.Offset = 2
//		testLen(0)
//	})
//}

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
