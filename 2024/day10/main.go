package main

import (
	"adventofcode/utils"
	"fmt"
)

func GetTestInput() []string {
	return []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day10")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 10, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 10, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	total := len(lines)

	return total
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
