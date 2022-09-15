package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	s, e := sqr(-5)
	if e != nil {
		fmt.Println("error : ", e)
	} else {
		fmt.Println("Result", s)
	}

	prnt1("lol")
	defer prnt2("lolo")
}

func prnt1(s string) {
	fmt.Println("first")
}

func prnt2(s string) {

	defer prnt1("j")
	fmt.Println("second")
}

func sqr(val float64) (float64, error) {
	if val < 0 {
		return 0, errors.New("Negative number")
	} else {
		return math.Sqrt(val), nil
	}
}
