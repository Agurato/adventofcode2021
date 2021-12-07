package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7(t *testing.T) {
	values, err := ReadLinesSep("inputs/day7.txt", ",")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 336120, day7Part1(values))
	assert.Equal(t, 96864235, day7Part2(values))
}

var resultDay7 int

func BenchmarkDay7Part1(b *testing.B) {
	values, _ := ReadLinesSep("inputs/day7.txt", ",")
	var r int
	for i := 0; i < b.N; i++ {
		r = day7Part1(values)
	}
	resultDay7 = r
}

func BenchmarkDay7Part2(b *testing.B) {
	values, _ := ReadLinesSep("inputs/day7.txt", ",")
	var r int
	for i := 0; i < b.N; i++ {
		r = day7Part2(values)
	}
	resultDay7 = r
}
