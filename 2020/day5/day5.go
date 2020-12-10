package day5

import (
	"sort"
	"strings"
)

// DecodeSeat converts binary space partitioning
// into a row, column and a seat ID
func DecodeSeat(seat string) (int64, int64, int64) {
	rowL := int64(0)
	rowH := int64(128)
	columnL := int64(0)
	columnH := int64(8)

	for i := 0; i < len(seat); i++ {
		switch seat[i] {
		case 'B':
			rowL = rowL + ((rowH - rowL) / 2)
		case 'F':
			rowH = rowH - ((rowH - rowL) / 2)
		case 'R':
			columnL = columnL + ((columnH - columnL) / 2)
		case 'L':
			columnH = columnH - ((columnH - columnL) / 2)
		}
	}
	return rowL, columnL, rowL*8 + columnL
}

// SolvePart1 takes input from a string
// and then solves the Day 5, Part 1 challenge
func SolvePart1(input string) int64 {
	passes := strings.Split(input, "\n")
	highest := int64(0)

	for i := 0; i < len(passes); i++ {
		_, _, seatID := DecodeSeat(passes[i])
		if seatID > highest {
			highest = seatID
		}
	}

	return highest
}

// SolvePart2 takes input from a string
// and then solves the Day 5, Part 2 challenge
func SolvePart2(input string) int64 {
	passes := strings.Split(input, "\n")
	seatIDs := make([]int64, len(passes), len(passes))

	for i := 0; i < len(passes); i++ {
		_, _, seatID := DecodeSeat(passes[i])
		seatIDs[i] = seatID
	}

	sort.Slice(seatIDs, func(i, j int) bool { return seatIDs[i] < seatIDs[j] })
	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i]+2 == seatIDs[i+1] {
			return seatIDs[i] + 1
		}
	}

	return 0
}
