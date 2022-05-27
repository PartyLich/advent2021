// Day 12: Passage Pathing
package day12

import (
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = map[string][]string

func parseLines(in string) (_ParseResult, error) {
	result := make(_ParseResult)

	return result, nil
}

// PartOne returns the number of paths from start to end, where lowercase nodes
// may be visited multiple times
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
