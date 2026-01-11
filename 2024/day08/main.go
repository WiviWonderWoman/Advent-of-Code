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

/*
......#....# 0,6 + //* 0,11
...#....0... //* 1,3
....#0....#. 2,10 + //* 2,4
..#....0.... //* 3,2
....0....#.. //* 4,9
.#....A..... 5,1 + //* 5,6
...#........ 6,3
#......#.... 7,0 + 7,7
........A...
.........A..
..........#. 10,10
..........#. //* 11,10
*/

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
	fmt.Println(antennaMap)

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
	total := len(lines)

	return total
}
