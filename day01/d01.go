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

var Solution = runner.Solution{
	Parse: func(i string) (interface{}, error) { return Parse(i) },
	One:   func(i interface{}) interface{} { return PartOne(i.([]uint)) },
	Two:   func(i interface{}) interface{} { return runner.Unimpl(i) },
}
