package main

import (
	"adventofcode2025/utils"
	"fmt"
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
	fmt.Println(reports)

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
	total := len(lines)

	reports := getParsedReports(lines)
	fmt.Println(reports)

	return total
}
