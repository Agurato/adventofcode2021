package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	values, err := ReadLines("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 376194, day6Part1(values))
	assert.Equal(t, 1693022481538, day6Part2(values))
}

var resultDay6 int

func BenchmarkDay6Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day6.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day6Part1(values)
	}
	resultDay6 = r
}

func BenchmarkDay6Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day6.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day6Part2(values)
	}
	resultDay6 = r
}
