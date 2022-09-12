package main

import (
	"fmt"
)

func main() {
	//var mapp map[int]string

	mapp := map[int]string{
		1: "mee",
		2: "hddhd",
	}
	mapp[3] = "jdjd"
	fmt.Println(mapp)
	delete(mapp, 2)

	for i, j := range mapp {
		fmt.Println(i, " ", j)
	}

}
