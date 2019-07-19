package main

import (
	"encoding/json"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getAuth(t *testing.T) {
	tf := func(code int, param string, status int, id int, role user.UserInfo_Role) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", "/auth?code="+param, nil, nil)
		So(r.Code, ShouldEqual, code)
		if r.Code != 500 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldBeNil)
			So(data["status"], ShouldEqual, status)
			if status == 1 {
				t, err := utils.JWTVerify(data["token"].(string), os.Getenv("JJ_JWTSECRET"))
				So(err, ShouldBeNil)
				So(utils.JWTParse(t, "id"), ShouldEqual, id)
				So(utils.JWTParse(t, "role"), ShouldEqual, role)
			}
		}
	}
	Convey("Auth router test", t, func() {
		r := utils.StartTestServer(setupRouter, "GET", "/auth", nil, nil)
		So(r.Code, ShouldEqual, 301)

		tf(200, "invalid", 2, 0, 0)
		tf(200, "valid_user", 1, 1, user.UserInfo_USER)
		tf(200, "valid_admin", 1, 2, user.UserInfo_ADMIN)

		tf(200, "frozen_user", 3, 0, 0)
		tf(500, "down", 0, 0, 0)
		tf(500, "userdown", 0, 0, 0)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
