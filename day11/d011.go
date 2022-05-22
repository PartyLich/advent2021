// Day 11: Dumbo Octopus
package day11

import (
	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = [][]int

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := make([][]int, len(lines))

	return result, nil
}

// PartOne returns
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
