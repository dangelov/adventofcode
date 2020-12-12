package main

import (
	"io/ioutil"
	"log"

	"dangelov.com/adventofcode/2020/day1"
	"dangelov.com/adventofcode/2020/day2"
	"dangelov.com/adventofcode/2020/day3"
	"dangelov.com/adventofcode/2020/day4"
	"dangelov.com/adventofcode/2020/day5"
	"dangelov.com/adventofcode/2020/day6"
	"dangelov.com/adventofcode/2020/day7"
	"dangelov.com/adventofcode/2020/day8"
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

	// # Day 5
	content, err = ioutil.ReadFile("day5-part1.input")
	if err != nil {
		log.Fatal(err)
	}
	input = string(content)

	print("Day 5, Part 1: ")
	println(day5.SolvePart1(input))
	print("Day 5, Part 2: ")
	println(day5.SolvePart2(input))

	// # Day 6
	content, err = ioutil.ReadFile("day6-part1.input")
	if err != nil {
		log.Fatal(err)
	}
	input = string(content)

	print("Day 6, Part 1: ")
	println(day6.SolvePart1(input))
	print("Day 6, Part 2: ")
	println(day6.SolvePart2(input))

	// # Day 7
	content, err = ioutil.ReadFile("day7-part1.input")
	if err != nil {
		log.Fatal(err)
	}
	input = string(content)

	print("Day 7, Part 1: ")
	println(day7.SolvePart1(input))
	print("Day 7, Part 2: ")
	println(day7.SolvePart2(input))

	// # Day 8
	content, err = ioutil.ReadFile("day8-part1.input")
	if err != nil {
		log.Fatal(err)
	}
	input = string(content)

	print("Day 8, Part 1: ")
	println(day8.SolvePart1(input))
	print("Day 8, Part 2: ")
	println(day8.SolvePart2(input))
}
