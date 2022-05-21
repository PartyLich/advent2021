// Day 10:
package day10

import (
	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = [][]string

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := make([][]string, len(lines))

	return result, nil
}

// PartOne returns the total syntax error score for the input
func PartOne(in _ParseResult) int {
	sum := 0

	return sum
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
