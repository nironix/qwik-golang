package main

import "fmt"

//goto and continue and switch

func main() {

	/*
	   	// goto
	   	var i int = 0
	   	for {
	   		i++
	   		if i == 5 {
	   			goto ke
	   		}
	   	}
	   ke:
	   	fmt.Println(i)*/

	/*
		//continue
		for i := 0; i < 9; i++ {
			if i%2 == 0 {
				continue
			}
			fmt.Println(i)
		}*/

	i := 0
	switch i {
	case 1:
		fmt.Println("true")
	case 0:
		fmt.Println("false")
	default:
		fmt.Println("Invalid")
	}
}
