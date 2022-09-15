package main

import (
	"fmt"
	"time"
)

func main() {

	res1 := make(chan string)

	go lol(res1)
	for {
		select {
		case op1 := <-res1:
			fmt.Println(op1)
			return
		default:
			fmt.Println("Signal Not Received")

		}
	}

}

func lol(cha chan string) {

	time.Sleep(4 * time.Second)
	cha <- "*****signal Received*****"

}
