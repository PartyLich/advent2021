// Day 9: Smoke Basin
package day09

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
	"github.com/partylich/go/iter"
)

type _ParseResult = [][]string

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := make([][]string, len(lines))

	for r, l := range lines {
		pts := strings.Split(l, "")
		result[r] = pts
	}

	return result, nil
}

// return true if grid[r][c] is smaller than any of its neighbors
func smallest(grid _ParseResult, r, c int) bool {
	width, height := len(grid[0]), len(grid)

	for y := r - 1; y <= r+1; y++ {
		if y < 0 || y >= height || (y == r) {
			continue
		}

		if x := c; grid[r][c] >= grid[y][x] {
			return false
		}
	}

	for x := c - 1; x <= c+1; x++ {
		if x < 0 || x >= width || (x == c) {
			continue
		}

		if y := r; grid[r][c] >= grid[y][x] {
			return false
		}
	}

	return true
}

// PartOne returns the sum of the risk levels of all low points on your
// heightmap.
func PartOne(in _ParseResult) int {
	sum := 0

	// TODO: avoid checking positions whos neighbor has been determined the lowest
	for r, row := range in {
		for c, col := range row {
			if smallest(in, r, c) {
				height, err := strconv.Atoi(col)
				if err != nil {
					panic("unable to convert string to int")
				}

				sum += height + 1
			}
		}
	}

	return sum
}

func toKey(r, c int) string {
	return fmt.Sprintf("%v,%v", c, r)
}

func traverse(m map[string]string, r, c int) int {
	key := toKey(r, c)
	if height, ok := m[key]; !ok || height == "9" {
		return 0
	}

	delete(m, key)

	return 1 +
		traverse(m, r-1, c) +
		traverse(m, r+1, c) +
		traverse(m, r, c-1) +
		traverse(m, r, c+1)
}

// PartTwo returns the product of the three largest basins
//
// A basin is all locations that eventually flow downward to a single low point.
// Therefore, every low point has a basin, although some basins are very small.
// Locations of height 9 do not count as being in any basin, and all other
// locations will always be part of exactly one basin.
//
// The size of a basin is the number of locations within the basin, including
// the low point.
func PartTwo(in _ParseResult) int {
	// map from pos to val. could have been done in one pass during the parse
	m := make(map[string]string)
	for r, row := range in {
		for c, height := range row {
			m[toKey(r, c)] = height
		}
	}

	var basins []int
	for r, row := range in {
		for c := range row {
			key := toKey(r, c)

			if height, ok := m[key]; !ok || height == "9" {
				continue
			}

			size := traverse(m, r, c)
			if size == 0 {
				continue
			}

			basins = append(basins, size)
		}
	}

	sort.Ints(basins)
	it := iter.New(basins).Rev().Take(3)
	return iter.Reduce[int, int](it, 1, func(prev, next int) int {
		return prev * next
	})
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.(_ParseResult)) },
			func(i interface{}) interface{} { return PartTwo(i.(_ParseResult)) },
		},
	}
}
