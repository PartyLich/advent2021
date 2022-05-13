// Day 3: Binary Diagnostic
package day03

import (
	"strconv"
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

// partition splits a list of strings representing binary numbers based on the
// value of digit at place
func partition(list []string, place int) ([]string, []string) {
	var (
		temp0 []string
		temp1 []string
	)

	for _, num := range list {
		// we definitely have only '0' and '1' in the strings, no utf-8 runes
		if string(num[place]) == "0" {
			temp0 = append(temp0, num)
		} else {
			temp1 = append(temp1, num)
		}
	}

	return temp0, temp1
}

// PartTwo multiplies the oxygen generator rating by the CO2 scrubber rating
func PartTwo(in []string) int {
	oxy := in
	co2 := in

	for i := 0; len(oxy) > 1; i++ {
		temp0, temp1 := partition(oxy, i)

		if len(temp0) > len(temp1) {
			// most common is 0
			oxy = temp0
		} else {
			// most common is 1
			oxy = temp1
		}
	}

	for i := 0; len(co2) > 1; i++ {
		temp0, temp1 := partition(co2, i)

		if len(temp0) > len(temp1) {
			// most common is 0
			co2 = temp1
		} else {
			// most common is 1
			co2 = temp0
		}
	}

	const binary int = 2
	const size int = 64
	oxyRate, err := strconv.ParseInt(oxy[0], binary, size)
	if err != nil {
		panic(err)
	}
	co2Rate, err := strconv.ParseInt(co2[0], binary, size)
	if err != nil {
		panic(err)
	}

	return int(oxyRate * co2Rate)
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i), nil },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.([]string)) },
			func(i interface{}) interface{} { return PartTwo(i.([]string)) },
		},
	}
}
