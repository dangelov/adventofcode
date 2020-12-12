package day8

import (
	"errors"
	"strconv"
	"strings"
)

// SolvePart1 takes input from a string
// and then solves the Day 7, Part 1 challenge
func SolvePart1(input string) int64 {
	instructions := strings.Split(strings.TrimRight(input, "\n"), "\n")
	val := int64(0)
	visited := map[int64]bool{}

	cur := int64(0)
	for cur < int64(len(instructions)) {
		if visited[cur] {
			return val
		}
		visited[cur] = true

		parts := strings.Split(instructions[cur], " ")
		num, _ := strconv.ParseInt(parts[1], 10, 64)
		switch parts[0] {
		case "nop":
			cur++
		case "acc":
			val += num
			cur++
		case "jmp":
			cur += num
		}
	}

	return val
}

func runProgram(instructions []string) (int64, error) {
	val := int64(0)
	visited := map[int64]bool{}

	cur := int64(0)
	for cur < int64(len(instructions)) {
		if visited[cur] {
			return val, errors.New("Loop detected")
		}
		visited[cur] = true

		parts := strings.Split(instructions[cur], " ")
		num, _ := strconv.ParseInt(parts[1], 10, 64)
		switch parts[0] {
		case "nop":
			cur++
		case "acc":
			val += num
			cur++
		case "jmp":
			cur += num
		}
	}

	return val, nil
}

// SolvePart2 takes input from a string
// and then solves the Day 7, Part 2 challenge
func SolvePart2(input string) int64 {
	instructions := strings.Split(strings.TrimRight(input, "\n"), "\n")
	tmp := make([]string, len(instructions))

	// Change each instruction once and try to run the program
	for i := 0; i < len(instructions); i++ {
		// "acc" never changes (can't be corrupted)
		if strings.Index(instructions[i], "acc") > -1 {
			continue
		}
		copy(tmp, instructions)

		// Switch jmp to nop and vice versa
		if strings.Index(tmp[i], "jmp") > -1 {
			tmp[i] = strings.ReplaceAll(tmp[i], "jmp", "nop")
		} else {
			tmp[i] = strings.ReplaceAll(tmp[i], "nop", "jmp")
		}
		res, err := runProgram(tmp)
		if err == nil {
			return res
		}
	}

	return 0
}
