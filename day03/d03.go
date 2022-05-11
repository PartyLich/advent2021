// Day 3: Binary Diagnostic
package day03

import (
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

var Parse = parse.Lines

// returns the frequency of 0 and 1 in each digit
func getFreq(in []string) [][2]int {
	digitCount := len(in[0])
	freq := make([][2]int, digitCount)

	for _, line := range in {
		digits := strings.Split(line, "")
		for i, d := range digits {
			switch d {
			case "0":
				freq[i][0] += 1
			case "1":
				freq[i][1] += 1
			}
		}
	}

	return freq
}

// PartOne uses the binary numbers in your diagnostic report to calculate the gamma rate and epsilon
// rate, then multiply them together.
func PartOne(in []string) int {
	freq := getFreq(in)

	gamma := 0
	epsilon := 0
	for i := len(freq) - 1; i >= 0; i-- {
		offset := len(freq) - 1 - i
		if freq[i][0] > freq[i][1] {
			// most common is 0
			epsilon |= 1 << offset
		} else {
			// most common is 1
			gamma |= 1 << offset
		}
	}

	return gamma * epsilon
}

// PartTwo multiplies the oxygen generator rating by the CO2 scrubber rating
func PartTwo(in []string) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i), nil },
		One:   func(i interface{}) interface{} { return PartOne(i.([]string)) },
		Two:   runner.Unimpl,
	}
}
