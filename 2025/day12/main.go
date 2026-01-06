package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"0:",
		"###",
		"##.",
		"##.",
		"",
		"1:",
		"###",
		"##.",
		".##",
		"",
		"2:",
		".##",
		"###",
		"##.",
		"",
		"3:",
		"##.",
		"###",
		"##.",
		"",
		"4:",
		"###",
		"#..",
		"###",
		"",
		"5:",
		"###",
		".#.",
		"###",
		"",
		"4x4: 0 0 0 0 2 0",
		"12x5: 1 0 1 0 2 2",
		"12x5: 1 0 1 0 3 2",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day12")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 12, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 12, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	total := len(lines)

	trees, shapes := parseInput(lines)
	treesWithShapes := getTreesWithShapes(trees, shapes)

	// fmt.Println("trees: ", trees)
	// fmt.Println("shapes: ", shapes)
	fmt.Println("treesWithShapes: ", treesWithShapes)

	return total
}

type Tree struct {
	index    int
	region   []string
	presents [][]string
}

func getTreesWithShapes(trees []string, shapes map[int][]string) []Tree {
	treesWithShapes := make([]Tree, 0, len(trees))
	for i, tree := range trees {
		parts := strings.Split(tree, ": ")
		amounts := utils.StringToIntArr(strings.Split(parts[1], " "))
		regionStr := strings.Split(parts[0], "x")
		regionRows := utils.StringToInt(regionStr[1])
		regionColumns := utils.StringToInt(regionStr[0])

		output := Tree{
			index:    i,
			region:   []string{},
			presents: [][]string{},
		}

		for j, value := range amounts {
			if value == 0 {
				continue
			}

			for range value {
				output.presents = append(output.presents, shapes[j])
			}
		}

		for range regionRows {
			output.region = append(output.region, strings.Repeat(".", regionColumns))
		}

		treesWithShapes = append(treesWithShapes, output)
	}

	return treesWithShapes
}

func parseInput(lines []string) ([]string, map[int][]string) {
	trees := make([]string, 0, len(lines))

	shapes := make(map[int][]string)
	shapeCount := 0
	for i, line := range lines {
		if len(line) > 3 {
			trees = append(trees, line)
		}

		if strings.HasPrefix(line, strconv.Itoa(shapeCount)) {
			shapes[shapeCount] = []string{lines[i+1], lines[i+2], lines[i+3]}
			shapeCount++
		}
	}
	return trees, shapes
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
