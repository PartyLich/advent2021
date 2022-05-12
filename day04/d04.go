// Day 4: Giant Squid
package day04

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
		One:   func(i interface{}) interface{} { return PartOne(i.([]string)) },
		Two:   runner.Unimpl,
	}
}
