package day09

import (
	"testing"
)

const ex string = `2199943210
3987894921
9856789892
8767896789
9899965678
`

func TestPartOne(t *testing.T) {
	in, err := parseLines(ex)
	if err != nil {
		t.Fatalf("Parse failure")
	}

	want := 15

	have := PartOne(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}

func TestPartTwo(t *testing.T) {
	in, err := parseLines(ex)
	if err != nil {
		t.Fatalf("Parse failure")
	}

	want := 1134

	have := PartTwo(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}
