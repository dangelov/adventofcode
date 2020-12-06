# Advent of Code

This repo contains solutions to [Advent of Code events](https://adventofcode.com/events), organized by year.

## How to run

There is a separate Go code and file for each day. Each file expects input in the form of a file in the same directory, named after the day + part + `.input`. For example, to run the Day 1 solution, create two files named `day1-part1.input` and `day1-part2.input` and then execute `go run day1.go`. Make sure you `cd` into the year directory first.

## Tests

The code was written using TDD. Each day has a `_test.go` containing tests for both parts of the puzzle, using the sample input provided from Advent of Code. You can run all tests by executing `go test ./...` in the root of the repo.

All code is licensed under the MIT license.
