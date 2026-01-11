package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
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
	total := 0

	antennaMap := make(map[string][]Position)

	for row := range lines {
		line := strings.Split(lines[row], "")
		for col := range line {
			if line[col] != "." {
				antennaMap[line[col]] = append(antennaMap[line[col]], Position{row, col})
			}
		}
	}

	maxRow := len(lines)
	maxCol := len(lines[0])

	all := mapset.NewSet[Position]()
	for _, frequencies := range antennaMap {
		for i := range frequencies {

			positions := compareAntennas(i, frequencies, maxRow, maxCol)

			for _, pos := range positions {
				if !all.Contains(pos) {
					all.Add(pos)
				}
			}

		}
	}

	total += all.Cardinality()

	return total
}

func compareAntennas(idx int, frequencies []Position, maxRow, maxCol int) []Position {
	compare := frequencies[idx]
	found := make([]Position, 0, len(frequencies)*2)

	for i := 0; i < len(frequencies)-1; i++ {
		if i == idx {
			continue
		}
		current := frequencies[i]

		diffRow := current.row - compare.row
		diffCol := compare.col - current.col

		// above
		above := Position{
			row: compare.row - diffRow,
			col: compare.col + diffCol,
		}

		if above.row >= 0 && above.row < maxRow && above.col >= 0 && above.col < maxCol {
			if !slices.Contains(found, above) {
				found = append(found, above)
			}
		}

		// below
		below := Position{
			row: current.row + diffRow,
			col: current.col - diffCol,
		}

		if below.row >= 0 && below.row < maxRow && below.col >= 0 && below.col < maxCol {
			if !slices.Contains(found, below) {
				found = append(found, below)
			}
		}
	}

	return found
}

func partTwo(lines []string) int {
	total := 0
	antennaMap := make(map[string][]Position)

	for row := range lines {
		line := strings.Split(lines[row], "")
		for col := range line {
			if line[col] != "." {
				antennaMap[line[col]] = append(antennaMap[line[col]], Position{row, col})
			}
		}
	}

	maxRow := len(lines) - 1
	maxCol := len(lines[0]) - 1

	all := mapset.NewSet[Position]()
	for _, frequencies := range antennaMap {
		for i := range frequencies {
			all.Add(frequencies[i])
			if len(frequencies) < 2 {
				continue
			}

			positions := compareAntennasPartTwo(i, frequencies, maxRow, maxCol)
			for _, pos := range positions {
				if !all.Contains(pos) {
					all.Add(pos)
				}
			}

		}
	}

	total += all.Cardinality()

	return total
}

func compareAntennasPartTwo(idx int, frequencies []Position, maxRow, maxCol int) []Position {
	compare := frequencies[idx]
	found := make([]Position, 0, len(frequencies)*2)

	for i := 0; i < len(frequencies); i++ {
		if i == idx {
			continue
		}
		current := frequencies[i]

		diffRow := current.row - compare.row
		diffCol := compare.col - current.col

		// above
		found = append(found, getAbovePositions(compare.row, compare.col, diffRow, diffCol, maxRow, maxCol)...)

		// below
		found = append(found, getBelowPositions(current.row, current.col, diffRow, diffCol, maxRow, maxCol)...)

	}

	return found
}

func getAbovePositions(row, col, diffRow, diffCol, maxRow, maxCol int) []Position {
	found := make([]Position, 0, row)
	initial := Position{row, col}

	for range row {
		above := Position{
			row: initial.row - diffRow,
			col: initial.col + diffCol,
		}

		if above.row < 0 || above.col < 0 || above.row > maxRow || above.col > maxCol {
			break
		}

		found = append(found, above)
		initial = above
	}
	return found
}

func getBelowPositions(row, col, diffRow, diffCol, maxRow, maxCol int) []Position {
	found := make([]Position, 0, maxRow-row)
	initial := Position{row, col}

	for range maxRow {
		below := Position{
			row: initial.row + diffRow,
			col: initial.col - diffCol,
		}

		if below.row < 0 || below.col < 0 || below.row > maxRow || below.col > maxCol {
			break
		}

		found = append(found, below)
		initial = below
	}
	return found
}
