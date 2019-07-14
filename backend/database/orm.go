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
}

func InitORM(dbName string, m ...interface{}) orm.Ormer {
	utils.LogPanic(orm.RegisterDataBase("default", "mysql",
		os.Getenv("JJ_MARIADBUSER")+":"+os.Getenv("JJ_MARIADBPWD")+"@"+
			utils.GetStringConfig("srv_config", dbName,utils.LocalConf.Deploy)+
			utils.GetStringConfig("srv_config", dbName, "suffix")))
	orm.RegisterModel(m...)
	utils.LogPanic(orm.RunSyncdb("default", utils.LocalConf.Deploy == "develop", false))
	return orm.NewOrm()
}
