package day16

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Run("executes task", func(t *testing.T) {
		cases := []struct {
			in   string
			want int
		}{
			{"8A004A801A8002F478", 16},
			{"620080001611562C8802118E34", 12},
			{"C0015000016115A2E0802F182340", 23},
			{"A0016C880162017C3686B18A3D4780", 31},
		}

		for _, c := range cases {
			have := PartOne(_ParseResult(c.in))
			if have != c.want {
				t.Errorf("PartOne(%v) == %v, want %v", c.in, have, c.want)
			}
		}
	})
}
