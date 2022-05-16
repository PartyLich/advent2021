package day07

import "testing"

const ex string = `16,1,2,0,4,2,7,1,2,14`

func TestPartOne(t *testing.T) {
	in, err := parsePos(ex)
	if err != nil {
		t.Fatalf("Parse failure")
	}

	want := 37

	have := PartOne(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}
