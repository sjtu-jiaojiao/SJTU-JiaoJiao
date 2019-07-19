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
		if rsp.Status == user.UserCreateResponse_SUCCESS || rsp.Status == user.UserCreateResponse_USER_EXIST {
			return rsp.User.UserId
		}
		return 0
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

		So(db.Ormer.Delete(&db.User{ID: id}).Error, ShouldBeNil)
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
		err := db.Ormer.Create(&db.User{
			ID:          1000,
			UserName:    "jiang",
			AvatarId:    "5d23ea2c32311335f935cd14",
			Telephone:   "12345678901",
			StudentId:   "1234",
			StudentName: "jiang",
		}).Error

		So(err, ShouldBeNil)
		req.UserId = 1000
		tf(1000, "jiang", "5d23ea2c32311335f935cd14",
			"12345678901", "1234", "jiang", 1)

		req.UserId = 1001
		tf(0, "", "", "", "", "", 0)

		So(db.Ormer.Delete(&db.User{ID: 1000}).Error, ShouldBeNil)
	})
}

func TestUserUpdate(t *testing.T) {
	var s srvUser
	var req user.UserInfo
	var rsp user.UserUpdateResponse
	tf := func(uid int, uname string, avatar string,
		telephone string, sid string, sname string, status user.UserInfo_Status, role user.UserInfo_Role) {
		info := db.User{
			Id: 2000,
		}
		err := db.Ormer.Read(&info)
		So(err, ShouldBeNil)
		So(info.Id, ShouldEqual, uid)
		So(info.UserName, ShouldEqual, uname)
		So(info.AvatarId, ShouldEqual, avatar)
		So(info.Telephone, ShouldEqual, telephone)
		So(info.StudentId, ShouldEqual, sid)
		So(info.StudentName, ShouldEqual, sname)
		So(info.Status, ShouldEqual, int32(status))
		So(info.Role, ShouldEqual, int32(role))
	}
	Convey("Test User Update", t, func() {
		_, err := db.Ormer.Insert(&db.User{
			Id:          2000,
			UserName:    "test1",
			AvatarId:    "5d23ea2c32311335f935cd14",
			Telephone:   "12345678901",
			StudentId:   "1234",
			StudentName: "jiang",
			Status:      int32(user.UserInfo_NORMAL),
			Role:        int32(user.UserInfo_USER),
		})
		So(err, ShouldBeNil)
		tf(2000, "test1", "5d23ea2c32311335f935cd14", "12345678901", "1234", "jiang", user.UserInfo_NORMAL, user.UserInfo_USER)

		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_INVALID_PARAM)

		req.UserId = 3000
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_NOT_FOUND)

		req.UserId = 2000
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_SUCCESS)
		tf(2000, "test1", "5d23ea2c32311335f935cd14", "12345678901", "1234", "jiang", user.UserInfo_NORMAL, user.UserInfo_USER)

		req.UserName = "test2"
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_SUCCESS)
		tf(2000, "test2", "5d23ea2c32311335f935cd14", "12345678901", "1234", "jiang", user.UserInfo_NORMAL, user.UserInfo_USER)

		req.AvatarId = "1111ea2c32311335f935cd14"
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_SUCCESS)
		tf(2000, "test2", "1111ea2c32311335f935cd14", "12345678901", "1234", "jiang", user.UserInfo_NORMAL, user.UserInfo_USER)

		req.Telephone = "56781234678"
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_SUCCESS)
		tf(2000, "test2", "1111ea2c32311335f935cd14", "56781234678", "1234", "jiang", user.UserInfo_NORMAL, user.UserInfo_USER)

		req.StudentId = "5689"
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_SUCCESS)
		tf(2000, "test2", "1111ea2c32311335f935cd14", "56781234678", "5689", "jiang", user.UserInfo_NORMAL, user.UserInfo_USER)

		req.StudentName = "xiaowang"
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_SUCCESS)
		tf(2000, "test2", "1111ea2c32311335f935cd14", "56781234678", "5689", "xiaowang", user.UserInfo_NORMAL, user.UserInfo_USER)

		req.Status = user.UserInfo_FROZEN
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_SUCCESS)
		tf(2000, "test2", "1111ea2c32311335f935cd14", "56781234678", "5689", "xiaowang", user.UserInfo_FROZEN, user.UserInfo_USER)

		req.Role = user.UserInfo_ADMIN
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, user.UserUpdateResponse_SUCCESS)
		tf(2000, "test2", "1111ea2c32311335f935cd14", "56781234678", "5689", "xiaowang", user.UserInfo_FROZEN, user.UserInfo_ADMIN)
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
		err := db.Ormer.Create(&db.User{
			ID:          2000,
			UserName:    "test1",
			AvatarId:    "5d23ea2c32311335f935cd14",
			Telephone:   "12345678901",
			StudentId:   "1234",
			StudentName: "jiang",
			Status:      1,
		}).Error
		So(err, ShouldBeNil)

		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.User), ShouldEqual, 1)
		tf(0, 2000, "test1", "5d23ea2c32311335f935cd14", "12345678901",
			"1234", "jiang", 1)
		rsp.User = nil

		err = db.Ormer.Create(&db.User{
			ID:          2001,
			UserName:    "test2",
			AvatarId:    "5d23ea2c32311335f935cd15",
			Telephone:   "12345678902",
			StudentId:   "12345",
			StudentName: "jiangzm",
			Status:      1,
		}).Error
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

		So(db.Ormer.Delete(&db.User{ID: 2000}).Error, ShouldBeNil)
		So(db.Ormer.Delete(&db.User{ID: 2001}).Error, ShouldBeNil)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
