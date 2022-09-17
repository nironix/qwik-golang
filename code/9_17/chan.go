package main

import "fmt"

func main() {
	// ub := make(chan string)
	// ub <- "ubb"
	// fmt.Println(<-ub)

	// bu := make(chan string, 2)
	// bu <- "buff"
	// fmt.Println(<-bu)

	pi := make(chan string)
	po := make(chan string)
	go pgs(pi, "hello")
	go pos(pi, po)
	fmt.Println(<-po)
}

func pgs(pi chan<- string, m string) {
	pi <- m
}
func pos(pi chan string, po chan<- string) {
	msg := <-pi
	po <- msg
}
