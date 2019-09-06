package main

import (
	avatar "jiaojiao/srv/avatar/proto"
	"jiaojiao/utils"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_addAvatar(t *testing.T) {
	tf := func(code int, userID string, file string, status avatar.AvatarCreateResponse_Status) {
		c, d := utils.GetTestData(setupRouter, "POST", "/avatar", url.Values{
			"file":   {file},
			"userID": {userID},
		}, "self")

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
		}
	}
	Convey("Add avatar test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  false,
			Self:  true,
			Admin: true,
		}, "POST", "/avatar", url.Values{
			"file":   {"valid"},
			"userID": {"1000"},
		}), ShouldBeZeroValue)

		tf(400, "0", "valid", 0)
		tf(400, "1000", "", 0)
		tf(200, "1000", "valid", avatar.AvatarCreateResponse_SUCCESS)
		tf(200, "1000", "invalid", avatar.AvatarCreateResponse_INVALID_TYPE)
		tf(500, "1000", "error", 0)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
