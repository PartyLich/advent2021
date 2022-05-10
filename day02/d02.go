package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type dir int

const (
	Forward dir = iota
	Down
	Up
)

type command struct {
	dir dir
	val int
}

type pos struct {
	horz  int
	depth int
}

// parseCom parses a command list from a string
func parseCom(in string) ([]command, error) {
	lines := parse.Lines(in)
	commands := make([]command, len(lines))

	for idx, line := range lines {
		parts := strings.Split(line, " ")
		dirStr, v := parts[0], parts[1]

		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("unable to parse command from %v: %w", line, err)
		}

		var direction dir

		switch dirStr {
		case "forward":
			direction = Forward
		case "down":
			direction = Down
		case "up":
			direction = Up
		default:
			return nil, fmt.Errorf("unable to parse command from %v: %w", line, err)
		}

		commands[idx] = command{
			direction,
			val,
		}
	}

	return commands, nil
}

// PartOne calculates the horizontal position and depth after following the planned course.
// Returns final horizontal position multiplied by final depth
func PartOne(commands []command) int {
	var pos pos

	for _, v := range commands {
		switch v.dir {
		case Forward:
			pos.horz += v.val
		case Down:
			pos.depth += v.val
		case Up:
			pos.depth -= v.val
		}
	}

	return pos.horz * pos.depth
}

// PartTwo calculates the horizontal position and depth after following the planned course.
// Returns final horizontal position multiplied by final depth
func PartTwo(commands []command) int {
	var pos pos
	var aim int

	for _, v := range commands {
		switch v.dir {
		case Forward:
			pos.horz += v.val
			pos.depth += aim * v.val
		case Down:
			aim += v.val
		case Up:
			aim -= v.val
		}
	}

	return pos.horz * pos.depth
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseCom(i) },
		One:   func(i interface{}) interface{} { return PartOne(i.([]command)) },
		Two:   func(i interface{}) interface{} { return PartTwo(i.([]command)) },
	}
}
