package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrmLoad(t *testing.T) {
	Convey("Load orm test", t, func() {
		test := Test{TestName: "tester"}

		// insert
		_, err := db.Insert(&test)
		So(err, ShouldEqual, nil)

		// update
		test.TestName = "jiang"
		_, err = db.Update(&test)
		So(err, ShouldEqual, nil)

		// read one
		u := Test{ID: test.ID}
		So(db.Read(&u), ShouldEqual, nil)

		// delete
		_, err = db.Delete(&u)
		So(err, ShouldEqual, nil)
	})
}
