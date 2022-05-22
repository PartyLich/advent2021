package day11

import (
	"fmt"
	"os"
	"testing"
)

const ex string = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

func TestPartOne(t *testing.T) {
	in, err := parseLines(ex)
	if err != nil {
		t.Fatalf("Parse failure")
	}

	want := 1656

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

	want := 195

	have := PartTwo(in)
	if have != want {
		t.Fatalf("PartTwo(%v) == %v, want %v", in, have, want)
	}
}

func BenchmarkPartOne(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "11")
	input, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}

	in, err := parseLines(string(input))
	if err != nil {
		b.Fatalf("Parse failure")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartOne(in)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "11")
	input, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}

	in, err := parseLines(string(input))
	if err != nil {
		b.Fatalf("Parse failure")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartTwo(in)
	}
}
