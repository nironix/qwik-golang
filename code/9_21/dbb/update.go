package dbb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Update() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("error", err)

	}
	uc := client.Database("UserDB").Collection("usrc")
	uf := bson.D{{"$set", bson.D{{"name", "nix"}}}}
	ru, err := uc.UpdateOne(context.TODO(), bson.D{{"age", 21}}, uf)
	fmt.Println(ru.ModifiedCount)
}
