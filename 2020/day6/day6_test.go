package day6

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	test := `abc

a
b
c

ab
ac

a
a
a
a

b`

	answer := int64(11)

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
	test := `abc

a
b
c

ab
ac

a
a
a
a

b`

	answer := int64(6)

	result := SolvePart2(test)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}
