// Day 18: Snailfish
package day18

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type SnailNum string
type _ParseResult []SnailNum

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := make([]SnailNum, len(lines))
	for i, l := range lines {
		result[i] = SnailNum(l)
	}

	return result, nil
}

func canExplode(in SnailNum) (int, bool) {
	return 0, false
}

func explode(in SnailNum, i int) SnailNum {
	return ""
}

var reSplit = regexp.MustCompile(`\d{2,}`)

func canSplit(in SnailNum) bool {
	return reSplit.MatchString(string(in))
}

func split(in SnailNum) SnailNum {
	count := 0
	split := func(s string) string {
		if count > 0 {
			return s
		}
		count += 1

		i, _ := strconv.Atoi(s)
		l := i / 2
		r := l + (i % 2)

		return fmt.Sprintf("[%v,%v]", l, r)
	}

	s := reSplit.ReplaceAllStringFunc(string(in), split)
	return SnailNum(s)
}

func reduce(in string) string {
	return ""
}

func add(a, b string) string {
	return ""
}

func magnitude(in string) int {
	return 0
}

// PartOne returns the magnitude of the final sum.
func PartOne(in _ParseResult) int {
	if len(in) == 0 {
		return 0
	}

	result := in[0]
	for i := 1; i < len(in); i++ {
		result = add(result, in[i])
	}

	return magnitude(result)
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
