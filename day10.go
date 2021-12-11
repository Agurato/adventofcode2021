package main

import (
	"fmt"
	"sort"
)

var (
	close = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	pointsPart1 = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	pointsPart2 = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
)

func PrintFifo(fifo []rune) {
	for _, char := range fifo {
		fmt.Printf("%c", char)
	}
	fmt.Println()
}

// Day10 Day 10
func (aoc AOC) Day10() {
	values, err := ReadLines("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day10Part1(values))
	fmt.Println(day10Part2(values))
}

func day10Part1(values []string) int {
	var fifo []rune
	sum := 0

	for _, line := range values {
		for _, char := range line {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				fifo = append(fifo, char)
			} else if char == ')' || char == ']' || char == '}' || char == '>' {
				if close[fifo[len(fifo)-1]] == char {
					fifo = fifo[:len(fifo)-1]
				} else {
					sum += pointsPart1[char]
					break
				}
			}
		}
		fifo = nil
	}

	return sum
}

func day10Part2(values []string) int {
	var fifo []rune
	var scores []int

forLine:
	for _, line := range values {
		fifo = nil
		for _, char := range line {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				fifo = append(fifo, char)
			} else if char == ')' || char == ']' || char == '}' || char == '>' {
				if close[fifo[len(fifo)-1]] == char {
					fifo = fifo[:len(fifo)-1]
				} else {
					continue forLine
				}
			}
		}

		score := 0
		for i := len(fifo) - 1; i >= 0; i-- {
			score = score*5 + pointsPart2[fifo[i]]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}
