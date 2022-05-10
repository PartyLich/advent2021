package day01

import (
	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

var Parse = parse.UintList

// PartOne counts the number of times a depth measurement increases from the previous measurement.
// (There is no measurement before the first measurement.)
func PartOne(depths []uint) uint {
	var count uint = 0
	for i, depth := range depths[1:] {
		if depth > depths[i] {
			count += 1
		}
	}

	return count
}

func sum(values []uint) uint {
	var s uint = 0
	for _, v := range values {
		s += v
	}

	return s
}

// PartTwo counts the number of times the sum of measurements in each 3 value sliding window
// increases from the previous sum.
func PartTwo(depths []uint) uint {
	var count uint = 0

	end := len(depths) - 3
	// day one and I already deeply miss Option types
	var prev *uint
	for i := range depths[:end] {
		if prev == nil {
			prev = new(uint)
			*prev = sum(depths[i : i+3])
		}
		next := sum(depths[i+1 : i+4])
		if *prev < next {
			count += 1
		}
		*prev = next
	}

	return count
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i) },
		One:   func(i interface{}) interface{} { return PartOne(i.([]uint)) },
		Two:   func(i interface{}) interface{} { return PartTwo(i.([]uint)) },
	}
}
