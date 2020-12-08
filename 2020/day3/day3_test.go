package day3

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	test := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	answer := int64(7)
	shiftX := int64(3)
	shiftY := int64(1)

	result := SolvePart1(test, shiftX, shiftY)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestSolvePart2(t *testing.T) {
	test := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	answer := int64(336)

	result := SolvePart2(test)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}
