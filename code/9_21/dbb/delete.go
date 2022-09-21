package dbb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Delete() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("error", err)

	}
	uc := client.Database("UserDB").Collection("usrc")
	dr, err := uc.DeleteOne(context.TODO(), bson.D{{"name", "niranjan"}})
	fmt.Println(dr.DeletedCount)
}
