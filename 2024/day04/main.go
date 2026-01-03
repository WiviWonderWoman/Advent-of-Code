package main

import (
	"adventofcode2025/utils"
	"fmt"
)

func GetTestInput() []string {
	return []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
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
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 04, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 04, Part 2 Answer: ", partTwo(lines))
}

const (
	word = "XMAS"
)

func partOne(lines []string) int {
	total := 0

	total += findHorizontal(lines)
	total += findVertical(lines)
	total += findDiagonal(lines)

	return total
}

func findHorizontal(lines []string) int {
	total := 0
	for _, line := range lines {
		for col, char := range line {
			columns := len(line) - 3

			if char == 'X' {
				if col >= 3 {
					// check backwards
					check := string(char) + string(+line[col-1]) + string(+line[col-2]) + string(+line[col-3])
					if check == word {
						total++
					}
				}
				if col < columns {
					// check forward
					check := string(char) + string(+line[col+1]) + string(+line[col+2]) + string(+line[col+3])
					if check == word {
						total++
					}
				}
			}
		}
	}
	return total
}

func findVertical(lines []string) int {
	total := 0
	for row, line := range lines {
		rows := len(lines) - 3

		for col, char := range line {
			if char == 'X' {
				if row >= 3 {
					// check up "backwards"
					check := string(
						char,
					) + string(
						+lines[row-1][col],
					) + string(
						+lines[row-2][col],
					) + string(
						+lines[row-3][col],
					)
					if check == word {
						total++
					}
				}
				if row < rows {
					// check down "forward"
					check := string(
						char,
					) + string(
						+lines[row+1][col],
					) + string(
						+lines[row+2][col],
					) + string(
						+lines[row+3][col],
					)
					if check == word {
						total++
					}
				}
			}
		}
	}
	return total
}

func findDiagonal(lines []string) int {
	total := 0

	for row, line := range lines {
		rows := len(lines) - 3

		for col, char := range line {
			columns := len(line) - 3

			if char == 'X' {
				// check up right
				if row >= 3 && col < columns {
					check := string(
						char,
					) + string(
						+lines[row-1][col+1],
					) + string(
						+lines[row-2][col+2],
					) + string(
						+lines[row-3][col+3],
					)
					if check == word {
						total++
					}
				}
				// check down right
				if row < rows && col < columns {
					check := string(
						char,
					) + string(
						+lines[row+1][col+1],
					) + string(
						+lines[row+2][col+2],
					) + string(
						+lines[row+3][col+3],
					)
					if check == word {
						total++
					}
				}
				// check down left
				if row < rows && col >= 3 {
					check := string(
						char,
					) + string(
						+lines[row+1][col-1],
					) + string(
						+lines[row+2][col-2],
					) + string(
						+lines[row+3][col-3],
					)
					if check == word {
						total++
					}
				}
				// check up left
				if row >= 3 && col >= 3 {
					check := string(
						char,
					) + string(
						+lines[row-1][col-1],
					) + string(
						+lines[row-2][col-2],
					) + string(
						+lines[row-3][col-3],
					)
					if check == word {
						total++
					}
				}
			}
		}
	}
	return total
}

func partTwo(lines []string) int {
	total := 0

	for row, line := range lines {
		rows := len(lines) - 2

		for col, char := range line {
			columns := len(line) - 2

			if char == 'A' {
				if row < 1 || row > rows || col > columns || col < 1 {
					continue
				}

				check := string(
					lines[row-1][col-1], // check top-left
				) + string(
					lines[row-1][col+1], // check top-right
				) + string(
					lines[row+1][col-1], // check bottom-left
				) + string(
					lines[row+1][col+1], // check bottom-right
				)

				// check if those are two 'M' and two 'S'
				if check == "MSMS" || check == "SMSM" || check == "SSMM" || check == "MMSS" {
					total++
				}

			}
		}
	}
	return total
}
