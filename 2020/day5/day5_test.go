package day5

import (
	"testing"
)

func TestDecodeSeat(t *testing.T) {
	tests := map[string][3]int64{
		"FBFBBFFRLR": {44, 5, 357},
		"BFFFBBFRRR": {70, 7, 567},
		"FFFBBBFRRR": {14, 7, 119},
		"BBFFBBFRLL": {102, 4, 820},
	}

	for input, output := range tests {
		row, column, seatID := DecodeSeat(input)
		if row != output[0] || column != output[1] || seatID != output[2] {
			t.Fail()
			t.Errorf("Test failed. Expected %v, got=%v", output, []int64{row, column, seatID})
		}
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestSolvePart1(t *testing.T) {
	test := `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

	answer := int64(820)

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
	test := `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
FFFBBBFRLR
BBFFBBFRLL`

	answer := int64(118)

	result := SolvePart2(test)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}
