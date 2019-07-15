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

func Test_main(t *testing.T) {
	main()
}
