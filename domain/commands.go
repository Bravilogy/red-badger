package domain

import (
	"math"
)

var commandsLookupTable = map[string]Command{
	"F": &ForwardCommand{},
	"L": &LeftCommand{},
	"R": &RightCommand{},
}

type Command interface {
	Run(robot *robot, world *world)
}

type LeftCommand struct{}

func (c *LeftCommand) Run(robot *robot, _ *world) {
	robot.orientation += 90
	if robot.orientation > 270 {
		robot.orientation = 0
	}
}

type RightCommand struct{}

func (c *RightCommand) Run(robot *robot, _ *world) {
	robot.orientation -= 90
	if robot.orientation < 0 {
		robot.orientation = 270
	}
}

type ForwardCommand struct{}

func (c *ForwardCommand) Run(robot *robot, world *world) {
	r := float64(robot.orientation) * math.Pi / 180
	x := int(math.Round(float64(robot.coords.x) + math.Cos(r)))
	y := int(math.Round(float64(robot.coords.y) + math.Sin(r)))

	if x < 0 || x > world.width || y < 0 || y > world.height {
		cs := coords{x, y}
		if world.HasScent(cs) {
			return
		}

		world.AddScent(cs)
		robot.lost = true
		return
	}

	robot.coords.x = x
	robot.coords.y = y
}
