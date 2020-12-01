package addr_parser

import (
	"fmt"
	parser "github.com/openvenues/gopostal/parser"
)

func ParseAddress(addr string, c chan string) {
	parsed := parser.ParseAddress(addr)
	fmt.Println(parsed)
	c <- "finished"
}