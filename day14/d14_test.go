package day14

import (
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
