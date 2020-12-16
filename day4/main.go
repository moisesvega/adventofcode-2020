package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/moisesvega/adventofcode-2020/utils"
)

const byr = "byr" // (Birth Year)
const iyr = "iyr" // (Issue Year)
const eyr = "eyr" // (Expiration Year)
const hgt = "hgt" // (Height)
const hcl = "hcl" // (Hair Color)
const ecl = "ecl" // (Eye Color)
const pid = "pid" // (Passport ID)
const cid = "cid" // (Country ID)

var rules = []string{byr, iyr, eyr, hgt, hcl, ecl, pid}

func main() {

	passports := pasePassports()
	// log.Println(passports)
	count := 0
	for _, passport := range passports {
		// log.Println(passport)
		if isPassportValid(rules, passport) {
			count++
		}
	}

	log.Println(count)

}

func isPassportValid(rules []string, passport map[string]string) bool {
	for _, rule := range rules {
		if _, ok := passport[rule]; !ok {
			return false
		}

		// part2
		if !valid(passport) {
			return false
		}
	}
	return true
}

func pasePassports() (out []map[string]string) {
	out = make([]map[string]string, 0)

	lines := utils.ReadLinesFromFile("/Users/moises/go/src/github.com/moisesvega/adventofcode-2020/day4/input.in")
	passport := map[string]string{}
	for _, line := range lines {

		if len(line) == 0 {
			out = append(out, passport)
			passport = map[string]string{}
			continue
		}

		fields := strings.Split(line, " ")

		for _, field := range fields {
			values := strings.Split(field, ":")
			passport[values[0]] = values[1]
		}

	}
	out = append(out, passport)
	return out
}

func valid(passport map[string]string) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.

	if len(passport[byr]) != 4 {
		return false
	}
	if len(passport[iyr]) != 4 {
		return false
	}
	if len(passport[eyr]) != 4 {
		return false
	}

	byrInt, err := strconv.Atoi(passport[byr])
	if err != nil {
		return false
	}
	iyrInt, err := strconv.Atoi(passport[iyr])
	if err != nil {
		return false
	}
	eyrInt, err := strconv.Atoi(passport[eyr])
	if err != nil {
		return false
	}

	if byrInt > 2002 || byrInt < 1920 {
		return false
	}

	if iyrInt > 2020 || iyrInt < 2010 {
		return false
	}

	if eyrInt > 2030 || eyrInt < 2020 {
		return false
	}

	heightInCm := strings.Split(passport[hgt], "cm")
	if len(heightInCm) > 1 {
		// If cm, the number must be at least 150 and at most 193.
		cm, err := strconv.Atoi(heightInCm[0])
		if err != nil {
			return false
		}

		if cm < 150 || cm > 193 {
			return false
		}
	}

	heightIn := strings.Split(passport[hgt], "in")
	if len(heightIn) > 1 {
		// If in, the number must be at least 59 and at most 76.
		in, err := strconv.Atoi(heightIn[0])
		if err != nil {
			return false
		}

		if in < 59 || in > 76 {
			return false
		}
	}

	if len(heightInCm) != 2 && len(heightIn) != 2 {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hairColor := strings.Split(passport[hcl], "#")
	if len(hairColor) != 2 || len(hairColor[1]) != 6 {
		return false
	}

	matched, _ := regexp.MatchString(`[a-f0-9]`, hairColor[1])
	if !matched {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.

	validColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if _, ok := validColors[passport[ecl]]; !ok {
		return false
	}
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	matched, _ = regexp.MatchString(`[0-9]`, passport[pid])
	if !matched || len(passport[pid]) != 9 {
		return false
	}

	// cid (Country ID) - ignored, missing or not.

	return true
}
