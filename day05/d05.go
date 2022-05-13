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
		}
		if l.vert() && startPt[1] > endPt[1] {
			l = Line{endPt, startPt}
		}

		lines[i] = l
	}

	return parseResult{
		slice.Filter(
			func(l Line) bool {
				return l.horz() || l.vert()
			},
			lines,
		),
		maxX, maxY,
	}, nil
}

// PartOne
func PartOne(in []string) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i), nil },
		Fn: [2]func(i interface{}) interface{}{
			runner.Unimpl,
			runner.Unimpl,
		},
	}
}
