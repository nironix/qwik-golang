package main

import "fmt"

func main() {
	swi(20)
}

func swi(e interface{}) {
	switch e.(type) {
	case int:
		fmt.Println("int ", e.(int))
	case string:
		fmt.Println("string ", e.(string))
	default:
		fmt.Println("nope")
	}
}
