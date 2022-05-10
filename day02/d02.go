package day02

import (
	"errors"

	"github.com/partylich/advent2021/runner"
)

type dir int

const (
	Forward dir = iota
	Down
	Up
)

type command struct {
	dir dir
	val int
}

func Parse(in string) ([]command, error) {
	return nil, errors.New("parse not implemented")
}

var Solution = runner.Solution{
	Parse: func(i string) (interface{}, error) { return Parse(i) },
	One:   func(i interface{}) interface{} { return runner.Unimpl(i) },
	Two:   func(i interface{}) interface{} { return runner.Unimpl(i) },
}
