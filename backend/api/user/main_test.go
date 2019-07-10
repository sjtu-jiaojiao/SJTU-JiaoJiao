package main

import (
	"encoding/json"
	"jiaojiao/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getUserInfo(t *testing.T){
	tf := func(code int, path string) map[string]interface{} {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", path, nil)
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		}
		return data
	}
	Convey("GetUserInfo router test", t, func() {
		data := tf(200, "/user/")
		So(len(data["user"].([]interface{})), ShouldEqual, 1)

		data = tf(200, "/user?userName=test2")
		So(len(data["user"].([]interface{})), ShouldEqual, 2)

		data = tf(200, "/user")
		So(len(data["user"].([]interface{})), ShouldEqual, 3)

		data = tf(200, "/user?limit=2")
		So(len(data["user"].([]interface{})), ShouldEqual, 2)

		data = tf(200, "/user?limit=2&offset=1")
		So(len(data["user"].([]interface{})), ShouldEqual, 2)

		data = tf(200, "/user?limit=2&offset=2")
		So(len(data["user"].([]interface{})), ShouldEqual, 1)

		tf(500, "/user?user=test3")
	})
}

func Test_findUser(t *testing.T) {
	tf := func(code int, path string) map[string]interface{} {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", path, nil)
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		}
		return data
	}
	Convey("FindUser router test", t, func() {
		data := tf(200, "/user?userName=test1")
		So(len(data["user"].([]interface{})), ShouldEqual, 1)

		data = tf(200, "/user?userName=test2")
		So(len(data["user"].([]interface{})), ShouldEqual, 2)

		data = tf(200, "/user")
		So(len(data["user"].([]interface{})), ShouldEqual, 3)

		data = tf(200, "/user?limit=2")
		So(len(data["user"].([]interface{})), ShouldEqual, 2)

		data = tf(200, "/user?limit=2&offset=1")
		So(len(data["user"].([]interface{})), ShouldEqual, 2)

		data = tf(200, "/user?limit=2&offset=2")
		So(len(data["user"].([]interface{})), ShouldEqual, 1)

		tf(500, "/user?user=test3")
	})
}

func Test_main(t *testing.T) {
	main()
}
