package main

import (
	"context"
	db "jiaojiao/database"
	user "jiaojiao/srv/user/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	var s srv
	var req user.UserCreateRequest

	tf := func(status user.UserCreateResponse_Status) int32 {
		var rsp user.UserCreateResponse
		So(s.Create(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(rsp.Status, ShouldEqual, status)
		return rsp.UserId
	}
	Convey("Test User Create", t, func() {
		tf(user.UserCreateResponse_INVALID_PARAM)

		req.StudentId = "1234"
		tf(user.UserCreateResponse_INVALID_PARAM)

		req.StudentId = ""
		req.StudentName = "jiang"
		tf(user.UserCreateResponse_INVALID_PARAM)

		req.StudentId = "1234"
		id := tf(user.UserCreateResponse_SUCCESS)
		So(id, ShouldBeGreaterThan, 0)

		id2 := tf(user.UserCreateResponse_USER_EXIST)
		So(id, ShouldEqual, id2)
	})
}

func TestQuery(t *testing.T) {
	var s srv
	var req user.UserQueryRequest

	tf := func(uid int, uname string, avatar string,
		telephone string, sid string, sname string) {
		var rsp user.UserInfo
		So(s.Query(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(rsp.UserId, ShouldEqual, uid)
		So(rsp.UserName, ShouldEqual, uname)
		So(rsp.AvatarId, ShouldEqual, avatar)
		So(rsp.Telephone, ShouldEqual, telephone)
		So(rsp.StudentId, ShouldEqual, sid)
		So(rsp.StudentName, ShouldEqual, sname)
	}
	Convey("Test User Query", t, func() {
		tf(0, "", "", "", "", "")
		_, err := o.Insert(&db.User{
			Id:          1000,
			UserName:    "jiang",
			AvatarId:    "5d23ea2c32311335f935cd14",
			Telephone:   "12345678901",
			StudentId:   "1234",
			StudentName: "jiang",
		})
		So(err, ShouldEqual, nil)
		req.UserId = 1000
		tf(1000, "jiang", "5d23ea2c32311335f935cd14",
			"12345678901", "1234", "jiang")

		req.UserId = 1001
		tf(0, "", "", "", "", "")
	})
}

func TestFind(t *testing.T) {
	var s srv
	var req user.UserFindRequest
	var rsp user.UserFindResponse
	tf := func(index int, uid int, uname string, avatar string,
		telephone string, sid string, sname string) {
		So(rsp.User[index].UserId, ShouldEqual, uid)
		So(rsp.User[index].UserName, ShouldEqual, uname)
		So(rsp.User[index].AvatarId, ShouldEqual, avatar)
		So(rsp.User[index].Telephone, ShouldEqual, telephone)
		So(rsp.User[index].StudentId, ShouldEqual, sid)
		So(rsp.User[index].StudentName, ShouldEqual, sname)
	}
	Convey("Test User Find", t, func() {
		req.UserName = "test"
		_, err := o.Insert(&db.User{
			Id:          2000,
			UserName:    "test1",
			AvatarId:    "5d23ea2c32311335f935cd14",
			Telephone:   "12345678901",
			StudentId:   "1234",
			StudentName: "jiang",
		})
		So(err, ShouldEqual, nil)

		So(s.Find(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(len(rsp.User), ShouldEqual, 1)
		tf(0, 2000, "test1", "5d23ea2c32311335f935cd14", "12345678901",
			"1234", "jiang")
		rsp.User = nil

		_, err = o.Insert(&db.User{
			Id:          2001,
			UserName:    "test2",
			AvatarId:    "5d23ea2c32311335f935cd15",
			Telephone:   "12345678902",
			StudentId:   "12345",
			StudentName: "jiangzm",
		})
		So(err, ShouldEqual, nil)

		So(s.Find(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(len(rsp.User), ShouldEqual, 2)
		tf(0, 2000, "test1", "5d23ea2c32311335f935cd14", "12345678901",
			"1234", "jiang")
		tf(1, 2001, "test2", "5d23ea2c32311335f935cd15", "12345678902",
			"12345", "jiangzm")
		rsp.User = nil

		req.Limit = 1
		So(s.Find(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(len(rsp.User), ShouldEqual, 1)
		tf(0, 2000, "test1", "5d23ea2c32311335f935cd14", "12345678901",
			"1234", "jiang")
		rsp.User = nil

		req.Offset = 1
		So(s.Find(context.TODO(), &req, &rsp), ShouldEqual, nil)
		So(len(rsp.User), ShouldEqual, 1)
		tf(0, 2001, "test2", "5d23ea2c32311335f935cd15", "12345678902",
			"12345", "jiangzm")
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
