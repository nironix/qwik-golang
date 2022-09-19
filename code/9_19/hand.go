package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("port open 9090")
	http.ListenAndServe(":9090", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w)
}
