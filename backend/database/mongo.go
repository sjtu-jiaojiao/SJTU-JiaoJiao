package db

import (
	"context"
	"jiaojiao/utils"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDatabase is mongodb database
var MongoDatabase *mongo.Database

// InitMongoDB init mongodb
func InitMongoDB(dbName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+os.Getenv("JJ_MONGODBUSER")+":"+os.Getenv("JJ_MONGODBPWD")+
		"@"+utils.GetStringConfig("srv_config", dbName, utils.LocalConf.Deploy)+"/"))

	utils.LogPanic(err)
	utils.LogPanic(client.Ping(ctx, readpref.Primary()))

	MongoDatabase = client.Database(utils.GetStringConfig("srv_config", dbName, "dbname"))
}
