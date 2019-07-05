package utils

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOrmLoad(t *testing.T) {
	Convey("Load orm test", t, func() {
		o := orm.NewOrm()

		user := User{UserName: "slene"}
	
		// insert
		id, err := o.Insert(&user)
		fmt.Printf("ID: %d, ERR: %v\n", id, err)
	
		// update
		user.UserName = "astaxie"
		num, err := o.Update(&user)
		fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	
		// read one
		u := User{Id: user.Id}
		err = o.Read(&u)
		fmt.Printf("ERR: %v\n", err)
	
		// delete
		num, err = o.Delete(&u)
		fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	})
}
