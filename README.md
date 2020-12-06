# Advent of Code

This repo contains solutions to [Advent of Code events](https://adventofcode.com/events), organized by year.

## How to run

There is a separate Go package for each day. All packages (days) are run from the `main.go` file, which provides input to each package in the format of "day" + "part" + ".input". For example, to run the Day 1 solution, create a file named `day1-part1.input` and then execute `go run main.go`. Make sure you `cd` into the year directory first.

## Tests

The code was written using TDD. Each day has a `_test.go` containing tests for both parts of the puzzle, using the sample input provided from Advent of Code. You can run all tests by executing `go test ./...` in the root of the repo.

All code is licensed under the MIT license.
