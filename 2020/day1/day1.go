package day1

import (
	"strconv"
	"strings"
)

// SolvePart1 takes input from a string
// and then solves the Day 1, Part 1 challenge
func SolvePart1(input string) int64 {
	lines := strings.Split(input, "\n")
	numbers := make([]int64, 0)

	for i := 0; i < len(lines); i++ {
		num, err := strconv.ParseInt(lines[i], 10, 64)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	for i := 0; i < len(numbers); i++ {
		for o := 0; o < len(numbers); o++ {
			if i != o && numbers[i]+numbers[o] == 2020 {
				return numbers[i] * numbers[o]
			}
		}
	}

	return 0
}

// SolvePart2 takes input from a string
// and then solves the Day 1, Part 2 challenge
func SolvePart2(input string) int64 {
	lines := strings.Split(input, "\n")
	numbers := make([]int64, 0)

	for i := 0; i < len(lines); i++ {
		num, err := strconv.ParseInt(lines[i], 10, 64)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	for i := 0; i < len(numbers); i++ {
		for o := 0; o < len(numbers); o++ {
			for p := 0; p < len(numbers); p++ {
				if i != o && i != p && o != p && numbers[i]+numbers[o]+numbers[p] == 2020 {
					return numbers[i] * numbers[o] * numbers[p]
				}
			}
		}
	}

	return 0
}
