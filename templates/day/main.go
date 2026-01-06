package main

import (
	"adventofcode/utils"
	"fmt"
)

func GetTestInput() []string {
	return []string{}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day<day>")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day <day>, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day <day>, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	total := len(lines)

	return total
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
