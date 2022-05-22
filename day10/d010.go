// Day 10:
package day10

import (
	"strings"

	"golang.org/x/exp/slices"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = [][]string

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := make([][]string, len(lines))

	for r, l := range lines {
		syms := strings.Split(l, "")
		result[r] = syms
	}

	return result, nil
}

var (
	scoreTable = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	openers = []string{"(", "[", "{", "<"}
)

const incomplete string = "incomplete"

func pop[T any](slice *[]T) *T {
	if len(*slice) == 0 {
		return nil
	}

	i := len(*slice) - 1
	v := (*slice)[i]
	*slice = append((*slice)[:i], (*slice)[i+1:]...)

	return &v
}

func checkLine(syms []string) (string, bool) {
	stack := make([]string, 0, len(syms))

	for _, s := range syms {
		if slices.Contains(openers, s) {
			stack = append(stack, s)
		} else {
			last := pop(&stack)
			if last == nil {
				return incomplete, false
			}

			switch s {
			case ")":
				if *last != "(" {
					return s, false
				}
			case "]":
				if *last != "[" {
					return s, false
				}
			case "}":
				if *last != "{" {
					return s, false
				}
			case ">":
				if *last != "<" {
					return s, false
				}
			}
		}
	}

	return "", true
}

// PartOne returns the total syntax error score for the input
func PartOne(in _ParseResult) int {
	sum := 0

	for _, syms := range in {
		a, ok := checkLine(syms)
		if ok {
			continue
		}

		// nb: zero value is 0, so missing table entries will be 0
		sum += scoreTable[a]
	}

	return sum
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
