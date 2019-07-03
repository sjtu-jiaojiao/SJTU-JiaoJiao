package utils

import (
	"jiaojiao/utils"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetConfig(t *testing.T) {
	So(utils.GetConfig("test", "test"), ShouldEqual, "Do_not_modify_this")
}
