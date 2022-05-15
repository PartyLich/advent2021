// Day 6: Lanternfish
package day06

import (
	"github.com/partylich/advent2021/runner"
)

func parseFish(in string) ([]int, error) {
	result := make([]int, 9)

	return result, nil
}

// PartOne returns the number of fish after 80 days
func PartOne(in []int) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseFish(i) },
		Fn: [2]func(i interface{}) interface{}{
			runner.Unimpl,
			runner.Unimpl,
		},
	}
}
