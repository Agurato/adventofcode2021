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
		for j := 0; j < 8; j++ {
			tmpFishesPerAge[j] = fishesPerAge[j+1]
		}
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
		tmpFishesPerAge[8] = fishesPerAge[0]
		tmpFishesPerAge[7] = fishesPerAge[8]
		tmpFishesPerAge[6] = fishesPerAge[7] + fishesPerAge[0]
		tmpFishesPerAge[5] = fishesPerAge[6]
		tmpFishesPerAge[4] = fishesPerAge[5]
		tmpFishesPerAge[3] = fishesPerAge[4]
		tmpFishesPerAge[2] = fishesPerAge[3]
		tmpFishesPerAge[1] = fishesPerAge[2]
		tmpFishesPerAge[0] = fishesPerAge[1]
		fishesPerAge = tmpFishesPerAge
	}
	count := 0
	for _, nb := range fishesPerAge {
		count += nb
	}
	return count
}
