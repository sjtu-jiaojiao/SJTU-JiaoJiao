package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadLog(t *testing.T) {
	Convey("Load log test", t, func() {
		So(LoadLog, ShouldNotPanic)
	})
}

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

func TestLog(t *testing.T) {
	tf := func(e interface{}, f func(f interface{}, v ...interface{}), v ...interface{}) string {
		_, s := LogContinueS(e, f, v...)
		return s
	}
	Convey("Log nil test", t, func() {
		var e interface{}
		So(func() { LogPanic(e) }, ShouldNotPanic)
		So(func() { LogPanic(e, "123%d", 4) }, ShouldNotPanic)
		So(tf(e, Error), ShouldEqual, "")
		So(tf(e, Error, "123%d", 4), ShouldEqual, "")
		So(LogContinue(e, Error), ShouldEqual, false)
		So(LogContinue(e, Error, "123%d", 4), ShouldEqual, false)
	})
	Convey("Log true test", t, func() {
		e := true
		So(func() { LogPanic(e) }, ShouldNotPanic)
		So(func() { LogPanic(e, "123%d", 4) }, ShouldNotPanic)
		So(tf(e, Error), ShouldEqual, "")
		So(tf(e, Error, "123%d", 4), ShouldEqual, "")
		So(LogContinue(e, Error), ShouldEqual, false)
		So(LogContinue(e, Error, "123%d", 4), ShouldEqual, false)
	})
	Convey("Log false test", t, func() {
		e := false
		So(func() { LogPanic(e) }, ShouldPanicWith, "false")
		So(func() { LogPanic(e, "123%d", 4) }, ShouldPanicWith, "1234")
		So(tf(e, Error), ShouldEqual, "false")
		So(tf(e, Error, "123%d", 4), ShouldEqual, "1234")
		So(LogContinue(e, Error), ShouldEqual, true)
		So(LogContinue(e, Error, "123%d", 4), ShouldEqual, true)
	})
	Convey("Log text test", t, func() {
		e := "Panic test"
		So(func() { LogPanic(e) }, ShouldPanicWith, e)
		So(func() { LogPanic(e, "123%d", 4) }, ShouldPanicWith, "1234")
		So(tf(e, Error), ShouldEqual, "Panic test")
		So(tf(e, Error, "123%d", 4), ShouldEqual, "1234")
		So(LogContinue(e, Error), ShouldEqual, true)
		So(LogContinue(e, Error, "123%d", 4), ShouldEqual, true)
	})
}
