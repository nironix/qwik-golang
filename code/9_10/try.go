package main

import (
	"bytes"
	"fmt"
)

func main() {

	as := []int{1, 2, 3, 4, 5, 6}
	ps := as
	fmt.Println(as)
	as[1] = 5
	fmt.Println(cap(as))
	fmt.Println(ps)
	sum(as)

	/*
		for i := 0; i <= len(as); i++ {
			fmt.Println(as[i], " ", ps[i])
		}*/
	aa := [][]int{{1, 2}, {2, 3}, {3, 4}}
	for _, j := range aa {
		fmt.Println(j[0])
	}
	var s = make([]int, 3)
	n := copy(s, []int{0, 1, 2, 3})
	fmt.Println(n, s)
	a := bytes.Trim([]byte("lalal"), "a")
	fmt.Println(string(a))

}
func sum(arr []int) {
	sum := 0
	for _, j := range arr {
		sum += j
	}
	fmt.Println(sum)
}
