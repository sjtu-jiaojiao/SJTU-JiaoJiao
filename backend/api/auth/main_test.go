package main

import (
	"encoding/json"
	"jiaojiao/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getAuth(t *testing.T) {
	tf := func(code int, path string) map[string]interface{} {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", path, nil)
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		}
		return data
	}
	Convey("Auth router test", t, func() {
		r := utils.StartTestServer(setupRouter, "GET", "/auth", nil)
		So(r.Code, ShouldEqual, 301)

		data := tf(200, "/auth?code=123456")
		So(data["status"], ShouldEqual, 3)

		data = tf(200, "/auth?code=valid")
		So(data["status"], ShouldEqual, 1)
		So(data["token"], ShouldEqual, "test_token")

		tf(500, "/auth?code=down")
	})
}

func Test_main(t *testing.T) {
	main()
}
