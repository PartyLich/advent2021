// Day 7: The Treachery of Whales
package day07

import (
	"math"
	"strconv"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

func parsePos(in string) ([]int, error) {
	return parse.Csv[int](in, strconv.Atoi)
}

type pos = int

func dist(from, to pos) int {
	return runner.Max(from, to) - runner.Min(from, to)
}

// cost returns the cost of moving all entries to position p
func cost(m map[pos]int, p pos) int {
	sum := 0

	for k, w := range m {
		if k == p {
			continue
		}

		sum += dist(p, k) * w
	}

	return sum
}

// PartOne returns the smallest total fuel spent required to align all crabs.
func PartOne(in []int) int {
	var (
		maxPos int
		minPos int
	)
	m := make(map[pos]int)

	for _, p := range in {
		if _, ok := m[p]; !ok {
			m[p] = 0
		}

		m[p] += 1

		maxPos = runner.Max(p, maxPos)
		minPos = runner.Min(p, minPos)
	}

	minCost := math.MaxInt
	for pos := minPos; pos <= maxPos; pos++ {
		if fuelc := cost(m, pos); fuelc < minCost {
			minCost = fuelc
		}
	}

	return minCost
}

func costTwo(m map[pos]int, p pos) int {
	sum := 0

	for k, w := range m {
		if k == p {
			continue
		}

		d := dist(p, k)
		// arithmetic series * weight
		sum += ((d * (d + 1)) / 2) * w
	}

	return sum
}

// PartTwo returns the smallest total fuel spent required to align all crabs.
// Cost is now like an arithemetic sequence or something I guess
func PartTwo(in []int) int {
	var (
		maxPos int
		minPos int
	)
	m := make(map[pos]int)

	for _, p := range in {
		if _, ok := m[p]; !ok {
			m[p] = 0
		}

		m[p] += 1

		maxPos = runner.Max(p, maxPos)
		minPos = runner.Min(p, minPos)
	}

	minCost := math.MaxInt
	for pos := minPos; pos <= maxPos; pos++ {
		minCost = runner.Min(minCost, costTwo(m, pos))
	}

	return minCost
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parsePos(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.([]int)) },
			func(i interface{}) interface{} { return PartTwo(i.([]int)) },
		},
	}
}
