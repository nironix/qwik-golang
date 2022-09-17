package main

import (
	"fmt"
	"os"
)

func main() {
	// da := "lalalal"
	// se := base64.StdEncoding.EncodeToString([]byte(da))
	// fmt.Println(se)

	// de, _ := base64.StdEncoding.DecodeString(se)
	// fmt.Println(string(de))
	//os.Setenv("h", "qc")
	fmt.Println(os.Getenv("h"))
}
