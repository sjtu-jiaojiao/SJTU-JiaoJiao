package main

import (
	"encoding/json"
	content "jiaojiao/srv/content/proto"
	"jiaojiao/utils"
	"net/http"
	"net/url"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_deleteContent(t *testing.T) {
	tf := func(code int, cid string, ctoken string, status content.ContentDeleteResponse_Status) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "DELETE", "/content?contentId="+cid+"&contentToken="+ctoken, nil,
			func(r *http.Request) {
				r.Header.Set("Authorization", "admin")
			})
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}
	Convey("GetSellInfo router test", t, func() {
		r := utils.StartTestServer(setupRouter, "DELETE", "/content?contentId=1000&contentToken=valid_token", nil, nil)
		So(r.Code, ShouldEqual, 403)

		tf(400, "", "", 0)
		tf(400, "1000", "", 0)
		tf(400, "", "valid_token", 0)
		tf(200, "1000", "valid_token", content.ContentDeleteResponse_SUCCESS)
		tf(200, "1000", "invalid_token", content.ContentDeleteResponse_INVALID_TOKEN)
		tf(200, "1001", "valid_token", content.ContentDeleteResponse_INVALID_TOKEN)
		tf(200, "1001", "invalid_token", content.ContentDeleteResponse_INVALID_TOKEN)
		tf(500, "2000", "valid_token", 0)
	})
}

func Test_addContent(t *testing.T) {
	v := url.Values{}
	tf := func(code int, status content.ContentCreateResponse_Status) {
		var data map[string]interface{}
		r := utils.StartTestServer(setupRouter, "POST", "/content",
			strings.NewReader(v.Encode()), func(r *http.Request) {
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				r.Header.Set("Authorization", "user")
			})
		So(r.Code, ShouldEqual, code)
		if r.Code == 200 {
			So(json.Unmarshal(r.Body.Bytes(), &data), ShouldEqual, nil)
			So(data["status"], ShouldEqual, status)
		}
	}
	Convey("AddContent router test", t, func() {
		tf(400, 0)
		v.Set("content", "123")
		tf(400, 0)
		v.Set("type", "1")
		tf(200, content.ContentCreateResponse_SUCCESS)
		r := utils.StartTestServer(setupRouter, "POST", "/content", strings.NewReader(v.Encode()),
			func(r *http.Request) {
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			})
		So(r.Code, ShouldEqual, 403)
		v.Set("contentId", "123")
		tf(400, 0)
		v.Set("contentToken", "valid_token")
		v.Del("contentId")
		tf(400, 0)
		v.Set("contentId", "123")
		tf(200, content.ContentCreateResponse_SUCCESS)
		v.Set("contentToken", "invalid_token")
		tf(200, content.ContentCreateResponse_INVALID_TOKEN)
		v.Set("contentId", "error")
		tf(500, 0)
		v.Set("contentId", "invalid")
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
