package main

import (
	"adventofcode2025/utils"
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
	// fmt.Println("diskMap: ", diskMap)

	decompressed := []string{}
	fileID := 0
	emptySpaces := 0
	for i, value := range diskMap {
		if i%2 == 0 {
			// even indecies are files => multiply "{index}" by the value
			file := strings.Repeat(strconv.Itoa(fileID), value)
			decompressed = append(decompressed, strings.Split(file, "")...)

			// increase fileID
			fileID++
		} else {
			// odd indecies are spaces => multply "." by the value
			space := strings.Repeat(".", value)
			decompressed = append(decompressed, strings.Split(space, "")...)

			// keep track of empty spaces
			emptySpaces += value
		}
	}

	// fmt.Println("emptySpaces: ", emptySpaces)
	// fmt.Println("decompressed: ", decompressed)

	// repeat * len(decompressed)
	for i := 0; i < len(decompressed); i++ {

		// find index first "."
		idx := slices.Index(decompressed, ".")

		// "pop" last
		popped, remaining := decompressed[len(decompressed)-1], decompressed[:len(decompressed)-1]
		// skip free space
		if popped == "." {
			fmt.Println(i, "EMPTY END")

			decompressed = remaining
			continue
		}

		// insert/replace "poped" into index
		decompressed = append(remaining[:idx], append([]string{popped}, remaining[idx+1:]...)...)

		// break if we are done
		if !slices.Contains(decompressed, ".") {
			fmt.Println("BREAK")
			break
		}

	}

	compacted := utils.StringToIntArr(decompressed)

	// loop thru compacted and add: {index * value} to total
	for i, value := range compacted {
		res := i * value
		total += res
	}

	return total
}

func partTwo(line string) int {
	total := len(line)

	return total
}
