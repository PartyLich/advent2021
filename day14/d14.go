// Day 14:
package day14

import (
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult struct {
	Tmpl  string
	Rules map[string]string
}

func parseLines(in string) (_ParseResult, error) {
	var result _ParseResult
	p := strings.Split(in, "\n\n")
	result.Tmpl = p[0]

	lines := parse.Lines(p[1])
	result.Rules = make(map[string]string)
	for _, l := range lines {
		parts := strings.Split(l, " -> ")
		result.Rules[parts[0]] = parts[1]
	}

	return result, nil
}

// PartOne returns the quantity of the most common element minus the quantity of
// the least common element after 10 steps.
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
