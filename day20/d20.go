// Day 20: Trench Map
package day20

import (
	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult struct {
	algo  []string
	image map[string]bool
}

func parseLines(in string) (_ParseResult, error) {
	var result _ParseResult

	return result, nil
}

// PartOne returns how many pixels are lit after enhancing the image twice.
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
