package day15

import (
	"testing"
)

const ex string = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
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

		want := 40

		have := PartOne(in)
		if have != want {
			t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
		}
	})
}
