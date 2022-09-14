package main

import (
	"fmt"
)

func main() {
	//fmt.Println("hhh")
	//var intro chan int
	//time.Sleep(1 * time.Second)

	intro := make(chan int)
	go baddd(intro)
	intro <- 20
	//time.Sleep(1 * time.Second)

	//fmt.Println(intro)
}

func baddd(inp chan int) {
	fmt.Println(100 + <-inp)

}
