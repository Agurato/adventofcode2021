package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	values, err := ReadLines("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 392421, day10Part1(values))
	assert.Equal(t, 2769449099, day10Part2(values))
}

var resultDay10 int

func BenchmarkDay10Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day10.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day10Part1(values)
	}
	resultDay10 = r
}

func BenchmarkDay10Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day10.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day10Part2(values)
	}
	resultDay10 = r
}
