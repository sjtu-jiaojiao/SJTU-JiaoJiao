package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrmLoad(t *testing.T) {
	Convey("Load orm test", t, func() {
		test := Test{TestName: "tester"}

		// insert
		_, err := Orm.Insert(&test)
		So(err, ShouldEqual, nil)

		// update
		test.TestName = "jiang"
		_, err = Orm.Update(&test)
		So(err, ShouldEqual, nil)

		// read one
		u := Test{Id: test.Id}
		So(Orm.Read(&u), ShouldEqual, nil)

		// delete
		_, err = Orm.Delete(&u)
		So(err, ShouldEqual, nil)
	})
}
