package day12

import (
	"fmt"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`start-A
start-b
A-c
A-b
b-d
A-end
b-end
`, 10},
		{`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`, 19},
		{`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`, 226},
	}

	for _, c := range cases {
		in, err := parseLines(c.in)
		if err != nil {
			t.Fatalf("Parse failure")
		}

		have := PartOne(in)
		if have != c.want {
			t.Errorf("PartOne(%v): \n\thave %v, \n\twant %v", in, have, c.want)
		}
	}
}

func TestPartTwo(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`start-A
start-b
A-c
A-b
b-d
A-end
b-end
`, 36},
		{`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`, 103},
		{`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`, 3509},
	}

	for _, c := range cases {
		in, err := parseLines(c.in)
		if err != nil {
			t.Fatalf("Parse failure")
		}

		have := PartTwo(in)
		if have != c.want {
			t.Errorf("PartTwo(%v): \n\thave %v, \n\twant %v", in, have, c.want)
		}
	}
}

func BenchmarkPartOne(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "12")
	input, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}

	in, err := parseLines(string(input))
	if err != nil {
		b.Fatalf("Parse failure")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartOne(in)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "12")
	input, err := os.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}

	in, err := parseLines(string(input))
	if err != nil {
		b.Fatalf("Parse failure")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PartTwo(in)
	}
}
