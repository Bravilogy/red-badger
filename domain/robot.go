package domain

type Robot struct {
	Coords      *Coords
	Orientation int
	Commands    []Command
	Lost        bool
}
