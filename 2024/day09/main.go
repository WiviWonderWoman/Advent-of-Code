package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func GetTestInput() string {
	return "2333133121414131402"
}

func GetInput() string {
	line, err := utils.ReadInputAsString("day09")
	if err != nil {
		panic(err)
	}

	return line
}

func main() {
	// lines:=GetInput()
	line := GetTestInput()

	fmt.Println("Day 09, Part 1 Answer: ", partOne(line))
	fmt.Println("Day 09, Part 2 Answer: ", partTwo(line))
}

func partOne(line string) int {
	total := 0

	diskMap := utils.StringToIntArr(strings.Split(line, ""))

	decompressed := []string{}
	fileID := 0

	for i, value := range diskMap {
		if i%2 == 0 {
			// even indecies are files => multiply "{fileID}" by the value
			for range value {
				decompressed = append(decompressed, strconv.Itoa(fileID))
			}

			fileID++
		} else {
			// odd indecies are spaces => multiply "." by the value
			space := strings.Repeat(".", value)
			decompressed = append(decompressed, strings.Split(space, "")...)
		}
	}

	compacted := utils.StringToIntArr(compactFiles(decompressed))

	// loop thru compacted and add: {index * value} to total
	for i, value := range compacted {
		total += i * value
	}

	return total
}

func compactFiles(input []string) []string {
	output := input

	for i := 0; i < len(output); i++ {
		// break if we are done
		if !slices.Contains(output, ".") {
			return output
		}
		// find index first "."
		idx := slices.Index(output, ".")

		// "pop" last
		popped, remaining := output[len(output)-1], output[:len(output)-1]
		// skip free space
		if popped == "." {
			output = remaining
			continue
		}

		// insert/replace "." with number
		output = append(remaining[:idx], append([]string{popped}, remaining[idx+1:]...)...)

	}

	return output
}

func partTwo(line string) int {
	total := len(line)

	return total
}
