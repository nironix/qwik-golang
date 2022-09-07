package main

import (
	"fmt"
	"reflect"
)

var date int = 101

func main() {
	fmt.Println("q1")
	name := "niro"
	age := 22
	city := "vellore"
	fmt.Printf("Name : ", name, "  %T", name)
	fmt.Print("age : ", age, " ", reflect.TypeOf(age), "\ncity : ", city, " ", reflect.TypeOf(city))

	fmt.Println("\n\nq2")

	type dat string
	var day dat
	fmt.Print("day : ", date, " ", reflect.TypeOf(day))
	day = "wednesday"
	fmt.Print("\nday : ", day)
	fmt.Println("date : ", date, " ", reflect.TypeOf(date))
	//day = "wednesday"
	//fmt.Print("day : ", day)

	fmt.Println("\n\nq3")
	a := 10
	b := 2
	fmt.Println("Arithmetic operators")
	fmt.Println("addition : ", a+b)
	fmt.Println("difference : ", a-b)
	fmt.Println("product :  ", a*b)
	fmt.Println("Relational operators")
	fmt.Println("equal :  ", a == b)
	fmt.Println("greater :  ", a > b)
	fmt.Println("lesser :  ", a < b)
	op1 := true
	op2 := false
	fmt.Println("Logical operators")
	fmt.Println("AND :  ", op1 && op2)
	fmt.Println("OR :  ", op1 || op2)
	fmt.Println("NOT :  ", !op1)
	num := 5
	fmt.Println("assignment operators: num=5")
	num += 2
	fmt.Println("+=2 :  ", num)
	num -= 3
	fmt.Println("-=3 :  ", num)
	num *= 4
	fmt.Println("*=4 :  ", num)

	fmt.Println("\n\nq4")
	t1 := 9
	t2 := 6
	fmt.Println("t1 : ", t1, " t2 : ", t2)
	t2, t1 = t1, t2
	fmt.Println("t1 : ", t1, " t2 : ", t2)

}
