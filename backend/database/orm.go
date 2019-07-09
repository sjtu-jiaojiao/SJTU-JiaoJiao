package db

import (
	"jiaojiao/utils"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// LoadORM load orm map
func LoadORM() {
	utils.LogPanic(orm.RegisterDriver("mysql", orm.DRMySQL), utils.Error)
	utils.LogPanic(orm.RegisterDataBase("default", "mysql",
		os.Getenv("JJ_MARIADBUSER")+":"+os.Getenv("JJ_MARIADBPWD")+"@/"+
			utils.GetStringConfig("sys_config", "maria_dbname")+
			"?"+utils.GetStringConfig("sys_config", "maria_arg")))
}

func InitORM(m ...interface{}) orm.Ormer {
	orm.RegisterModel(m...)
	utils.LogPanic(orm.RunSyncdb("default", utils.LocalConf.Deploy == "develop", false))
	return orm.NewOrm()
}
