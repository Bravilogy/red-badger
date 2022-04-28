package domain

import (
	"fmt"
)

type orientation int

var orientationLookupTable = map[string]int{
	"E": 0,
	"N": 90,
	"W": 180,
	"S": 270,
}

func (o orientation) String() string {
	switch o {
	case 0:
		return "E"
	case 90:
		return "N"
	case 180:
		return "W"
	case 270:
		return "S"
	}

	return ""
}

type robot struct {
	Commands    []Command
	coords      *coords
	orientation orientation
	lost        bool
}

func (r robot) IsLost() bool {
	return r.lost
}

func (r robot) Report() string {
	l := ""
	if r.IsLost() {
		l = "LOST"
	}

	return fmt.Sprintf("%s %s %s", r.coords, r.orientation, l)
}
