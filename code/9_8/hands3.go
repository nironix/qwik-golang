package main

import "fmt"

func main() {
	var num int = 1
	fmt.Println("Enter num to check divisible")
	fmt.Scan(&num)
	for i := 1; i <= 50; i++ {
		if i%num == 0 {
			fmt.Print(" ", i)
		}
	}
	fmt.Println("\nEnter num to add 50")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		fmt.Print(" ", i+50)
	}
	rs := []int{45, 99, 75}
	sum := 0
	for _, j := range rs {
		sum += j
	}
	fmt.Println("\nSUM is ", sum)
}
