package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrm(t *testing.T) {
	InitORM("testdb", new(Test))
	Convey("Load orm test", t, func() {
		test := Test{TestName: "tester"}

		// insert
		_, err := Ormer.Insert(&test)
		So(err, ShouldBeNil)

		// update
		test.TestName = "jiang"
		_, err = Ormer.Update(&test)
		So(err, ShouldBeNil)

		// read one
		u := Test{Id: test.Id}
		So(Ormer.Read(&u), ShouldBeNil)

		// delete
		_, err = Ormer.Delete(&u)
		So(err, ShouldBeNil)
	})
}
