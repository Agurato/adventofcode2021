package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day6 Day 6
func (aoc AOC) Day6() {
	values, err := ReadLines("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day6Part1(values))
	fmt.Println(day6Part2(values))
}

func day6Part1(values []string) int {
	var fishesPerAge [9]int
	for _, val := range strings.Split(values[0], ",") {
		fishesPerAge[First(strconv.Atoi(val))]++
	}
	for i := 0; i < 80; i++ {
		var tmpFishesPerAge [9]int
		// Aging
		for j := 0; j < 8; j++ {
			tmpFishesPerAge[j] = fishesPerAge[j+1]
		}
		// Reproduction
		tmpFishesPerAge[8] = fishesPerAge[0]
		tmpFishesPerAge[6] += fishesPerAge[0]
		fishesPerAge = tmpFishesPerAge
	}
	count := 0
	for _, nb := range fishesPerAge {
		count += nb
	}
	return count
}

func day6Part2(values []string) int {
	var fishesPerAge [9]int
	for _, val := range strings.Split(values[0], ",") {
		fishesPerAge[First(strconv.Atoi(val))]++
	}
	for i := 0; i < 256; i++ {
		var tmpFishesPerAge [9]int
		// Aging
		for j := 0; j < 8; j++ {
			tmpFishesPerAge[j] = fishesPerAge[j+1]
		}
		// Reproduction
		tmpFishesPerAge[8] = fishesPerAge[0]
		tmpFishesPerAge[6] += fishesPerAge[0]
		fishesPerAge = tmpFishesPerAge
	}
	count := 0
	for _, nb := range fishesPerAge {
		count += nb
	}
	return count
}
