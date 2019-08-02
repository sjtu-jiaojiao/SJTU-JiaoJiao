package main

import (
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_addUser(t *testing.T) {
	tf := func(code int, status user.UserCreateResponse_Status, studentID string, studentName string, id int, user string) {
		c, d := utils.GetTestData(setupRouter, "POST", "/user", url.Values{
			"studentID":   {studentID},
			"studentName": {studentName},
		}, user)

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
			So(d["user"].(map[string]interface{})["userID"], ShouldEqual, id)
		}
	}
	Convey("Add user router test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  false,
			Self:  false,
			Admin: true,
		}, "POST", "/user", url.Values{
			"studentID":   {"10000"},
			"studentName": {"test"},
		}), ShouldBeZeroValue)

		tf(400, 0, "", "test", 1000, "admin")
		tf(400, 0, "10000", "", 1000, "admin")
		tf(200, user.UserCreateResponse_SUCCESS, "10000", "test", 1000, "admin")
		tf(200, user.UserCreateResponse_SUCCESS, "10001", "test", 1001, "admin")
		tf(500, 0, "30000", "test", 0, "admin")
		tf(200, user.UserCreateResponse_USER_EXIST, "30001", "test", 1000, "admin")
	})
}

func Test_getUserInfo(t *testing.T) {
	tf := func(code int, path string, admin bool, id int, studentID interface{}) {
		c, d := utils.GetTestData(setupRouter, "GET", "/user/"+path, nil, utils.If(admin, "admin", "").(string))

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["userID"], ShouldEqual, id)
			So(d["studentID"], ShouldEqual, studentID)
		}
	}
	Convey("GetUserInfo router test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: true,
			User:  true,
			Self:  true,
			Admin: true,
		}, "GET", "/user/1000", nil), ShouldBeZeroValue)

		tf(400, "0", false, 0, nil)
		tf(200, "1000", false, 1000, nil)
		tf(200, "1000", true, 1000, "10000")
		tf(500, "3000", false, 0, nil)
	})
}

func Test_updateUser(t *testing.T) {
	v := url.Values{}
	tf := func(code int, status user.UserUpdateResponse_Status, user string) {
		c, d := utils.GetTestData(setupRouter, "PUT", "/user", v, user)

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
		}
	}

	Convey("Update user router test", t, func() {
		tf(400, 0, "")
		tf(400, 0, "admin")
		v.Set("userID", "1000")
		tf(200, user.UserUpdateResponse_SUCCESS, "self")

		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  false,
			Self:  true,
			Admin: true,
		}, "PUT", "/user", v), ShouldBeZeroValue)

		v.Set("userID", "1001")
		tf(200, user.UserUpdateResponse_NOT_FOUND, "self")
		v.Set("userID", "3000")
		tf(500, 0, "self")
	})
}

func Test_findUser(t *testing.T) {
	v := url.Values{}
	tf := func(code int, admin bool, cnt int, studentID interface{}) {
		c, d := utils.GetTestData(setupRouter, "GET", "/user?"+v.Encode(), nil, utils.If(admin, "admin", "").(string))

		So(c, ShouldEqual, code)
		if d != nil {
			So(len(d["user"].([]interface{})), ShouldEqual, cnt)
			So(d["user"].([]interface{})[0].(map[string]interface{})["studentID"], ShouldEqual, studentID)
		}
	}
	Convey("FindUser router test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: true,
			User:  true,
			Self:  true,
			Admin: true,
		}, "GET", "/user?userName=test1", nil), ShouldBeZeroValue)

		tf(403, false, 0, nil)
		v.Set("userName", "test1")
		tf(200, false, 1, nil)
		tf(200, true, 1, "10000")
		v.Set("userName", "test2")
		tf(200, false, 2, nil)
		v.Set("userName", "test3")
		tf(200, true, 3, "10000")
		v.Set("userName", "error")
		tf(500, false, 0, nil)
		v.Del("userName")
		tf(200, true, 3, "10000")
		v.Set("limit", "2")
		tf(200, true, 2, "10000")
		v.Set("offset", "1")
		tf(200, true, 2, "10001")
		v.Set("offset", "2")
		tf(200, true, 1, "10002")
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
