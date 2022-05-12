// Day 4: Giant Squid
package day04

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

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

func updatePlayer(p *player, num string) {
	for r, row := range p.board {
		if c := slices.Index(row, num); c != -1 {
			p.rowSum[r] += 1
			p.colSum[c] += 1
			// why the actual fuck doesnt the math package have a max fn for int?
			// ...especially when int is the DEFAULT number type. jfc
			if p.rowSum[r] > p.max {
				p.max = p.rowSum[r]
			}
			if p.colSum[c] > p.max {
				p.max = p.colSum[c]
			}

			p.marked[r][c] = true

			return
		}
	}
}

func getScore(p *player, lastCall string) int {
	var score int

	lastNum, err := strconv.Atoi(lastCall)
	if err != nil {
		panic(err)
	}

	// find the sum of all unmarked numbers
	for r, row := range p.board {
		for c, num := range row {
			if p.marked[r][c] {
				continue
			}
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			score += val
		}
	}

	return score * lastNum
}

// PartOne finds the sum of all unmarked numbers on a winning bingo board, then
// multiplies that sum by the last number called
func PartOne(game bingo) int {
	var score int

play:
	for _, num := range game.draws {
		for i := range game.players {
			updatePlayer(&game.players[i], num)

			if game.players[i].max == gridSize {
				score = getScore(&game.players[i], num)
				break play
			}
		}
	}

	return score
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return Parse(i) },
		One:   func(i interface{}) interface{} { return PartOne(i.(bingo)) },
		Two:   runner.Unimpl,
	}
}
