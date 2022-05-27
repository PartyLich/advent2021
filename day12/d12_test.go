package day12

import "testing"

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
