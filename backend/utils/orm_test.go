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
		o.Using("default") // 默认使用 default，你可以指定为其他数据库

		user := new(User)
		user.user_name = "jiaojiao"
		user.password = "123456"

		release := new(Release)
		release.user = user
		release.is_saled = true

		o.Read(user)
		fmt.Println(user)
		fmt.Println("test finished")
	})
}
