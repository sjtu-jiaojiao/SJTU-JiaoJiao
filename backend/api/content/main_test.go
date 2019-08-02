package main

import (
	content "jiaojiao/srv/content/proto"
	"jiaojiao/utils"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_deleteContent(t *testing.T) {
	tf := func(code int, cid string, cToken string, status content.ContentDeleteResponse_Status) {
		c, d := utils.GetTestData(setupRouter, "DELETE", "/content?contentID="+cid+"&contentToken="+cToken,
			nil, "admin")

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
		}
	}
	Convey("GetSellInfo router test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  true,
			Self:  true,
			Admin: true,
		}, "DELETE", "/content?contentID=1000&contentToken=valid_token", nil), ShouldBeZeroValue)

		tf(400, "", "", 0)
		tf(400, "012345678901234567890123", "", 0)
		tf(400, "", "valid_token", 0)
		tf(200, "012345678901234567890123", "valid_token", content.ContentDeleteResponse_SUCCESS)
		tf(200, "012345678901234567890123", "invalid_token", content.ContentDeleteResponse_INVALID_TOKEN)
		tf(200, "12345", "valid_token", content.ContentDeleteResponse_INVALID_TOKEN)
		tf(200, "12345", "invalid_token", content.ContentDeleteResponse_INVALID_TOKEN)
		tf(500, "987654321098765432109876", "valid_token", 0)
	})
}

func Test_addContent(t *testing.T) {
	v := url.Values{}
	tf := func(code int, status content.ContentCreateResponse_Status) {
		c, d := utils.GetTestData(setupRouter, "POST", "/content",
			v, "user")

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
		}
	}
	Convey("AddContent router test", t, func() {
		tf(400, 0)
		v.Set("content", "123")
		tf(400, 0)
		v.Set("type", "1")
		tf(200, content.ContentCreateResponse_SUCCESS)

		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  true,
			Self:  true,
			Admin: true,
		}, "POST", "/content", v), ShouldBeZeroValue)

		v.Set("contentID", "123")
		tf(400, 0)
		v.Set("contentToken", "valid_token")
		v.Del("contentID")
		tf(400, 0)
		v.Set("contentID", "123")
		tf(200, content.ContentCreateResponse_SUCCESS)
		v.Set("contentToken", "invalid_token")
		tf(200, content.ContentCreateResponse_INVALID_TOKEN)
		v.Set("contentID", "error")
		tf(500, 0)
		v.Set("contentID", "invalid")
		tf(200, content.ContentCreateResponse_INVALID_TYPE)
	})
}

func Test_updateContent(t *testing.T) {
	// TODO
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
