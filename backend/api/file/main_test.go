package main

import (
	"jiaojiao/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getFile(t *testing.T) {
	tf := func(code int, id string) {
		c, _ := utils.GetTestData(setupRouter, "GET", "/file/"+id, nil, "")

		So(c, ShouldEqual, code)
	}
	Convey("Get file test", t, func() {
		So(utils.RoleTest(setupRouter, utils.Role{
			Guest: true,
			User:  true,
			Self:  true,
			Admin: true,
		}, "GET", "/file/1", nil), ShouldBeZeroValue)

		tf(404, "invalid")
		tf(403, "invalid_type")
		tf(200, "valid")
		tf(500, "error")
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
