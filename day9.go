package main

import "fmt"

func makeGrid(values []string) [][]uint8 {
	width := len(values[0])
	grid := make([][]uint8, len(values))

	for i, line := range values {
		grid[i] = make([]uint8, width)
		for j, x := range line {
			grid[i][j] = uint8(x) - 48
		}
	}

	return grid
}

// Day9 Day 9
func (aoc AOC) Day9() {
	values, err := ReadLines("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day9Part1(values))
	fmt.Println(day9Part2(values))
}

func day9Part1(values []string) int {
	grid := makeGrid(values)
	width := len(grid[0])
	height := len(grid)

	count := 0
	for y := range grid {
		for x := range grid[y] {
			if x > 0 && grid[y][x-1] <= grid[y][x] {
				continue
			} else if x < width-1 && grid[y][x+1] <= grid[y][x] {
				continue
			} else if y > 0 && grid[y-1][x] <= grid[y][x] {
				continue
			} else if y < height-1 && grid[y+1][x] <= grid[y][x] {
				continue
			}
			count += int(grid[y][x] + 1)
		}
	}

	return count
}

func day9Part2(values []string) int {
	return 0
}
