// Day 5: Hydrothermal Venture
package day05

import (
	"strconv"
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
	"github.com/partylich/advent2021/slice"
)

type Line struct {
	Start []int
	End   []int
}

func (l Line) horz() bool {
	return l.Start[1] == l.End[1]
}

func (l Line) vert() bool {
	return l.Start[0] == l.End[0]
}

var Parse = parse.Lines

type parseResult struct {
	lines      []Line
	maxX, maxY int
}

func toNum(in []string) ([]int, error) {
	result := make([]int, len(in))
	for i, v := range in {
		n, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}

		result[i] = n
	}

	return result, nil
}

func parseLines(in []string) (parseResult, error) {
	lines := make([]Line, len(in))
	var (
		maxX, maxY int
	)

	for i, lineStr := range in {
		endPoints := strings.Split(lineStr, " -> ")
		startStr := strings.Split(endPoints[0], ",")
		endStr := strings.Split(endPoints[1], ",")
		startPt, err := toNum(startStr)
		if err != nil {
			return parseResult{}, err
		}
		maxX = runner.Max(maxX, startPt[0])
		maxY = runner.Max(maxY, startPt[1])

		endPt, err := toNum(endStr)
		if err != nil {
			return parseResult{}, err
		}
		maxX = runner.Max(maxX, endPt[0])
		maxY = runner.Max(maxY, endPt[1])

		l := Line{startPt, endPt}
		if l.horz() && startPt[0] > endPt[0] {
			l = Line{endPt, startPt}
		} else if startPt[1] > endPt[1] {
			l = Line{endPt, startPt}
		}

		lines[i] = l
	}

	return parseResult{
		lines,
		maxX, maxY,
	}, nil
}

// PartOne returns the number of points where at least two lines overlap,
// considering only horizontal and vertical lines
func PartOne(in []string) int {
	parsed, err := parseLines(in)
	if err != nil {
		panic(err)
	}
	parsed.lines = slice.Filter(
		func(l Line) bool {
			return l.horz() || l.vert()
		},
		parsed.lines,
	)

	grid := runner.NewGrid[int](parsed.maxY+1, parsed.maxX+1)

	var count int
	for _, l := range parsed.lines {
		for r := l.Start[1]; r <= l.End[1]; r++ {
			for c := l.Start[0]; c <= l.End[0]; c++ {
				if grid[r][c] <= 1 && grid[r][c] > -1 {
					grid[r][c] += 1
				}
				if grid[r][c] > 1 {
					count += 1
					grid[r][c] = -1
				}
			}
		}
	}

	return count
}

// PartTwo returns the number of points where at least two lines overlap,
// considering horizontal, vertical, and diagonal (45deg) lines
func PartTwo(in []string) int {
	parsed, err := parseLines(in)
	if err != nil {
		panic(err)
	}

	grid := runner.NewGrid[int](parsed.maxY+1, parsed.maxX+1)
	count := 0

	update := func(r, c int) {
		if grid[r][c] <= 1 && grid[r][c] > -1 {
			grid[r][c] += 1
		}
		if grid[r][c] > 1 {
			count += 1
			grid[r][c] = -1
		}
	}

	for _, l := range parsed.lines {
		if l.horz() || l.vert() {
			for r := l.Start[1]; r <= l.End[1]; r++ {
				for c := l.Start[0]; c <= l.End[0]; c++ {
					update(r, c)
				}
			}
		} else {
			step := 1
			pred := func(i int) bool {
				return i <= l.End[0]
			}
			if l.Start[0] > l.End[0] {
				step = -1
				pred = func(i int) bool {
					return i >= l.End[0]
				}
			}

			for r, c := l.Start[1], l.Start[0]; r <= l.End[1] && pred(c); r, c = r+1, c+step {
				update(r, c)
			}
		}
	}

	return count
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i), nil },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.([]string)) },
			func(i interface{}) interface{} { return PartTwo(i.([]string)) },
		},
	}
}
