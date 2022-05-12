// Day 4: Giant Squid
package day04

import (
	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

const gridSize int = 5

// int makes more sense, but i dont need them parsed to do comparisons
type board = [][]string

type player struct {
	board board
	// max number of filled spaces in any row or column
	max int
	// current sum of marked spaces in each row
	rowSum [gridSize]int
	// current sum of marked spaces in each col
	colSum [gridSize]int
	// marked status of each space on board
	marked [gridSize][gridSize]bool
}

func newPlayer(board board) player {
	var (
		rowSum [gridSize]int
		colSum [gridSize]int
		marked [gridSize][gridSize]bool
	)

	return player{
		board,
		0,
		rowSum,
		colSum,
		marked,
	}
}

// bingo represents a bingo game state
type bingo struct {
	draws   []string
	players []player
}

var Parse = parse.Lines

// PartOne
func PartOne(in []string) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i), nil },
		One:   func(i interface{}) interface{} { return PartOne(i.([]string)) },
		Two:   runner.Unimpl,
	}
}
