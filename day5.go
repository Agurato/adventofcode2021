package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

func PrintGrid(grid [10][10]int) {
	for y := range grid {
		for x := range grid[y] {
			if grid[x][y] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(grid[x][y])
			}
		}
		fmt.Println()
	}
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func getCoordFromString(line string) (start, end Coord) {
	startEndStr := strings.Split(line, " -> ")
	startStr := strings.Split(startEndStr[0], ",")
	endStr := strings.Split(startEndStr[1], ",")
	start = Coord{
		x: First(strconv.Atoi(startStr[0])),
		y: First(strconv.Atoi(startStr[1])),
	}
	end = Coord{
		x: First(strconv.Atoi(endStr[0])),
		y: First(strconv.Atoi(endStr[1])),
	}

	return
}

// Day5 Day 5
func (aoc AOC) Day5() {
	values, err := ReadLines("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day5Part1(values))
	fmt.Println(day5Part2(values))
}

func day5Part1(values []string) int {
	var grid [1000][1000]int
	for _, line := range values {
		start, end := getCoordFromString(line)
		if start.x == end.x {
			startY, endY := start.y, end.y
			if start.y > end.y {
				startY, endY = end.y, start.y
			}
			for y := startY; y <= endY; y++ {
				grid[start.x][y]++
			}
		} else if start.y == end.y {
			startX, endX := start.x, end.x
			if start.x > end.x {
				startX, endX = end.x, start.x
			}
			for x := startX; x <= endX; x++ {
				grid[x][start.y]++
			}
		}
	}

	var count int
	for _, col := range grid {
		for y := range col {
			if col[y] > 1 {
				count++
			}
		}
	}
	return count
}

func day5Part2(values []string) int {
	var grid [1000][1000]int
	for _, line := range values {
		start, end := getCoordFromString(line)
		if start.x == end.x {
			startY, endY := start.y, end.y
			if start.y > end.y {
				startY, endY = end.y, start.y
			}
			for y := startY; y <= endY; y++ {
				grid[start.x][y]++
			}
		} else if start.y == end.y {
			startX, endX := start.x, end.x
			if start.x > end.x {
				startX, endX = end.x, start.x
			}
			for x := startX; x <= endX; x++ {
				grid[x][start.y]++
			}
		} else {
			var x, y int
			for i := 0; i <= abs(end.x-start.x); i++ {
				grid[x+start.x][y+start.y]++
				if end.x > start.x {
					x++
				} else {
					x--
				}
				if end.y > start.y {
					y++
				} else {
					y--
				}
			}
		}
	}

	var count int
	for _, col := range grid {
		for y := range col {
			if col[y] > 1 {
				count++
			}
		}
	}
	return count
}
