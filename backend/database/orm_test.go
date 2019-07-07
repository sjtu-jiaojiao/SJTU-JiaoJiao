package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrmLoad(t *testing.T) {
	Convey("Load orm test", t, func() {
		test := Test{TestName: "tester"}

		// insert
		_, err := dbo.Insert(&test)
		So(err, ShouldEqual, nil)

		// update
		test.TestName = "jiang"
		_, err = dbo.Update(&test)
		So(err, ShouldEqual, nil)

		// read one
		u := Test{Id: test.Id}
		So(dbo.Read(&u), ShouldEqual, nil)

		// delete
		_, err = dbo.Delete(&u)
		So(err, ShouldEqual, nil)
	})
}
