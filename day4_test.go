package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	values, err := ReadLines("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 60368, day4Part1(values))
	assert.Equal(t, 17435, day4Part2(values))
}

var resultDay4 int

func BenchmarkDay4Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day4.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day4Part1(values)
	}
	resultDay4 = r
}

func BenchmarkDay4Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day4.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day4Part2(values)
	}
	resultDay4 = r
}
