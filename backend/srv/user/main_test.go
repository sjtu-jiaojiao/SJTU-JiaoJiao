package main

import (
	"context"
	db "jiaojiao/database"
	user "jiaojiao/srv/user/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserCreate(t *testing.T) {
	var s srv
	var req user.UserCreateRequest

	tf := func(status user.UserCreateResponse_Status) int32 {
		var rsp user.UserCreateResponse
		So(s.Create(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
		if rsp.Status == user.UserCreateResponse_SUCCESS || rsp.Status == user.UserCreateResponse_USER_EXIST {
			return rsp.User.UserID
		}
		return 0
	}
	Convey("Test User Create", t, func() {
		tf(user.UserCreateResponse_INVALID_PARAM)

		req.StudentID = "1234"
		tf(user.UserCreateResponse_INVALID_PARAM)

		req.StudentID = ""
		req.StudentName = "jiang"
		tf(user.UserCreateResponse_INVALID_PARAM)

		req.StudentID = "1234"
		id := tf(user.UserCreateResponse_SUCCESS)
		So(id, ShouldBeGreaterThan, 0)
		defer func() { So(db.Ormer.Delete(&db.User{ID: id}).Error, ShouldBeNil) }()

		id2 := tf(user.UserCreateResponse_USER_EXIST)
		So(id, ShouldEqual, id2)
	})
}

func TestUserQuery(t *testing.T) {
	var s srv
	var req user.UserQueryRequest

	tf := func(uid int, sid string) {
		var rsp user.UserInfo
		So(s.Query(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.UserID, ShouldEqual, uid)
		So(rsp.StudentID, ShouldEqual, sid)
	}
	Convey("Test User Query", t, func() {
		tf(0, "")
		err := db.Ormer.Create(&db.User{
			ID:          1000,
			UserName:    "test",
			AvatarID:    "012345678901234567890123",
			Telephone:   "12345678901",
			StudentID:   "10000",
			StudentName: "jiang",
		}).Error

		So(err, ShouldBeNil)
		req.UserID = 1000
		tf(1000, "10000")
		defer func() { So(db.Ormer.Delete(&db.User{ID: 1000}).Error, ShouldBeNil) }()

		req.UserID = 1001
		tf(0, "")
	})
}

func TestUserUpdate(t *testing.T) {
	var req user.UserInfo

	tf := func(status user.UserUpdateResponse_Status, name string, telephone string) {
		var s srv
		var rsp user.UserUpdateResponse
		So(s.Update(context.TODO(), &req, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)

		info := db.User{
			ID: 1100,
		}
		err := db.Ormer.First(&info).Error
		So(err, ShouldBeNil)
		So(info.UserName, ShouldEqual, name)
		So(info.Telephone, ShouldEqual, telephone)
	}
	Convey("Test User Update", t, func() {
		err := db.Ormer.Create(&db.User{
			ID:          1100,
			UserName:    "test",
			AvatarID:    "012345678901234567890123",
			Telephone:   "12345678901",
			StudentID:   "11000",
			StudentName: "jiang",
			Status:      int32(user.UserInfo_NORMAL),
			Role:        int32(user.UserInfo_USER),
		}).Error
		So(err, ShouldBeNil)
		defer func() { So(db.Ormer.Delete(&db.User{ID: 1100}).Error, ShouldBeNil) }()

		tf(user.UserUpdateResponse_INVALID_PARAM, "test", "12345678901")

		req.UserID = 1101
		tf(user.UserUpdateResponse_NOT_FOUND, "test", "12345678901")

		req.UserID = 1100
		tf(user.UserUpdateResponse_SUCCESS, "test", "12345678901")

		req.UserName = "test1"
		tf(user.UserUpdateResponse_SUCCESS, "test1", "12345678901")

		req.Telephone = "56781234678"
		tf(user.UserUpdateResponse_SUCCESS, "test1", "56781234678")

		req.ClearEmpty = true
		req.Telephone = ""
		tf(user.UserUpdateResponse_SUCCESS, "test1", "")
	})
}

func TestUserFind(t *testing.T) {
	var req user.UserFindRequest
	tf := func(cnt int, index int, uid int, sid string) {
		var s srv
		var rsp user.UserFindResponse
		So(s.Find(context.TODO(), &req, &rsp), ShouldBeNil)
		So(len(rsp.User), ShouldEqual, cnt)
		So(rsp.User[index].UserID, ShouldEqual, uid)
		So(rsp.User[index].StudentID, ShouldEqual, sid)
	}
	Convey("Test User Find", t, func() {
		err := db.Ormer.Create(&db.User{
			ID:          1200,
			UserName:    "test1",
			AvatarID:    "012345678901234567890123",
			Telephone:   "12345678901",
			StudentID:   "12000",
			StudentName: "jiang",
			Status:      1,
		}).Error
		So(err, ShouldBeNil)
		defer func() {
			So(db.Ormer.Delete(&db.User{ID: 1200}).Error, ShouldBeNil)
		}()

		tf(1, 0, 1200, "12000")

		err = db.Ormer.Create(&db.User{
			ID:          1201,
			UserName:    "test2",
			AvatarID:    "012345678901234567890123",
			Telephone:   "12345678902",
			StudentID:   "12010",
			StudentName: "jiangzm",
			Status:      1,
		}).Error
		So(err, ShouldBeNil)
		defer func() {
			So(db.Ormer.Delete(&db.User{ID: 1201}).Error, ShouldBeNil)
		}()

		req.Limit = 200
		tf(2, 1, 1201, "12010")

		req.Limit = 1
		tf(1, 0, 1200, "12000")

		req.Offset = 1
		tf(1, 0, 1201, "12010")
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
