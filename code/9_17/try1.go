package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
)

type t struct {
	age  int
	name string
}
type tt []t

func main() {

	fmt.Println(math.Sqrt(4))
	fmt.Println(math.Pow10(5))

	s := []int{2, 8, 9, 1, 101, 34}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)

	var b bool = true
	fmt.Println(reflect.TypeOf(b), " ", b)
	var ss string = strconv.FormatBool(b)
	fmt.Println(reflect.TypeOf(ss), " ", ss)

}
