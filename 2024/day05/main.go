package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day05")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 05, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 05, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	total := 0

	idx := slices.Index(lines, "")
	rules := lines[0:idx]
	updates := lines[idx:]

	// find valid updates
	valid := checkUpadates(updates, rules)

	// add middle page-numbers together
	for _, update := range valid {
		pages := strings.Split(update, ",")
		length := len(pages)

		middle := pages[length/2]
		if middle != "" {
			total += utils.StringToInt(middle)
		}

	}

	return total
}

func checkUpadates(updates, rules []string) []string {
	valid := []string{}

	for _, update := range updates {
		pages := strings.Split(update, ",")
		ok := true
		current := ""
		var before []string
		var after []string

		for i := range pages {
			current = pages[i]
			if current == "" {
				continue
			}

			before = pages[:i]
			after = pages[i+1:]

			ok = checkRules(rules, current, before, after)
			if !ok {
				break
			}
		}

		if !ok {
			continue
		}

		valid = append(valid, update)
	}

	return valid
}

func checkRules(rules []string, page string, before, after []string) bool {
	for _, rule := range rules {
		ok := true
		current := strings.Split(rule, "|")

		if strings.HasPrefix(rule, page) && len(before) > 0 {
			if slices.Contains(before, current[1]) {
				ok = false
			}
		} else if strings.HasSuffix(rule, page) && len(after) > 0 {
			if slices.Contains(after, current[0]) {
				ok = false
			}
		}

		if !ok {
			return false
		}
	}

	return true
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
