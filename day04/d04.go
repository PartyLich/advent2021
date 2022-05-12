// Day 4: Giant Squid
package day04

import (
	"strings"

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

func Parse(in string) (bingo, error) {
	groups := strings.Split(strings.TrimSpace(in), "\n\n")
	draws := strings.Split(groups[0], ",")
	boards := make([]player, len(groups)-1)

	for idx, group := range groups[1:] {
		rows := parse.Lines(group)
		board := make(board, len(rows))
		for r, row := range rows {
			board[r] = strings.Fields(row)
		}

		boards[idx] = newPlayer(board)
	}

	return bingo{
		draws,
		boards,
	}, nil
}

// PartOne
func PartOne(in []string) int {
	return 0
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i) },
		One:   func(i interface{}) interface{} { return PartOne(i.([]string)) },
		Two:   runner.Unimpl,
	}
}
