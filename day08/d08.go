// Day 8: Seven Segment Search
package day08

import (
	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = [][][]string

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := runner.NewGrid[[]string](len(lines), 2)

	return result, nil
}

// PartOne returns how many times the digits 1, 4, 7, or 8 appear in output
// values
func PartOne(in _ParseResult) int {
	count := 0
	return count
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
