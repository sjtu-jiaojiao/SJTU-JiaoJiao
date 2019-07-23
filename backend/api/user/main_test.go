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
	tf := func(code int, status user.UserCreateResponse_Status, studentID string, studentName string, id int, user string) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "POST", "/user", strings.NewReader(url.Values{
			"studentID":   {studentID},
			"studentName": {studentName},
		}.Encode()),
			func(r *http.Request) {
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				r.Header.Set("Authorization", user)
			})
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
			So(data["user"].(map[string]interface{})["userID"], ShouldEqual, id)
		}
	}
	Convey("Add user router test", t, func() {
		tf(403, 0, "1000", "test", 0, "")
		tf(403, 0, "1000", "test", 0, "user")
		tf(403, 0, "1000", "test", 0, "self")

		tf(400, 0, "", "test", 1, "admin")
		tf(400, 0, "1000", "", 1, "admin")
		tf(200, user.UserCreateResponse_SUCCESS, "1000", "test", 1, "admin")
		tf(200, user.UserCreateResponse_SUCCESS, "1001", "test", 2, "admin")
		tf(500, 0, "2000", "test", 0, "admin")
		tf(200, user.UserCreateResponse_USER_EXIST, "3001", "test", 1, "admin")
	})
}

func Test_getUserInfo(t *testing.T) {
	tf := func(code int, path string, admin bool, id int, studentID interface{}) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", "/user/"+path, nil,
			func(r *http.Request) {
				if admin {
					r.Header.Set("Authorization", "admin")
				}
			})
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 && r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["userID"], ShouldEqual, id)
			So(data["studentID"], ShouldEqual, studentID)
		}
	}
	Convey("GetUserInfo router test", t, func() {
		r := utils.StartTestServer(setupRouter, "GET", "/user/0", nil, nil)
		So(r.Code, ShouldEqual, 400)

		tf(200, "1000", false, 1000, nil)
		tf(200, "1000", true, 1000, "1000")
		tf(500, "2000", false, 0, nil)
	})
}

func Test_updateUser(t *testing.T) {
	v := url.Values{}
	tf := func(code int, status user.UserUpdateResponse_Status, user string) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "PUT", "/user", strings.NewReader(v.Encode()), func(r *http.Request) {
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			r.Header.Set("Authorization", user)
		})
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}

	Convey("Update user router test", t, func() {
		tf(400, 0, "")
		tf(400, 0, "admin")
		v.Set("userID", "1000")
		tf(403, 0, "")
		tf(403, 0, "user")
		tf(200, user.UserUpdateResponse_SUCCESS, "self")
		tf(200, user.UserUpdateResponse_SUCCESS, "admin")
		v.Set("userID", "1001")
		tf(200, user.UserUpdateResponse_NOT_FOUND, "self")
		v.Set("userID", "2000")
		tf(500, 0, "self")
	})
}

func Test_findUser(t *testing.T) {
	tf := func(code int, path string, admin bool, cnt int, studentID interface{}) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", "/user?"+path, nil,
			func(r *http.Request) {
				if admin {
					r.Header.Set("Authorization", "admin")
				}
			})
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 && r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(len(data["user"].([]interface{})), ShouldEqual, cnt)
			So(data["user"].([]interface{})[0].(map[string]interface{})["studentID"], ShouldEqual, studentID)
		}
	}
	Convey("FindUser router test", t, func() {
		tf(200, "userName=test1", false, 1, nil)
		tf(200, "userName=test1", true, 1, "1")
		tf(200, "userName=test2", false, 2, nil)
		tf(200, "userName=test3", true, 3, "1")
		tf(200, "userName=", true, 3, "1")
		tf(200, "limit=2", true, 2, "1")
		tf(200, "limit=2&offset=1", true, 2, "2")
		tf(200, "limit=2&offset=2", true, 1, "3")
		tf(500, "userName=down", false, 0, nil)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
