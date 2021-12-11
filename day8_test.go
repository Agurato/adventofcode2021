package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8(t *testing.T) {
	values, err := ReadLines("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 278, day8Part1(values))
	assert.Equal(t, 0, day8Part2(values))
}

var resultDay8 int

func BenchmarkDay8Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day8.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day8Part1(values)
	}
	resultDay8 = r
}

func BenchmarkDay8Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day8.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day8Part2(values)
	}
	resultDay8 = r
}
