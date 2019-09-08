package main

import (
	buyinfo "jiaojiao/srv/buyinfo/proto"
	"jiaojiao/utils"
	"net/url"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getBuyInfo(t *testing.T) {
	tf := func(code int, path string, cid string) {
		c, d := utils.GetTestData(setupRouter, "GET", "/buyInfo/"+path, nil, "")

		So(c, ShouldEqual, code)
		if d != nil {
			v, _ := strconv.Atoi(path)
			So(d["buyInfoID"], ShouldEqual, v)
			So(d["contentID"], ShouldEqual, cid)
		}
	}
	Convey("GetBuyInfo router test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: true,
			User:  true,
			Self:  true,
			Admin: true,
		}, "GET", "/buyInfo/1000", nil), ShouldBeZeroValue)

		tf(400, "0", "")
		tf(200, "1000", "012345678901234567890123")
		tf(500, "3000", "")
	})
}

func Test_addBuyInfo(t *testing.T) {
	v := url.Values{}
	tf := func(code int, status buyinfo.BuyInfoCreateResponse_Status, user string) {
		c, d := utils.GetTestData(setupRouter, "POST", "/buyInfo", v, user)

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
		}
	}
	Convey("AddBuyInfo router test", t, func() {
		tf(400, 0, "")
		v.Set("userID", "1")
		tf(400, 0, "")
		v.Set("validTime", "12345")
		tf(400, 0, "")
		v.Set("goodName", "good")
		tf(200, buyinfo.BuyInfoCreateResponse_SUCCESS, "self")

		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  false,
			Self:  true,
			Admin: true,
		}, "POST", "/buyInfo", v), ShouldBeZeroValue)

		v.Set("contentID", "1234")
		tf(400, 0, "admin")
		v.Del("contentID")
		v.Set("contentToken", "valid")
		tf(400, 0, "admin")
		v.Set("contentID", "1234")
		tf(200, buyinfo.BuyInfoCreateResponse_SUCCESS, "self")
		v.Set("contentToken", "invalid_token")
		tf(200, buyinfo.BuyInfoCreateResponse_INVALID_TOKEN, "self")
		v.Set("contentID", "error")
		tf(500, 0, "admin")
	})
}

func Test_findBuyInfo(t *testing.T) {
	tf := func(code int, value url.Values, length int) {
		c, d := utils.GetTestData(setupRouter, "GET", "/buyInfo?"+value.Encode(), nil, "")

		So(c, ShouldEqual, code)
		if d != nil {
			So(len(d["buyInfo"].([]interface{})), ShouldEqual, length)
		}
	}
	Convey("GetBuyInfo router test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: true,
			User:  true,
			Self:  true,
			Admin: true,
		}, "GET", "/buyinfo", url.Values{"userID": {"1001"}}), ShouldBeZeroValue)

		tf(200, url.Values{"userID": {"1000"}}, 2)
		tf(200, url.Values{"userID": {"1000"}, "status": {"1"}}, 1)
		tf(500, url.Values{"userID": {"1001"}}, 0)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
