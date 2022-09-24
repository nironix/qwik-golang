package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dat struct {
	Name  string `bson:"name" json:"name"`
	Age   string `bson:"age" json:"age"`
	Email string `bson:"email" json:"email"`
	Addr  string `bson:"addr" json:"addr"`
	Qual  string `bson:"qual" json:"qual"`
	Fb    string `bson:"fb" json:"fb"`
}

var uc *mongo.Collection
var cntx context.Context

func main() {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("error", err)
	}
	cntx = context.Background()
	uc = client.Database("assessment").Collection("col1")
	fil := http.FileServer(http.Dir("D:/GoLang/qwik-golang/code/9_23/ass/page"))
	http.Handle("/", fil)
	http.HandleFunc("/fcreate", create)
	http.HandleFunc("/fread", read)
	http.HandleFunc("/fupdate", update)
	http.HandleFunc("/fdelete", delete)

	fmt.Println("starting server at 9090...")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("error", err)
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}

	del, err := uc.DeleteOne(cntx, bson.D{{"name", r.FormValue("name")}})
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Fprintln(w, "Deleted documents : ", del.DeletedCount)
	myvar := map[string]interface{}{"res": "Deleted Documents " + strconv.FormatInt(int64(del.DeletedCount), 10), "head": "DELETING RECORD"}
	outputHTML(w, "page/result.html", myvar)
}

func update(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}
	fil := bson.D{{"name", r.FormValue("name")}}
	uf := bson.D{{"$set", bson.D{{"email", r.FormValue("email")}}}}
	ru, err := uc.UpdateOne(cntx, fil, uf)
	if err != nil {
		fmt.Println(err)

	}
	//fmt.Fprintln(w, "Updated documents : ", ru.ModifiedCount)

	myvar := map[string]interface{}{"res": "Updated Documents " + strconv.FormatInt(int64(ru.ModifiedCount), 10), "head": "UPDATING RECORD"}
	outputHTML(w, "page/result.html", myvar)
}

func read(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}

	cur, err := uc.Find(cntx, bson.D{{"name", r.FormValue("name")}})
	if err != nil {
		fmt.Println(err)

	}

	var res []bson.M

	for cur.Next(cntx) {
		var r bson.M
		if err = cur.Decode(&r); err != nil {
			fmt.Println(err)
		}
		res = append(res, r)

	}
	//fmt.Fprintln(w, res)
	myvar := map[string]interface{}{"res": "Displaying data of " + r.FormValue("name"), "head": "READING RECORD"}
	outputHTML(w, "page/result.html", myvar)
}
func create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}
	//var f dat
	// f.Name = r.FormValue("name")
	// f.Age = r.FormValue("age")
	// f.Email = r.FormValue("email")
	// f.Addr = r.FormValue("addr")
	// f.Qual = r.FormValue("qual")
	// f.Fb = r.FormValue("fb")

	iname := r.FormValue("name")
	iage := r.FormValue("age")
	iem := r.FormValue("email")
	iaddr := r.FormValue("addr")
	iqual := r.FormValue("qual")
	ifb := r.FormValue("fb")
	// fmt.Fprintln(w, "name : ", iname, "\nage : ", iage,
	// 	"\norg : ", iem, "\naddress : ", iaddr,
	// 	"\norg : ", iqual, "\norg : ", ifb)
	f := dat{iname, iage, iem, iaddr, iqual, ifb}
	//st :=[]interface{}{ bson.D{{"name", iname}, {"age", iage}, {"email", iem}, {"addr", iaddr}, {"qual", iqual}, {"fn", ifb}}}
	//st=toDoc(f)
	//fmt.Println(f)
	_, err := uc.InsertOne(cntx, f)
	if err != nil {
		fmt.Println("error", err)
	}

	//fmt.Fprintln(w, res.InsertedID)
	myvar := map[string]interface{}{"res": "Inserted Documents ", "head": "DELETING RECORD"}
	outputHTML(w, "page/result.html", myvar)
}

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
