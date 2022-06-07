// Day 16: Packet Decoder
package day16

import (
	"math"
	"strconv"
	"strings"

	"github.com/partylich/advent2021/runner"
)

type _ParseResult struct {
	msg    string
	msgLen int
	used   int
}

func parseLines(in string) (_ParseResult, error) {
	s := strings.TrimSpace(in)
	return _ParseResult{s, len(s), 0}, nil
}

type Ptype int
type Pversion int

type _Packet interface {
	version() Pversion
	value() int
	subpackets() []_Packet
}

// Packets with type ID 4 represent a literal value. Literal value packets encode a single binary number.
type _PacketLiteral struct {
	ver Pversion
	val int
}

func (p _PacketLiteral) version() Pversion {
	return p.ver
}

func (p _PacketLiteral) value() int {
	return p.val
}

func (p _PacketLiteral) subpackets() []_Packet {
	return []_Packet{}
}

// _PacketOperator performs some calculation on one or more sub-packets contained
// within.
type _PacketOperator struct {
	ver     Pversion
	val     int
	packets []_Packet
}

func (p _PacketOperator) version() Pversion {
	return p.ver
}

func (p _PacketOperator) value() int {
	return p.val
}

func (p _PacketOperator) subpackets() []_Packet {
	return p.packets
}

// Every packet begins with a standard header: the first three bits encode the
// packet version, and the next three bits encode the packet type ID.
func header(in *_ParseResult) (Ptype, Pversion) {
	i := in.used / 4
	skip := in.used % 4
	nibbles := int(math.Ceil(float64(6+skip) / 4))
	s := in.msg[i : i+nibbles]

	bits, _ := strconv.ParseInt(s, 16, 64)
	bits >>= (nibbles * 4) - skip - 6

	pType := Ptype((bits) & 0b111)
	pVer := Pversion(bits>>3) & 0b111

	in.used += 6

	return pType, pVer
}

func parseLiteral(in *_ParseResult, ver Pversion) _PacketLiteral {
	idx := in.used / 4
	skip := in.used % 4
	end := runner.Min(len(in.msg), idx+4)
	s := in.msg[idx:end]

	bitLen := 4*len(s) - skip
	msg, _ := strconv.ParseUint(s, 16, 64)
	if skip != 0 {
		msg &^= (0xF << bitLen) // clear used bits
	}

	var value uint64

	bits := skip
	for read, i := uint64(1), 0; read == 1; i++ {
		shift := (bitLen - 5 - (i * 5))
		group := msg >> shift
		group &= 0b11111
		read = group >> 4

		v := group & 0xF
		value = v | (value << 4)

		in.used += 5
		bits += 5

		if bits >= 11 {
			idx += bits / 4
			skip = bits % 4
			end := runner.Min(len(in.msg), idx+4)
			s = in.msg[idx:end]

			bitLen = 4*len(s) - skip
			msg, _ = strconv.ParseUint(s, 16, 64)
			if skip != 0 {
				msg &^= (0xF << bitLen) // clear used bits
			}

			i = -1
			bits = skip
		}
	}

	return _PacketLiteral{ver, int(value)}
}

func parseOperator(in *_ParseResult, ver Pversion) _PacketOperator {
	idx := in.used / 4
	skip := in.used % 4
	end := runner.Min(len(in.msg), idx+6)
	s := in.msg[idx:end]

	bitLen := 4*len(s) - skip

	msg, _ := strconv.ParseUint(s, 16, 64)
	if skip != 0 {
		msg &^= (0xF << bitLen) // clear used bits
	}

	lenType := msg >> (bitLen - 1)
	in.used += 1

	var subpacks []_Packet

	switch lenType {
	case 0:
		subLen := msg >> (bitLen - 16)
		subLen &^= (1 << 16)
		in.used += 15

		start := in.used
		for done := 0; done < int(subLen); done = in.used - start {
			p := parsePacket(in)
			subpacks = append(subpacks, p)
		}

	default:
		subCount := msg >> (bitLen - 12)
		subCount &^= (1 << 11)
		in.used += 11

		for i := 0; i < int(subCount); i++ {
			p := parsePacket(in)
			subpacks = append(subpacks, p)
		}
	}

	return _PacketOperator{ver, 0, subpacks}
}

func parsePacket(in *_ParseResult) _Packet {
	pType, pVer := header(in)

	switch pType {
	case 4:
		return parseLiteral(in, pVer)
	default:
		return parseOperator(in, pVer)
	}
}

func sumVersions(p _Packet) int {
	sum := int(p.version())
	for _, packet := range p.subpackets() {
		sum += sumVersions(packet)
	}

	return sum
}

// PartOne returns the sum of the version numbers in all parsed packets
func PartOne(in _ParseResult) int {
	p := parsePacket(&in)

	return sumVersions(p)
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.(_ParseResult)) },
			runner.Unimpl,
		},
	}
}
