package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/partylich/advent2021/day01"
	"github.com/partylich/advent2021/day02"
	"github.com/partylich/advent2021/runner"
)

// handleErr logs and exits on error
func handleErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	var userVal = new(string)

	flag.StringVar(userVal, "d", "all", "run only the specified day, eg \"01\"")
	flag.StringVar(userVal, "day", "all", "run only the specified day, eg \"01\"")
	flag.Parse()

	fmt.Println("Advent of Code 2021")
	m := map[string]runner.Solution{
		"01": day01.Solution,
		"02": day02.Solution,
	}

	switch *userVal {
	case "all":
		// run all days
		for k, v := range m {
			handleErr(runner.RunDay(k, v))
		}
		break
	default:
		// run single day based on user input
		s, ok := m[*userVal]
		if !ok {
			fmt.Printf("Invalid solution requested: %v\n", *userVal)
			os.Exit(1)
		}

		handleErr(runner.RunDay(*userVal, s))
	}

	os.Exit(0)
}
