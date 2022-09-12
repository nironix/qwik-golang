package main

import "fmt"

type inn struct {
	gg  intro
	num int
	yy  *int
}

type intro struct {
	age  int
	name string
}

func main() {
	var x int
	y := inn{
		intro{age: 8, name: "niro"}, 10, &x,
	}
	x = y.gg.age
	fmt.Println(x)
	t := struct {
		a int
		b int
	}{
		5, 10,
	}

	fmt.Println(t)

}
