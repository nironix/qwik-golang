package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type stu_str struct {
	//ID   bson. `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Qual string `bson:"qual" json:"qual"`
	Age  string `bson:"age" json:"age"`
}

func main() {
	http.HandleFunc("/api/v1/student", stu)
	http.ListenAndServe(":9090", nil)

}

func stu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("error", err)
	}
	uc := client.Database("student").Collection("studata")

	var stud stu_str
	if err := json.NewDecoder(r.Body).Decode(&stud); err != nil {
		http.Error(w, "error", http.StatusNotFound)
	}
	cntx := context.Background()
	switch r.Method {
	case "GET":
		err = createRecord(w, uc, cntx, stud)

	case "POST":
		err = getrec(w, uc, cntx)
	case "PUT":
		err = updaterec(w, uc, cntx, stud)
	case "DELETE":
		err = delrec(w, uc, cntx, stud)
	}

}
func delrec(w http.ResponseWriter, uc *mongo.Collection, cntx context.Context, stud stu_str) error {
	del, err := uc.DeleteOne(cntx, bson.D{{"age", "21"}, {"qual", "MSC"}})
	if err != nil {
		fmt.Println(err)
		return err

	}
	fmt.Fprintln(w, stud, del.DeletedCount)
	return err
}
func updaterec(w http.ResponseWriter, uc *mongo.Collection, cntx context.Context, stud stu_str) error {

	fil := bson.D{{"age", stud.Age}}
	uf := bson.D{{"$set", bson.D{{"name", "nix"}}}}
	ru, err := uc.UpdateOne(cntx, fil, uf)

	if err != nil {
		fmt.Println(err)
		return err

	}
	fmt.Fprintln(w, ru.ModifiedCount)
	return err
}
func createRecord(w http.ResponseWriter, uc *mongo.Collection, cntx context.Context, stud stu_str) error {

	res, err := uc.InsertOne(cntx, stud)
	if err != nil {
		fmt.Println("error", err)

	}

	fmt.Fprintln(w, res.InsertedID)

	return err
}

func getrec(w http.ResponseWriter, uc *mongo.Collection, cntx context.Context) error {

	cur, err := uc.Find(cntx, bson.D{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	var res []bson.M

	for cur.Next(cntx) {
		var r bson.M
		if err = cur.Decode(&r); err != nil {
			return nil
		}
		res = append(res, r)

	}
	fmt.Fprintln(w, res)
	return nil

}
