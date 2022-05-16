package runner

import (
	"fmt"
	"os"
	"time"
)

// Unimpl returns the string UNIMPLEMENTED
func Unimpl(_i interface{}) interface{} { return "UNIMPLEMENTED" }

// Solution defines methods required for each AoC puzzle solution
type Solution struct {
	Parse func(string) (interface{}, error)
	Fn    [2]func(interface{}) interface{}
}

func duration(val interface{}, start time.Time) (interface{}, string) {
	return val, fmt.Sprintf("%v", time.Since(start))
}

// runDay executes the provided solution using the supplied input filename
func RunDay(day string, solution Solution) error {
	fileName := fmt.Sprintf("./input/%v.txt", day)
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	for part, fn := range solution.Fn {
		start := time.Now()

		inp, err := solution.Parse(string(input))
		if err != nil {
			return err
		}

		res, dur := duration(fn(inp), start)
		fmt.Printf("\tDay %v.%v: %v\t%v\n", day, part, res, dur)
	}

	return nil
}

// Max returns the largest of two integers
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Min returns the smallest of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// NewGrid creates a grid with r rows and c columns
func NewGrid[T any](r, c int) [][]T {
	grid := make([][]T, r)
	for i := range grid {
		grid[i] = make([]T, c)
	}

	return grid
}
