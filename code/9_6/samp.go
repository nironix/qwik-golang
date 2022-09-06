package main

import (
	"fmt"
	"reflect"
)

var i int = 10

func main() {
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))
	a, _ := lol()
	fmt.Println(a)
}
func lol() (string, int) {
	return "a", 1
}
