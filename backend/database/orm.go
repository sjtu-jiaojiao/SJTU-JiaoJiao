package db

import (
	"jiaojiao/utils"
	"os"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var Ormer orm.Ormer

// LoadORM load orm map
func LoadORM() {
	utils.LogPanic(orm.RegisterDriver("mysql", orm.DRMySQL), utils.Error)
	orm.DefaultTimeLoc = time.UTC
}

func InitORM(dbName string, m ...interface{}) {
	utils.LogPanic(orm.RegisterDataBase("default", "mysql",
		os.Getenv("JJ_MARIADBUSER")+":"+os.Getenv("JJ_MARIADBPWD")+"@"+
			utils.GetStringConfig("srv_config", dbName, utils.LocalConf.Deploy)+
			utils.GetStringConfig("srv_config", dbName, "suffix")))
	orm.RegisterModel(m...)
	utils.LogPanic(orm.RunSyncdb("default", utils.LocalConf.Deploy == "develop", false))
	Ormer = orm.NewOrm()
}
