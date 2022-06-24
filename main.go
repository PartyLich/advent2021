package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/partylich/advent2021/day01"
	"github.com/partylich/advent2021/day02"
	"github.com/partylich/advent2021/day03"
	"github.com/partylich/advent2021/day04"
	"github.com/partylich/advent2021/day05"
	"github.com/partylich/advent2021/day06"
	"github.com/partylich/advent2021/day07"
	"github.com/partylich/advent2021/day08"
	"github.com/partylich/advent2021/day09"
	"github.com/partylich/advent2021/day10"
	"github.com/partylich/advent2021/day11"
	"github.com/partylich/advent2021/day12"
	"github.com/partylich/advent2021/day13"
	"github.com/partylich/advent2021/day14"
	"github.com/partylich/advent2021/day15"
	"github.com/partylich/advent2021/day16"
	"github.com/partylich/advent2021/day17"
	"github.com/partylich/advent2021/day18"
	"github.com/partylich/advent2021/day20"
	"github.com/partylich/advent2021/day21"
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
		"01": day01.Solution(),
		"02": day02.Solution(),
		"03": day03.Solution(),
		"04": day04.Solution(),
		"05": day05.Solution(),
		"06": day06.Solution(),
		"07": day07.Solution(),
		"08": day08.Solution(),
		"09": day09.Solution(),
		"10": day10.Solution(),
		"11": day11.Solution(),
		"12": day12.Solution(),
		"13": day13.Solution(),
		"14": day14.Solution(),
		"15": day15.Solution(),
		"16": day16.Solution(),
		"17": day17.Solution(),
		"18": day18.Solution(),
		"20": day20.Solution(),
		"21": day21.Solution(),
	}

	start := time.Now()

	switch *userVal {
	case "all":
		// run all days
		for k, v := range m {
			handleErr(runner.RunDay(k, v))
		}
	default:
		// run single day based on user input
		s, ok := m[*userVal]
		if !ok {
			fmt.Printf("Invalid solution requested: %v\n", *userVal)
			os.Exit(1)
		}

		handleErr(runner.RunDay(*userVal, s))
	}

	fmt.Printf("Total elapsed: %v\n", time.Since(start))

	os.Exit(0)
}
