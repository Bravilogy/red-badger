package utils

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/bravilogy/robots/domain"
)

var (
	invalidWorldParamsError = errors.New("invalid world parameters specified")
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

func ParseUniverseFromFile(file *os.File) (*domain.Universe, error) {
	var (
		universe = &domain.Universe{}
		scanner  = bufio.NewScanner(file)
	)

	var (
		line               int
		robotFieldsCounter int
	)
	for scanner.Scan() {
		if line == 0 {
			world, err := parseWorld(scanner.Text())
			if err != nil {
				return nil, err
			}

			universe.World = world
		}

	}

	return nil, invalidWorldParamsError
}
