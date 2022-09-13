package main

import (
	"encoding/json"
	"fmt"
)

type tt struct {
	A int    `json:"n1"`
	B string `json:"n2"`
	C int    `json:"n3"`
}

func main() {
	mapp := map[int]string{
		1: "lala",
		2: "kaka",
		3: "jaja",
		4: "haha",
	}
	fmt.Println(mapp)
	aa, _ := json.Marshal(mapp)
	fmt.Println(string(aa))

	att := tt{1, "2", 3}
	var e error
	//btt, _ := json.MarshalIndent(att, " ", "\t")
	btt, _ := json.Marshal(att)
	fmt.Println(e)
	fmt.Println(string(btt))

	back := `{"n1":1,"n2":"2","n3":3}` //or use previously created
	var jss tt

	e = json.Unmarshal([]byte(back), &jss)
	fmt.Println(e)
	fmt.Println(jss)
	var jmap map[int]string
	_ = json.Unmarshal([]byte(aa), &jmap)
	fmt.Println(jmap)

}
