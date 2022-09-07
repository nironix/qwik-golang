package main

import (
	"fmt"
)

const co = 50

func main() {
	const co1 = 60
	//to find default value
	var i bool
	var j int
	var k string = "igjyfs"
	var l float32
	var str string = "wednesday"
	fmt.Println("bool : ", i)
	fmt.Println("string : ", k)
	fmt.Println("int : ", j)
	fmt.Println("float : ", l)
	fmt.Printf("%T", k) //prints data type
	fmt.Printf("%s", k) //prints string
	type newt int
	var p newt
	fmt.Printf("%T\n", p)
	fmt.Println(str + " " + k)
}
