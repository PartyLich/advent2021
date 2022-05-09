package day01

import "testing"

func TestPartOne(t *testing.T) {
	in := []uint{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}
	var want uint = 7

	have := PartOne(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}

func TestPartTwo(t *testing.T) {
	in := []uint{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}
	var want uint = 5

	have := PartTwo(in)
	if have != want {
		t.Fatalf("PartTwo(%v) == %v, want %v", in, have, want)
	}
}
