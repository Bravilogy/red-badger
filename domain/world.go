package domain

// world struct holds the dimensions of the world and scent coordinates
// instead of exposing fields, we expose methods to interact with the world
type world struct {
	width  int
	height int
	scents map[coords]bool
}

func (w *world) AddScent(coords coords) {
	w.scents[coords] = true
}

func (w world) HasScent(coords coords) bool {
	return w.scents[coords]
}
