package utils

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    Id   int `orm:"column(user_id)"`
    UserName string `orm:"size(100)"`
}

func ormLoad() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "sjtujj:sjtujj@/test?charset=utf8&loc=Asia%2FShanghai")

	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
