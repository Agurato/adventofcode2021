package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	values, err := ReadLines("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 4139586, day3Part1(values))
	assert.Equal(t, 1800151, day3Part2(values))
}

var resultDay3 int

func BenchmarkDay3Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day3.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day3Part1(values)
	}
	resultDay3 = r
}

func BenchmarkDay3Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day3.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day3Part2(values)
	}
	resultDay3 = r
}
