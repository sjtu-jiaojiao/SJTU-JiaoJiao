package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAssignNotEmpty(t *testing.T) {
	Convey("Assign function test", t, func() {
		s1 := "S1_Test_String"
		s2 := ""
		AssignNotEmpty(&s2, &s1)
		So(s1, ShouldNotEqual, s2)
		s2 = "S2_Test_String"
		AssignNotEmpty(&s2, &s1)
		So(s1, ShouldEqual, s2)
	})
}

func TestAssignNotZero(t *testing.T) {
	Convey("Assign function test1", t, func() {
		var s1, s2 int32
		s1 = 123
		s2 = 0
		AssignNotZero(&s2, &s1)
		So(s1, ShouldNotEqual, s2)
		s2 = 456
		AssignNotZero(&s2, &s1)
		So(s1, ShouldEqual, s2)
	})
	Convey("Assign function test2", t, func() {
		var s1, s2 float32
		s1 = 123.45
		s2 = 0
		AssignNotZero(&s2, &s1)
		So(s1, ShouldNotEqual, s2)
		s2 = 456.78
		AssignNotZero(&s2, &s1)
		So(s1, ShouldEqual, s2)
	})
	Convey("Assign function test3", t, func() {
		var s1, s2 float64
		s1 = 123.45
		s2 = 0
		AssignNotZero(&s2, &s1)
		So(s1, ShouldNotEqual, s2)
		s2 = 456.78
		AssignNotZero(&s2, &s1)
		So(s1, ShouldEqual, s2)
	})
	Convey("Assign function test4", t, func() {
		var s1, s2 int
		s1 = 123
		s2 = 0
		So(func() { AssignNotZero(&s2, &s1) }, ShouldPanic)
	})
}
