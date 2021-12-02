package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	horizontal int
	depth      int
	aim        int
}

// Day2 Day 2
func (aoc AOC) Day2() {
	values, err := ReadLines("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day2Part1(values))
	fmt.Println(day2Part2(values))
}

func day2Part1(values []string) int {
	position := Position{}
	for _, line := range values {
		instr := strings.Split(line, " ")
		value, _ := strconv.Atoi(instr[1])
		switch instr[0] {
		case "up":
			position.depth -= value
		case "down":
			position.depth += value
		case "forward":
			position.horizontal += value
		}
	}
	return position.horizontal * position.depth
}

func day2Part2(values []string) int {
	position := Position{}
	for _, line := range values {
		instr := strings.Split(line, " ")
		value, _ := strconv.Atoi(instr[1])
		switch instr[0] {
		case "up":
			position.aim -= value
		case "down":
			position.aim += value
		case "forward":
			position.horizontal += value
			position.depth += position.aim * value
		}
	}
	return position.horizontal * position.depth
}
