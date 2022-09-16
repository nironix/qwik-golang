package main

import "fmt"

type sal struct {
	basic float64
	tax   float64
}
type total interface {
	calc() float64
}

func main() {

	var bas float64
	var tax float64
	fmt.Println("enter basic : ")
	fmt.Scan(&bas)
	fmt.Println("enter tax : ")
	fmt.Scan(&tax)
	var s total
	s = sal{bas, tax}
	fmt.Println("Total Salary : ", s.calc())

}

func (m sal) calc() float64 {
	return m.basic + m.tax
}
