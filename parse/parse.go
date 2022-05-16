package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// Lines splits a string at newline characters
func Lines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

// UintList parses a list of uint from a newline delimited string
func UintList(input string) ([]uint, error) {
	lines := Lines(input)
	parsed := make([]uint, len(lines))

	for i, line := range lines {
		v, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse uint from %v: %w", line, err)
		}

		parsed[i] = uint(v)
	}

	return parsed, nil
}

// IntList parses a list of uint from a newline delimited string
func IntList(input string) ([]int, error) {
	lines := Lines(input)
	parsed := make([]int, len(lines))

	for i, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse int from %v: %w", line, err)
		}

		parsed[i] = v
	}

	return parsed, nil
}

// Csv parses a single line of comma separated values into a slice of T.
func Csv[T any](input string, conv func(string) (T, error)) ([]T, error) {
	records := strings.Split(strings.TrimSpace(input), ",")

	result := make([]T, len(records))
	for i, r := range records {
		v, err := conv(r)
		if err != nil {
			return result, err
		}

		result[i] = v
	}

	return result, nil
}
