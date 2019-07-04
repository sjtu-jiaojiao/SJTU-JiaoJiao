package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestError(t *testing.T) {
	Convey("Error log test", t, func() {
		Error("This is test error!")
	})
}

// Warning logs a message at warning level.
func TestWarning(t *testing.T) {
	Convey("Warning log test", t, func() {
		Warning("This is test warning!")
	})
}

// Info logs a message at information level.
func TestInfo(t *testing.T) {
	Convey("Info log test", t, func() {
		Info("This is test info!")
	})
}

func TestLogPanic(t *testing.T) {
	Convey("LogPanic nil test", t, func() {
		var e interface{}
		So(func() { LogPanic(e) }, ShouldNotPanic)
		So(func() { LogPanic(e, "123%d", 4) }, ShouldNotPanic)
	})
	Convey("LogPanic true test", t, func() {
		e := true
		So(func() { LogPanic(e) }, ShouldNotPanic)
		So(func() { LogPanic(e, "123%d", 4) }, ShouldNotPanic)
	})
	Convey("LogPanic false test", t, func() {
		e := false
		So(func() { LogPanic(e) }, ShouldPanicWith, false)
		So(func() { LogPanic(e, "123%d", 4) }, ShouldPanicWith, "1234")
	})
	Convey("LogPanic text test", t, func() {
		e := "Panic test"
		So(func() { LogPanic(e) }, ShouldPanicWith, e)
		So(func() { LogPanic(e, "123%d", 4) }, ShouldPanicWith, "1234")
	})
}
