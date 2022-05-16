// Day 7: The Treachery of Whales
package day07

import (
	"github.com/partylich/advent2021/runner"
)

func parsePos(in string) ([]int, error) {
	return []int{}, nil
}

// PartOne returns the smallest total fuel spent required to align all crabs.
func PartOne(in []int) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parsePos(i) },
		Fn: [2]func(i interface{}) interface{}{
			runner.Unimpl,
			runner.Unimpl,
		},
	}
}
