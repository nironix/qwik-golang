package main

import "fmt"

type definter interface {
	meth1() int
	//meth2() string
}

type val struct {
	val1 int
	val2 int
}

func main() {

	var m definter
	m = val{10, 20}
	fmt.Println(m.meth1())
	fmt.Printf("%T", m)
}

func (m val) meth1() int {
	return m.val1 + m.val2
}
