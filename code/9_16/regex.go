package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(".com")
	fmt.Println(re.FindString("niro@vit.com"))

	re = regexp.MustCompile("a([a-z]+)")
	fmt.Println(re.FindString("alphbet are good"))
}
