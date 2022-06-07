// Day 16: Packet Decoder
package day16

import (
	"math"
	"strconv"
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

type Ptype int
type Pversion int

// Every packet begins with a standard header: the first three bits encode the
// packet version, and the next three bits encode the packet type ID.
func header(in *_ParseResult) (Ptype, Pversion) {
	i := in.used / 4
	skip := in.used % 4
	nibbles := int(math.Ceil(float64(6+skip) / 4))
	s := in.msg[i : i+nibbles]

	bits, _ := strconv.ParseInt(s, 16, 64)

	bits >>= (nibbles * 4) - skip - 6

	pType := Ptype((bits) & 0b111)
	pVer := Pversion(bits>>3) & 0b111

	in.used += 6

	return pType, pVer
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
