package parser

import (
	"fmt"

	parser "github.com/openvenues/gopostal/parser"
)

type address struct {
	house         string
	category      string
	near          string
	houseNumber   string
	road          string
	unit          string
	level         string
	staircase     string
	entrance      string
	poBox         string
	postcode      string
	suburb        string
	cityDistrict  string
	city          string
	island        string
	stateDistrict string
	state         string
	countryRegion string
	country       string
	worldRegion   string
}

// ParseAddress takes an address and a channel to coommunicate on
func ParseAddress(addr string, c chan string) {
	parsed := parser.ParseAddress(addr)
	fmt.Println(parsed)
	c <- "finished"
}
