package domain

var CommandsLookupTable = map[string]Command{
	"F": &ForwardCommand{},
	"L": &LeftCommand{},
	"R": &RightCommand{},
}

type Command interface {
	Run()
}

type LeftCommand struct{}

func (c *LeftCommand) Run() {
}

type RightCommand struct{}

func (c *RightCommand) Run() {
}

type ForwardCommand struct{}

func (c *ForwardCommand) Run() {
}

type World struct {
	Width, Height int
}

type Robot struct {
	X, Y, Orientation int
	Commands          []*Command
}

type Universe struct {
	World  *World
	Robots []*Robot
}
