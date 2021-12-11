package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9(t *testing.T) {
	values, err := ReadLines("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 524, day9Part1(values))
	assert.Equal(t, 0, day9Part2(values))
}

var resultDay9 int

func BenchmarkDay9Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day9.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day9Part1(values)
	}
	resultDay9 = r
}

func BenchmarkDay9Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day9.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day9Part2(values)
	}
	resultDay9 = r
}
