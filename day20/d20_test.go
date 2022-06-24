package day20

import (
	"testing"
)

const ex string = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###
`

func TestPartOne(t *testing.T) {
	t.Run("parses input", func(t *testing.T) {
		_, err := parseLines(ex)
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}
	})
	t.Run("gets index", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("parse failure: %v", err)
		}

		want := 34
		have := getIndex(in.image, 2, 2)
		if have != want {
			t.Fatalf("getIndex  \n\thave %v \n\t want %v", have, want)
		}
	})
	t.Run("executes task", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("parse failure: %v", err)
		}

		want := 35

		have := PartOne(in)
		if have != want {
			t.Fatalf("PartOne(%#v) == %v, want %v", in, have, want)
		}
	})
}

func TestPartTwo(t *testing.T) {
	t.Run("executes task x50", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("parse failure: %v", err)
		}

		want := 3351

		have := PartTwo(in)
		if have != want {
			t.Fatalf("PartOne(%#v) == %v, want %v", in, have, want)
		}
	})
}
