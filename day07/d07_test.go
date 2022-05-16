package day07

import (
	"fmt"
	"os"
	"testing"
)

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

func TestPartTwo(t *testing.T) {
	in, err := parsePos(ex)
	if err != nil {
		t.Fatalf("Parse failure")
	}

	want := 168

	have := PartTwo(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}

func BenchmarkPartOne(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "07")
	input, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}

	in, err := parsePos(string(input))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartOne(in)
	}
}
