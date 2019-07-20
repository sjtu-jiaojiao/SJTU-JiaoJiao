package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrm(t *testing.T) {
	InitORM("testdb", new(Test), new(User), new(SellInfo), new(Good))
	defer CloseORM()
	Convey("Load orm test", t, func() {
		test := Test{TestName: "tester"}

		// insert
		So(Ormer.Create(&test).Error, ShouldBeNil)

		// update
		test.TestName = "jiang"
		So(Ormer.Save(&test).Error, ShouldBeNil)

		// read one
		u := Test{ID: test.ID}
		So(Ormer.First(&u).Error, ShouldBeNil)

		// delete
		So(Ormer.Delete(&u).Error, ShouldBeNil)
	})
}
