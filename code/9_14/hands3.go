package main

import (
	"fmt"
	"time"
)

func main() {
	res1 := make(chan string)
	res2 := make(chan string)
	res3 := make(chan string)

	go ffunc(res1)
	go ffunc2(res2)
	go ffunc3(res3)
	select {
	case op1 := <-res1:
		fmt.Println(op1)
	case op2 := <-res2:
		fmt.Println(op2)
	case op3 := <-res3:
		fmt.Println(op3)
	}
}

func ffunc(cha chan string) {
	time.Sleep(3 * time.Second)
	cha <- "hello from 1"

}

func ffunc2(cha chan string) {
	time.Sleep(2 * time.Second)
	cha <- "hello from 2"
}
func ffunc3(cha chan string) {
	time.Sleep(1 * time.Second)
	cha <- "hello from 3"
}
