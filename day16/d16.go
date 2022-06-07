// Day 16: Packet Decoder
package day16

import (
	"strings"

	"github.com/partylich/advent2021/runner"
)

type _ParseResult string

// PartOne returns the sum of the version numbers in all parsed packets
func PartOne(in _ParseResult) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return strings.TrimSpace(i), nil },
		Fn: [2]func(i interface{}) interface{}{
			runner.Unimpl,
			runner.Unimpl,
		},
	}
}
