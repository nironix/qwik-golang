package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type sal struct {
	Basic int `json:"basic"`
	Hra   int `json:"hra"`
	Da    int `json:"da"`
}

type det struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Id     int    `json:"id"`
	Salary sal    `json:"salary"`
}

func main() {
	fmt.Println("---q1---")
	user := det{"niro", 22, 110, sal{25000, 2000, 3000}}
	fmt.Println("data :\n", user)
	user.Name = "Niranjan"
	//fmt.Println(str)
	fmt.Println("\n before entering file\n", user)
	fmt.Println("---q2---")

	str, _ := json.Marshal(user)              //to string format
	_ = ioutil.WriteFile("lo.txt", str, 0644) //write in file
	fc, _ := ioutil.ReadFile("lo.txt")        // read file
	fmt.Println("\ngetting from file\n", string(fc))
	var jmap det
	_ = json.Unmarshal(fc, &jmap) //to json format
	jmap.Name = "niro"

	nstr, _ := json.Marshal(jmap)              //to string format
	_ = ioutil.WriteFile("lo.txt", nstr, 0644) //write in file again
	fc, _ = ioutil.ReadFile("lo.txt")          // read file
	fmt.Println("\ngetting from file\n", string(fc))

}
