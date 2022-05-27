// Day 12: Passage Pathing
package day12

import (
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = map[string]map[string]bool

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := make(_ParseResult)

	for _, l := range lines {
		nodes := strings.Split(l, "-")
		a, b := nodes[0], nodes[1]

		if result[a] == nil {
			result[a] = make(map[string]bool)
		}
		if result[b] == nil {
			result[b] = make(map[string]bool)
		}
		// consider start->node edges unidirectional
		if a == "start" {
			result[a][b] = true
			continue
		}
		if b == "start" {
			result[b][a] = true
			continue
		}

		result[a][b] = true
		result[b][a] = true
	}

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
