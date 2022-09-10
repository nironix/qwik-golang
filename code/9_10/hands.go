package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("---q1---")
	var num int = 1
	fmt.Scan(&num)
	arr := [10]int{}
	for i := 0; i < num; i++ {
		arr[i] = i + 50
	}

	fmt.Println(arr)

	fmt.Println("\n---q2---")
	pp := []int{2, 4, 5, 1, 9, 16, 7, 82, 18, 45, 11, 78}
	var s = make([]int, 12)
	n := copy(s, pp)
	sort.Ints(s)
	fmt.Println(n, s)

	fmt.Println("\n---q3---")

	for _, j := range s {
		if j%2 == 0 {
			fmt.Print(j, " ")
		}
	}

	fmt.Println("\n---q4---")
	var nu int = 1
	fmt.Println("enter num to search :")
	fmt.Scan(&nu)
	flag := 0
	for _, j := range s {
		if j == nu {
			fmt.Print(j, " found ")
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Print(nu, " not found ")

	}

}
