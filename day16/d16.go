// Day 16: Packet Decoder
package day16

import (
	"strings"

	"github.com/partylich/advent2021/runner"
)

type _ParseResult struct {
	msg    string
	msgLen int
	used   int
}

func parseLines(in string) (_ParseResult, error) {
	s := strings.TrimSpace(in)
	return _ParseResult{s, len(s), 0}, nil
}

// PartOne returns the sum of the version numbers in all parsed packets
func PartOne(in _ParseResult) int {
	return 0
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
