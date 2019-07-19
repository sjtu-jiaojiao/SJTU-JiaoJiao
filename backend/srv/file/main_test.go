package main

import (
	"context"
	file "jiaojiao/srv/file/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSrvFileQuery(t *testing.T) {
	var s srvFile
	var req file.FileQueryRequest
	var rsp file.FileQueryResponse
	Convey("Test", t, func() {
		s.Query(context.TODO(), &req, &rsp)
	})
}

func TestMain(m *testing.M) {
	main()
	m.Run()
}
