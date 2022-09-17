package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// b := []byte("lala lala")
	// s := bytes.Split(b, []byte{'l'})
	// fmt.Println(s)
	// for _, j := range s {
	// 	fmt.Println(string(j))
	// }

	// j := bytes.Join(s, []byte("**"))
	// fmt.Println(string(j))

	v := crc32.NewIEEE()
	v.Write([]byte("test"))
	h := v.Sum32()
	fmt.Println(h, " ", v)

}
