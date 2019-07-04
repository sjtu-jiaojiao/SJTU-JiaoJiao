package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetConfig(t *testing.T) {
	Convey("Consul config test", t, func() {
		So(GetConfig("test", "test"), ShouldEqual, "Do_not_modify_this")
	})
}
