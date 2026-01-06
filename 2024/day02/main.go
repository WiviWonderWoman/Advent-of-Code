package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day02")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 02, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 02, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	total := 0

	reports := getParsedReports(lines)

	for _, report := range reports {
		ok := checkReport(report)
		if ok {
			total++
		}
	}

	return total
}

func getParsedReports(lines []string) [][]int {
	reports := [][]int{}
	for _, raw := range lines {
		line := strings.Fields(raw)
		report := utils.StringToIntArr(line)
		reports = append(reports, report)
	}

	return reports
}

func checkReport(report []int) bool {
	safe := true

	if !findDuplicates(report, 0) {
		return false
	}

	increasing := report[0] < report[1]
	switch increasing {
	case true:
		safe = checkIncreasing(report)
	case false:
		safe = checkDecreasing(report)
	}

	return safe
}

func findDuplicates(report []int, rep int) bool {
	j := 0
	duplicates := 0

	for i := 1; i < len(report); i++ {
		if report[j] == report[i] {
			duplicates++
			if rep == 0 {
				break
			}
		}
		j++
	}

	return duplicates <= rep
}

func checkIncreasing(report []int) bool {
	safe := true
	for i := 1; i < len(report); i++ {
		if report[i]-report[i-1] > 3 || report[i]-report[i-1] < 1 {
			safe = false
			break
		}
	}

	return safe
}

func checkDecreasing(report []int) bool {
	safe := true
	for i := 1; i < len(report); i++ {
		if report[i-1]-report[i] > 3 || report[i-1]-report[i] < 1 {
			safe = false
			break
		}
	}

	return safe
}

func partTwo(lines []string) int {
	total := 0

	reports := getParsedReports(lines)

	for _, report := range reports {
		ok := checkReportPartTwo(report)
		if ok {
			total++
		}
	}

	return total
}

func checkReportPartTwo(report []int) bool {
	safe := true
	increasing := slices.IsSorted(report)

	switch increasing {
	case true:
		safe = checkIncreasing(report)
		if !safe {
			safe = checkEveryLevel(report, 0)
		}
	case false:
		safe = checkDecreasing(report)
		if !safe {
			safe = checkEveryLevel(report, 0)
		}
	}

	return safe
}

func checkEveryLevel(report []int, index int) bool {
	output := []int{}
	fixed := false

	increasing := slices.IsSorted(report)
	for i := index; i < len(report); i++ {
		level := report[i]
		if i != 0 {
			if !fixed && slices.Contains(output, level) {
				fixed = true
				index = i
				continue
			}

			switch increasing {
			case true:
				if !fixed && level-report[i-1] > 3 || level-report[i-1] < 1 {
					fixed = true
					continue
				}

			case false:
				if !fixed && level > report[i-1] || report[i-1]-level > 3 || report[i-1]-level < 1 {
					fixed = true
					continue
				}

			}
		}

		output = append(output, level)
	}

	increasing = slices.IsSorted(output)

	safe := false
	switch increasing {
	case true:
		safe = checkIncreasing(output)
	case false:
		safe = checkDecreasing(output)
	}

	if index != len(report)-1 {
		return safe
	}

	if !safe {
		safe = checkEveryLevel(report, index)
	}

	return safe
}
