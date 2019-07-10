package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
)

func TestLoadMongoDB(t *testing.T) {
	Convey("Insert mongodb test", t, func() {
		collection := MongoDatabase.Collection("test")
		_, err := collection.InsertOne(MongoContext, bson.M{"name": "pi", "value": 3.14159})
		So(err, ShouldBeNil)
	})
	Convey("Select mongodb test", t, func() {
		collection := MongoDatabase.Collection("test")
		cur, err := collection.Find(MongoContext, bson.D{})
		So(err, ShouldBeNil)

		for cur.Next(MongoContext) {
			var result bson.M
			So(cur.Decode(&result), ShouldBeNil)
		}
		So(cur.Err(), ShouldBeNil)
	})
	Convey("Delete mongodb test", t, func() {
		collection := MongoDatabase.Collection("test")
		_, err := collection.DeleteMany(MongoContext, bson.D{{"name", "pi"}})
		So(err, ShouldBeNil)
	})
}
