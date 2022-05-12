package day04

import (
	"testing"

	"github.com/partylich/advent2021/runner"
)

func TestParse(t *testing.T) {
	in := `7,4,9,5

22 13 17
 8  2 23
21  9 14
`

	have, err := Parse(in)
	if err != nil {
		t.Fatalf("Parse(%v) error %v", in, err)
	}

	draws := []string{"7", "4", "9", "5"}
	for i, d := range draws {
		if d != have.draws[i] {
			t.Fatalf("Parse draws\n\thave %v\n\twant %v", have.draws, draws)
		}
	}

	player := newPlayer(
		[][]string{
			{"22", "13", "17"},
			{"8", "2", "23"},
			{"21", "9", "14"},
		},
	)
	for i, row := range player.board {
		if !runner.SliceCompare[string](row, have.players[0].board[i]) {
			t.Fatalf("Parse board(%v) == %v, want %v", in, have.players[0].board[i], row)
		}
	}
}

func TestPartOne(t *testing.T) {
	in, err := Parse(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`)
	if err != nil {
		t.Fatalf("Parse(%v) error %v", in, err)
	}

	want := 4512

	have := PartOne(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}

func TestPartTwo(t *testing.T) {
	in, err := Parse(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`)
	if err != nil {
		t.Fatalf("Parse(%v) error %v", in, err)
	}

	want := 1924

	have := PartTwo(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}
