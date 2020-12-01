package main

import (
	parser "github.com/openvenues/gopostal/parser"
)

func pA(addr string) []parser.ParsedComponent {
	parsed := parser.ParseAddress(addr)
	return parsed
}