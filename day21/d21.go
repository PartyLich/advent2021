// Day 21: Dirac Dice
package day21

import (
	"errors"
	"strconv"
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type _ParseResult = []Player

type Player struct {
	pos, score int
}

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := make([]Player, 2)

	for r, l := range lines {
		row := strings.Split(l, " ")
		posStr := row[len(row)-1]
		pos, err := strconv.Atoi(posStr)
		if err != nil {
			return nil, errors.New("parse failure")
		}

		result[r] = Player{pos - 1, 0}
	}

	return result, nil
}

// PartOne returns the product of the score of the losing player and the number
// of times the die was rolled.
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
