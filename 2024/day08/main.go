package main

import (
	"adventofcode/utils"
	"fmt"
)

func GetTestInput() []string {
	return []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day08")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 08, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 08, Part 2 Answer: ", partTwo(lines))
}

type Position struct {
	row int
	col int
}

func partOne(lines []string) int {
	total := len(lines)

	return total
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
