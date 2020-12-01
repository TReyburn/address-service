package main

import (
	"fmt"
	"address-service/addr-parser"
)

func main() {
	someStr := "20544 CR 138 Goshen IN 46526"
	res := pA(someStr)
	fmt.Println(res)
}
