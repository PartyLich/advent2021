package day14

import (
	"fmt"
	"os"
	"testing"
)

const ex string = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`

func TestPartOne(t *testing.T) {
	t.Run("parses input", func(t *testing.T) {
		_, err := parseLines(ex)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("executes task", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("Parse failure")
		}

		want := 1588

		have := PartOne(in)
		if have != want {
			t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
		}
	})
}

func TestPartTwo(t *testing.T) {
	t.Run("executes task", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("Parse failure")
		}

		want := 2188189693529

		have := PartTwo(in)
		if have != want {
			t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
		}
	})
}

func BenchmarkPartOne(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "14")
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
	fileName := fmt.Sprintf("../input/%v.txt", "14")
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
