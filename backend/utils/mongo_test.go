package utils

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMongoLoad(t *testing.T) {
	Convey("Insert mongodb test", t, func() {
		collection := mongoDatabase.Collection("test")
		res, err := collection.InsertOne(mongoContext, bson.M{"name": "pi", "value": 3.14159})

		if err != nil {
			Error("Insert into mongodb failed.")
		}

		id := res.InsertedID
		fmt.Println("insert into database test of id: ", id)
	})
	Convey("Select mongodb test", t, func(){
		collection := mongoDatabase.Collection("test")
		cur, err := collection.Find(mongoContext, bson.D{})
		if err != nil { 
			Error("Select from mongodb failed.")
		}

		for cur.Next(mongoContext) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil { 
				Error("Select from mongodb failed.")
			}	
			fmt.Println(result)
		}

		if err := cur.Err(); err != nil {
			Error("Select from mongodb failed.")
		}
	})
	Convey("Delete mongodb test", t, func() {
		collection := mongoDatabase.Collection("test")
		res, err := collection.DeleteMany(mongoContext, bson.D{{"name", "pi"},},)

		if err != nil {
			Error("Insert into mongodb failed.")
		}

		fmt.Println("insert into database test of id: ", res)
	})
}
