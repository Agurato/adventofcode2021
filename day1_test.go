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

func BenchmarkDay1Part1(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day1.txt")
	for i := 0; i < b.N; i++ {
		day1Part1(values)
	}
}

func BenchmarkDay1Part2(b *testing.B) {
	values, _ := ReadLinesToInt("inputs/day1.txt")
	for i := 0; i < b.N; i++ {
		day1Part2(values)
	}
}
