// Day 17: Trick Shot
package day17

import (
	"errors"
	"math"
	"regexp"
	"strconv"

	"github.com/partylich/advent2021/runner"
)

type Bound struct {
	Min, Max int
}

func (b Bound) InBounds(i int) bool {
	return i <= b.Max && i >= b.Min
}

type _ParseResult = []Bound

func parseLines(in string) (_ParseResult, error) {
	re := regexp.MustCompile(`x=(-?\d+)\.{2}(-?\d+), y=(-?\d+)\.{2}(-?\d+)`)
	m := re.FindSubmatch([]byte(in))

	bounds := make([]int, 4)
	for idx := 1; idx < len(m); idx++ {
		b, err := strconv.Atoi(string(m[idx]))
		if err != nil {
			return nil, errors.New("parse failure")
		}

		bounds[idx-1] = b
	}

	xB := Bound{bounds[0], bounds[1]}
	yB := Bound{bounds[2], bounds[3]}

	return []Bound{xB, yB}, nil
}

type Sol struct {
	// v is the initial velocity.
	v []int
	// top is the highest point [x, y].
	top []int
}

func step(b _ParseResult, vx, vy int) (Sol, bool) {
	var pX, pY int
	result := Sol{[]int{vx, vy}, []int{0, 0}}

	for pX <= b[0].Max && pY >= b[1].Min {
		pX += vx
		pY += vy

		if pY > result.top[1] {
			result.top[0] = pX
			result.top[1] = pY
		}
		if b[0].InBounds(pX) && b[1].InBounds(pY) {
			return result, true
		}

		vx = runner.Max(0, vx-1)
		vy -= 1
	}

	return result, false
}

// PartOne returns the highest y position of any trajectory that has a discrete
// point in the target area.
func PartOne(in _ParseResult) int {
	max := 0
	minVy := 0

	// third equation of motion v² = v₀² + 2a∆s
	// 0 = v₀² + 2(-1)(start of target range - 0)
	// 0 = v₀² + 2(-1)(start of target range - 0)
	// √(2 * start) = v₀
	minVx := int(math.Floor(math.Sqrt(float64(2 * in[0].Min))))

loop:
	for vx, hit := minVx, false; !hit; vx++ {
		for minVy = 0; !hit && minVy < 100; minVy++ {
			s, ok := step(in, vx, minVy)
			hit = ok
			if hit {
				max = runner.Max(max, s.top[1])
				minVx = vx
				break loop
			}
		}
	}

	for vy, missY := minVy, 0; missY < 50; vy++ {
		hits := 0

		for vx, missX := minVx, 0; missX < 50; vx++ {
			s, ok := step(in, vx, vy)
			if ok {
				hits += 1
				max = runner.Max(max, s.top[1])

				missX, missY = 0, 0
			} else {
				missX += 1
			}
		}

		if hits == 0 {
			missY += 1
		}
	}

	return max
}

// PartTwo returns the number of  distinct initial velocity values cause the
// probe to be within the target area after any step.
func PartTwo(in _ParseResult) int {
	// third equation of motion v² = v₀² + 2a∆s
	// 0 = v₀² + 2(-1)(start of target range - 0)
	// 0 = v₀² + 2(-1)(start of target range - 0)
	// √(2 * start) = v₀
	minVx := int(math.Floor(math.Sqrt(float64(2 * in[0].Min))))
	minVy := 0

loop:
	for vx, hit := minVx, false; !hit; vx++ {
		for minVy = 0; !hit && minVy < 100; minVy++ {
			_, hit = step(in, vx, minVy)
			if hit {
				minVx = vx
				break loop
			}
		}
	}

	result := 0
	for vy, missY := in[1].Min, 0; missY < 30; vy++ {
		hits := 0

		for vx, missX := minVx, 0; missX < 130; vx++ {
			_, ok := step(in, vx, vy)
			if ok {
				hits += 1
				missX, missY = 0, 0
			} else {
				missX += 1
			}
		}

		result += hits
		if hits == 0 {
			missY += 1
		}
	}

	return result
}

func Solution() runner.Solution {
	return runner.Solution{
		Parse: func(i string) (interface{}, error) { return parseLines(i) },
		Fn: [2]func(i interface{}) interface{}{
			func(i interface{}) interface{} { return PartOne(i.(_ParseResult)) },
			func(i interface{}) interface{} { return PartTwo(i.(_ParseResult)) },
		},
	}
}
