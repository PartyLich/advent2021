package day06

import "testing"

const ex string = `3,4,3,1,2`

func TestParse(t *testing.T) {
	have, err := parseFish(ex)
	if err != nil {
		t.Fatalf("Parse failure")
	}

	want := []int{0, 1, 1, 2, 1, 0, 0, 0, 0}

	for i, v := range want {
		if v != have[i] {
			t.Fatalf("Parse(%v) == %v, want %v", ex, have, want)
		}
	}
}

func TestPartOne(t *testing.T) {
	in, err := parseFish(ex)
	if err != nil {
		t.Fatalf("Parse failure")
	}

	want := 5934

	have := PartOne(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}
