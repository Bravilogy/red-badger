package domain

import "fmt"

type coords struct {
	x, y int
}

func (c coords) String() string {
	return fmt.Sprintf("%d %d", c.x, c.y)
}
