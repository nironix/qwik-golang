package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("sleep")
	//time.Sleep(time.Second * 2)
	fmt.Println("woke up")
	tl, _ := time.Parse(time.RFC3339, "2022-01-02T12:22:22Z")
	loc, _ := time.LoadLocation("America/Los_Angeles")
	tl = tl.In(loc)
	fmt.Println(tl.Format(time.RFC822))

}
