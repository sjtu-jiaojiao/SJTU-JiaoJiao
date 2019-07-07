package main

import (
	"context"
	auth "jiaojiao/srv/auth/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuth(t *testing.T) {
	var s srv
	var req auth.AuthRequest
	var rsp auth.AuthResponse
	Convey("Test Auth", t, func() {
		So(s.Auth(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(rsp.Status, ShouldEqual, -1)
		So(rsp.Token, ShouldEqual, "")

		req.Code = "123456"
		So(s.Auth(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(rsp.Status, ShouldEqual, 2)
		So(rsp.Token, ShouldEqual, "")

		// No test for valid code
	})
}

func Test_main(t *testing.T) {
	main()
}
