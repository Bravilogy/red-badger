package domain

type World struct {
	Width, Height int
}

type Robot struct {
	X, Y, Orientation int
}

type Universe struct {
	World  *World
	Robots []*Robot
}
