package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int = 5
	var b int = 6
	fmt.Print("a : ", a, " ", reflect.TypeOf(a), "\nb : ", b, " ", reflect.TypeOf(b))

	var c float64 = (float64(a) + float64(b)) / 2
	fmt.Println(c)

}
