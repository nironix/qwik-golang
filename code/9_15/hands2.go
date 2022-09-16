package main

import (
	"errors"
	"fmt"
)

func main() {
	var up float64
	var down float64
	fmt.Println("enter numerator : ")
	fmt.Scan(&up)
	fmt.Println("enter denominator : ")
	fmt.Scan(&down)
	s, e := div(up, down)
	if e != nil {
		fmt.Println("error : ", e)
	} else {
		fmt.Println("Result", s)
	}

}

func div(up float64, down float64) (float64, error) {
	if down == 0 {
		return 0, errors.New("denominator zero")
	} else {
		return up / down, nil
	}
}
