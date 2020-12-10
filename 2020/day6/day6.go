package day6

import (
	"strings"
)

// SolvePart1 takes input from a string
// and then solves the Day 6, Part 1 challenge
func SolvePart1(input string) int64 {
	groups := strings.Split(input, "\n\n")
	total := int64(0)

	for i := 0; i < len(groups); i++ {
		answers := map[byte]bool{}

		all := strings.ReplaceAll(groups[i], "\n", "")
		for p := 0; p < len(all); p++ {
			answers[all[p]] = true
		}

		total += int64(len(answers))
	}

	return total
}

// SolvePart2 takes input from a string
// and then solves the Day 6, Part 2 challenge
func SolvePart2(input string) int64 {
	groups := strings.Split(input, "\n\n")
	total := int64(0)

	for i := 0; i < len(groups); i++ {
		answers := map[byte]int{}

		all := strings.ReplaceAll(groups[i], "\n", "")
		for p := 0; p < len(all); p++ {
			answers[all[p]]++
		}

		groupCount := len(strings.Split(strings.TrimRight(groups[i], "\n"), "\n"))
		for _, count := range answers {
			if count == groupCount {
				total++
			}
		}
	}

	return total
}
