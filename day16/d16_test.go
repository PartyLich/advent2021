package day16

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Run("parse packet header", func(t *testing.T) {
		cases := []struct {
			ex _ParseResult
			t  Ptype
			v  Pversion
		}{
			{_ParseResult{`D2`, 2, 0}, 4, 6},
			{_ParseResult{`38006F45291200`, 14, 22}, 4, 6},
			{_ParseResult{`38006F45291200`, 14, 33}, 4, 2},
		}

		for _, c := range cases {
			pt, v := header(&c.ex)

			if pt != c.t {
				t.Errorf("parse header type(%v) \n\thave %b\n\twant %b", c.ex, pt, c.t)
			}
			if v != c.v {
				t.Errorf("parse header version(%v) \n\thave %b\n\twant %b", c.ex, v, c.v)
			}
		}
	})
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
			in, err := parseLines(c.in)
			if err != nil {
				t.Fatalf("Parse failure")
			}

			have := PartOne(in)
			if have != c.want {
				t.Errorf("PartOne(%v) == %v, want %v", c.in, have, c.want)
			}
		}
	})
}
