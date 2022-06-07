// Day 17: Trick Shot
package day17

import (
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = string

func parseLines(in string) (_ParseResult, error) {
	return "", runner.ErrUnimplemented
}

// PartOne returns the highest y position of any trajectory that has a discrete
// point in the target area.
func PartOne(in _ParseResult) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			runner.Unimpl,
			runner.Unimpl,
		},
	}
}
