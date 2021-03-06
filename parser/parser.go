package parser

import (
	"log"

	"github.com/TReyburn/address-service/addresspb"
	postal "github.com/openvenues/gopostal/parser"
)

// InitParser is a simple init to load our large C dll
func InitParser() {
	t := "test"
	a := postal.ParseAddress(t)
	log.Println("Parser loaded", a)
}

// ParseAddress takes an address and a channel to coommunicate on
func ParseAddress(addr string, ch chan *addresspb.APResponse) {
	parsed := postal.ParseAddress(addr)
	go convertAddress(parsed, ch)
}

func convertAddress(pc []postal.ParsedComponent, ch chan *addresspb.APResponse) {
	addr := addresspb.APResponse{}
	for _, c := range pc {
		switch c.Label {
		case "house":
			addr.House = c.Value
		case "category":
			addr.Category = c.Value
		case "near":
			addr.Near = c.Value
		case "house_number":
			addr.HouseNumber = c.Value
		case "road":
			addr.Road = c.Value
		case "unit":
			addr.Unit = c.Value
		case "level":
			addr.Level = c.Value
		case "staircase":
			addr.Staircase = c.Value
		case "entrance":
			addr.Entrance = c.Value
		case "po_box":
			addr.PoBox = c.Value
		case "postcode":
			addr.Postcode = c.Value
		case "suburb":
			addr.Suburb = c.Value
		case "city_district":
			addr.CityDistrict = c.Value
		case "city":
			addr.City = c.Value
		case "island":
			addr.Island = c.Value
		case "state_district":
			addr.StateDistrict = c.Value
		case "state":
			addr.State = c.Value
		case "country_region":
			addr.CountryRegion = c.Value
		case "country":
			addr.Country = c.Value
		case "world_region":
			addr.WorldRegion = c.Value
		default:
			// What should we do here? How can make this unit testable without adding to our protobuf?
			log.Println("Did not recognize", c.Label)
		}
	}
	ch <- &addr
}
