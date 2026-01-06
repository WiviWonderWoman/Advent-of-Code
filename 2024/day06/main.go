package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day06")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 06, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 06, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	total := 0

	up := "^"
	right := ">"
	down := "v"
	left := "<"
	mark := "X"

	grid := getGrid(lines)

	for row := 0; row < len(grid); row++ {
		line := grid[row]

		for col := 0; col < len(line); col++ {
			str := line[col]

			switch str {
			case up:
				if grid[row-1][col] == "#" {
					// turn => right
					turned := replaceString(line, []string{right}, col)
					grid[row] = turned

				} else {
					marked := replaceString(line, []string{mark}, col)
					grid[row] = marked

					moved := replaceString(grid[row-1], []string{up}, col)
					grid[row-1] = moved
				}
				row = 0
			case right:
				if len(line) == col+1 {
					break
				}

				if line[col+1] == "#" {
					// turn => down
					turned := replaceString(line, []string{down}, col)
					grid[row] = turned

					col = 0
				} else {
					marked := replaceString(line, []string{mark}, col)
					moved := replaceString(marked, []string{right}, col+1)

					grid[row] = moved
				}

			case down:
				if len(grid) == row+1 {
					marked := replaceString(line, []string{mark}, col)
					grid[row] = marked
					break
				}

				if grid[row+1][col] == "#" {
					// turn => left
					turned := replaceString(line, []string{left}, col)
					grid[row] = turned

					col = 0
				} else {
					marked := replaceString(line, []string{mark}, col)
					grid[row] = marked

					moved := replaceString(grid[row+1], []string{down}, col)
					grid[row+1] = moved
				}

			case left:
				if line[col-1] == "#" {
					// turn => up
					turned := replaceString(line, []string{up}, col)
					grid[row] = turned

				} else {
					marked := replaceString(line, []string{mark}, col)
					moved := replaceString(marked, []string{left}, col-1)

					grid[row] = moved

				}

				col = 0
			}

			if len(line) == col+1 {
				break
			}
		}
	}

	for _, line := range grid {
		for _, str := range line {
			if str == "X" {
				total++
			}
		}
	}
	return total
}

func replaceString(input, replace []string, idx int) []string {
	var output []string

	output = append(input[:idx], append(replace, input[idx+1:]...)...)

	return output
}

func getGrid(lines []string) [][]string {
	output := [][]string{}
	for _, str := range lines {
		output = append(output, strings.Split(str, ""))
	}
	return output
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
