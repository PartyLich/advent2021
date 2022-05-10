package parse

import "testing"

func TestLines(t *testing.T) {
	in := `foo
bar
baz`
	var want = []string{"foo", "bar", "baz"}

	have := Lines(in)
	for i, s := range have {
		if s != want[i] {
			t.Fatalf("Lines(%v) == %v, want %v", in, have, want)
		}
	}
}

func TestPartTwo(t *testing.T) {
	in := `199
200
208
210
200
207
240
269
260
263`
	want := []uint{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	have, err := UintList(in)
	if err != nil {
		t.Fatalf("UintList(%v) error %v, want %v", in, err, want)
	}
	for i, s := range have {
		if s != want[i] {
			t.Fatalf("Lines(%v) == %v, want %v", in, have, want)
		}
	}
}

func TestIntList(t *testing.T) {
	in := `199
200
208
210
200
207
240
269
260
263`
	want := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	have, err := IntList(in)
	if err != nil {
		t.Fatalf("IntList(%v) error %v, want %v", in, err, want)
	}
	for i, s := range have {
		if s != want[i] {
			t.Fatalf("Lines(%v) == %v, want %v", in, have, want)
		}
	}
}
