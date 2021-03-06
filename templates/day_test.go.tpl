package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay{{.day}}(t *testing.T) {
	values, err := ReadLines("inputs/day{{.day}}.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 0, day{{.day}}Part1(values))
	assert.Equal(t, 0, day{{.day}}Part2(values))
}

var resultDay{{.day}} int

func BenchmarkDay{{.day}}Part1(b *testing.B) {
	values, _ := ReadLines("inputs/day{{.day}}.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day{{.day}}Part1(values)
	}
	resultDay{{.day}} = r
}

func BenchmarkDay{{.day}}Part2(b *testing.B) {
	values, _ := ReadLines("inputs/day{{.day}}.txt")
	var r int
	for i := 0; i < b.N; i++ {
		r = day{{.day}}Part2(values)
	}
	resultDay{{.day}} = r
}
