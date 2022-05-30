// Day 14:
package day14

import (
	"math"
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

func inc(m *map[string]int, k string) {
	i, ok := (*m)[k]
	if !ok {
		i = 0
	}

	(*m)[k] = i + 1
}

// PartOne returns the quantity of the most common element minus the quantity of
// the least common element after 10 steps.
func PartOne(in _ParseResult) int {
	counts := make(map[string]int)
	last := in.Tmpl
	res := ""

	for idx := 0; idx < 10; idx++ {
		res = string(last[0])

		for i := 0; i < len(last)-1; i++ {
			b := in.Rules[last[i:i+2]]
			c := string(last[i+1])

			res += b + c
		}

		last = res
	}

	max, min := math.MinInt, math.MaxInt
	for _, v := range res {
		inc(&counts, string(v))
	}
	for _, v := range counts {
		max = runner.Max(max, v)
		min = runner.Min(min, v)
	}

	return max - min
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
