package dbb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Read() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("error", err)

	}
	uc := client.Database("UserDB").Collection("usrc")
	rd, err := uc.Find(context.TODO(), bson.D{})

	var res []bson.M

	if err := rd.All(context.TODO(), &res); err != nil {
		fmt.Println(err)
	}

	for _, d := range res {
		fmt.Println(d)
	}
	fil := bson.D{{"name", "niro"}}
	rf, err := uc.Find(context.TODO(), fil)

	if err := rf.All(context.TODO(), &res); err != nil {
		fmt.Println(err)
	}
	for _, d := range res {
		fmt.Println(d)
	}

	var res3 bson.M
	if err := uc.FindOne(context.TODO(), fil).Decode(&res3); err != nil {
		fmt.Println(err)
	}

}
