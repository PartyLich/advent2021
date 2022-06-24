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
