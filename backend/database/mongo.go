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

var MongoDatabase *mongo.Database
var MongoContext context.Context

// LoadMongoDB init mongodb
func LoadMongoDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+os.Getenv("JJ_MONGODBUSER")+":"+os.Getenv("JJ_MONGODBPWD")+
		"@"+utils.GetStringConfig("sys_config", "mongo_uri")+"/"))

	utils.LogPanic(err)
	utils.LogPanic(client.Ping(ctx, readpref.Primary()))

	MongoDatabase = client.Database(utils.GetStringConfig("sys_config", "mongo_dbname"))
	MongoContext, _ = context.WithTimeout(context.Background(), 10*time.Second)
}