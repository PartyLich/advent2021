// Day 15: Chiton
package day15

import (
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = [][]int

func parseLines(in string) (_ParseResult, error) {
	return nil, runner.ErrUnimplemented
}

// PartOne returns the lowest total risk of any path from top left to bottom
// right.
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
