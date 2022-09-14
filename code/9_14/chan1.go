package main

import "fmt"

func main() {
	loop := make(chan string)

	go func() {
		loop <- "n1"
		loop <- "n2"
		loop <- "n3"
		loop <- "n4"
		loop <- "n5"
		loop <- "n6"
		close(loop)
	}()

	for i := range loop {
		fmt.Println(i)
	}
	loop1 := make(chan string, 8)
	loop1 <- "sjjsjs"
	fmt.Println(len(loop1))
}
