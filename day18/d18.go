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
	open := 0
	idx := -1
	lastDigit := false
	for _, r := range in {
		switch r {
		case '[':
			open += 1
			if open == 5 {
				idx += 1
				return idx, true
			}
			lastDigit = false
		case ']':
			open -= 1
			lastDigit = false
		case ',':
			lastDigit = false
		default:
			if !lastDigit {
				idx += 1
			}
			lastDigit = true
		}
	}

	return 0, false
}

func explode(in SnailNum, i int) SnailNum {
	if len(in) == 0 {
		panic("explode requires valid SnailNum, received empty string")
	}

	result := string(in)
	re := regexp.MustCompile(`(-?\d+)`)
	digits := re.FindAllString(result, -1)

	d := make([]int, len(digits))
	for i, digit := range digits {
		d[i], _ = strconv.Atoi(digit)
	}

	// pair's left value is added to the first regular number to the left of the
	// exploding pair (if any)
	if i-1 >= 0 {
		d[i-1] += d[i]
	}
	// pair's right value is added to the first regular number to the right of
	// the exploding pair (if any).
	if i+2 < len(digits) {
		d[i+2] += d[i+1]
	}
	d[i], d[i+1] = -42, -42

	j := 0
	subsDigit := func(s string) string {
		result := fmt.Sprintf("%v", d[j])
		j++

		return result
	}
	result = re.ReplaceAllStringFunc(result, subsDigit)

	reZero := regexp.MustCompile(`\[-42,-42\]`)

	return SnailNum(reZero.ReplaceAllString(result, "0"))
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
