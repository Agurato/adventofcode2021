package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Day7 Day 7
func (aoc AOC) Day7() {
	values, err := ReadLinesSep("inputs/test.txt", ",")
	if err != nil {
		panic(err)
	}
	fmt.Println(day7Part1(values))
	fmt.Println(day7Part2(values))
}

func day7Part1(values []string) int {
	positions := make([]int, len(values))
	for i, val := range values {
		positions[i] = First(strconv.Atoi(val))
	}
	sort.Ints(positions)
	sum := 0
	median := positions[int(len(positions)/2)]
	for _, pos := range positions {
		sum += Abs(pos - median)
	}
	return sum
}

func day7Part2(values []string) int {
	positions := make([]int, len(values))
	for i, val := range values {
		positions[i] = First(strconv.Atoi(val))
	}
	sum := 0
	for _, pos := range positions {
		sum += pos
	}
	average := sum / len(positions)
	sumFloor := 0
	for _, pos := range positions {
		distance := Abs(average - pos)
		sumFloor += distance * (distance + 1) / 2
	}
	sumCeiling := 0
	for _, pos := range positions {
		distance := Abs(average + 1 - pos)
		sumCeiling += distance * (distance + 1) / 2
	}
	return Min(sumFloor, sumCeiling)
}
