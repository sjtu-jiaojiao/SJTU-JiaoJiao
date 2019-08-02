package db

import (
	"fmt"
	"jiaojiao/utils"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql orm
)

// Ormer is opened orm object
var Ormer *gorm.DB

// InitORM init the orm object
func InitORM(dbName string, m ...interface{}) {
	conn := fmt.Sprintf("%s:%s@%s%s", os.Getenv("JJ_MARIADBUSER"), os.Getenv("JJ_MARIADBPWD"),
		utils.GetStringConfig("srv_config", dbName, utils.LocalConf.Deploy),
		utils.GetStringConfig("srv_config", dbName, "suffix"))
	var err error
	Ormer, err = gorm.Open("mysql", conn)
	utils.LogPanic(err)

	Ormer.AutoMigrate(m...)

	Ormer.DB().SetMaxIdleConns(0)
	Ormer.DB().SetMaxOpenConns(256)
	Ormer.DB().SetConnMaxLifetime(time.Second * 16)
}

// CloseORM close orm object
func CloseORM() {
	if utils.LocalConf.Deploy != "develop" {
		utils.LogPanic(Ormer.Close())
	}
}
