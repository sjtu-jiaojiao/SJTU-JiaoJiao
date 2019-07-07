package utils

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoDatabase *mongo.Database
var mongoContext context.Context

// LoadMongoDB init mongodb
func LoadMongoDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+os.Getenv("JJ_MONGODBUSER")+":"+os.Getenv("JJ_MONGODBPWD")+
		"@"+GetStringConfig("sys_config", "mongo_uri")+"/"))

	LogPanic(err)
	LogPanic(client.Ping(ctx, readpref.Primary()))

	mongoDatabase = client.Database(GetStringConfig("sys_config", "mongo_dbname"))
	mongoContext, _ = context.WithTimeout(context.Background(), 10*time.Second)
}
