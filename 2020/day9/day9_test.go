package day9

import (
	"testing"
)

func TestIsSumOfNumberPair(t *testing.T) {
	tests := map[int64]bool{
		25: true,
		38: true,
		26: false,
		39: false,
	}

	numbers := []int64{12, 13, 17, 21}

	for n, answer := range tests {
		result := isSumOfNumberPair(n, numbers)
		if result != answer {
			t.Fail()
			t.Errorf("Test failed. Expected %v, got=%v", answer, result)
		}
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestFindInvalidNumber(t *testing.T) {
	numbers := []int64{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
	}
	answer := int64(127)

	result := findInvalidNumber(numbers, 5)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestSumOfContiguousNumbers(t *testing.T) {
	numbers := []int64{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		350,
		421,
		682,
	}
	answer := [3]int64{15, 47, 62}

	low, high, sum := sumOfContiguousNumbers(127, numbers)
	if low != answer[0] || high != answer[1] || sum != answer[2] {
		t.Fail()
		t.Errorf("Test failed. Expected %d+%d=%d, got=%d+%d=%d", answer[0], answer[1], answer[2], low, high, sum)
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestSolvePart1(t *testing.T) {
	test := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

	answer := int64(127)

	result := SolvePart1(test)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestSolvePart2(t *testing.T) {
	test := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

	answer := int64(62)

	result := SolvePart2(test, 5)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}
