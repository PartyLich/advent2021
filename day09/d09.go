// Day 9: Smoke Basin
package day09

import (
	"strconv"
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
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

// PartTwo returns the product of the three largest basins
func PartTwo(in _ParseResult) int {
	product := 0

	return product
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
