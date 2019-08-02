package main

import (
	auth "jiaojiao/srv/auth/proto"
	user "jiaojiao/srv/user/proto"
	"jiaojiao/utils"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getAuth(t *testing.T) {
	tf := func(code int, param string, status auth.AuthResponse_Status, id int, role user.UserInfo_Role) {
		c, d := utils.GetTestData(setupRouter, "GET", "/auth?code="+param, nil, "")

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
			if status == auth.AuthResponse_SUCCESS {
				t, err := utils.JWTVerify(d["token"].(string), os.Getenv("JJ_JWTSECRET"))
				So(err, ShouldBeNil)
				So(utils.JWTParse(t, "id"), ShouldEqual, id)
				So(utils.JWTParse(t, "role"), ShouldEqual, role)
			}
		}
	}
	Convey("Auth router test", t, func() {
		r := utils.StartTestServer(setupRouter, "GET", "/auth", nil, nil)
		So(r.Code, ShouldEqual, 301)

		tf(200, "invalid", auth.AuthResponse_INVALID_CODE, 0, 0)
		tf(200, "valid_user", auth.AuthResponse_SUCCESS, 1000, user.UserInfo_USER)
		tf(200, "valid_admin", auth.AuthResponse_SUCCESS, 1001, user.UserInfo_ADMIN)
		tf(200, "frozen_user", auth.AuthResponse_FROZEN_USER, 0, 0)
		tf(500, "error", 0, 0, 0)
		tf(500, "user_error", 0, 0, 0)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
