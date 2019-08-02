package db

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
)

func TestLoadMongoDB(t *testing.T) {
	InitMongoDB("testmongo")
	collection := MongoDatabase.Collection("test")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	Convey("Insert mongodb test", t, func() {
		_, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		So(err, ShouldBeNil)
	})
	Convey("Select mongodb test", t, func() {
		cur, err := collection.Find(ctx, bson.D{})
		So(err, ShouldBeNil)

		for cur.Next(ctx) {
			var result bson.M
			So(cur.Decode(&result), ShouldBeNil)
		}
		So(cur.Err(), ShouldBeNil)
	})
	Convey("Delete mongodb test", t, func() {
		_, err := collection.DeleteMany(ctx, bson.D{{"name", "pi"}})
		So(err, ShouldBeNil)
	})
}
