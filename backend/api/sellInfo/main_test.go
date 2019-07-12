package main

import (
	"encoding/json"
	"jiaojiao/utils"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getSellInfo(t *testing.T) {
	tf := func(code int, path string, admin bool) map[string]interface{} {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", path, nil,
			func(r *http.Request) {
				if admin {
					r.Header.Set("Authorization", "valid_user")
				}
			})
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		}
		return data
	}
	Convey("GetSellInfo router test", t, func() {
		r := utils.StartTestServer(setupRouter, "GET", "/sellInfo", nil, nil)
		So(r.Code, ShouldEqual, 404)

		r = utils.StartTestServer(setupRouter, "GET", "/sellInfo/0", nil, nil)
		So(r.Code, ShouldEqual, 400)

		data := tf(200, "/sellInfo/1000", true)
		So(data["sellInfoId"], ShouldEqual, 1000)
		So(data["goodName"], ShouldEqual, "good")
		So(data["validTime"], ShouldEqual, 1234567890)
		So(data["contentId"], ShouldEqual, "123456789abc123456789abc")

		data = tf(200, "/sellInfo/1000", false)
		So(data["sellInfoId"], ShouldEqual, nil)
		So(data["goodName"], ShouldEqual, nil)
		So(data["validTime"], ShouldEqual, nil)
		So(data["contentId"], ShouldEqual, nil)

		r = utils.StartTestServer(setupRouter, "GET", "/sellInfo/2000", nil, nil)
		So(r.Code, ShouldEqual, 500)
	})
}