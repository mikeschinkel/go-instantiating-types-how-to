package insure

import (
	"fmt"
	"strings"
)

type Locations []*Location
type Location struct {
	Indicator        bool
	Address1         string
	City             string
	State            string
	ZIPCode          string
	FullAddress      string
	Type             string
	StringBuilderLoc strings.Builder
}

func NewLocation(addr string) *Location {
	return &Location{Address1: addr}
}

func (l Location) StringWithTabs(tabs string) (s string) {
	return fmt.Sprintf("%sLocation: %s", tabs+tab, l.Address1)
}
