package main

import (
	"encoding/json"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getSellInfo(t *testing.T) {
	tf := func(code int, path string, id int, cid string) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", "/sellInfo/"+path, nil, nil)
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["sellInfoId"], ShouldEqual, id)
			So(data["contentId"], ShouldEqual, cid)
		}
	}
	Convey("GetSellInfo router test", t, func() {
		tf(400, "0", 0, "")
		tf(200, "1000", 1000, "123456789abc123456789abc")
		tf(500, "2000", 0, "")
	})
}

func Test_addSellInfo(t *testing.T) {
	v := url.Values{}
	tf := func(code int, status sellinfo.SellInfoCreateResponse_Status, user string) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "POST", "/sellInfo", strings.NewReader(v.Encode()),
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
	Convey("AddSellInfo router test", t, func() {
		tf(400, 0, "")
		v.Set("userId", "1")
		tf(400, 0, "")
		v.Set("validTime", "12345")
		tf(400, 0, "")
		v.Set("goodName", "good")
		tf(403, 0, "")
		tf(403, 0, "user")
		tf(200, sellinfo.SellInfoCreateResponse_SUCCESS, "self")
		tf(200, sellinfo.SellInfoCreateResponse_SUCCESS, "admin")
		v.Set("contentId", "1234")
		tf(400, 0, "admin")
		v.Del("contentId")
		v.Set("contentToken", "valid")
		tf(400, 0, "admin")
		v.Set("contentId", "1234")
		tf(200, sellinfo.SellInfoCreateResponse_SUCCESS, "admin")
		v.Set("contentToken", "invalid_token")
		tf(200, sellinfo.SellInfoCreateResponse_INVALID_TOKEN, "admin")
		v.Set("contentId", "error")
		tf(500, 0, "admin")
	})
}

func Test_findSellInfo(t *testing.T) {
	// TODO
}

func Test_deleteContent(t *testing.T) {
	tf := func(code int, cid string, ctoken string, status sellinfo.ContentDeleteResponse_Status) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "DELETE", "/content?contentId="+cid+"&contentToken="+ctoken, nil,
			func(r *http.Request) {
				r.Header.Set("Authorization", "admin")
			})
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}
	Convey("GetSellInfo router test", t, func() {
		r := utils.StartTestServer(setupRouter, "DELETE", "/content?contentId=1000&contentToken=valid_token", nil, nil)
		So(r.Code, ShouldEqual, 403)

		tf(400, "", "", 0)
		tf(400, "1000", "", 0)
		tf(400, "", "valid_token", 0)
		tf(200, "1000", "valid_token", sellinfo.ContentDeleteResponse_SUCCESS)
		tf(200, "1000", "invalid_token", sellinfo.ContentDeleteResponse_INVALID_TOKEN)
		tf(200, "1001", "valid_token", sellinfo.ContentDeleteResponse_INVALID_TOKEN)
		tf(200, "1001", "invalid_token", sellinfo.ContentDeleteResponse_INVALID_TOKEN)
		tf(500, "2000", "valid_token", 0)
	})
}

func Test_addContent(t *testing.T) {
	v := url.Values{}
	tf := func(code int, status sellinfo.ContentCreateResponse_Status) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "POST", "/content",
			strings.NewReader(v.Encode()), func(r *http.Request) {
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				r.Header.Set("Authorization", "user")
			})
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}
	Convey("AddContent router test", t, func() {
		tf(400, 0)
		v.Set("content", "123")
		tf(400, 0)
		v.Set("type", "1")
		tf(200, sellinfo.ContentCreateResponse_SUCCESS)
		r := utils.StartTestServer(setupRouter, "POST", "/content", strings.NewReader(v.Encode()),
			func(r *http.Request) {
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			})
		So(r.Code, ShouldEqual, 403)
		v.Set("contentId", "123")
		tf(400, 0)
		v.Set("contentToken", "valid_token")
		v.Del("contentId")
		tf(400, 0)
		v.Set("contentId", "123")
		tf(200, sellinfo.ContentCreateResponse_SUCCESS)
		v.Set("contentToken", "invalid_token")
		tf(200, sellinfo.ContentCreateResponse_INVALID_TOKEN)
		v.Set("contentId", "error")
		tf(500, 0)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
