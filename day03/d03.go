// Day 3: Binary Diagnostic
package day03

import (
	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

var Parse = parse.Lines

// PartOne uses the binary numbers in your diagnostic report to calculate the gamma rate and epsilon
// rate, then multiply them together.
func PartOne(in []string) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i), nil },
		One:   runner.Unimpl,
		Two:   runner.Unimpl,
	}
}
