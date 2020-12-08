package main

import (
	"io/ioutil"
	"log"

	"dangelov.com/adventofcode/2020/day1"
	"dangelov.com/adventofcode/2020/day2"
	"dangelov.com/adventofcode/2020/day3"
	"dangelov.com/adventofcode/2020/day4"
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

	// # Day 2
	content, err = ioutil.ReadFile("day2-part1.input")
	if err != nil {
		log.Fatal(err)
	}
	input = string(content)

	print("Day 2, Part 1: ")
	println(day2.SolvePart1(input))
	print("Day 2, Part 2: ")
	println(day2.SolvePart2(input))

	// # Day 3
	content, err = ioutil.ReadFile("day3-part1.input")
	if err != nil {
		log.Fatal(err)
	}
	input = string(content)

	print("Day 3, Part 1: ")
	println(day3.SolvePart1(input, 3, 1))
	print("Day 3, Part 2: ")
	println(day3.SolvePart2(input))

	// # Day 4
	content, err = ioutil.ReadFile("day4-part1.input")
	if err != nil {
		log.Fatal(err)
	}
	input = string(content)

	print("Day 4, Part 1: ")
	println(day4.SolvePart1(input))
	print("Day 4, Part 2: ")
	println(day4.SolvePart2(input))
}
