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
	t.Run("parse literal value packet", func(t *testing.T) {
		cases := []struct {
			ex  _ParseResult
			ver Pversion
			val int
		}{
			{_ParseResult{`D2FE28`, 6, 0}, 6, 2021},
			{_ParseResult{`38006F45291200`, 14, 22}, 6, 10},
			{_ParseResult{`38006F45291200`, 14, 33}, 2, 20},
		}

		for _, c := range cases {
			_, v := header(&c.ex)
			have := parseLiteral(&c.ex, v)

			if have.version() != c.ver {
				t.Errorf("parse literal(%v) version\n\thave %#v\n\twant %v", c.ex, have, c.ver)
			}
			if have.value() != c.val {
				t.Errorf("parse literal(%v) value\n\thave %#v\n\twant %v", c.ex, have, c.val)
			}
		}
	})
	t.Run("parse operator packet", func(t *testing.T) {
		cases := []struct {
			ex   _ParseResult
			ver  Pversion
			subs int
		}{
			{_ParseResult{`38006F45291200`, 14, 0}, 1, 2},
			{_ParseResult{`EE00D40C823060`, 14, 0}, 7, 3},
		}

		for _, c := range cases {
			typ, v := header(&c.ex)
			have := parseOperator(&c.ex, v, typ)

			if have.version() != c.ver {
				t.Errorf("parse operator(%v) version\n\thave %#v\n\twant %v", c.ex, have, c.ver)
			}
			if len(have.subpackets()) != c.subs {
				t.Errorf("parse operator(%v) value\n\thave %#v\n\twant %v", c.ex, have, c.subs)
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

func TestPartTwo(t *testing.T) {
	t.Run("executes task", func(t *testing.T) {
		cases := []struct {
			in   string
			want int
		}{
			{"C200B40A82", 3},
			{"04005AC33890", 54},
			{"880086C3E88112", 7},
			{"CE00C43D881120", 9},
			{"D8005AC2A8F0", 1},
			{"F600BC2D8F", 0},
			{"9C005AC2F8F0", 0},
			{"9C0141080250320F1802104A08", 1},
		}

		for _, c := range cases {
			in, err := parseLines(c.in)
			if err != nil {
				t.Fatalf("Parse failure")
			}

			have := PartTwo(in)
			if have != c.want {
				t.Errorf("PartTwo(%v) == %v, want %v", c.in, have, c.want)
			}
		}
	})
}
