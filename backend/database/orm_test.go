package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrm(t *testing.T) {
	o := InitORM(new(Test))
	Convey("Load orm test", t, func() {
		test := Test{TestName: "tester"}

		// insert
		_, err := o.Insert(&test)
		So(err, ShouldEqual, nil)

		// update
		test.TestName = "jiang"
		_, err = o.Update(&test)
		So(err, ShouldEqual, nil)

		// read one
		u := Test{Id: test.Id}
		So(o.Read(&u), ShouldEqual, nil)

		// delete
		_, err = o.Delete(&u)
		So(err, ShouldEqual, nil)
	})
}
