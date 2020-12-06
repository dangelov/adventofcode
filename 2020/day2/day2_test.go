package day2

import (
	"strings"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	test := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	testInput := strings.Join(test, "\n")
	answer := int64(2)

	result := SolvePart1(testInput)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestSolvePart2(t *testing.T) {
	test := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	testInput := strings.Join(test, "\n")
	answer := int64(1)

	result := SolvePart2(testInput)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}
