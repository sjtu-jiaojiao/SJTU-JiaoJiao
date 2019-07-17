package main

import (
	"encoding/json"
	"jiaojiao/utils"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getSellInfo(t *testing.T) {
	tf := func(code int, path string) map[string]interface{} {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", path, nil, nil)
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		}
		return data
	}
	Convey("GetSellInfo router test", t, func() {
		r := utils.StartTestServer(setupRouter, "GET", "/sellInfo/0", nil, nil)
		So(r.Code, ShouldEqual, 400)

		data := tf(200, "/sellInfo/1000")
		So(data["sellInfoId"], ShouldEqual, 1000)
		So(data["goodName"], ShouldEqual, "good")
		So(data["validTime"], ShouldEqual, 1234567890)
		So(data["contentId"], ShouldEqual, "123456789abc123456789abc")
		So(data["userId"], ShouldEqual, 1000)

		r = utils.StartTestServer(setupRouter, "GET", "/sellInfo/2000", nil, nil)
		So(r.Code, ShouldEqual, 500)
	})
}

func Test_addSellInfo(t *testing.T) {
	tf := func(code int, path string, status int) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "PUT", path, nil,
			func(r *http.Request) {
				r.Header.Set("Authorization", "valid_user")
			})
		So(r.Code, ShouldEqual, code)
		if r.Code != 500 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}
	Convey("AddSellInfo router test", t, func() {
		r := utils.StartTestServer(setupRouter, "PUT", "/sellInfo?validTime=12345&goodName=good&userId=1", nil, nil)
		So(r.Code, ShouldEqual, 403)
		r = utils.StartTestServer(setupRouter, "PUT", "/sellInfo?validTime=12345&goodName=good&userId=1", nil,
			func(r *http.Request) {
				r.Header.Set("Authorization", "user")
			})
		So(r.Code, ShouldEqual, 403)

		tf(200, "/sellInfo", -1)
		tf(200, "/sellInfo?validTime=12345&goodName=good&userId=1", 1)
		tf(200, "/sellInfo?validTime=12345&goodName=good&userId=1&contentId=1234&contentToken=invalid_token", 2)
		tf(200, "/sellInfo?validTime=12345&goodName=good&userId=1&contentId=1234&contentToken=valid", 1)
		tf(500, "/sellInfo?validTime=12345&goodName=good&userId=1&contentId=error", 0)
		tf(200, "/sellInfo?validTime=12345&goodName=good&userId=1&contentId=1234", -1)
		tf(200, "/sellInfo?validTime=12345&goodName=good&userId=1&contentToken=invalid_token", -1)
	})
}

func Test_deleteContent(t *testing.T) {
	tf := func(code int, cid string, ctoken string, status int32) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "DELETE", "/content?contentId="+cid+"&contentToken="+ctoken, nil,
			func(r *http.Request) {
				r.Header.Set("Authorization", "admin")
			})
		So(r.Code, ShouldEqual, code)
		if r.Code != 500 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}
	Convey("GetSellInfo router test", t, func() {
		r := utils.StartTestServer(setupRouter, "DELETE", "/content?contentId=1000&contentToken=valid_token", nil, nil)
		So(r.Code, ShouldEqual, 403)

		tf(200, "", "", -1)
		tf(200, "1000", "", -1)
		tf(200, "", "valid_token", -1)
		tf(200, "1000", "valid_token", 1)
		tf(200, "1000", "invalid_token", 2)
		tf(200, "1001", "valid_token", 2)
		tf(200, "1001", "invalid_token", 2)
		tf(500, "2000", "valid_token", 0)
	})
}

func Test_addContent(t *testing.T) {
	tf := func(code int, cid string, ctoken string, content string, ctype string, status int32) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "PUT", "/content?contentId="+cid+"&contentToken="+ctoken+"&content="+content+"&type="+ctype,
			nil, func(r *http.Request) {
				r.Header.Set("Authorization", "user")
			})
		So(r.Code, ShouldEqual, code)
		if r.Code != 500 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}
	Convey("AddContent router test", t, func() {
		r := utils.StartTestServer(setupRouter, "PUT", "/content", nil, nil)
		So(r.Code, ShouldEqual, 403)
		tf(200, "", "", "123", "", -1)
		tf(200, "", "", "", "1", -1)
		tf(200, "", "", "123", "1", 1)
		tf(200, "123", "invalid_token", "123", "1", 2)
		tf(200, "123", "valid_token", "123", "1", 1)
		tf(500, "error", "", "123", "1", 0)
		tf(200, "123", "", "123", "1", -1)
		tf(200, "", "token", "123", "1", -1)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
