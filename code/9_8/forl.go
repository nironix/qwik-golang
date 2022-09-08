package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	/*
		//regular for loop
		for i := 0; i < n; i++ {
			fmt.Println(i)
		}*/

	/*
		//infinite for loop
		i := 0
		for {
			fmt.Println(i)
			if i > n {
				break
			}
			i++
		}
	*/
	/*
		//while loop
		i:=0
		for i < n {
			fmt.Println(i)
			i++
		}
	*/
	//rs := []string{"a", "b", "c", "d"}
	rs := "kdkdkd"
	for i, j := range rs {
		fmt.Println(i, " ", rune(j))
	}
}
