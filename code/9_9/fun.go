package main

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("Welcome ")
}

type numval int

func main() {
	i, j, k := prin(5, 6)
	fmt.Println(i, " ", j, " ", k)
	fmt.Println(join_str("hhhd", "ssjjs", "sjsjs"))

	x := numval(10)
	y := numval(20)
	res := x.mult(y)
	fmt.Println(res)

}

func prin(i, j int) (int, int, int) {
	fmt.Println("done", i)
	return i * i, i + i, i + j
}
func init() {
	fmt.Println("good morning")
}

func join_str(strs ...string) string {
	return strings.Join(strs, "-")
}

func (d1 numval) mult(d2 numval) numval {
	return d1 * d2
}
