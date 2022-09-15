package main

import (
	"fmt"
)

func main() {
	// _, err := os.Open("ttry.go")
	// if err != nil {
	// 	panic("ksks")
	// } else {
	// 	fc, _ := ioutil.ReadFile("try.go") //Read a file
	// 	fmt.Println(string(fc))
	// }

	fmt.Println(divv(10, 10))

	fmt.Println(divv(10, 0))
}

func divv(i, j int) int {
	defer func() {
		fmt.Println(recover())
	}()
	q := i / j
	return q
}
