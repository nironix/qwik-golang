package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 string = "22"
	var str2 string = "22.25"

	i, _ := strconv.ParseInt(str1, 0, 32)
	j, _ := strconv.ParseFloat(str2, 0)
	fmt.Println(i)
	fmt.Println(j)
}
