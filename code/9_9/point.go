package main

import "fmt"

func main() {
	var x int = 100
	var y *int
	y = &x

	fmt.Println(y)
	fmt.Printf("%T", y)
}
