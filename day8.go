package main

import (
	"fmt"
	"strings"
)

func getNumberFromString(val string) int {
	switch len(val) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 5: // 2 or 3 or 5
		if strings.ContainsRune(val, 'e') {
			return 5
		}
		if strings.ContainsRune(val, 'g') {
			return 2
		}
		return 3
	case 6: // 0 or 6 or 9
		if !strings.ContainsRune(val, 'f') {
			return 0
		}
		if strings.ContainsRune(val, 'g') {
			return 6
		}
		return 9
	case 7:
		return 8
	}
	return 0
}

// Day8 Day 8
func (aoc AOC) Day8() {
	values, err := ReadLines("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day8Part1(values))
	fmt.Println(day8Part2(values))
}

func day8Part1(values []string) int {
	count := 0
	for _, line := range values {
		inout := strings.Split(line, " | ")
		for _, out := range strings.Split(inout[1], " ") {
			outLength := len(out)
			if outLength == 2 || outLength == 3 || outLength == 4 || outLength == 7 {
				count++
			}
		}
	}
	return count
}

func day8Part2(values []string) int {
	count := 0
	for _, line := range values {
		inout := strings.Split(line, " | ")
		var nums [4]int
		fmt.Print(inout[1] + " ")
		for i, out := range strings.Split(inout[1], " ") {
			nums[i] = getNumberFromString(out)
		}
		fmt.Println(nums)
		count += nums[0]*1000 + nums[1]*100 + nums[2]*10 + nums[3]
	}
	// return count
	return 0
}
