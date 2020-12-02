package parser

import (
	"fmt"

	postal "github.com/openvenues/gopostal/parser"
)

// Address struct containing all possible address fields
type Address struct {
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
	err           []string
}

// ParseAddress takes an address and a channel to coommunicate on
func ParseAddress(addr string, c chan Address) {
	parsed := postal.ParseAddress(addr)
	go convertAddress(parsed, c)
}

func convertAddress(pc []postal.ParsedComponent, c chan Address) {
	addr := Address{}
	for _, c := range pc {
		switch c.Label {
		case "house":
			addr.house = c.Value
		case "category":
			addr.category = c.Value
		case "near":
			addr.near = c.Value
		case "house_number":
			addr.houseNumber = c.Value
		case "road":
			addr.road = c.Value
		case "unit":
			addr.unit = c.Value
		case "level":
			addr.level = c.Value
		case "staircase":
			addr.staircase = c.Value
		case "entrance":
			addr.entrance = c.Value
		case "po_box":
			addr.poBox = c.Value
		case "postcode":
			addr.postcode = c.Value
		case "suburb":
			addr.suburb = c.Value
		case "city_district":
			addr.cityDistrict = c.Value
		case "city":
			addr.city = c.Value
		case "island":
			addr.island = c.Value
		case "state_district":
			addr.stateDistrict = c.Value
		case "state":
			addr.state = c.Value
		case "country_region":
			addr.countryRegion = c.Value
		case "country":
			addr.country = c.Value
		case "world_region":
			addr.worldRegion = c.Value
		default:
			fmt.Println("Did not recognize", c.Label)
			addr.err = append(addr.err, c.Label)
		}
	}
	c <- addr
}
