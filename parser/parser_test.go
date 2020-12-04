package parser

import (
	"testing"

	"github.com/TReyburn/address-service/addresspb"
	postal "github.com/openvenues/gopostal/parser"
)

func TestConvertAddressHouse(t *testing.T) {
	pc := postal.ParsedComponent{Label: "house", Value: "test house"}
	spc := []postal.ParsedComponent{pc}
	c := make(chan *addresspb.APResponse)

	go convertAddress(spc, c)

	res := <-c

	if res.House != "test house" {
		t.Errorf("Expected 'test house' value for 'house' but got %+v", res.House)
	}
}

func TestConvertAddressCategory(t *testing.T) {
	pc := postal.ParsedComponent{Label: "category", Value: "test cat"}
	spc := []postal.ParsedComponent{pc}
	c := make(chan *addresspb.APResponse)

	go convertAddress(spc, c)

	res := <-c

	if res.Category != "test cat" {
		t.Errorf("Expected 'test cat' value for 'category' but got %+v", res.Category)
	}
}

// At this time not sure how to test
// func TestConvertAddressBadStruct(t *testing.T) {
// 	pc := postal.ParsedComponent{Label: "unexpected", Value: "bad label value"}
// 	spc := []postal.ParsedComponent{pc}
// 	c := make(chan *addresspb.APResponse)

// 	go convertAddress(spc, c)

// 	res := <-c

// 	if len(res.err) != 1 {
// 		t.Errorf("Received incorrect error(s): %+v", res.err)
// 	}
// }
