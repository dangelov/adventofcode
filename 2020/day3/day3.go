package day3

import (
	"strings"
)

// SolvePart1 takes input from a string
// and then solves the Day 3, Part 1 challenge
func SolvePart1(input string, shiftX, shiftY int64) int64 {
	lines := strings.Split(input, "\n")
	trees := int64(0)
	wrapAt := int64(len(lines[0]))
	tree := byte('#')

	x := int64(0)

	for y := int64(0); y < int64(len(lines)); y = y + shiftY {
		// Make sure we're not reading empty lines
		if len(lines[y]) == 0 {
			continue
		}

		if lines[y][x] == tree {
			trees++
		}

		x = (x + shiftX) % wrapAt
	}

	return trees
}

// SolvePart2 takes input from a string
// and then solves the Day 2, Part 2 challenge
func SolvePart2(input string) int64 {
	result := int64(1)

	tests := [][]int64{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for i := 0; i < len(tests); i++ {
		result *= SolvePart1(input, tests[i][0], tests[i][1])
	}

	return result
}
