package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	x.PushFront(8)
	for i := x.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
