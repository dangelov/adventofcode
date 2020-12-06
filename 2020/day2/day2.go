package day2

import (
	"regexp"
	"strconv"
	"strings"
)

// SolvePart1 takes input from a string
// and then solves the Day 2, Part 1 challenge
func SolvePart1(input string) int64 {
	lines := strings.Split(input, "\n")
	valid := int64(0)
	re := regexp.MustCompile(`([0-9]*)-([0-9]*)\s([a-z]):\s(.*)`)

	for i := 0; i < len(lines); i++ {
		matches := re.FindStringSubmatch(lines[i])

		if matches != nil {
			min, _ := strconv.Atoi(matches[1])
			max, _ := strconv.Atoi(matches[2])
			char := matches[3]
			pass := matches[4]

			c := strings.Count(pass, char)
			if c >= min && c <= max {
				valid++
			}
		}
	}

	return valid
}

// SolvePart2 takes input from a string
// and then solves the Day 2, Part 2 challenge
func SolvePart2(input string) int64 {
	lines := strings.Split(input, "\n")
	valid := int64(0)
	re := regexp.MustCompile(`([0-9]*)-([0-9]*)\s([a-z]):\s(.*)`)

	for i := 0; i < len(lines); i++ {
		matches := re.FindStringSubmatch(lines[i])

		if matches != nil {
			pos1, _ := strconv.Atoi(matches[1])
			pos2, _ := strconv.Atoi(matches[2])
			char := matches[3][0]
			pass := matches[4]

			matches := 0
			if pass[pos1-1] == char {
				matches++
			}
			if pass[pos2-1] == char {
				matches++
			}

			if matches == 1 {
				valid++
			}
		}
	}

	return valid
}
