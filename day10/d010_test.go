package day10

import (
	"strings"
	"testing"
)

const ex string = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

func TestPartOne(t *testing.T) {
	in, err := parseLines(ex)
	if err != nil {
		t.Fatalf("Parse failure")
	}

	want := 26397

	have := PartOne(in)
	if have != want {
		t.Fatalf("PartOne(%v) == %v, want %v", in, have, want)
	}
}

func TestPop(t *testing.T) {
	s := []int{1, 2, 3}
	want := 3

	have := pop(&s)
	if have == nil {
		t.Fatalf("Pop(%v) == <nil>, want %v", s, have)
	}

	if *have != want {
		t.Fatalf("Pop(%v) == %v, want %v", s, *have, want)
	}

	if len(s) != 2 {
		t.Fatalf("Pop(%v) \n\tlen == %v, want %v", s, len(s), 2)
	}
}

func TestCheckLine(t *testing.T) {
	cases := []struct{ in, want string }{
		{"{([(<{}[<>[]}>{[]{[(<()>", "}"},
		{"[[<[([]))<([[{}[[()]]]", ")"},
		{"[<(<(<(<{}))><([]([]()", ")"},
		{"<{([([[(<>()){}]>(<<{{", ">"},
	}

	for _, c := range cases {
		in := strings.Split(c.in, "")
		have, ok := checkLine(in)
		if ok {
			t.Fatalf("CheckLine(%v) should be corrupt, got ok", in)
		}

		if have != c.want {
			t.Fatalf("CheckLine(%v) \n\thave %v, want %v", c.in, have, c.want)
		}
	}
}
