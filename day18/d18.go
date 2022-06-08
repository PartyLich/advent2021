// Day 18: Snailfish
package day18

import (
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = string

func parseLines(in string) (_ParseResult, error) {
	return "", runner.ErrUnimplemented
}

func reduce(in string) string {
	return ""
}

func add(a, b string) string {
	return ""
}

func magnitude(in string) int {
	return 0
}

// PartOne returns the magnitude of the final sum.
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
