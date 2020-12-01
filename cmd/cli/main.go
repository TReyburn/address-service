package main

import (
	"fmt"
	parser "github.com/openvenues/gopostal/parser"
)

func main() {
	parsed := parser.ParseAddress("781 Franklin Ave Crown Heights Brooklyn NY 11216 USA")
	fmt.Println(parsed)
}
