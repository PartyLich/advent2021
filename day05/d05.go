// Day 5: Hydrothermal Venture
package day05

import (
	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

var Parse = parse.Lines

// PartOne
func PartOne(in []string) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i), nil },
		Fn: [2]func(i interface{}) interface{}{
			runner.Unimpl,
			runner.Unimpl,
		},
	}
}
