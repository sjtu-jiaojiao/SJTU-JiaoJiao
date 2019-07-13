package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrm(t *testing.T) {
	o := InitORM("testdb",new(Test))
	Convey("Load orm test", t, func() {
		test := Test{TestName: "tester"}

		// insert
		_, err := o.Insert(&test)
		So(err, ShouldBeNil)

		// update
		test.TestName = "jiang"
		_, err = o.Update(&test)
		So(err, ShouldBeNil)

		// read one
		u := Test{Id: test.Id}
		So(o.Read(&u), ShouldBeNil)

		// delete
		_, err = o.Delete(&u)
		So(err, ShouldBeNil)
	})
}
