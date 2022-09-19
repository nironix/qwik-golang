package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	fmt.Println("response ", res.Status)

	sc := bufio.NewScanner(res.Body)
	for i := 0; sc.Scan() && i < 5; i++ {
		fmt.Println(sc.Text())
	}
}
