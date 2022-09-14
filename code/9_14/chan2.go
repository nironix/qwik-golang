package main

import (
	"fmt"
	"time"
)

func main() {
	res1 := make(chan string)
	res2 := make(chan string)

	go ffunc(res1)
	go ffunc2(res2)

	select {
	case op1 := <-res1:
		fmt.Println(op1)
	case op2 := <-res2:
		fmt.Println(op2)
	}
}

func ffunc(cha chan string) {
	time.Sleep(4 * time.Second)
	cha <- "hello"

}

func ffunc2(cha chan string) {
	time.Sleep(2 * time.Second)
	cha <- "hello123"
}
