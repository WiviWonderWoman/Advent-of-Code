package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strconv"
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
	valid, _ := sortUpdates(updates, rules)

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

func sortUpdates(updates, rules []string) ([]string, []string) {
	valid := []string{}
	invalid := []string{}

	for _, update := range updates {
		pages := strings.Split(update, ",")
		ok := true

		for i, current := range pages {
			if current == "" {
				continue
			}

			before := pages[:i]
			after := pages[i+1:]

			ok = checkRules(rules, current, before, after)
			if !ok {
				break
			}

		}

		if !ok {
			if !slices.Contains(invalid, update) {
				invalid = append(invalid, update)
			}
		}

		if ok {
			if !slices.Contains(valid, update) {
				valid = append(valid, update)
			}
		}

	}

	return valid, invalid
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
	idx := slices.Index(lines, "")
	rules := lines[0:idx]
	slices.Sort(rules)

	updates := lines[idx:]

	// find invalid updates
	_, invalid := sortUpdates(updates, rules)
	fmt.Println("invalid: ", len(invalid))

	// get sorted Rule map
	ruleMap := make(map[string]Rule)
	for _, rule := range rules {
		r := getStructuredRule(rule, rules)
		ruleMap[strconv.Itoa(r.value)] = r
	}
	// fmt.Println("ruleMap: ", ruleMap)

	return total
}

type Rule struct {
	value       int
	comesBefore []int
	comesAfter  []int
}

func getStructuredRule(input string, rules []string) Rule {
	inputParts := strings.Split(input, "|")
	structured := Rule{
		value:       utils.StringToInt(inputParts[0]),
		comesBefore: []int{utils.StringToInt(inputParts[1])},
		comesAfter:  []int{},
	}

	for _, rule := range rules {
		if rule == input {
			continue
		}

		if strings.Contains(rule, inputParts[0]) {
			ruleParts := strings.Split(rule, "|")
			if ruleParts[0] == inputParts[0] {
				// comesBefore append [1]
				structured.comesBefore = append(structured.comesBefore, utils.StringToInt(ruleParts[1]))
			}

			if ruleParts[1] == inputParts[0] {
				// comesAfter append [0]
				structured.comesAfter = append(structured.comesAfter, utils.StringToInt(ruleParts[0]))
			}
		}
	}

	slices.Sort(structured.comesAfter)
	slices.Sort(structured.comesBefore)

	return structured
}
