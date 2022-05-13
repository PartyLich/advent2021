package slice

import (
	"testing"
)

func TestCompare(t *testing.T) {
	cases := []struct {
		a, b []int
		want bool
	}{
		{
			[]int{},
			[]int{},
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3},
			false,
		},
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
			false,
		},
	}

	for _, c := range cases {
		have := Compare(c.a, c.b)

		if have != c.want {
			t.Errorf("Compare(\n\t%v\n\t%v) \n\thave %v\n\twant %v", c.a, c.b, have, c.want)
		}
	}
}

func TestMap(t *testing.T) {
	inc := func(i int) int {
		i += 1
		return i
	}

	cases := []struct {
		a    []int
		f    func(int) int
		want []int
	}{
		{
			[]int{},
			inc,
			[]int{},
		},
		{
			[]int{1, 2, 3},
			inc,
			[]int{2, 3, 4},
		},
	}

	for _, c := range cases {
		have := Map(inc)(c.a)

		for i, v := range c.want {
			if have[i] != v {
				t.Errorf("Map(%T)(%v) \n\thave %v\n\twant %v", c.f, c.a, have, c.want)
			}
		}
	}
}

func TestReduce(t *testing.T) {
	sum := func(a, b int) int {
		return a + b
	}

	cases := []struct {
		a    []int
		fn   func(int, int) int
		init int
		want int
	}{
		{
			[]int{},
			sum,
			0,
			0,
		},
		{
			[]int{1, 2, 3},
			sum,
			0,
			6,
		},
		{
			[]int{1, 2, 3},
			func(a, b int) int {
				return a * b
			},
			1,
			6,
		},
	}

	for _, c := range cases {
		have := Reduce(c.fn)(c.a, c.init)

		if have != c.want {
			t.Errorf(
				"Reduce(%T)(%v, %v) \n\thave %v\n\twant %v",
				c.fn, c.a, c.init, have, c.want,
			)
		}
	}

	for _, c := range cases {
		have := Fold(c.fn)(c.a, c.init)

		if have != c.want {
			t.Errorf(
				"Reduce(%T)(%v, %v) \n\thave %v\n\twant %v",
				c.fn, c.a, c.init, have, c.want,
			)
		}
	}
}

func TestReduceRight(t *testing.T) {
	sub := func(a, b int) int {
		return a - b
	}

	cases := []struct {
		a    []int
		fn   func(int, int) int
		init int
		want int
	}{
		{
			[]int{},
			sub,
			0,
			0,
		},
		{
			[]int{1, 2, 3},
			sub,
			0,
			-6,
		},
		{
			[]int{1, 2, 3},
			func(a, b int) int {
				return a * b
			},
			1,
			6,
		},
	}

	for _, c := range cases {
		have := ReduceRight(c.fn)(c.a, c.init)

		if have != c.want {
			t.Errorf(
				"Reduce(%T)(%v, %v) \n\thave %v\n\twant %v",
				c.fn, c.a, c.init, have, c.want,
			)
		}
	}
}
