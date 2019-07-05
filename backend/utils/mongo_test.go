package utils

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMongoLoad(t *testing.T) {
	Convey("Load mongodb test", t, func() {
		collection := mongoDatabase.Collection("picture")
		res, err := collection.InsertOne(mongoContext, bson.M{"name": "pi", "value": 3.14159})
		if err != nil {
			Error("Insert into mongodb failed.")
		}

		id := res.InsertedID
		fmt.Println(id)
	})
}
