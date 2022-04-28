package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bravilogy/robots/domain"
)

var (
	invalidWorldParamsError = errors.New("invalid world parameters specified")
	invalidRobotParamsError = errors.New("invalid robot parameters specified")
)

func parseWorld(w string) (*domain.World, error) {
	parts := strings.Split(w, " ")

	if len(parts) != 2 {
		return nil, invalidWorldParamsError
	}

	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, invalidWorldParamsError
	}

	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, invalidWorldParamsError
	}

	return &domain.World{
		Width:  width,
		Height: height,
	}, nil
}

func parseRobot(buf []string) (*domain.Robot, error) {
	if len(buf) != 2 {
		return nil, invalidRobotParamsError
	}

	params := strings.Split(buf[0], " ")
	if len(params) != 3 {
		return nil, invalidRobotParamsError
	}

	x, err := strconv.Atoi(params[0])
	if err != nil {
		return nil, invalidRobotParamsError
	}

	y, err := strconv.Atoi(params[1])
	if err != nil {
		return nil, invalidRobotParamsError
	}

	var orientation int
	switch params[2] {
	case "W":
		orientation = 180
	case "N":
		orientation = 90
	case "E":
		orientation = 0
	case "S":
		orientation = 270
	default:
		return nil, invalidRobotParamsError
	}

	return &domain.Robot{
		X:           x,
		Y:           y,
		Orientation: orientation,
	}, nil

}

func ParseUniverseFromString(input string) (*domain.Universe, error) {
	var (
		universe = &domain.Universe{}
		chunks   = strings.Split(input, "\n\n")
	)

	if len(chunks) < 1 {
		return nil, invalidInputFormatError
	}

	for i, chunk := range chunks {
		lines := strings.Split(chunk, "\n")

		if len(lines) < 2 {
			return nil, invalidInputFormatError
		}

		if i == 0 {
			world, err := parseWorld(lines[0])
			if err != nil {
				return nil, err
			}

			universe.World = world

			robot, err := parseRobot(lines[1:])
			if err != nil {
				return nil, err
			}
			universe.Robots = append(universe.Robots, robot)

			continue
		}

		robot, err := parseRobot(lines)
		if err != nil {
			return nil, err
		}
		universe.Robots = append(universe.Robots, robot)
	}

	return universe, nil
}
