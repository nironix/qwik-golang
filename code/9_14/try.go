package main

import (
	"fmt"
	"time"
)

func tto(ss string) {
	for i := 0; i < 5; i++ {
		fmt.Println(ss)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	//go tto("one")
	fmt.Println("sksk")
	//go tto("two")

	time.Sleep(1 * time.Second)

	go func() {
		fmt.Println("ksks")
	}()

}
