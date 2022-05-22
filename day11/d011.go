// Day 11: Dumbo Octopus
package day11

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

type state struct {
	grid    _ParseResult
	flashes int
}

// Idx1D given a width, row and column for a 2D array, returns equivalent index
// for a 1D array
func Idx1D(w, r, c int) int {
	return r*w + c
}

func flash(s state, f []bool, r, c int) (state, []bool) {
	w, h := len(s.grid[0]), len(s.grid)
	idx := Idx1D(w, r, c)
	if f[idx] {
		return s, f
	}

	f[idx] = true
	s.flashes += 1
	// any octupus that flashed has its energy set to 0
	s.grid[r][c] = 0

	// increases the energy level of all adjacent octopuses by 1, including
	// octopuses that are diagonally adjacent
	// If this causes an octopus to have an energy level greater than 9, it also
	// flashes.
	for y := r - 1; y <= r+1; y++ {
		if y < 0 || y >= h {
			continue
		}

		for x := c - 1; x <= c+1; x++ {
			if x < 0 || x >= w ||
				f[Idx1D(w, y, x)] {
				continue
			}

			s.grid[y][x] += 1
			if s.grid[y][x] > 9 {
				s, f = flash(s, f, y, x)
			}
		}
	}

	return s, f
}

func step(s state) state {
	w, h := len(s.grid[0]), len(s.grid)
	flashed := make([]bool, w*h)

	// First, the energy level of each octopus increases by 1.
	for r, row := range s.grid {
		for c := range row {
			s.grid[r][c] += 1
		}
	}

	for r, row := range s.grid {
		for c, nrg := range row {
			// any octopus with an energy level greater than 9 flashes.
			if nrg > 9 {
				s, flashed = flash(s, flashed, r, c)
			}
		}
	}

	return s
}

// PartOne returns total flashes after 100 steps.
func PartOne(in _ParseResult) int {
	steps := 100
	state := state{in, 0}

	for ; steps > 0; steps-- {
		state = step(state)
	}

	return state.flashes
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
