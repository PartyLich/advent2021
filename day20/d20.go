// Day 20: Trench Map
package day20

import (
	"fmt"
	"strings"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type Image struct {
	minX, minY int
	maxX, maxY int
	pixels     map[string]bool
	// value of pixels outside the image focus
	void bool
}

func NewImage() Image {
	return Image{
		0, 0,
		0, 0,
		make(map[string]bool),
		false,
	}
}

type _ParseResult struct {
	algo  []bool
	image Image
}

func toKey(r, c int) string {
	return fmt.Sprintf("%v,%v", c, r)
}

func parseLines(in string) (_ParseResult, error) {
	var result _ParseResult
	lines := parse.Lines(in)

	algoStr := lines[0]
	algo := strings.Split(algoStr, "")
	result.algo = make([]bool, len(algo))
	for i, v := range algo {
		result.algo[i] = (v == "#")
	}

	result.image = NewImage()
	for r, l := range lines[2:] {
		for c, pixel := range strings.Split(l, "") {
			result.image.pixels[toKey(r, c)] = (pixel == "#")
			result.image.maxX = runner.Max(result.image.maxX, c)
			result.image.maxY = runner.Max(result.image.maxY, r)
		}
	}

	return result, nil
}

// getIndex returns enhancement algorithm index for the pixel at r,c.
func getIndex(img Image, r, c int) int {
	var result, idx int

	for i := r + 1; i >= r-1; i-- {
		for j := c + 1; j >= c-1; j, idx = j-1, idx+1 {
			var pixel bool
			if (i < img.minY || i > img.maxY) ||
				(j < img.minX || j > img.maxX) {
				pixel = img.void
			} else {
				pixel = img.pixels[toKey(i, j)]
			}

			if pixel {
				result |= (1 << idx)
			}
		}
	}

	return result
}

func enhance(algo []bool, image Image) Image {
	next := NewImage()
	visited := make(map[string]bool)

	for r := image.minY - 1; r <= image.maxY+1; r++ {
		for c := image.minX - 1; c <= image.maxX+1; c++ {
			k := toKey(r, c)
			visited[k] = true

			idx := getIndex(image, r, c)
			if algo[idx] {
				next.pixels[k] = true
				next.maxX = runner.Max(next.maxX, c)
				next.maxY = runner.Max(next.maxY, r)
				next.minX = runner.Min(next.minX, c)
				next.minY = runner.Min(next.minY, r)
			}
		}
	}

	// when algo[0] is #, anything outside the current lit pixels becomes
	// lit. Like, infinitely many pixels
	if image.void {
		next.void = algo[511]
	} else {
		next.void = algo[0]
	}

	return next
}

// PartOne returns how many pixels are lit after enhancing the image twice.
func PartOne(in _ParseResult) int {
	const enhancements = 2
	var (
		count int
		image Image
	)

	image = in.image
	for i := 0; i < enhancements; i++ {
		image = enhance(in.algo, image)
	}

	for _, v := range image.pixels {
		if v {
			count += 1
		}
	}

	return count
}

// PartTwo returns how many pixels are lit after enhancing the image 50 times.
func PartTwo(in _ParseResult) int {
	return 0
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
