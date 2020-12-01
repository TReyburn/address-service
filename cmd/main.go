package main

import (
	"fmt"
	"github.com/TReyburn/address-service/addr-parser"
)

func main() {
	someStr := "20544 CR 138 Goshen IN 46526"
	res := addr_parser.ParseAddress(someStr)
	fmt.Println(res)
}
