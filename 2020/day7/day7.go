package day7

import (
	"strconv"
	"strings"
)

// bagsContained counts the number of bags
// contained within a bag
func bagsContained(bag string, rules map[string]map[string]int64) int64 {
	// This bag can't contain any other bags
	if len(rules[bag]) == 0 {
		return 0
	}

	total := int64(0)
	for b, count := range rules[bag] {
		total += count
		total += count * bagsContained(b, rules)
	}

	return total
}

// canContainBag checks whether an outerBag
// can contain a bag
func canContainBag(outerBag, bag string, rules map[string]map[string]int64) bool {
	// This bag can't contain any other bags
	if len(rules[outerBag]) == 0 {
		return false
	}

	canContain := false
	for b := range rules[outerBag] {
		// Is this our bag?
		if b == bag {
			return true
		}

		// If not, keep looking
		canContain = canContain || canContainBag(b, bag, rules)
	}

	return canContain
}

// parseRules gets a string of rules as input and produces
// a map describing which bags can contain which other bags
func parseRules(input string) map[string]map[string]int64 {
	result := map[string]map[string]int64{}

	// Cleanup the input of irrelevant words & punctuation
	input = strings.ReplaceAll(input, "bags", "")
	input = strings.ReplaceAll(input, "bag", "")
	input = strings.ReplaceAll(input, ".", "")

	rules := strings.Split(strings.TrimRight(input, "\n"), "\n")
	for i := 0; i < len(rules); i++ {
		mapping := strings.Split(rules[i], "  contain ")
		// If this bag can contain other bags, they'd be split by a comma
		contents := strings.Split(mapping[1], ", ")

		if strings.Contains(contents[0], "no other") {
			// This bag can't contain any other bags
			result[mapping[0]] = map[string]int64{}
		} else {
			// Can contain other bags, find which
			val := map[string]int64{}
			for _, content := range contents {
				// Get two parts, splitting on first whitespace
				parts := strings.SplitN(content, " ", 2)
				count, _ := strconv.ParseInt(parts[0], 10, 64)
				val[strings.Trim(parts[1], " ")] = count
			}
			result[mapping[0]] = val
		}
	}

	return result
}

// SolvePart1 takes input from a string
// and then solves the Day 7, Part 1 challenge
func SolvePart1(input string) int64 {
	total := int64(0)

	rules := parseRules(input)
	bag := "shiny gold"

	for outerBag := range rules {
		if canContainBag(outerBag, bag, rules) {
			total++
		}
	}

	return total
}

// SolvePart2 takes input from a string
// and then solves the Day 7, Part 2 challenge
func SolvePart2(input string) int64 {
	rules := parseRules(input)

	return bagsContained("shiny gold", rules)
}
