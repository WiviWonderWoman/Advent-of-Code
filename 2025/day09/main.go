package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

func GetTestInput() []string {
	return []string{"7,1", "11,1", "11,7", "9,7", "9,5", "2,5", "2,3", "7,3"}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day09")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 09, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 09, Part 2 Answer: ", partTwo(lines))
}

type Corner struct {
	row    float64
	column float64
}

func partOne(lines []string) int {
	total := 0

	corners := getCorners(lines)

	combos := make([][]Corner, 0, len(corners))
	for _, corner := range corners {
		res := findPossibleCorners(corner, corners)
		if len(res) > 1 {
			combos = append(combos, findPossibleCorners(corner, corners))
		}
	}

	for _, combo := range combos {
		startCorner := combo[0]

		for i := 1; i < len(combo); i++ {
			compareCorner := combo[i]
			x := compareCorner.row - startCorner.row + 1
			y := compareCorner.column - startCorner.column + 1
			total = int(math.Max((x * y), float64(total)))

		}
	}
	// fmt.Println("1 combos: ", combos)
	return total
}

func findPossibleCorners(comp Corner, corners []Corner) []Corner {
	combos := make([]Corner, 0, len(corners))

	for _, corner := range corners {
		if corner.row == comp.row && corner.column == comp.column {
			continue
		}

		if corner.row > comp.row && corner.column > comp.column {
			if len(combos) < 1 {
				combos = append(combos, comp)
			}

			combos = append(combos, corner)
		}
	}

	return combos
}

func getCorners(lines []string) []Corner {
	corners := make([]Corner, 0, len(lines))

	for _, line := range lines {
		corner := Corner{
			row:    utils.StringToFloat64(strings.Split(line, ",")[1]),
			column: utils.StringToFloat64(strings.Split(line, ",")[0]),
		}

		corners = append(corners, corner)
	}

	slices.SortFunc(corners, func(a, b Corner) int {
		result := int(a.row - b.row)
		if result == 0 {
			return int(a.column - b.column)
		}
		return result
	})

	return corners
}

/*
	0 1 2 3 4 5 6 7 8 9 1011

0: . . . . . . . . . . . . . .
1: . . . . . . . 7 8 9 1011. . 1, 7 - 11
2: . . . . . . . 7 8 9 1011. . 2, 7 - 11
3: . . 2 3 4 5 6 7 8 9 1011. . 3, 2 - 11
4: . . 2 3 4 5 6 7 8 9 1011. . 4, 2 - 11
5: . . 2 3 4 5 6 7 8 9 1011. . 5, 2 - 11
6: . . . . . . . . . 9 1011. . 6, 9 - 11
7: . . . . . . . . . 9 1011. . 7, 9 - 11
8: . . . . . . . . . . . . . .
*/
func partTwo(lines []string) int {
	total := 0

	corners := getCorners(lines)
	// fmt.Println("corners: ", corners)

	edges := getEdges(corners)
	// fmt.Println("edges: ", edges)

	filled := getEdges(edges)
	fmt.Println("filled: ", filled)
	// total = len(filled)

	combos := make([][]Corner, 0, len(corners))
	for _, corner := range corners {
		res := findPossibleFilledCorners(corner, corners, filled)
		if len(res) > 1 {
			combos = append(combos, findPossibleCorners(corner, corners))
		}
	}
	fmt.Println("combos: ", combos)

	for _, combo := range combos {
		startCorner := combo[0]

		for i := 1; i < len(combo); i++ {
			compareCorner := combo[i]
			x := compareCorner.row - startCorner.row + 1
			y := compareCorner.column - startCorner.column + 1
			total = int(math.Max((x * y), float64(total)))

		}
	}

	return total
}

func findPossibleFilledCorners(comp Corner, corners []Corner, filled []Corner) []Corner {
	combos := make([]Corner, 0, len(corners))

	for _, corner := range corners {
		if corner.row == comp.row && corner.column == comp.column {
			continue
		}

		if corner.row > comp.row && corner.column > comp.column {
			for i := comp.row; i < corner.row; i++ {
				if slices.Contains(filled, Corner{row: i, column: corner.column}) &&
					slices.Contains(filled, Corner{row: i, column: comp.column}) {

					if len(combos) < 1 {
						combos = append(combos, comp)
					}

					combos = append(combos, corner)
				}
			}
		}

	}

	return combos
}

func getEdges(corners []Corner) []Corner {
	edges := make([]Corner, 0, len(corners)*10)
	for _, corner := range corners {
		current := corner
		for _, c := range corners {
			if current == c {
				continue
			}

			if current.row == c.row && c.column > current.column {
				cols := NewSlice(current.column, c.column, 1)

				for _, col := range cols {
					edges = append(edges, Corner{row: c.row, column: col})
				}
			}

			if current.row == c.row && c.column < current.column {
				cols := NewSlice(c.column, current.column, 1)

				for _, col := range cols {
					edges = append(edges, Corner{row: c.row, column: col})
				}
			}

			if current.column == c.column && c.row > current.row {
				rows := NewSlice(current.row, c.row, 1)

				for _, row := range rows {
					edges = append(edges, Corner{row: row, column: c.column})
				}
			}

			if current.column == c.column && c.row < current.row {
				rows := NewSlice(c.row, current.row, 1)

				for _, row := range rows {
					edges = append(edges, Corner{row: row, column: c.column})
				}
			}

		}
	}

	slices.SortFunc(edges, func(a, b Corner) int {
		resultRow := int(a.row - b.row)
		resultColumn := int(a.column - b.column)

		if resultRow == 0 {
			return resultColumn
		}
		return resultRow
	})

	edges = slices.Compact(edges)

	return edges
}

func NewSlice(start, end float64, step int) []float64 {
	if step <= 0 || end < start {
		return []float64{}
	}
	s := make([]float64, 0, 1+int((end-start))/step)
	for start <= end {
		s = append(s, start)
		start += float64(step)
	}
	return s
}
