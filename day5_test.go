package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5(t *testing.T) {
	values, err := ReadLines("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 7674, day5Part1(values))
	assert.Equal(t, 20898, day5Part2(values))
}

var resultDay5 int

func BenchmarkDay5Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day5.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day5Part1(values)
	}
	resultDay5 = r
}

func BenchmarkDay5Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day5.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day5Part2(values)
	}
	resultDay5 = r
}
