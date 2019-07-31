package main

import (
	"context"
	message "jiaojiao/srv/message/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	var s srv
	var req message.MessageCreateRequest
	Convey("Test Create Message", t, func() {
		err := s.Create(context.TODO(), &req, &message.MessageCreateResponse{})
		ShouldBeNil(err)
	})
}

func TestQuery(t *testing.T) {
	var s srv
	var req message.MessageQueryRequest
	Convey("Test Query Message", t, func() {
		err := s.Query(context.TODO(), &req, &message.MessageQueryResponse{})
		ShouldBeNil(err)
	})

}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
