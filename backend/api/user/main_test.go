package main

import (
	"encoding/json"
	"jiaojiao/utils"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getUserInfo(t *testing.T) {
	tf := func(code int, path string, admin bool) map[string]interface{} {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", path, nil,
			func(r *http.Request) {
				if admin {
					r.Header.Set("Authorization", "admin")
				}
			})
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		}
		return data
	}
	Convey("GetUserInfo router test", t, func() {
		r := utils.StartTestServer(setupRouter, "GET", "/user", nil, nil)
		So(r.Code, ShouldEqual, 403)

		r = utils.StartTestServer(setupRouter, "GET", "/user/0", nil, nil)
		So(r.Code, ShouldEqual, 400)

		data := tf(200, "/user/1000", false)
		So(data["userName"], ShouldEqual, "test")

		r = utils.StartTestServer(setupRouter, "GET", "/user/2000", nil, nil)
		So(r.Code, ShouldEqual, 500)
	})
}

func Test_updateUser(t *testing.T) {
	var v = url.Values{
		"userId": {"1001"},
	}
	var data map[string]interface{}

	Convey("UpdateUser router test", t, func() {
		r := utils.StartTestServer(setupRouter, "POST", "/user", nil, nil)
		So(r.Code, ShouldEqual, 400)

		r = utils.StartTestServer(setupRouter, "POST", "/user/1001", nil, nil)
		So(r.Code, ShouldEqual, 404)

		r = utils.StartTestServer(setupRouter, "POST", "/user", strings.NewReader(v.Encode()), nil)
		So(r.Code, ShouldEqual, 403)

		r = utils.StartTestServer(setupRouter, "POST", "/user", strings.NewReader(v.Encode()),
			func(r *http.Request) {
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				r.Header.Set("Authorization", "valid_user")
			})
		So(r.Code, ShouldEqual, 200)
		So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		So(data["status"], ShouldEqual, 1)

		v.Set("userId", "2001")
		r = utils.StartTestServer(setupRouter, "POST", "/user", strings.NewReader(v.Encode()),
			func(r *http.Request) {
				r.Header.Set("Authorization", "valid_user")
			})
		So(r.Code, ShouldEqual, 200)
		So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		So(data["status"], ShouldEqual, -1)

		v.Set("userId", "3001")
		r = utils.StartTestServer(setupRouter, "POST", "/user", strings.NewReader(v.Encode()),
			func(r *http.Request) {
				r.Header.Set("Authorization", "valid_user")
			})
		So(r.Code, ShouldEqual, 200)
		So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		So(data["status"], ShouldEqual, 2)
	})
}

func Test_findUser(t *testing.T) {
	tf := func(code int, path string, admin bool) map[string]interface{} {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", path, nil,
			func(r *http.Request) {
				if admin {
					r.Header.Set("Authorization", "admin")
				}
			})
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		}
		return data
	}
	Convey("FindUser router test", t, func() {
		data := tf(200, "/user?userName=test1", false)
		So(len(data["user"].([]interface{})), ShouldEqual, 1)

		data = tf(200, "/user?userName=test2", false)
		So(len(data["user"].([]interface{})), ShouldEqual, 2)

		data = tf(200, "/user", true)
		So(len(data["user"].([]interface{})), ShouldEqual, 3)

		data = tf(200, "/user?limit=2", true)
		So(len(data["user"].([]interface{})), ShouldEqual, 2)

		data = tf(200, "/user?limit=2&offset=1", true)
		So(len(data["user"].([]interface{})), ShouldEqual, 2)

		data = tf(200, "/user?limit=2&offset=2", true)
		So(len(data["user"].([]interface{})), ShouldEqual, 1)

		tf(200, "/user?userName=test3", false)
		tf(500, "/user?userName=down", false)
	})
}

func Test_main(t *testing.T) {
	main()
}
