package domain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const MaxCoordinateValues = 50

var (
	invalidWorldParamsError     = errors.New("invalid world parameters specified")
	invalidWorldDimensionsError = errors.New("invalid world dimensions specified. the maximum w/h is 50")
	invalidRobotParamsError     = errors.New("invalid robot parameters specified")
	invalidRobotCoordsError     = errors.New("invalid robot coordinates specified. the maximum x/y is 50")
	invalidInputFormatError     = errors.New("input format seems to be invalid")
)

func parseWorld(w string) (*world, error) {
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

	if width > MaxCoordinateValues || height > MaxCoordinateValues {
		return nil, invalidWorldDimensionsError
	}

	return &world{
		width:  width,
		height: height,
		scents: make(map[coords]bool),
	}, nil
}

func parseRobot(buf []string) (*robot, error) {
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

	if x > MaxCoordinateValues || y > MaxCoordinateValues {
		return nil, invalidRobotCoordsError
	}

	var commands []Command
	for _, commandString := range strings.Split(buf[1], "") {
		c, ok := commandsLookupTable[commandString]
		if !ok {
			return nil, errors.New(fmt.Sprintf("invalid command specified - %s", commandString))
		}

		commands = append(commands, c)
	}

	o, ok := orientationLookupTable[params[2]]
	if !ok {
		return nil, invalidRobotParamsError
	}

	return &robot{
		coords:      &coords{x: x, y: y},
		orientation: orientation(o),
		Commands:    commands,
	}, nil

}

func NewUniverseFromString(input string) (*Universe, error) {
	var (
		universe = &Universe{}
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
