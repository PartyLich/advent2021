// Day 6: Lanternfish
package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/partylich/advent2021/runner"
)

func parseFish(in string) ([]int, error) {
	fishStr := strings.Split(strings.TrimSpace(in), ",")
	result := make([]int, 9)

	for _, v := range fishStr {
		n, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
		if n > 8 {
			return []int{}, fmt.Errorf("input outside of valid range (0,8): %v", n)
		}

		result[n] += 1
	}

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
