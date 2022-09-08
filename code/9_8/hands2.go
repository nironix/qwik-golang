package main

import "fmt"

func main() {
	var a int = 3
	var b int = 6

	fmt.Println("---a---")
	if a%b == 0 && b%a == 0 {
		fmt.Println("divisible by each other")
	}
	fmt.Println("---b & c---")
	var n int = 202
	if n%2 == 0 {
		fmt.Print(n, " is EVEN")
		if n > 200 {
			fmt.Print(" & Greater than 200")
		}
	} else {
		fmt.Print(n, " is ODD")
		if n > 200 {
			fmt.Print(" & Greater than 200")
		}
	}
	fmt.Println("\n---d---")
	var mark int
	fmt.Scan(&mark)
	if mark > 90 {
		fmt.Println("S Grade")
	} else if mark > 80 {
		fmt.Println("A Grade")
	} else if mark > 70 {
		fmt.Println("B Grade")
	} else if mark > 60 {
		fmt.Println("C Grade")
	} else if mark > 50 {
		fmt.Println("D Grade")
	} else {
		fmt.Println("F Grade")
	}

}
