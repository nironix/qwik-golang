package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hellohandler)
	fil := http.FileServer(http.Dir("D:/GoLang/qwik-golang/code/9_20/dyn/page"))
	http.Handle("/", fil)
	http.HandleFunc("/form", formhandler)

	fmt.Println("starting server at 9090...")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("error", err)
	}

}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}
	iname := r.FormValue("name")
	iage := r.FormValue("age")
	iorg := r.FormValue("org")
	iaddr := r.FormValue("addr")
	fmt.Fprintln(w, "name : ", iname, "\nage : ", iage,
		"\norg : ", iorg, "\naddress : ", iaddr)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "no", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "we are in hello")
}
