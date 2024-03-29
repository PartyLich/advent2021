// Day 21: Dirac Dice
package day21

import (
	"errors"
	"fmt"
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

func movePlayer(p Player, move int) Player {
	p.pos = (p.pos + move) % 10
	p.score += p.pos + 1

	return p
}

func victory(win int) func(Player) bool {
	return func(p Player) bool {
		return p.score >= win
	}
}

func play(state _ParseResult, turn int) [2]int {
	win := victory(21)
	cache := make(map[string][2]int)
	var helper func(_ParseResult, int) [2]int

	helper = func(state _ParseResult, turn int) [2]int {
		if win(state[0]) {
			return [2]int{1, 0}
		}
		if win(state[1]) {
			return [2]int{0, 1}
		}

		result := [2]int{0, 0}
		moves := [7][]int{
			{3, 1},
			{4, 3},
			{5, 6},
			{6, 7},
			{7, 6},
			{8, 3},
			{9, 1},
		}

		for _, m := range moves {
			s := make(_ParseResult, 2)
			copy(s, state)

			s[turn] = movePlayer(state[turn], m[0])

			key := fmt.Sprintf("%v,%v", s, (turn+1)%2)
			var r [2]int
			if c, ok := cache[key]; ok {
				r = c
			} else {
				r = helper(s, (turn+1)%2)
				cache[key] = r
			}

			result[0], result[1] = result[0]+(r[0]*m[1]), result[1]+(r[1]*m[1])
		}

		return result
	}

	return helper(state, turn)
}

// PartTwo returns the number of universes the winning player wins in.
func PartTwo(in _ParseResult) int {
	// 27 universes each turn
	// 3   1         = 1
	// 4   2 + 1 + 0 = 3
	// 5   3 + 2 + 1 = 6
	// 6   2 + 3 + 2 = 7
	// 7   1 + 2 + 3 = 6
	// 8   0 + 1 + 2 = 3
	// 9   1         = 1
	r := play(in, 0)

	return runner.Max(r[0], r[1])
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
