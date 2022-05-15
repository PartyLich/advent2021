package day06

import (
	"fmt"
	"os"
	"testing"
)

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

func TestPartTwo(t *testing.T) {
	in, err := parseFish(ex)
	if err != nil {
		t.Fatalf("Parse failure")
	}

	want := 26984457539

	have := PartTwo(in)
	if have != want {
		t.Fatalf("PartTwo(%v) == %v, want %v", in, have, want)
	}
}

func BenchmarkPartOne(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "06")
	input, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}

	in, err := parseFish(string(input))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartOne(in)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "06")
	input, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}

	in, err := parseFish(string(input))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartTwo(in)
	}
}
