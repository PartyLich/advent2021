package day17

import (
	"testing"
)

const ex string = `target area: x=20..30, y=-10..-5
`

func TestPartOne(t *testing.T) {
	t.Run("parses input", func(t *testing.T) {
		_, err := parseLines(ex)
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}
	})
	t.Run("step", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("parse failure: %v", err)
		}

		cases := []struct {
			vx, vy int
			top    int
			ok     bool
		}{
			{17, -4, 0, false},
			{6, 9, 45, true},
		}

		for _, c := range cases {
			have, ok := step(in, c.vx, c.vy)
			if ok != c.ok {
				t.Errorf("Step(%#v, %v, %v) \n\tok? %v, want %v", ex, c.vx, c.vy, ok, c.ok)
			}
			if have.top[1] != c.top {
				t.Errorf("Step(%#v, %v, %v) \n\ttop %v, want %v", ex, c.vx, c.vy, have.top[1], c.top)
			}
		}
	})
	t.Run("executes task", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("parse failure: %v", err)
		}

		want := 45

		have := PartOne(in)
		if have != want {
			t.Fatalf("PartOne(%#v) == %v, want %v", in, have, want)
		}
	})
}
