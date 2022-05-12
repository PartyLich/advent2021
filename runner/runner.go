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
	One   func(interface{}) interface{}
	Two   func(interface{}) interface{}
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

	inp, err := solution.Parse(string(input))
	if err != nil {
		return err
	}

	start := time.Now()
	res, dur := duration(solution.One(inp), start)
	fmt.Printf("\tDay %v.1: %v\t%v\n", day, res, dur)

	start = time.Now()
	res, dur = duration(solution.Two(inp), start)
	fmt.Printf("\tDay %v.2: %v\t%v\n", day, res, dur)

	return nil
}

func SliceCompare[T comparable](a []T, b []T) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
