package main

import (
	message "jiaojiao/srv/message/proto"
	"jiaojiao/utils"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_addMessage(t *testing.T) {
	tf := func(code int, f string, to string, t string, m string, status message.MessageCreateResponse_Status) {
		c, d := utils.GetTestData(setupRouter, "POST", "/message", url.Values{
			"fromUser": {f},
			"toUser":   {to},
			"type":     {t},
			"msg":      {m},
		}, "self")

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
		}
	}
	Convey("Add message test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  false,
			Self:  true,
			Admin: false,
		}, "POST", "/message", url.Values{
			"fromUser": {"1000"},
			"toUser":   {"2000"},
			"type":     {"1"},
			"msg":      {"msg"},
		}), ShouldBeZeroValue)

		tf(400, "0", "2000", "1", "valid", 0)
		tf(400, "1000", "0", "1", "valid", 0)
		tf(400, "1000", "2000", "0", "valid", 0)
		tf(400, "1000", "2000", "1", "", 0)
		tf(400, "1000", "2000", "2", "", 0)
		tf(400, "1000", "2000", "3", "", 0)
		tf(400, "1000", "2000", "2", "", 0)
		tf(200, "1000", "2000", "1", "text msg", message.MessageCreateResponse_SUCCESS)
		tf(200, "1000", "2000", "2", "valid_file", message.MessageCreateResponse_SUCCESS)
		tf(200, "1000", "2000", "2", "valid", message.MessageCreateResponse_INVALID_TYPE)
		tf(500, "3000", "2000", "1", "error", 0)
	})
}

func Test_findMessage(t *testing.T) {
	tf := func(code int, f string, to string, w string, status message.MessageQueryResponse_Status) {
		c, d := utils.GetTestData(setupRouter, "GET", "/message?"+url.Values{
			"fromUser": {f},
			"toUser":   {to},
			"way":      {w},
		}.Encode(), nil, "self")

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
		}
	}
	Convey("Find message test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  false,
			Self:  true,
			Admin: true,
		}, "GET", "/message?"+url.Values{
			"fromUser": {"1000"},
			"toUser":   {"2000"},
			"way":      {"1"},
		}.Encode(), nil), ShouldBeZeroValue)

		tf(400, "0", "2000", "1", 0)
		tf(400, "1000", "0", "1", 0)
		tf(400, "1000", "2000", "0", 0)
		tf(200, "1000", "2000", "1", message.MessageQueryResponse_SUCCESS)
		tf(500, "3000", "2000", "1", 0)
	})
}

func Test_getMessage(t *testing.T) {
	tf := func(code int, id string, status message.MessageQueryResponse_Status) {
		c, d := utils.GetTestData(setupRouter, "GET", "/message/"+id, nil, "self")

		So(c, ShouldEqual, code)
		if d != nil {
			So(d["status"], ShouldEqual, status)
		}
	}
	Convey("Get message test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: false,
			User:  false,
			Self:  true,
			Admin: true,
		}, "GET", "/message/1000", nil), ShouldBeZeroValue)

		tf(400, "0", 0)
		tf(200, "1000", message.MessageQueryResponse_SUCCESS)
		tf(500, "3000", 0)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
