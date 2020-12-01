package

import (
	parser "github.com/openvenues/gopostal/parser"
)

func ParseAddress(addr string) []parser.ParsedComponent {
	parsed := parser.ParseAddress(addr)
	return parsed
}