package day18

import (
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
