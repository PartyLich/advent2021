// Day 9: Smoke Basin
package day09

import (
	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = [][]string

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := runner.NewGrid[string](len(lines), 2)

	return result, nil
}

// PartOne returns the sum of the risk levels of all low points on your
// heightmap.
func PartOne(in _ParseResult) int {
	sum := 0

	return sum
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.(_ParseResult)) },
			runner.Unimpl,
		},
	}
}
