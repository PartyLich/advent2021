// Day 15: Chiton
package day15

import (
	"strconv"
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = [][]int

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := runner.NewGrid[int](len(lines), len(lines[0]))

	for r, l := range lines {
		row := strings.Split(l, "")
		for c, nrgStr := range row {
			nrg, err := strconv.Atoi(nrgStr)
			if err != nil {
				panic("parse failure")
			}

			result[r][c] = nrg
		}
	}

	return result, nil
}

// PartOne returns the lowest total risk of any path from top left to bottom
// right.
func PartOne(in _ParseResult) int {
	width := len(in[0])
	height := len(in)
	result := runner.NewGrid[int](height, width)
	result[height-1][width-1] = in[height-1][width-1]

	for r := height - 1; r >= 0; r-- {
		for c := width - 1; c > 0; c-- {
			// last column, up only
			if c == width-1 && r < height-1 {
				result[r][c] = in[r][c] + result[r+1][c]
			}

			left := in[r][c-1] + result[r][c]
			up := left
			if r < height-1 {
				up = in[r][c-1] + result[r+1][c-1]
			}

			result[r][c-1] = runner.Min(left, up)
		}
	}

	return runner.Min(result[0][1], result[1][0])
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.(_ParseResult)) },
			runner.Unimpl,
		},
	}
}
