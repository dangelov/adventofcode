package day9

import (
	"strconv"
	"strings"
)

// isSumOfNumberPair checks whether a number is the sum
// of any two distinct numbers in a list
func isSumOfNumberPair(num int64, numbers []int64) bool {
	if len(numbers) < 2 {
		return false
	}

	for i := 0; i < len(numbers)-1; i++ {
		for o := 1; o < len(numbers); o++ {
			if numbers[i] == numbers[o] { // Can't be the same number
				continue
			}

			if numbers[i]+numbers[o] == num {
				return true
			}
		}
	}

	return false
}

// sumOfContiguousNumbers checks whether a number is the sum
// of any two distinct numbers in a list
func sumOfContiguousNumbers(num int64, numbers []int64) (int64, int64, int64) {
	if len(numbers) < 2 {
		return 0, 0, 0
	}

	for i := 0; i < len(numbers)-1; i++ {
		lowest := numbers[i]
		highest := numbers[i]
		interim := numbers[i]
		for o := i + 1; o < len(numbers); o++ {
			// The cumulative sum is too high, start again
			if interim+numbers[o] > num {
				break
			}

			// Remember the highest/lowest numbers
			if numbers[o] < lowest {
				lowest = numbers[o]
			}
			if numbers[o] > highest {
				highest = numbers[o]
			}

			// We found the numbers adding up to our desired sum
			if interim+numbers[o] == num {
				return lowest, highest, highest + lowest
			}

			interim += numbers[o]
		}
	}

	return 0, 0, 0
}

func findInvalidNumber(numbers []int64, preambleLen int) int64 {
	for i := preambleLen; i < len(numbers); i++ {
		if !isSumOfNumberPair(numbers[i], numbers[i-preambleLen:i]) {
			return numbers[i]
		}
	}

	return 0
}

// SolvePart1 takes input from a string
// and then solves the Day 7, Part 1 challenge
func SolvePart1(input string) int64 {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	numbers := make([]int64, len(lines))
	for i := 0; i < len(lines); i++ {
		numbers[i], _ = strconv.ParseInt(lines[i], 10, 64)
	}

	return findInvalidNumber(numbers, 25)
}

// SolvePart2 takes input from a string
// and then solves the Day 7, Part 2 challenge
func SolvePart2(input string, preambleLen int) int64 {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	numbers := make([]int64, len(lines))
	for i := 0; i < len(lines); i++ {
		numbers[i], _ = strconv.ParseInt(lines[i], 10, 64)
	}

	invNum := findInvalidNumber(numbers, preambleLen)
	_, _, sum := sumOfContiguousNumbers(invNum, numbers)

	return sum
}
