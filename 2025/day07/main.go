package main

import (
	"adventofcode2025/utils"
	"fmt"
	"slices"
	"strings"
)

func GetTestInput() []string {
	return []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day07")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 07, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 07, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	total := 0
	previous := []int{}

	for row, line := range lines {
		current := []int{}

		for col, char := range line {
			if row == 0 && char == 'S' {
				previous = []int{col}
				continue
			}

			if slices.Contains(previous, col) {
				if char == '^' {
					total++
					current = append(current, col-1, col+1)
				} else {
					current = append(current, col)
				}
			}
		}

		if len(current) > 0 && slices.Compare(current, previous) != 0 {
			previous = current
		}
	}

	return total
}

func partTwo(lines []string) int {
	total := 0
	lines = removeLines(lines)

	previous := []int{}

	for row, line := range lines {
		current := []int{}
		for col, char := range line {
			if row == 0 && char == 'S' {
				previous = []int{col}
				continue
			}

			if slices.Contains(previous, col) {
				if char == '^' {
					current = append(current, col-1, col+1)
				} else {
					current = append(current, col)
				}
			}
		}

		if len(current) > 0 && slices.Compare(current, previous) != 0 {
			previous = current
			filteredBeams := []int{}
			for index, item := range current {
				if slices.Index(current, item) == index {
					filteredBeams = append(filteredBeams, item)
				}
			}

			total += len(filteredBeams)
		}
	}
	return total
}

func removeLines(input []string) []string {
	lines := []string{}
	for _, line := range input {
		if strings.Contains(line, "^") || strings.Contains(line, "S") {
			lines = append(lines, line)
		}
	}

	return lines
}
