package main

import (
	"fmt"
	"net/url"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(time.Unix(now.Unix(), 0))

	s := "postgress://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

}
