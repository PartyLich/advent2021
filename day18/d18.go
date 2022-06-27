// Day 18: Snailfish
package day18

import (
	"strconv"

	"github.com/partylich/advent2021/parse"
	"github.com/partylich/advent2021/runner"
)

type Node struct {
	val   int
	depth int
	left  *Node
	right *Node
}

func (n *Node) Copy() *Node {
	result := &Node{}
	*result = *n
	node := result

	for p := n.right; p != nil; p = p.right {
		temp := &Node{p.val, p.depth, node, nil}
		node.right = temp
		node = temp
	}

	return result
}

type _ParseResult []*Node

func parseNum(in string) (*Node, error) {
	var result, tail *Node

	depth := 0
	addNode := func(val int) {
		n := &Node{val, depth, tail, nil}
		if tail == nil {
			result = n
			tail = n
		} else {
			tail.right = n
			tail = n
		}
	}

	lastDigit := false
	s := ""
	for _, r := range in {
		switch r {
		case '[':
			depth += 1
			lastDigit = false
		case ']':
			if lastDigit {
				val, e := strconv.Atoi(s)
				if e != nil {
					return nil, e
				}
				addNode(val)
			}
			depth -= 1
			lastDigit = false
		case ',':
			if lastDigit {
				val, e := strconv.Atoi(s)
				if e != nil {
					return nil, e
				}
				addNode(val)
			}
			lastDigit = false
		default:
			if !lastDigit {
				s = ""
			}

			lastDigit = true
			s += string(r)
		}
	}

	return result, nil
}

func parseLines(in string) (_ParseResult, error) {
	lines := parse.Lines(in)
	result := make(_ParseResult, len(lines))
	for i, l := range lines {
		r, err := parseNum(l)
		if err != nil {
			return nil, err
		}

		result[i] = r
	}

	return result, nil
}

func canExplode(in *Node) (*Node, bool) {
	maxD, start := maxDepth(in)
	if maxD < 5 {
		return nil, false
	}

	return start, true
}

func explode(in *Node, start *Node) *Node {
	node := &Node{0, start.depth - 1, start.left, start.right.right}

	if start.left != nil {
		start.left.val += start.val
		start.left.right = node
	}

	if start.right.right != nil {
		start.right.right.val += start.right.val
		start.right.right.left = node
	}

	*start = *node

	return in
}

func canSplit(in *Node) (*Node, bool) {
	for n := in; n != nil; n = n.right {
		if n.val >= 10 {
			return n, true
		}
	}

	return nil, false
}

func split(in *Node, start *Node) *Node {
	l := start.val / 2
	r := l + (start.val % 2)

	nodeR, nodeL := &Node{}, &Node{}
	*nodeL = Node{l, start.depth + 1, start.left, nodeR}
	*nodeR = Node{r, start.depth + 1, nodeL, start.right}

	if start.left != nil {
		start.left.right = nodeL
	} else {
		in = nodeL
	}
	if start.right != nil {
		start.right.left = nodeR
	}

	return in
}

func reduce(in *Node) *Node {
	result := in
	for {
		tmp := result.Copy()
		idx, explodable := canExplode(tmp)
		start, splittable := canSplit(tmp)

		switch {
		case explodable:
			result = explode(tmp, idx)
		case splittable:
			result = split(tmp, start)
		default:
			return result
		}
	}
}

func add(a, b *Node) *Node {
	var tail *Node
	list := a.Copy()

	for n := list; n != nil; n = n.right {
		n.depth += 1
		tail = n
	}

	tail.right = b.Copy()
	tail.right.left = tail

	for n := tail.right; n != nil; n = n.right {
		n.depth += 1
	}

	return reduce(list)
}

func maxDepth(in *Node) (int, *Node) {
	var start *Node
	maxD := -1

	for n := in; n != nil; n = n.right {
		if n.depth > maxD {
			maxD = n.depth
			start = n
		}
	}

	return maxD, start
}

func magnitude(in *Node) int {
	for maxD, start := maxDepth(in); maxD != 0; maxD, start = maxDepth(in) {
		for n := start; n != nil; n = n.right {
			if n.depth == maxD {
				val := (n.val * 3) + (n.right.val * 2)
				node := &Node{val, maxD - 1, n.left, n.right.right}

				if n.left != nil {
					n.left.right = node
				}
				if n.right.right != nil {
					n.right.right.left = node
				}

				*n = *node
			}
		}
	}

	return in.val
}

// PartOne returns the magnitude of the final sum.
func PartOne(in _ParseResult) int {
	if len(in) == 0 {
		return 0
	}

	result := in[0]
	for i := 1; i < len(in); i++ {
		result = add(result, in[i])
	}

	return magnitude(result)
}

// PartTwo returns the largest magnitude of any sum of two different snailfish
// numbers from the homework assignment.
func PartTwo(in _ParseResult) int {
	if len(in) == 0 {
		return 0
	}

	result := 0
	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			a := add(in[i], in[j])
			b := add(in[j], in[i])
			result = runner.Max(result, magnitude(a))
			result = runner.Max(result, magnitude(b))
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
