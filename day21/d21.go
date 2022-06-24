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

type Die struct {
	sides int
	val   int
}

func NewDie(sides int) *Die {
	return &Die{sides, 0}
}

func (d *Die) Next() int {
	result := d.val + 1
	d.val = (d.val + 1) % d.sides

	return result
}

// PartOne returns the product of the score of the losing player and the number
// of times the die was rolled.
func PartOne(in _ParseResult) int {
	var (
		rolls int
		p     int // current player
	)
	d := NewDie(100)

play:
	for ; ; p = (p + 1) % 2 {
		var move int
		for i := 0; i < 3; i++ {
			move += d.Next()
			rolls += 1
		}

		in[p].pos = (in[p].pos + move) % 10
		in[p].score += in[p].pos + 1

		if in[p].score >= 1000 {
			break play
		}
	}
	p = (p + 1) % 2

	return rolls * in[p].score
}

// PartTwo returns the number of universes the winning player wins in.
func PartTwo(in _ParseResult) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.(_ParseResult)) },
			runner.Unimpl,
		},
	}
}
