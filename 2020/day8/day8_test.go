package day8

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	test := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

	answer := int64(5)

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
	test := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

	answer := int64(8)

	result := SolvePart2(test)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}
