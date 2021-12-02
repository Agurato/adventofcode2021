package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	values, err := ReadLinesToInt("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1477, day1Part1(values))
	assert.Equal(t, 1523, day1Part2(values))
}

var resultDay1 int

func BenchmarkDay1Part1(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day1.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day1Part1(values)
	}
	resultDay1 = r
}

func BenchmarkDay1Part2(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day1.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day1Part2(values)
	}
	resultDay1 = r
}
