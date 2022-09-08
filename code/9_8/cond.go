package main

import "fmt"

func main() {
	var i int
	fmt.Println("enter a number : ")
	fmt.Scan(&i)
	if i == 0 {
		fmt.Println("NUMBER IS 0")

	} else if i%2 == 0 {
		fmt.Println("EVEN NUMBER")
		if i > 5 {
			fmt.Println("greater than 5")
		}
	} else {
		fmt.Println("ODD NUMBER")
	}

}
