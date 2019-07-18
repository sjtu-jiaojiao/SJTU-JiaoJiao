package db

import (
	"fmt"
	"jiaojiao/utils"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Ormer *gorm.DB

func InitORM(dbName string, m ...interface{}) {
	conn := fmt.Sprintf("%s:%s@%s%s", os.Getenv("JJ_MARIADBUSER"), os.Getenv("JJ_MARIADBPWD"),
		utils.GetStringConfig("srv_config", dbName, utils.LocalConf.Deploy),
		utils.GetStringConfig("srv_config", dbName, "suffix"))
	var err error
	Ormer, err = gorm.Open("mysql", conn)
	utils.LogPanic(err)

	if utils.LocalConf.Deploy == "develop" {
		Ormer.DropTableIfExists(m...)
	}
	Ormer.AutoMigrate(m...)
}

func CloseORM() {
	if utils.LocalConf.Deploy != "develop" {
		utils.LogPanic(Ormer.Close())
	}
}
