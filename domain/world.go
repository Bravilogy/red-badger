package domain

// World struct holds the dimensions of the world and scent coordinates
// instead of exposing fields, we expose methods to interact with the world
type World struct {
	width  int
	height int
	scents map[Coords]bool
}

func (w *World) AddScent(x, y int) {
	w.scents[Coords{x, y}] = true
}

func (w World) HasScent(coords Coords) bool {
	return w.scents[coords]
}
