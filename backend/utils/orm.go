package utils

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id           int
	user_name    string
	password     string
	portrait_url string
	telephone    string
	release      []*Release `orm:"reverse(many)"`
}

type Release struct {
	Id       int
	is_saled bool
	user     *User `orm:"rel(fk)"`
}

func ormLoad() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:khq799385272@/sjtujj?charset=utf8&loc=Asia%2FShanghai")

	orm.RegisterModel(new(User), new(Release))
}
