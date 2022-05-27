// Day 12: Passage Pathing
package day12

import (
	"strings"
	"unicode"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
	"golang.org/x/exp/slices"
)

type _ParseResult = map[string]map[string]bool

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := make(_ParseResult)

	for _, l := range lines {
		nodes := strings.Split(l, "-")
		a, b := nodes[0], nodes[1]

		if result[a] == nil {
			result[a] = make(map[string]bool)
		}
		if result[b] == nil {
			result[b] = make(map[string]bool)
		}
		// consider start->node edges unidirectional
		if a == "start" {
			result[a][b] = true
			continue
		}
		if b == "start" {
			result[b][a] = true
			continue
		}

		result[a][b] = true
		result[b][a] = true
	}

	return result, nil
}

func isLower(s string) bool {
	for _, c := range s {
		if unicode.IsUpper(c) {
			return false
		}
	}

	return true
}

func traverse(graph _ParseResult, visited []string, from string) [][]string {
	var paths [][]string

	for adj := range graph[from] {
		// dont revisit lowercase nodes
		if isLower(adj) && slices.Contains(visited, adj) {
			continue
		}

		v := make([]string, len(visited), len(visited)+2)
		copy(v, visited)
		v = append(v, adj)

		// completed path
		if adj == "end" {
			paths = append(paths, v)
			continue
		}

		paths = append(paths, traverse(graph, v, adj)...)
	}

	return paths
}

// PartOne returns the number of paths from start to end, where lowercase nodes
// may be visited multiple times
func PartOne(in _ParseResult) int {
	visited := []string{"start"}
	paths := traverse(in, visited, "start")

	return len(paths)
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.(_ParseResult)) },
			runner.Unimpl,
		},
	}
}
