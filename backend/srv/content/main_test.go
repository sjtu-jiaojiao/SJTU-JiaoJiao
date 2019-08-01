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
			So(rsp.ContentID, ShouldNotBeBlank)
			So(rsp.ContentToken, ShouldNotBeBlank)
		} else {
			So(rsp.ContentID, ShouldBeBlank)
			So(rsp.ContentToken, ShouldBeBlank)
		}
		return rsp.ContentID, rsp.ContentToken
	}

	Convey("Test SellInfo Content Create", t, func() {
		req.Content = []byte{0}
		tf(content.ContentCreateResponse_INVALID_PARAM, false)

		req.Type = content.Type_PICTURE
		tf(content.ContentCreateResponse_INVALID_PARAM, false)
		req.Type = 0
		req.Content = []byte{1, 2, 3, 4, 5, 6}
		tf(content.ContentCreateResponse_INVALID_PARAM, false)
		req.Type = content.Type_PICTURE
		req.ContentID = "1234"
		tf(content.ContentCreateResponse_INVALID_PARAM, false)
		req.ContentID = ""
		req.ContentToken = "12463-25897fsfs-5232"
		tf(content.ContentCreateResponse_INVALID_PARAM, false)

		req.ContentID = "1234"
		tf(content.ContentCreateResponse_INVALID_TOKEN, false)

		req.ContentID = ""
		req.ContentToken = ""
		id, token := tf(content.ContentCreateResponse_SUCCESS, true)

		req.ContentID = id
		req.ContentToken = token
		defer func() {
			var sc srv
			var rspc content.ContentDeleteResponse
			err := sc.Delete(context.TODO(), &content.ContentDeleteRequest{
				ContentID:    id,
				ContentToken: token,
			}, &rspc)
			So(err, ShouldBeNil)
		}()
		tf(content.ContentCreateResponse_SUCCESS, true)
		tf(content.ContentCreateResponse_SUCCESS, true)

		req.ContentToken = "12463-25897fsfs-5232"
		tf(content.ContentCreateResponse_INVALID_TOKEN, false)
	})
}

func TestSrvContentUpdate(t *testing.T) {
	// TODO
}

func TestSrvContentDelete(t *testing.T) {
	// TODO
}

func TestSrvContentQuery(t *testing.T) {
	// TODO
}

func TestSrvContentCheck(t *testing.T) {
	// TODO
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
