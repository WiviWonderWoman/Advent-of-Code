package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day01")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 01, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 01, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	total := 0

	leftSide, rightSide := getSortedLocationIDs(lines)

	for i, left := range leftSide {
		total += int(math.Abs(float64(left) - float64(rightSide[i])))
	}

	return total
}

func getSortedLocationIDs(lines []string) ([]int, []int) {
	leftSide := []int{}
	rightSide := []int{}
	for _, line := range lines {
		parts := strings.Fields(line)
		leftSide = append(leftSide, utils.StringToInt(parts[0]))
		rightSide = append(rightSide, utils.StringToInt(parts[1]))
	}

	slices.Sort(leftSide)
	slices.Sort(rightSide)

	return leftSide, rightSide
}

func partTwo(lines []string) int {
	total := 0

	leftSide, rightSide := getSortedLocationIDs(lines)

	for _, left := range leftSide {
		multiplier := 0
		for _, right := range rightSide {
			if right == left {
				multiplier++
			}
		}
		total += left * multiplier
	}

	return total
}
