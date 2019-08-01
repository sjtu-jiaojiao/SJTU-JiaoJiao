package main

import (
	"context"
	avatar "jiaojiao/srv/avatar/proto"
	"jiaojiao/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAvatarCreate(t *testing.T) {
	tf := func(userID int32, file string, status avatar.AvatarCreateResponse_Status) {
		var s srv
		var rsp avatar.AvatarCreateResponse
		So(s.Create(context.TODO(), &avatar.AvatarCreateRequest{
			UserID: userID,
			File:   utils.If(file != "", []byte(file), []byte{0}).([]byte),
		}, &rsp), ShouldBeNil)
		So(rsp.Status, ShouldEqual, status)
	}

	Convey("Test Avatar Create", t, func() {
		tf(0, "valid_file", avatar.AvatarCreateResponse_INVALID_PARAM)
		tf(1000, "", avatar.AvatarCreateResponse_INVALID_PARAM)
		tf(1000, "invalid_file", avatar.AvatarCreateResponse_INVALID_TYPE)
		tf(1000, "valid_file", avatar.AvatarCreateResponse_SUCCESS)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
