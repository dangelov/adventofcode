package day4

import (
	"regexp"
	"strconv"
	"strings"
)

// SolvePart1 takes input from a string
// and then solves the Day 4, Part 1 challenge
func SolvePart1(input string) int64 {
	passports := strings.Split(input, "\n\n")

	reqs := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid",
	}

	valid := int64(0)

	for i := 0; i < len(passports); i++ {
		// Normalize the way key:values are split
		data := strings.ReplaceAll(passports[i], "\n", " ")
		// Get all the key:value pairs
		pairs := strings.Split(data, " ")

		// Go through all the keys and see how many
		// required keys we can find
		numKeys := 0
		for o := 0; o < len(pairs); o++ {
			pair := strings.Split(pairs[o], ":")

			// Check the requirements for our current key
			for p := 0; p < len(reqs); p++ {
				if pair[0] == reqs[p] {
					numKeys++
				}
			}
		}

		// Having the same number of validated and required keys
		// means this passport is valid
		if numKeys == len(reqs) {
			valid++
		}
	}

	return valid
}

// SolvePart2 takes input from a string
// and then solves the Day 4, Part 2 challenge
func SolvePart2(input string) int64 {
	passports := strings.Split(input, "\n\n")

	reqs := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid",
	}

	valid := int64(0)

	for i := 0; i < len(passports); i++ {
		// Normalize the way key:values are split
		data := strings.ReplaceAll(passports[i], "\n", " ")
		// Get all the key:value pairs
		pairs := strings.Split(data, " ")

		// Go through all the keys and see how many
		// required keys we can find
		numKeys := 0
		for o := 0; o < len(pairs); o++ {
			pair := strings.Split(pairs[o], ":")

			// Check the requirements for our current key
			for p := 0; p < len(reqs); p++ {
				if pair[0] == reqs[p] {
					val := pair[1]
					switch pair[0] {
					case "byr": // Validate birth year
						num, err := strconv.ParseInt(val, 10, 64)
						if err != nil {
							continue
						}
						if num < 1920 || num > 2002 {
							continue
						}
					case "iyr": // Validate issue year
						num, err := strconv.ParseInt(val, 10, 64)
						if err != nil {
							continue
						}
						if num < 2010 || num > 2020 {
							continue
						}
					case "eyr": // Validate expiration year
						num, err := strconv.ParseInt(val, 10, 64)
						if err != nil {
							continue
						}
						if num < 2020 || num > 2030 {
							continue
						}
					case "hgt": // Validate height
						// Must be at least a 2-digit number + 2-letter unit
						if len(val) < 4 {
							continue
						}

						// Last 2 characters are the unit
						unit := val[len(val)-2:]

						// Only "in" or "cm" accepted
						if unit != "in" && unit != "cm" {
							continue
						}

						num, err := strconv.ParseInt(strings.Replace(val, unit, "", 1), 10, 64)
						if err != nil {
							continue
						}
						if unit == "cm" && (num < 150 || num > 193) {
							continue
						}
						if unit == "in" && (num < 59 || num > 76) {
							continue
						}
					case "hcl": // Validate hair color
						if len(val) != 7 { // "#ABCABC"
							continue
						}

						if val[0] != '#' { // Must start with a "#"
							continue
						}

						hex := val[1:] // Remove the leading #

						reg := regexp.MustCompile("[0-9a-fA-F]+")
						if !reg.Match([]byte(hex)) {
							continue
						}
					case "ecl":
						allowed := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
						valid := false
						for i := 0; i < len(allowed); i++ {
							if val == allowed[i] {
								valid = true
							}
						}

						if !valid {
							continue
						}
					case "pid":
						if len(val) != 9 {
							continue
						}

						reg := regexp.MustCompile("[0-9]+")
						if !reg.Match([]byte(val)) {
							continue
						}
					}
					numKeys++
				}
			}
		}

		// Having the same number of validated and required keys
		// means this passport is valid
		if numKeys == len(reqs) {
			valid++
		}
	}

	return valid
}
