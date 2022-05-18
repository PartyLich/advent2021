// Day 8: Seven Segment Search
package day08

import (
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = [][][]string

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := runner.NewGrid[[]string](len(lines), 2)

	for i, v := range lines {
		parts := strings.Split(v, " | ")
		signals := strings.Fields(strings.TrimSpace(parts[0]))
		output := strings.Fields(strings.TrimSpace(parts[1]))
		result[i][0], result[i][1] = signals, output
	}

	return result, nil
}

// PartOne returns how many times the digits 1, 4, 7, or 8 appear in output
// values
func PartOne(in _ParseResult) int {
	count := 0

	for _, v := range in {
		for _, digit := range v[1] {
			switch len(digit) {
			case 2:
				fallthrough
			case 4:
				fallthrough
			case 3:
				fallthrough
			case 7:
				count += 1
			}
		}
	}

	return count
}

// PartTwo returns the sum of all decoded outputs
func PartTwo(in _ParseResult) int {
	count := 0
	return count
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
