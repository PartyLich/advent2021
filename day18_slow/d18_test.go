package day18_slow

import (
	"fmt"
	"os"
	"testing"
)

const ex string = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`

func TestPartOne(t *testing.T) {
	t.Run("parses input", func(t *testing.T) {
		_, err := parseLines(ex)
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}
	})
	t.Run("explode number", func(t *testing.T) {
		cases := []struct {
			in   SnailNum
			want SnailNum
			idx  int
		}{
			{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]", 0},
			{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]", 4},
			{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]", 3},
			{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", 3},
			{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]", 7},
			{
				"[[[[4,0],[5,4]],[[7,0],[15,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]",
				"[[[[4,0],[5,4]],[[7,0],[15,5]]],[10,[[0,[11,3]],[[6,3],[8,8]]]]]",
				9,
			},
		}

		for _, c := range cases {
			idx, ok := canExplode(c.in)
			if !ok {
				t.Errorf("canExplode(%v) \n\thave %v\n\twant %v", c.in, ok, true)
			}
			if idx != c.idx {
				t.Errorf("canExplode(%v) index \n\thave %v\n\twant %v", c.in, idx, c.idx)
			}

			have := explode(c.in, idx)
			if have != c.want {
				t.Errorf("explode(%v) \n\thave %v\n\twant %v", c.in, have, c.want)
			}
		}
	})
	t.Run("split number", func(t *testing.T) {
		cases := []struct {
			in   SnailNum
			want SnailNum
		}{
			{"[[[[0,7],4],[15,[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"},
			{"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"},
		}

		for _, c := range cases {

			ok := canSplit(c.in)
			if !ok {
				t.Errorf("canSplit(%v) \n\thave %v\n\twant %v", c.in, ok, true)
			}

			have := split(c.in)
			if have != c.want {
				t.Errorf("split(%v) \n\thave %v\n\twant %v", c.in, have, c.want)
			}
		}
	})
	t.Run("reduce number", func(t *testing.T) {
		cases := []struct {
			in   SnailNum
			want SnailNum
		}{
			{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
			{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
			{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
			{"[[1,2],[[3,4],5]]", "[[1,2],[[3,4],5]]"},
		}

		for _, c := range cases {
			have := reduce(c.in)
			if have != c.want {
				t.Errorf("reduce(%v) \n\thave %v\n\twant %v", c.in, have, c.want)
			}
		}
	})
	t.Run("add snailfish numbers", func(t *testing.T) {
		cases := []struct {
			in   string
			want SnailNum
		}{
			{`[1,1]
[2,2]
[3,3]
[4,4]`,
				"[[[[1,1],[2,2]],[3,3]],[4,4]]"},
			{`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
`,
				"[[[[3,0],[5,3]],[4,4]],[5,5]]"},
			{`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]
`,
				"[[[[5,0],[7,4]],[5,5]],[6,6]]"},
			{`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
			`,
				"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"},
			{`[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
			`,
				"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]"},

			{`[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]`,
				"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]"},

			{`[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]
[7,[5,[[3,8],[1,4]]]]`,
				"[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]"},

			{`[[[[7,7],[7,8]],[[9,5],[8,7]]],[[[6,8],[0,8]],[[9,9],[9,0]]]]
[[2,[2,2]],[8,[8,1]]]`,
				"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]"},
			{`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`,
				"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
		}

		for _, c := range cases {
			ns, err := parseLines(c.in)
			if err != nil {
				t.Fatalf("parse failure: %v", err)
			}

			have := ns[0]
			for i := 1; i < len(ns); i++ {
				have = add(have, ns[i])
			}

			if have != c.want {
				t.Errorf("add(%v) \n\thave %v\n\twant %v", c.in, have, c.want)
			}
		}
	})
	t.Run("calc magnitude", func(t *testing.T) {
		cases := []struct {
			in   SnailNum
			want int
		}{
			{"[9,1]", 29},
			{"[1,9]", 21},
			{"[[9,1],[1,9]]", 129},
			{"[[1,2],[[3,4],5]]", 143},
			{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
			{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
			{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
			{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
			{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
		}

		for _, c := range cases {
			have := magnitude(c.in)
			if have != c.want {
				t.Errorf("magnitude(%v) \n\thave %v\n\twant %v", c.in, have, c.want)
			}
		}
	})
	t.Run("executes task", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("parse failure: %v", err)
		}

		want := 4140

		have := PartOne(in)
		if have != want {
			t.Fatalf("PartOne(%#v) == %v, want %v", in, have, want)
		}
	})
}

func TestPartTwo(t *testing.T) {
	t.Run("executes task", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("parse failure: %v", err)
		}

		want := 3993

		have := PartTwo(in)
		if have != want {
			t.Fatalf("PartOne(%#v) == %v, want %v", in, have, want)
		}
	})
}

func BenchmarkPartOne(b *testing.B) {
	fileName := fmt.Sprintf("../input/%v.txt", "18")
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
	fileName := fmt.Sprintf("../input/%v.txt", "18")
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
