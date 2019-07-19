package main

import (
	"encoding/json"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_addUser(t *testing.T) {
	tf := func(code int, status user.UserCreateResponse_Status, path string, admin bool, studentId string, studentName string) map[string]interface{} {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "PUT", path,
			strings.NewReader(url.Values{
				"studentId":   {studentId},
				"studentName": {studentName},
			}.Encode()),
			func(r *http.Request) {
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				if admin {
					r.Header.Set("Authorization", "admin")
				}
			})
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}" && r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
		return data
	}
	Convey("Add user router test", t, func() {
		tf(404, user.UserCreateResponse_UNKNOWN, "/user/1000", true, "1000", "www")

		tf(403, user.UserCreateResponse_UNKNOWN, "/user", false, "1000", "www")

		tf(200, user.UserCreateResponse_INVALID_PARAM, "/user", true, "1000", "")

		tf(200, user.UserCreateResponse_INVALID_PARAM, "/user", true, "", "www")

		tf(200, user.UserCreateResponse_SUCCESS, "/user", true, "1000", "www")

		tf(200, user.UserCreateResponse_SUCCESS, "/user", true, "1001", "www")

		tf(200, user.UserCreateResponse_UNKNOWN, "/user", true, "2000", "www")

		tf(200, user.UserCreateResponse_USER_EXIST, "/user", true, "3000", "www")
	})
}

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
	v := url.Values{
		"userId": {"1001"},
	}
	tf := func(code int, status int32) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "POST", "/user", strings.NewReader(v.Encode()), func(r *http.Request) {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			r.Header.Set("Authorization", "valid_user")
		})
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}

	Convey("UpdateUser router test", t, func() {
		r := utils.StartTestServer(setupRouter, "POST", "/user", nil, nil)
		So(r.Code, ShouldEqual, 400)

		r = utils.StartTestServer(setupRouter, "POST", "/user", strings.NewReader(v.Encode()), func(r *http.Request) {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		})
		So(r.Code, ShouldEqual, 403)

		tf(200, 1)
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

func TestMain(m *testing.M) {
	main()
	m.Run()
}
