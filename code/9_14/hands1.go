package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	odd := make(chan int)
	even := make(chan int)

	go eve(even)
	go od(odd)

	for _, i := range arr {

		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

}

func eve(cha chan int) {
	time.Sleep(1 * time.Second)
	for i := range cha {
		fmt.Println("eve ", i)
	}
}

func od(cha chan int) {
	time.Sleep(4 * time.Second)
	for i := range cha {
		fmt.Println("odd ", i)
	}
}
