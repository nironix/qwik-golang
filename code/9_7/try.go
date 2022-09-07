package main

import (
	"fmt"
	"reflect"
)

func main() {
	//to find default value
	var i bool
	var j int
	var k string = "igjyfs"
	var l float32
	fmt.Println("bool : ", i)
	fmt.Println("string : ", k)
	fmt.Println("int : ", j)
	fmt.Println("float : ", l)
	fmt.Printf("%T", k) //prints data type
	fmt.Printf("%s", k)
	type newt int
	var p newt
	fmt.Println(reflect.TypeOf(p))
	fmt.Printf("%T", p)
}
