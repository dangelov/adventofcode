package day7

import (
	"reflect"
	"testing"
)

func TestBagsContained(t *testing.T) {
	inputRules := `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`
	rules := parseRules(inputRules)

	answer := int64(126)

	result := bagsContained("shiny gold", rules)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestCanContainBag(t *testing.T) {
	inputRules := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`
	rules := parseRules(inputRules)

	answers := map[string]bool{
		"bright white": true,
		"muted yellow": true,
		"dark orange":  true,
		"light red":    true,
		"shiny gold":   false,
		"dark olive":   false,
		"vibrant plum": false,
		"faded blue":   false,
		"dotted black": false,
	}
	bag := "shiny gold"

	for outerBag, answer := range answers {
		result := canContainBag(outerBag, bag, rules)
		if result != answer {
			t.Fail()
			t.Errorf("Test failed. Expected %v for %s, got=%v", answer, outerBag, result)
		}
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestParseRules(t *testing.T) {
	test := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
faded blue bags contain no other bags.
`

	answer := map[string]map[string]int64{
		"light red":   {"bright white": 1, "muted yellow": 2},
		"dark orange": {"bright white": 3, "muted yellow": 4},
		"faded blue":  {},
	}

	result := parseRules(test)
	if !reflect.DeepEqual(result, answer) {
		t.Fail()
		t.Errorf("Test failed. Expected % v, got=%v", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}

func TestSolvePart1(t *testing.T) {
	test := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

	answer := int64(4)

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
	test := `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`

	answer := int64(126)

	result := SolvePart2(test)
	if result != answer {
		t.Fail()
		t.Errorf("Test failed. Expected %d, got=%d", answer, result)
	}

	if t.Failed() {
		t.FailNow()
	}
}
