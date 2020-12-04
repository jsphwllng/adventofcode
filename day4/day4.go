package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Input is the parsed text file
var Input = parse()

func parse() (parsed [][]string) {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		parsed = append(parsed, fields)
	}
	return
}

func main() {
	fmt.Println("Answer: ", part1(Input))
	fmt.Println("Answer: ", part2(Input))
}

func part1(input [][]string) (anwser int) {
loop:
	for _, row := range input {
		fields := map[string]string{}
		required := []string{
			"byr",
			"iyr",
			"eyr",
			"hgt",
			"hcl",
			"ecl",
			"pid",
		}
		for _, field := range row {
			s := strings.Split(field, ":")
			fields[s[0]] = s[1]
		}
		for _, k := range required {
			if fields[k] == "" {
				continue loop
			}
		}
		anwser++
	}
	return
}

func part2(input [][]string) (anwser int) {
	type rule struct {
		min, max int
		re       *regexp.Regexp
	}
loop:
	for _, row := range input {
		fields := map[string]string{}
		required := map[string]rule{
			"byr": {
				min: 1920, max: 2002,
				re: regexp.MustCompile(`^(\d{4})$`)},
			"iyr": {
				min: 2010, max: 2020,
				re: regexp.MustCompile(`^(\d{4})$`)},
			"eyr": {
				min: 2020, max: 2030,
				re: regexp.MustCompile(`^(\d{4})$`)},
			"hgt": {
				min: 150, max: 193,
				re: regexp.MustCompile(`^(\d+)(cm|in)$`)},
			"hcl": {
				re: regexp.MustCompile(`^#[0-9a-f]{6}$`)},
			"ecl": {re: regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)},
			"pid": {re: regexp.MustCompile(`^(\d{9})$`)},
			"cid": {re: regexp.MustCompile(`.*`)},
		}
		for _, field := range row {
			s := strings.Split(field, ":")
			rule := required[s[0]]
			if rule.re == nil {
				continue loop
			}
			if !rule.re.MatchString(s[1]) {
				continue loop
			}
			switch s[0] {
			case "byr":
				fallthrough
			case "iyr":
				fallthrough
			case "eyr":
				u, err := strconv.Atoi(rule.re.FindStringSubmatch(s[1])[1])
				if err != nil {
					panic(err)
				}
				if u < rule.min || u > rule.max {
					continue loop
				}
			case "hgt":
				u, err := strconv.Atoi(rule.re.FindStringSubmatch(s[1])[1])
				if err != nil {
					panic(err)
				}
				switch rule.re.FindStringSubmatch(s[1])[2] {
				case "in":
					if u < 59 || u > 76 {
						continue loop
					}
				default:
					if u < rule.min || u > rule.max {
						continue loop
					}
				}
			}
			fields[s[0]] = s[1]
		}
		delete(fields, "cid")
		fields["cid"] = "ignored"
		for k := range required {
			if fields[k] == "" {
				continue loop
			}
		}
		anwser++
	}
	return
}
