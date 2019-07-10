package main

import (
	"encoding/json"
	"jiaojiao/utils"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_getAuth(t *testing.T) {
	tf := func(code int, path string) map[string]interface{} {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "GET", path, nil)
		So(r.Code, ShouldEqual, code)
		if r.Body.String() != "{}" {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
		}
		return data
	}
	Convey("Auth router test", t, func() {
		r := utils.StartTestServer(setupRouter, "GET", "/auth", nil)
		So(r.Code, ShouldEqual, 301)

		data := tf(200, "/auth?code=invalid")
		So(data["status"], ShouldEqual, 2)

		data = tf(200, "/auth?code=valid_user")
		So(data["status"], ShouldEqual, 1)
		t, err := utils.JWTVerify(data["token"].(string), os.Getenv("JJ_JWTSECRET"))
		So(err, ShouldEqual, nil)
		So(utils.JWTParse(t, "id"), ShouldEqual, 1)
		So(utils.JWTParse(t, "role"), ShouldEqual, 1)

		data = tf(200, "/auth?code=valid_admin")
		So(data["status"], ShouldEqual, 1)
		t, err = utils.JWTVerify(data["token"].(string), os.Getenv("JJ_JWTSECRET"))
		So(err, ShouldEqual, nil)
		So(utils.JWTParse(t, "id"), ShouldEqual, 1)
		So(utils.JWTParse(t, "role"), ShouldEqual, 2)

		tf(500, "/auth?code=down")
		tf(500, "/auth?code=userdown")
		tf(500, "/auth?code=admindown")
	})
}

func Test_main(t *testing.T) {
	main()
}
