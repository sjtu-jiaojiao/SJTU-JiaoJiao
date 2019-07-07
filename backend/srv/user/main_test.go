package main

import (
	"context"
	user "jiaojiao/srv/user/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	var s srv
	var req user.UserCreateRequest
	var rsp user.UserCreateResponse
	Convey("Test User Create", t, func() {
		So(s.Create(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(rsp.Status, ShouldEqual, -1)
		So(rsp.UserId, ShouldEqual, 0)

		req.StudentId = 1234
		So(s.Create(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(rsp.Status, ShouldEqual, -1)
		So(rsp.UserId, ShouldEqual, 0)

		req.StudentId = 0
		req.StudentName = "test"
		So(s.Create(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(rsp.Status, ShouldEqual, -1)
		So(rsp.UserId, ShouldEqual, 0)

		req.StudentId = 1234
		So(s.Create(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(rsp.Status, ShouldEqual, 1)
		So(rsp.UserId, ShouldEqual, 1)

		So(s.Create(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(rsp.Status, ShouldEqual, 2)
		So(rsp.UserId, ShouldEqual, 1)
	})
}

func Test_main(t *testing.T) {
	main()
}
