// Day 17: Trick Shot
package day17

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/partylich/advent2021/runner"
)

type Bound struct {
	Min, Max int
}

func (b Bound) InBounds(i int) bool {
	return i <= b.Max && i >= b.Min
}

type _ParseResult = []Bound

func parseLines(in string) (_ParseResult, error) {
	re := regexp.MustCompile(`x=(-?\d+)\.{2}(-?\d+), y=(-?\d+)\.{2}(-?\d+)`)
	m := re.FindSubmatch([]byte(in))

	bounds := make([]int, 4)
	for idx := 1; idx < len(m); idx++ {
		b, err := strconv.Atoi(string(m[idx]))
		if err != nil {
			return nil, errors.New("parse failure")
		}

		bounds[idx-1] = b
	}

	xB := Bound{bounds[0], bounds[1]}
	yB := Bound{bounds[2], bounds[3]}

	return []Bound{xB, yB}, nil
}

// PartOne returns the highest y position of any trajectory that has a discrete
// point in the target area.
func PartOne(in _ParseResult) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			runner.Unimpl,
			runner.Unimpl,
		},
	}
}
