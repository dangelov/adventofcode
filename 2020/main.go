package main

import (
	"io/ioutil"
	"log"

	"dangelov.com/adventofcode/2020/day1"
)

func main() {
	// This function runs the solutions to all the days, it expects
	// each day to have a file named `dayN-partM.input` where N
	// is the day and M the part of that day. Some days (or all days?)
	// have the same input for both parts.

	// # Day 1
	content, err := ioutil.ReadFile("day1-part1.input")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)

	print("Day 1, Part 1: ")
	println(day1.SolvePart1(input))
	print("Day 1, Part 2: ")
	println(day1.SolvePart2(input))
}
