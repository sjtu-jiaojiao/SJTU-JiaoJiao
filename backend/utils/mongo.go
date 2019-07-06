package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoDatabase *mongo.Database
var mongoContext context.Context

func mongoLoad() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	mongoDatabase = client.Database("sjtujj")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

}
