package day21

import (
	"testing"
)

const ex string = `Player 1 starting position: 4
Player 2 starting position: 8
`

func TestPartOne(t *testing.T) {
	t.Run("parses input", func(t *testing.T) {
		_, err := parseLines(ex)
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}
	})
	t.Run("deterministic die", func(t *testing.T) {
		d := NewDie(100)

		for i := 1; i <= 100; i++ {
			have := d.Next()
			if have != i {
				t.Fatalf("Die(100) \n\thave %v \n\twant %v", have, i)
			}
		}

		have := d.Next()
		if have != 1 {
			t.Fatalf("Die(100) \n\thave %v \n\twant %v", have, 1)
		}
	})
	t.Run("executes task", func(t *testing.T) {
		in, err := parseLines(ex)
		if err != nil {
			t.Fatalf("parse failure: %v", err)
		}

		want := 739785

		have := PartOne(in)
		if have != want {
			t.Fatalf("PartOne(%#v) \n\thave %v \n\twant %v", in, have, want)
		}
	})
}
