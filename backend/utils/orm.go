package utils

import (
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Test is db test table map
type Test struct {
	ID       int    `orm:"auto;pk;column(user_id)"`
	TestName string `orm:"size(100)"`
}

var db orm.Ormer

// LoadORM load orm map
func LoadORM() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", os.Getenv("JJ_MARIADBUSER")+":"+os.Getenv("JJ_MARIADBPWD")+"@/"+
		GetStringConfig("sys_config", "maria_dbname")+"?"+GetStringConfig("sys_config", "maria_arg"))

	orm.RegisterModel(new(Test))
	orm.RunSyncdb("default", false, false)

	db = orm.NewOrm()
}
