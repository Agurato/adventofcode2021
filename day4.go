package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type BingoNumber struct {
	number int
	drawn  bool
}

type Bingo struct {
	numbers [25]BingoNumber
}

func (b Bingo) Print() {
	green := color.New(color.FgGreen, color.Bold)
	for i := 0; i < 25; i++ {
		if b.numbers[i].drawn {
			green.Printf("%3d", b.numbers[i].number)
		} else {
			fmt.Printf("%3d", b.numbers[i].number)
		}
		if (i+1)%5 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func (b Bingo) SumNotDrawn() (sum int) {
	for i := 0; i < 25; i++ {
		if !b.numbers[i].drawn {
			sum += b.numbers[i].number
		}
	}
	return
}

func (b *Bingo) Draw(val int) {
	for i := 0; i < 25; i++ {
		if b.numbers[i].number == val {
			b.numbers[i].drawn = true
		}
	}
}

func (b Bingo) IsComplete() bool {
	for i := 0; i < 5; i++ {
		rowComplete := true
		columnComplete := true
		// Check rows
		for j := 0; j < 5; j++ {
			if !b.numbers[i*5+j].drawn {
				rowComplete = false
				break
			}
		}
		if rowComplete {
			return true
		}
		// Check columns
		for j := 0; j < 5; j++ {
			if !b.numbers[i+j*5].drawn {
				columnComplete = false
				break
			}
		}
		if columnComplete {
			return true
		}
	}
	return false
}

// Day4 Day 4
func (aoc AOC) Day4() {
	values, err := ReadLinesSep("inputs/day4.txt", "\n\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(day4Part1(values))
	fmt.Println(day4Part2(values))
}

func day4Part1(values []string) int {
	var bingos []Bingo

	// Make bingo grids
	for i := 1; i < len(values); i++ {
		bingos = append(bingos, Bingo{})
		for j, v := range Chunks(values[i], 3) {
			bingos[i-1].numbers[j].number, _ = strconv.Atoi(strings.Trim(v, " \n"))
		}
	}

	var draw, winnerBoard int
drawing:
	for i, drawStr := range strings.Split(values[0], ",") {
		draw, _ = strconv.Atoi(drawStr)
		for j := range bingos {
			bingos[j].Draw(draw)
			if i > 5 {
				if bingos[j].IsComplete() {
					winnerBoard = j
					break drawing
				}
			}
		}
	}

	return draw * bingos[winnerBoard].SumNotDrawn()
}

func day4Part2(values []string) int {
	var bingos []Bingo

	// Make bingo grids
	for i := 1; i < len(values); i++ {
		bingos = append(bingos, Bingo{})
		for j, v := range Chunks(values[i], 3) {
			bingos[i-1].numbers[j].number, _ = strconv.Atoi(strings.Trim(v, " \n"))
		}
	}

	var draw int
drawing:
	for i, drawStr := range strings.Split(values[0], ",") {
		draw, _ = strconv.Atoi(drawStr)
		for j := 0; j < len(bingos); j++ {
			bingos[j].Draw(draw)
			if i > 5 {
				if bingos[j].IsComplete() {
					if len(bingos) == 1 {
						break drawing
					}
					bingos[j] = bingos[len(bingos)-1]
					bingos = bingos[:len(bingos)-1]
					j--
				}
			}
		}
	}

	return draw * bingos[0].SumNotDrawn()
}
