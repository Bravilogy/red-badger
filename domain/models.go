package domain

type World struct {
	Width, Height int
}

type Robot struct {
}

type Universe struct {
	World  *World
	Robots []*Robot
}
