package main

import "fmt"

func main() {

	fmt.Println("---q1---")
	sum := vadd(1, 2, 3, 4, 5)
	fmt.Println(sum)

	fmt.Println("\n---q2---")
	var num int
	fmt.Println("Enter number to find factorial : ")
	fmt.Scan(&num)
	fmt.Println("Factorial of ", num, " is ", fact(num))

	fmt.Println("\n---q3---")
	var l int
	var b int
	fmt.Println("Enter length : ")
	fmt.Scan(&l)
	fmt.Println("Enter breadth(0 for square) : ")
	fmt.Scan(&b)
	as, ar := area(l, b)
	fmt.Println("Area of square : ", as)
	fmt.Println("Area of Rectangle : ", ar)

	fmt.Println("\n---q5---")
	var num1 int = 6
	var num2 int = 9

	fmt.Println("num1 : ", num1, " num2 : ", num2)
	swap(&num1, &num2)
	fmt.Println("num1 : ", num1, " num2 : ", num2)
}

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func area(l int, b int) (int, int) {
	return l * l, l * b
}

func vadd(nums ...int) int {
	sum := 0
	for _, j := range nums {
		sum += j
	}
	return sum
}

func fact(n int) int {
	f, i := 1, 1
	for i <= n {
		f = f * i
		i++
	}
	return f
}
