package main

import (
	"fmt"
	"log"
	"os"
)

// handleErr logs and exits on error
func handleErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	fmt.Println("Advent of Code 2021")
	os.Exit(0)
}
