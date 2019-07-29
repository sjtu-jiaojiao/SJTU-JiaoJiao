package main

import (
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"jiaojiao/utils"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getSellInfo(t *testing.T) {
	tf := func(code int, path string, id int, cid string) {
		c, d := utils.GetTestData(setupRouter, "GET", "/sellInfo/"+path, nil, "")

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["sellInfoID"], ShouldEqual, id)
			So(d["contentID"], ShouldEqual, cid)
		}
	}
	Convey("GetSellInfo router test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: true,
			User:  true,
			Self:  true,
			Admin: true,
		}, "GET", "/sellinfo/1000", nil), ShouldBeZeroValue)

		tf(400, "0", 0, "")
		tf(200, "1000", 1000, "123456789abc123456789abc")
		tf(500, "2000", 0, "")
	})
}

func Test_addSellInfo(t *testing.T) {
	v := url.Values{}
	tf := func(code int, status sellinfo.SellInfoCreateResponse_Status, user string) {
		c, d := utils.GetTestData(setupRouter, "POST", "/sellInfo", v, user)

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
		}
	}
	Convey("AddSellInfo router test", t, func() {
		tf(400, 0, "")
		v.Set("userID", "1")
		tf(400, 0, "")
		v.Set("validTime", "12345")
		tf(400, 0, "")
		v.Set("goodName", "good")
		tf(200, sellinfo.SellInfoCreateResponse_SUCCESS, "self")

		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  false,
			Self:  true,
			Admin: true,
		}, "POST", "/sellInfo", v), ShouldBeZeroValue)

		v.Set("contentID", "1234")
		tf(400, 0, "admin")
		v.Del("contentID")
		v.Set("contentToken", "valid")
		tf(400, 0, "admin")
		v.Set("contentID", "1234")
		tf(200, sellinfo.SellInfoCreateResponse_SUCCESS, "self")
		v.Set("contentToken", "invalid_token")
		tf(200, sellinfo.SellInfoCreateResponse_INVALID_TOKEN, "self")
		v.Set("contentID", "error")
		tf(500, 0, "admin")
	})
}

func Test_findSellInfo(t *testing.T) {
	// TODO
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
