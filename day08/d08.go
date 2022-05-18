// Day 8: Seven Segment Search
package day08

import (
	"math"
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

// map the frequency of each character
func sigMap(signals []string) []int {
	m := make([]int, 7)
	for _, s := range signals {
		for _, c := range s {
			// 'a' is 97, and we want it at idx 0
			m[c-97] += 1
		}
	}

	return m
}

const (
	// flags for each display segment
	t  = 0b1000000
	b  = 0b0100000
	c  = 0b0010000
	tl = 0b0001000
	tr = 0b0000100
	bl = 0b0000010
	br = 0b0000001

	// digits formed by each flag combination
	zero  = t | b | tl | tr | bl | br
	one   = tr | br
	two   = t | c | b | tr | bl
	three = t | c | b | tr | br
	four  = tl | c | tr | br
	five  = t | b | c | tl | br
	six   = t | b | c | tl | bl | br
	seven = t | br | tr
	eight = t | b | c | tl | tr | bl | br
	nine  = t | b | c | tl | tr | br
)

// return wire->segment mapping given a set of 10 signals
func decodeSignals(signals []string) map[rune]int {
	runeToVal := make(map[rune]int)
	valToRune := make(map[int]rune)

	m := sigMap(signals)
	for i, freq := range m {
		switch freq {
		case 4:
			c := i + 97
			runeToVal[rune(c)] = bl
			valToRune[bl] = rune(c)
		case 9:
			c := i + 97
			runeToVal[rune(c)] = br
			valToRune[br] = rune(c)
		case 6:
			c := i + 97
			runeToVal[rune(c)] = tl
			valToRune[tl] = rune(c)
		}
	}

	for len(valToRune) < 6 {
		for _, signal := range signals {
			switch len(signal) {
			case 2:
				// 1 - we have br, so we can find tr
				for _, v := range signal {
					if _, ok := runeToVal[v]; ok {
						continue
					}
					valToRune[tr] = v
					runeToVal[v] = tr
				}
			case 4:
				// 4 - not tl and not in 1 => center
				if _, ok := valToRune[tr]; !ok {
					break
				}
				for _, v := range signal {
					if _, ok := runeToVal[v]; ok {
						continue
					}
					valToRune[c] = v
					runeToVal[v] = c
				}
			case 3:
				// 7 - we have br, and will have tr once we loop over 1
				// so 7 gives us a definite top value
				if _, ok := valToRune[tr]; !ok {
					break
				}
				for _, v := range signal {
					if _, ok := runeToVal[v]; ok {
						continue
					}
					valToRune[t] = v
					runeToVal[v] = t
				}
			}
		}
	}

	return runeToVal
}

// decode displayed digits given wire->segment mapping
func decodeOutput(runeToVal map[rune]int, outputs []string) int {
	// from binary rep to int
	conv := map[int]int{
		zero:  0,
		one:   1,
		two:   2,
		three: 3,
		four:  4,
		five:  5,
		six:   6,
		seven: 7,
		eight: 8,
		nine:  9,
	}

	var num int
	for i, output := range outputs {
		digit := 0
		for _, c := range output {
			val, ok := runeToVal[c]
			if !ok {
				val = b
			}

			digit |= val
		}

		num += conv[digit] * int(math.Pow10((3 - i)))
	}

	return num
}

// PartTwo returns the sum of all decoded outputs
func PartTwo(in _ParseResult) int {
	sum := 0

	for _, line := range in {
		runeMap := decodeSignals(line[0])
		num := decodeOutput(runeMap, line[1])
		sum += num
	}

	return sum
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.(_ParseResult)) },
			func(i interface{}) interface{} { return PartTwo(i.(_ParseResult)) },
		},
	}
}
