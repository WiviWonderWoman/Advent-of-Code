package main

import (
	"adventofcode2025/utils"
	"fmt"
)

func GetTestInput() []string {
	return []string{}
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
	total := len(lines)

	return total
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
