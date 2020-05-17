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
	tf := func(status auth.AuthResponse_Status, token string) {
		var rsp auth.AuthResponse
		So(s.Auth(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		So(rsp.Token, ShouldEqual, token)
	}
	Convey("Test Auth", t, func() {
		tf(auth.AuthResponse_INVALID_PARAM, "")

		req.Code = "123456"
		tf(auth.AuthResponse_INVALID_CODE, "")

		// No test for valid code
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
