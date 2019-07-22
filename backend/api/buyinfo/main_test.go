package main

import (
	"encoding/json"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	"jiaojiao/utils"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getBuyInfo(t *testing.T) {
	tf := func(code int, path string, id int, cid string) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", "/buyInfo/"+path, nil, nil)
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["buyInfoId"], ShouldEqual, id)
			So(data["contentId"], ShouldEqual, cid)
		}
	}
	Convey("GetBuyInfo router test", t, func() {
		tf(400, "0", 0, "")
		tf(200, "1000", 1000, "123456789abc123456789abc")
		tf(500, "2000", 0, "")
	})
}

func Test_addBuyInfo(t *testing.T) {
	v := url.Values{}
	tf := func(code int, status buyinfo.BuyInfoCreateResponse_Status, user string) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "POST", "/buyInfo", strings.NewReader(v.Encode()),
			func(r *http.Request) {
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				r.Header.Set("Authorization", user)
			})
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}
	Convey("AddBuyInfo router test", t, func() {
		tf(400, 0, "")
		v.Set("userId", "1")
		tf(400, 0, "")
		v.Set("validTime", "12345")
		tf(400, 0, "")
		v.Set("goodName", "good")
		tf(403, 0, "")
		tf(403, 0, "user")
		tf(200, buyinfo.BuyInfoCreateResponse_SUCCESS, "self")
		tf(200, buyinfo.BuyInfoCreateResponse_SUCCESS, "admin")
		v.Set("contentId", "1234")
		tf(400, 0, "admin")
		v.Del("contentId")
		v.Set("contentToken", "valid")
		tf(400, 0, "admin")
		v.Set("contentId", "1234")
		tf(200, buyinfo.BuyInfoCreateResponse_SUCCESS, "admin")
		v.Set("contentToken", "invalid_token")
		tf(200, buyinfo.BuyInfoCreateResponse_INVALID_TOKEN, "admin")
		v.Set("contentId", "error")
		tf(500, 0, "admin")
	})
}

func Test_findBuyInfo(t *testing.T) {
	// TODO
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
