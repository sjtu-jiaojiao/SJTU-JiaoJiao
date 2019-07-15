package main

import (
	"context"
	db "jiaojiao/database"
	user "jiaojiao/srv/user/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserCreate(t *testing.T) {
	var s srvUser
	var req user.UserCreateRequest

	tf := func(status user.UserCreateResponse_Status) int32 {
		var rsp user.UserCreateResponse
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
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

		_, err := db.Ormer.Delete(&db.User{
			Id: id,
		})
		So(err, ShouldBeNil)
	})
}

func TestUserQuery(t *testing.T) {
	var s srvUser
	var req user.UserQueryRequest

	tf := func(uid int, uname string, avatar string,
		telephone string, sid string, sname string, status int) {
		var rsp user.UserInfo
		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.UserId, ShouldEqual, uid)
		So(rsp.UserName, ShouldEqual, uname)
		So(rsp.AvatarId, ShouldEqual, avatar)
		So(rsp.Telephone, ShouldEqual, telephone)
		So(rsp.StudentId, ShouldEqual, sid)
		So(rsp.StudentName, ShouldEqual, sname)
		So(rsp.Status, ShouldEqual, status)
	}
	Convey("Test User Query", t, func() {
		tf(0, "", "", "", "", "", 0)
		_, err := db.Ormer.Insert(&db.User{
			Id:          1000,
			UserName:    "jiang",
			AvatarId:    "5d23ea2c32311335f935cd14",
			Telephone:   "12345678901",
			StudentId:   "1234",
			StudentName: "jiang",
			Status:      1,
		})

		So(err, ShouldBeNil)
		req.UserId = 1000
		tf(1000, "jiang", "5d23ea2c32311335f935cd14",
			"12345678901", "1234", "jiang", 1)

		req.UserId = 1001
		tf(0, "", "", "", "", "", 0)

		_, err = db.Ormer.Delete(&db.User{
			Id: 1000,
		})
		So(err, ShouldBeNil)
	})
}

func TestUserFind(t *testing.T) {
	var s srvUser
	var req user.UserFindRequest
	var rsp user.UserFindResponse
	tf := func(index int, uid int, uname string, avatar string,
		telephone string, sid string, sname string, status int) {
		So(rsp.User[index].UserId, ShouldEqual, uid)
		So(rsp.User[index].UserName, ShouldEqual, uname)
		So(rsp.User[index].AvatarId, ShouldEqual, avatar)
		So(rsp.User[index].Telephone, ShouldEqual, telephone)
		So(rsp.User[index].StudentId, ShouldEqual, sid)
		So(rsp.User[index].StudentName, ShouldEqual, sname)
		So(rsp.User[index].Status, ShouldEqual, status)
	}
	Convey("Test User Find", t, func() {
		req.UserName = "test"
		_, err := db.Ormer.Insert(&db.User{
			Id:          2000,
			UserName:    "test1",
			AvatarId:    "5d23ea2c32311335f935cd14",
			Telephone:   "12345678901",
			StudentId:   "1234",
			StudentName: "jiang",
			Status:      1,
		})
		So(err, ShouldBeNil)

		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.User), ShouldEqual, 1)
		tf(0, 2000, "test1", "5d23ea2c32311335f935cd14", "12345678901",
			"1234", "jiang", 1)
		rsp.User = nil

		_, err = db.Ormer.Insert(&db.User{
			Id:          2001,
			UserName:    "test2",
			AvatarId:    "5d23ea2c32311335f935cd15",
			Telephone:   "12345678902",
			StudentId:   "12345",
			StudentName: "jiangzm",
			Status:      1,
		})
		So(err, ShouldBeNil)

		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.User), ShouldEqual, 2)
		tf(0, 2000, "test1", "5d23ea2c32311335f935cd14", "12345678901",
			"1234", "jiang", 1)
		tf(1, 2001, "test2", "5d23ea2c32311335f935cd15", "12345678902",
			"12345", "jiangzm", 1)
		rsp.User = nil

		req.Limit = 1
		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.User), ShouldEqual, 1)
		tf(0, 2000, "test1", "5d23ea2c32311335f935cd14", "12345678901",
			"1234", "jiang", 1)
		rsp.User = nil

		req.Offset = 1
		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.User), ShouldEqual, 1)
		tf(0, 2001, "test2", "5d23ea2c32311335f935cd15", "12345678902",
			"12345", "jiangzm", 1)

		_, err = db.Ormer.Delete(&db.User{
			Id: 2000,
		})
		So(err, ShouldBeNil)
		_, err = db.Ormer.Delete(&db.User{
			Id: 2001,
		})
		So(err, ShouldBeNil)
	})
}

func TestAdminUserCreate(t *testing.T) {
	var s srvAdmin
	var req user.AdminUserRequest

	tf := func(status user.AdminUserResponse_Status) int32 {
		var rsp user.AdminUserResponse
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		return rsp.AdminId
	}
	Convey("Test Admin User Create", t, func() {
		tf(user.AdminUserResponse_INVALID_PARAM)

		req.StudentId = "1000"
		id := tf(user.AdminUserResponse_SUCCESS)
		So(id, ShouldBeGreaterThan, 0)

		id2 := tf(user.AdminUserResponse_USER_EXIST)
		So(id, ShouldEqual, id2)

		_, err := db.Ormer.Delete(&db.AdminUser{
			Id: id,
		})
		So(err, ShouldBeNil)
	})
}

func TestAdminUserFind(t *testing.T) {
	var s srvAdmin
	var req user.AdminUserRequest
	tf := func(status user.AdminUserResponse_Status, id int) {
		var rsp user.AdminUserResponse
		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		So(rsp.AdminId, ShouldEqual, id)
	}
	Convey("Test Admin User Find", t, func() {
		tf(user.AdminUserResponse_INVALID_PARAM, 0)

		req.StudentId = "2000"
		tf(user.AdminUserResponse_NOT_FOUND, 0)

		_, err := db.Ormer.Insert(&db.AdminUser{
			Id:        1000,
			StudentId: "2000",
		})
		So(err, ShouldBeNil)

		tf(user.AdminUserResponse_SUCCESS, 1000)
		_, err = db.Ormer.Delete(&db.AdminUser{
			Id: 1000,
		})
		So(err, ShouldBeNil)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
