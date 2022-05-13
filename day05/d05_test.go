package day05

import (
	"fmt"
	"os"
	"testing"
)

const ex string = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestPartOne(t *testing.T) {
	in := Parse(ex)

	want := 5

	have := PartOne(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}

func BenchmarkPartOne(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "05")
	input, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}

	in := Parse(string(input))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartOne(in)
	}
}

func TestPartTwo(t *testing.T) {
	in := Parse(ex)
	want := 12

	have := PartTwo(in)
	if have != want {
		t.Fatalf("PartTwo(%v) == %v, want %v", in, have, want)
	}
}
