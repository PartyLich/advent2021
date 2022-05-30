// Day 13: Transparent Origami
package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type dir int

const (
	Up dir = iota
	Left
)

type Instruction struct {
	Dir dir
	Pt  int
}

type _ParseResult struct {
	Dots  [][]bool
	Instr []Instruction
}

func parseLines(in string) (_ParseResult, error) {
	g := strings.Split(in, "\n\n")
	lines := parse.Lines(g[0])
	pts := make([][]int, len(lines))

	var mX, mY int
	for r, l := range lines {
		p, err := parse.Csv(l, strconv.Atoi)
		if err != nil {
			panic("parse failure")
		}

		pts[r] = p
		mX = runner.Max(mX, p[0])
		mY = runner.Max(mY, p[1])
	}

	grid := runner.NewGrid[bool](mY+1, mX+1)
	for _, p := range pts {
		x, y := p[0], p[1]
		grid[y][x] = true
	}

	lines = parse.Lines(g[1])
	is := make([]Instruction, len(lines))
	for r, i := range lines {
		parts := strings.Split(i[11:], "=")

		var d dir
		switch parts[0] {
		case "y":
			d = Up
		case "x":
			d = Left
		default:
			panic("parse failure")
		}

		p, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("parse failure")
		}

		is[r] = Instruction{d, p}
	}

	return _ParseResult{grid, is}, nil
}

func fold(dots [][]bool, i Instruction) [][]bool {
	var (
		rInit, cInit int
		up           bool
		next         [][]bool
	)

	switch i.Dir {
	case Up:
		rInit = i.Pt + 1
		next = make([][]bool, i.Pt)
		copy(next, dots)
		up = true
	case Left:
		cInit = i.Pt + 1
		next = runner.NewGrid[bool](len(dots), i.Pt)
		for r, row := range dots {
			copy(next[r], row)
		}
	}

	for r := rInit; r < len(dots); r++ {
		for c := cInit; c < len(dots[0]); c++ {
			nextR, nextC := r, c
			if up {
				delta := r - i.Pt
				nextR = i.Pt - delta
			} else {
				delta := c - i.Pt
				nextC = i.Pt - delta
			}

			next[nextR][nextC] = next[nextR][nextC] || dots[r][c]
		}
	}

	return next
}

// PartOne returns how many dots are visible after completing the first fold
// instruction.
func PartOne(in _ParseResult) int {
	next := fold(in.Dots, in.Instr[0])

	count := 0
	for _, r := range next {
		for _, hasDot := range r {
			if hasDot {
				count += 1
			}
		}
	}

	return count

}

// PartTwo returns nonsense, but prints out the result of following all fold
// instructions.
func PartTwo(in _ParseResult) int {
	next := fold(in.Dots, in.Instr[0])

	for _, i := range in.Instr[1:] {
		next = fold(next, i)
	}

	for _, r := range next {
		for _, hasDot := range r {
			if hasDot {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return -1
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
