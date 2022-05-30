// Day 13: Transparent Origami
package day13

import (
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = [][]int

func parseLines(in string) (_ParseResult, error) {
	var result _ParseResult

	return result, runner.ErrUnimplemented
}

// PartOne returns how many dots are visible after completing the first fold
// instruction.
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
