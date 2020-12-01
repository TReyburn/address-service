package addr_parser

import (
	"fmt"

	parser "github.com/openvenues/gopostal/parser"
)

// ParseAddress takes an address and a channel to coommunicate on
func ParseAddress(addr string, c chan string) {
	parsed := parser.ParseAddress(addr)
	fmt.Println(parsed)
	c <- "finished"
}
