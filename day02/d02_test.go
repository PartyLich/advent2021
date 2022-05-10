package day02

import "testing"

func TestParse(t *testing.T) {
	in := `forward 5
down 5
forward 8
up 3
down 8
forward 2`
	want := []command{
		{Forward, 5},
		{Down, 5},
		{Forward, 8},
		{Up, 3},
		{Down, 8},
		{Forward, 2},
	}

	have, err := parseCom(in)
	if err != nil {
		t.Fatalf("parseCom(%v) err: %v, want %v", in, err, want)
	}
	for i, s := range want {
		if s != have[i] {
			t.Fatalf("Lines(%v) == %v, want %v", in, have, want)
		}
	}
}

func TestPartOne(t *testing.T) {
	in := []command{
		{Forward, 5},
		{Down, 5},
		{Forward, 8},
		{Up, 3},
		{Down, 8},
		{Forward, 2},
	}
	var want int = 150

	have := PartOne(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}

func TestPartTwo(t *testing.T) {
	in := []command{
		{Forward, 5},
		{Down, 5},
		{Forward, 8},
		{Up, 3},
		{Down, 8},
		{Forward, 2},
	}
	var want int = 900

	have := PartTwo(in)
	if have != want {
		t.Fatalf("PartOne(%v) \n have %v, want %v", in, have, want)
	}
}
