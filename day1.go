package main

import "fmt"

// Day1 Day 1
func (aoc AOC) Day1() {
	values, err := ReadLinesToInt("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day1Part1(values))
	fmt.Println(day1Part2(values))
}

func day1Part1(values []int) int {
	var count int
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			count++
		}
	}
	return count
}

func day1Part2(values []int) int {
	var sliding []int = make([]int, len(values)-2)
	// Calculate sliding windows
	for i := 0; i < len(values)-2; i++ {
		sliding[i] = values[i] + values[i+1] + values[i+2]
	}
	return day1Part1(sliding)
}
