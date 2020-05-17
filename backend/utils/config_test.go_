package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetStringConfig(t *testing.T) {
	Convey("Consul string config test", t, func() {
		So(GetStringConfig("test", "test"), ShouldEqual, "Do_not_modify_this")
	})
	Convey("Consul string config test not exist", t, func() {
		So(GetStringConfig("test", "testnotexist"), ShouldEqual, "")
	})
}

func TestLoadLocalConfig(t *testing.T) {
	Convey("Local config load test", t, func() {
		So(LoadLocalConfig, ShouldNotPanic)
		So(LocalConf.Test, ShouldEqual, "Do_not_modify_this")
	})
}

func TestLoadConsulConfig(t *testing.T) {
	Convey("Consul config load test", t, func() {
		So(LoadConsulConfig, ShouldNotPanic)
		So(GetStringConfig("test", "test"), ShouldEqual, "Do_not_modify_this")
	})
}

func TestGetDeployHost(t *testing.T) {
	Convey("Deploy host load test", t, func() {
		So(GetDeployHost("test"), ShouldEqual, "Do_not_modify_this")
	})
}
