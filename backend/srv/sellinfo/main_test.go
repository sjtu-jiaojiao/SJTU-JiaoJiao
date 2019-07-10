package main

import (
	"context"
	sellinfo "jiaojiao/srv/sellinfo/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	var s srvContent
	var req sellinfo.ContentCreateRequest
	var rsp sellinfo.ContentCreateResponse
	Convey("Test SellInfo Create", t, func() {
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_PARAM)

		req.Content = []byte{1, 2, 3, 4, 5, 6}
		req.Type = sellinfo.ContentCreateRequest_PICTURE
		req.ContentId = "1234"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_PARAM)

		req.ContentId = ""
		req.ContentToken = "12463-25897fsfs-5232"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_PARAM)

		req.ContentId = "1234"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_TOKEN)

		req.ContentId = ""
		req.ContentToken = ""
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_SUCCESS)
		So(rsp.ContentId, ShouldNotBeBlank)
		So(rsp.ContentToken, ShouldNotBeBlank)

		req.ContentId = rsp.ContentId
		req.ContentToken = rsp.ContentToken
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_SUCCESS)
		So(rsp.ContentId, ShouldNotBeBlank)
		So(rsp.ContentToken, ShouldNotBeBlank)

		req.ContentToken = "12463-25897fsfs-5232"
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, sellinfo.ContentCreateResponse_INVALID_TOKEN)
	})
}
