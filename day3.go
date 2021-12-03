package main

import (
	"fmt"
	"strconv"
)

type Bits struct {
	zeroes int
	ones   int
}

// Day3 Day 3
func (aoc AOC) Day3() {
	values, err := ReadLines("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(day3Part1(values))
	fmt.Println(day3Part2(values))
}

func day3Part1(values []string) int {
	length := len(values[0])
	countBits := make([]Bits, length)
	gamma := make([]rune, length)
	epsilon := make([]rune, length)
	for _, line := range values {
		for i, char := range line {
			switch char {
			case '0':
				countBits[i].zeroes += 1
			case '1':
				countBits[i].ones += 1
			}
		}
	}

	for i, bits := range countBits {
		if bits.ones > bits.zeroes {
			gamma[i] = '1'
			epsilon[i] = '0'
		} else {
			gamma[i] = '0'
			epsilon[i] = '1'
		}
	}

	gammaVal, _ := strconv.ParseInt(string(gamma), 2, 64)
	epsilonVal, _ := strconv.ParseInt(string(epsilon), 2, 64)
	return int(gammaVal * epsilonVal)
}

func day3Part2(values []string) int {
	oxygenValues := make([]string, len(values))
	co2Values := make([]string, len(values))
	copy(oxygenValues, values)
	copy(co2Values, values)
	for criteriaIndex := 0; criteriaIndex < len(values[0]); criteriaIndex++ {
		var countBits Bits
		for _, value := range oxygenValues {
			switch value[criteriaIndex] {
			case '0':
				countBits.zeroes++
			case '1':
				countBits.ones++
			}
		}
		var tempValues []string
		for _, value := range oxygenValues {
			criteria := '0'
			if countBits.ones >= countBits.zeroes {
				criteria = '1'
			}
			if rune(value[criteriaIndex]) == criteria {
				tempValues = append(tempValues, value)
			}
		}
		oxygenValues = make([]string, len(tempValues))
		copy(oxygenValues, tempValues)
		if len(tempValues) == 1 {
			break
		}
	}
	for criteriaIndex := 0; criteriaIndex < len(values[0]); criteriaIndex++ {
		var countBits Bits
		for _, value := range co2Values {
			switch value[criteriaIndex] {
			case '0':
				countBits.zeroes++
			case '1':
				countBits.ones++
			}
		}
		var tempValues []string
		for _, value := range co2Values {
			criteria := '1'
			if countBits.zeroes <= countBits.ones {
				criteria = '0'
			}
			if rune(value[criteriaIndex]) == criteria {
				tempValues = append(tempValues, value)
			}
		}
		co2Values = make([]string, len(tempValues))
		copy(co2Values, tempValues)
		if len(tempValues) == 1 {
			break
		}
	}

	oxygenVal, _ := strconv.ParseInt(string(oxygenValues[0]), 2, 64)
	co2Val, _ := strconv.ParseInt(string(co2Values[0]), 2, 64)
	return int(oxygenVal * co2Val)
}
