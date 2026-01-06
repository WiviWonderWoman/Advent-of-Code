package main

import (
	"adventofcode/utils"
	"fmt"
)

func GetTestInput() []string {
	return []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day04")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	lines := GetInput()
	// lines := GetTestInput()

	fmt.Println("Day 04, Part 1 Answer: ", dayFour(lines, 0, false))
	fmt.Println("Day 04, Part 2 Answer: ", dayFour(lines, 0, true))
}

func dayFour(lines []string, total int, swap bool) int {
	current := 0
	currentLines := []string{}
	for row := 0; row < len(lines); row++ {
		currentLine := ""
		line := lines[row]
		up := row == 0
		down := row == len(lines)-1

		for col := 0; col < len(line); col++ {
			char := line[col]
			currentChar := char

			left := col == 0
			right := col == len(line)-1
			count := 0
			if char == '@' {
				toCheck := []string{}
				if !up {
					toCheck = append(toCheck, string(lines[row-1][col]))
					if !left {
						toCheck = append(toCheck, string(lines[row-1][col-1]))
					}
					if !right {
						toCheck = append(toCheck, string(lines[row-1][col+1]))
					}
				}

				if !down {
					toCheck = append(toCheck, string(lines[row+1][col]))
					if !left {
						toCheck = append(toCheck, string(lines[row+1][col-1]))
					}
					if !right {
						toCheck = append(toCheck, string(lines[row+1][col+1]))
					}
				}

				if !right {
					toCheck = append(toCheck, string(line[col+1]))
				}
				if !left {
					toCheck = append(toCheck, string(line[col-1]))
				}

				count = checkAdjacent(toCheck)
				if count < 4 {
					total++
					current++
					currentChar = 'x'
				}
			}
			currentLine += string(currentChar)

		}
		currentLines = append(currentLines, currentLine)
	}

	if swap && current > 0 {
		return dayFour(currentLines, total, swap)
	}

	return total
}

func checkAdjacent(chars []string) int {
	totalValue := 0
	for _, char := range chars {
		if char == "@" {
			totalValue++
		}
	}

	return totalValue
}
