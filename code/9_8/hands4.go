package main

import "fmt"

func main() {
	var day int
	fmt.Println("enter a number to get a day")
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("SUNDAY")
	case 2:
		fmt.Println("MONDAY")
	case 3:
		fmt.Println("TUESDAY")
	case 4:
		fmt.Println("WEDNESDAY")
	case 5:
		fmt.Println("THURSDAY")
	case 6:
		fmt.Println("FRIDAY")
	case 7:
		fmt.Println("SATURDAY")
	default:
		fmt.Println("Invalid INPUT")
	}

	var age int
	fmt.Println("enter age to check eligiblity for voting")
	fmt.Scan(&age)
	if age > 18 {
		goto eligiblee
	} else {
		goto no_eligible
	}
eligiblee:
	fmt.Println("eligible for voting")
no_eligible:
	fmt.Println("Not eligible for voting")
}
