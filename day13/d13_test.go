package day13

import (
	"testing"
)

const ex string = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
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

		want := 17

		have := PartOne(in)
		if have != want {
			t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
		}
	})
}
