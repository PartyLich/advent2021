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
	const steps int = 80

	for i := 0; i < steps; i++ {
		born := in[0]

		for j := 0; j < 8; j++ {
			in[j] = in[j+1]
		}

		in[8] = born
		in[6] += born
	}

	count := 0
	for _, v := range in {
		count += v
	}

	return count
}

// PartTwo returns the number of fish after 256 days
func PartTwo(in []int) int {
	const steps int = 256

	for i := 0; i < steps; i++ {
		born := in[0]

		for j := 0; j < 8; j++ {
			in[j] = in[j+1]
		}

		in[8] = born
		in[6] += born
	}

	count := 0
	for _, v := range in {
		count += v
	}

	return count
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseFish(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.([]int)) },
			func(i interface{}) interface{} { return PartTwo(i.([]int)) },
		},
	}
}
