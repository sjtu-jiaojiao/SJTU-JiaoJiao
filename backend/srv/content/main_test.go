package main

import (
	"context"
	content "jiaojiao/srv/content/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSrvContentCreate(t *testing.T) {
	var req content.ContentCreateRequest

	tf := func(status content.ContentCreateResponse_Status, success bool) (string, string) {
		var s srv
		var rsp content.ContentCreateResponse
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if success {
			So(rsp.ContentId, ShouldNotBeBlank)
			So(rsp.ContentToken, ShouldNotBeBlank)
		} else {
			So(rsp.ContentId, ShouldBeBlank)
			So(rsp.ContentToken, ShouldBeBlank)
		}
		return rsp.ContentId, rsp.ContentToken
	}

	Convey("Test SellInfo Content Create", t, func() {
		req.Content = []byte{0}
		tf(content.ContentCreateResponse_INVALID_PARAM, false)

		req.Type = content.ContentCreateRequest_PICTURE
		tf(content.ContentCreateResponse_INVALID_PARAM, false)
		req.Type = 0
		req.Content = []byte{1, 2, 3, 4, 5, 6}
		tf(content.ContentCreateResponse_INVALID_PARAM, false)
		req.Type = content.ContentCreateRequest_PICTURE
		req.ContentId = "1234"
		tf(content.ContentCreateResponse_INVALID_PARAM, false)
		req.ContentId = ""
		req.ContentToken = "12463-25897fsfs-5232"
		tf(content.ContentCreateResponse_INVALID_PARAM, false)

		req.ContentId = "1234"
		tf(content.ContentCreateResponse_INVALID_TOKEN, false)

		req.ContentId = ""
		req.ContentToken = ""
		id, token := tf(content.ContentCreateResponse_SUCCESS, true)

		req.ContentId = id
		req.ContentToken = token
		tf(content.ContentCreateResponse_SUCCESS, true)
		tf(content.ContentCreateResponse_SUCCESS, true)

		req.ContentToken = "12463-25897fsfs-5232"
		tf(content.ContentCreateResponse_INVALID_TOKEN, false)

		var sc srv
		var rspc content.ContentDeleteResponse
		err := sc.Delete(context.TODO(), &content.ContentDeleteRequest{
			ContentId:    id,
			ContentToken: token,
		}, &rspc)
		So(err, ShouldBeNil)
	})
}

func TestSrvContentDelete(t *testing.T) {
	// TODO
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
