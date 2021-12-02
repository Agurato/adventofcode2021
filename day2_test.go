package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	values, err := ReadLines("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 2117664, day2Part1(values))
	assert.Equal(t, 2073416724, day2Part2(values))
}

var resultDay2 int

func BenchmarkDay2Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day2.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day2Part1(values)
	}
	resultDay2 = r
}

func BenchmarkDay2Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day2.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day2Part2(values)
	}
	resultDay2 = r
}
