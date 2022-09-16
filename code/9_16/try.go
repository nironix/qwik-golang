package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hey\nla"
	b := `li\nlo`
	fmt.Println(a)
	fmt.Println(b)
	str := "aA"
	for i := range str {
		fmt.Println(str[i])
	}

	fmt.Println(len(str))
	fmt.Println(strings.Compare("A", "a"))
	res := fmt.Sprintf("%s %s", str, str)
	fmt.Println(res)
	arr := []string{"hi", "hlo", "bye"}
	fmt.Println(strings.Join(arr, " "))
	str = "hello there bye"

	fmt.Println(strings.Repeat(str, 3))

	fmt.Println(strings.Index(str, "th"))
}
