package utils

import (
	"errors"
	"strings"
)

var (
	invalidInputFormatError = errors.New("input format seems to be invalid")
	invalidRobotFormatError = errors.New("one of the robot inputs is invalid")
)

func ValidateInputFormat(input string) (string, error) {
	robots := strings.Split(input, "\n\n")
	if len(robots) < 1 {
		return "", invalidInputFormatError
	}

	if !validateLineCount(robots[0], 3) {
		return "", invalidInputFormatError
	}

	for _, r := range robots[1:] {
		if !validateLineCount(r, 2) {
			return "", invalidRobotFormatError
		}
	}

	return input, nil
}

func validateLineCount(s string, c int) bool {
	lines := strings.Split(s, "\n")
	return len(lines) == c
}
