package dbb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("error", err)

	}
	fmt.Println(client)
	uc := client.Database("UserDB").Collection("usrc")
	user := bson.D{{"name", "niro"}, {"age", 21}}
	res, err := uc.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("error", err)

	}
	fmt.Println(res.InsertedID)

	urs := []interface{}{
		bson.D{{"name", "niran"}, {"age", 22}},
		bson.D{{"name", "niranjan"}, {"age", 22}},
	}
	ress, err := uc.InsertMany(context.TODO(), urs)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(ress.InsertedIDs)
}
