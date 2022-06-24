// Day 21: Dirac Dice
package day21

import (
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = []int

func parseLines(in string) (_ParseResult, error) {
	result := make([]int, 2)

	return result, nil
}

// PartOne returns the product of the score of the losing player and the number
// of times the die was rolled.
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
