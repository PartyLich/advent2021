package main

import (
	"fmt"
	"log"
	"os"

	"github.com/partylich/advent2021/day01"
)

// handleErr logs and exits on error
func handleErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	fmt.Println("Advent of Code 2021")

	input, err := os.ReadFile("./input/01.txt")
	handleErr(err)

	inp01, err := day01.Parse(string(input))
	handleErr(err)

	fmt.Printf("\tDay 01.1: %v\n", day01.PartOne(inp01))
	os.Exit(0)
}
