package day1

import "testing"

func TestSolvePart1(t *testing.T) {
	test := "1721\n979\n366\n299\n675\n1456"
	answer := int64(514579)

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
	test := "1721\n979\n366\n299\n675\n1456"
	answer := int64(241861950)

	result := SolvePart2(test)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}
