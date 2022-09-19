package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	fil := http.FileServer(http.Dir("D:/GoLang/qwik-golang/code/9_19/page/stat"))
	http.Handle("/stat", fil)
	fmt.Println("started at 8090....")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Println("err : ", err)
	}

}

/*func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello there")
}*/

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		//http.Error(w, "404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "no", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "we are in")
}
