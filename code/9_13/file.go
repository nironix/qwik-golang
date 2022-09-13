package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	s := "lalal%ahiyfiyl"
	//fs, oc := os.Open("lo.txt")
	// if oc != nil {
	// 	fmt.Println("fail")
	// }
	oo := ioutil.WriteFile("lo.txt", []byte(s), 8522)
	fmt.Println(oo)
	fc, _ := ioutil.ReadFile("lo.txt") //Read a file
	fmt.Println(string(fc))

}
