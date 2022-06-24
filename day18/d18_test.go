package day18

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
		h, err := parseLines(ex)
		if err != nil || h == nil {
			t.Errorf("parse failure: %v", err)
		}
	})

	t.Run("parses list", func(t *testing.T) {
		ex := `[9,[8,7]]`
		head, err := parseNum(ex)
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}

		h, err := parseNum("[9,1]")
		if err != nil || h == nil {
			t.Errorf("parse failure: %v", err)
		}

		want := []struct {
			val, depth int
		}{
			{9, 1},
			{8, 2},
			{7, 2},
		}

		for i, n := 0, head; n != nil; i, n = i+1, n.right {
			if n.val != want[i].val {
				t.Fatalf("parseNum(%#v) \n\thave %v, \n\twant %v", ex, n.val, want[i].val)
			}
			if n.depth != want[i].depth {
				t.Fatalf("parseNum(%#v) \n\thave %v, \n\twant %v", ex, n.depth, want[i].depth)
			}
		}
	})
	t.Run("calc magnitude", func(t *testing.T) {
		cases := []struct {
			in   string
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
			num, _ := parseNum(c.in)
			have := magnitude(num)
			if have != c.want {
				t.Errorf("magnitude(%v) \n\thave %v\n\twant %v", c.in, have, c.want)
			}
		}
	})
	t.Run("can explode number", func(t *testing.T) {
		cases := []struct {
			in   string
			want bool
		}{
			{"[[[[[9,8],1],2],3],4]", true},
			{"[7,[6,[5,[4,[3,2]]]]]", true},
			{"[[6,[5,[4,[3,2]]]],1]", true},
			{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", true},
			{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", true},
			{
				"[[[[4,0],[5,4]],[[7,0],[15,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]",
				true,
			},
		}

		for _, c := range cases {
			num, _ := parseNum(c.in)
			_, ok := canExplode(num)
			if !ok {
				t.Errorf("canExplode(%v) \n\thave %v\n\twant %v", c.in, ok, true)
			}
		}
	})
	t.Run("explode number", func(t *testing.T) {
		cases := []struct {
			in   string
			want []Node
		}{
			{"[[[[[9,8],1],2],3],4]",
				[]Node{
					{0, 4, nil, nil},
					{9, 4, nil, nil},
					{2, 3, nil, nil},
					{3, 2, nil, nil},
					{4, 1, nil, nil},
				},
			},
			{"[7,[6,[5,[4,[3,2]]]]]",
				[]Node{
					{7, 1, nil, nil},
					{6, 2, nil, nil},
					{5, 3, nil, nil},
					{7, 4, nil, nil},
					{0, 4, nil, nil},
				},
			},
			{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",

				[]Node{
					{3, 2, nil, nil},
					{2, 3, nil, nil},
					{8, 4, nil, nil},
					{0, 4, nil, nil},
					{9, 2, nil, nil},
					{5, 3, nil, nil},
					{4, 4, nil, nil},
					{3, 5, nil, nil},
					{2, 5, nil, nil},
				},
			},
			{"[[[[0,[3,2]],[3,3]],[4,4]],[5,5]]",
				[]Node{
					{3, 4, nil, nil},
					{0, 4, nil, nil},
					{5, 4, nil, nil},
					{3, 4, nil, nil},
					{4, 3, nil, nil},
					{4, 3, nil, nil},
					{5, 2, nil, nil},
					{5, 2, nil, nil},
				},
			},
		}

		for _, c := range cases {
			num, _ := parseNum(c.in)
			start, _ := canExplode(num)

			explode(num, start)
			for i, n := 0, num; n != nil; i, n = i+1, n.right {
				if n.val != c.want[i].val {
					t.Errorf("explode(%v) \n\thave %v\n\twant %v", c.in, n.val, c.want[i].val)
				}
				if n.depth != c.want[i].depth {
					t.Errorf("explode(%v) \n\thave %v\n\twant %v", c.in, n.depth, c.want[i].depth)
				}
			}
		}
	})
	t.Run("split number", func(t *testing.T) {
		cases := []struct {
			in   string
			want []Node
		}{
			{"[[[[0,7],4],[15,[0,13]]],[1,1]]",
				[]Node{
					{0, 4, nil, nil},
					{7, 4, nil, nil},
					{4, 3, nil, nil},
					{7, 4, nil, nil},
					{8, 4, nil, nil},
					{0, 4, nil, nil},
					{13, 4, nil, nil},
					{1, 2, nil, nil},
					{1, 2, nil, nil},
				},
			},
			{"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
				[]Node{
					{0, 4, nil, nil},
					{7, 4, nil, nil},
					{4, 3, nil, nil},
					{7, 4, nil, nil},
					{8, 4, nil, nil},
					{0, 4, nil, nil},
					{6, 5, nil, nil},
					{7, 5, nil, nil},
					{1, 2, nil, nil},
					{1, 2, nil, nil},
				},
			},
		}

		for _, c := range cases {
			num, _ := parseNum(c.in)
			start, ok := canSplit(num)
			if !ok {
				t.Errorf("canSplit(%v) \n\thave %v\n\twant %v", c.in, ok, true)
			}

			split(num, start)
			for i, n := 0, num; n != nil; i, n = i+1, n.right {
				if n.val != c.want[i].val {
					t.Errorf("split(%v) \n\thave %v\n\twant %v", c.in, n.val, c.want[i].val)
				}
				if n.depth != c.want[i].depth {
					t.Errorf("split(%v) \n\thave %v\n\twant %v", c.in, n.depth, c.want[i].depth)
				}
			}
		}
	})
	t.Run("reduce number", func(t *testing.T) {
		cases := []struct {
			in   string
			want []Node
		}{
			{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
				[]Node{
					{0, 4, nil, nil},
					{7, 4, nil, nil},
					{4, 3, nil, nil},
					{7, 4, nil, nil},
					{8, 4, nil, nil},
					{6, 4, nil, nil},
					{0, 4, nil, nil},
					{8, 2, nil, nil},
					{1, 2, nil, nil},
				},
			},
			{"[[[[[9,8],1],2],3],4]",
				[]Node{
					{0, 4, nil, nil},
					{9, 4, nil, nil},
					{2, 3, nil, nil},
					{3, 2, nil, nil},
					{4, 1, nil, nil},
				},
			},
			{"[[6,[5,[4,[3,2]]]],1]",
				[]Node{
					{6, 2, nil, nil},
					{5, 3, nil, nil},
					{7, 4, nil, nil},
					{0, 4, nil, nil},
					{3, 1, nil, nil},
				},
			},
			{"[[1,2],[[3,4],5]]",
				[]Node{
					{1, 2, nil, nil},
					{2, 2, nil, nil},
					{3, 3, nil, nil},
					{4, 3, nil, nil},
					{5, 2, nil, nil},
				},
			},
		}

		for _, c := range cases {
			num, _ := parseNum(c.in)

			have := reduce(num)
			for i, n := 0, have; n != nil; i, n = i+1, n.right {
				if n.val != c.want[i].val {
					t.Errorf("reduce(%v) val \n\thave %v\n\twant %v", c.in, n.val, c.want[i].val)
				}
				if n.depth != c.want[i].depth {
					t.Errorf("reduce(%v) depth \n\thave %v\n\twant %v", c.in, n.depth, c.want[i].depth)
				}
			}
		}
	})
	t.Run("add snailfish numbers", func(t *testing.T) {
		cases := []struct {
			in   string
			want []Node
		}{
			{`[1,1]
[2,2]
[3,3]
[4,4]`,
				[]Node{
					{1, 4, nil, nil},
					{1, 4, nil, nil},
					{2, 4, nil, nil},
					{2, 4, nil, nil},
					{3, 3, nil, nil},
					{3, 3, nil, nil},
					{4, 2, nil, nil},
					{4, 2, nil, nil},
				},
			},
			{`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`,
				[]Node{
					{3, 4, nil, nil},
					{0, 4, nil, nil},
					{5, 4, nil, nil},
					{3, 4, nil, nil},
					{4, 3, nil, nil},
					{4, 3, nil, nil},
					{5, 2, nil, nil},
					{5, 2, nil, nil},
				},
			},
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

			for i, n := 0, have; n != nil; i, n = i+1, n.right {
				if n.val != c.want[i].val {
					t.Errorf("add(%v) val \n\thave %v\n\twant %v", c.in, n.val, c.want[i].val)
				}
				if n.depth != c.want[i].depth {
					t.Errorf("add(%v) depth \n\thave %v\n\twant %v", c.in, n.depth, c.want[i].depth)
				}
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
