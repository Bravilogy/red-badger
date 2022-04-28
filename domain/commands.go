package domain

import (
	"fmt"
	"math"
)

var CommandsLookupTable = map[string]Command{
	"F": &ForwardCommand{},
	"L": &LeftCommand{},
	"R": &RightCommand{},
}

type Command interface {
	Run(robot *Robot, world *World)
}

type LeftCommand struct{}

func (c *LeftCommand) Run(robot *Robot, _ *World) {
	robot.Orientation -= 90
}

type RightCommand struct{}

func (c *RightCommand) Run(robot *Robot, _ *World) {
	robot.Orientation += 90
}

type ForwardCommand struct{}

func (c *ForwardCommand) Run(robot *Robot, world *World) {
	orientationRadians := float64(robot.Orientation) * math.Pi / 180
	fmt.Println(orientationRadians, math.Cos(orientationRadians))
}
