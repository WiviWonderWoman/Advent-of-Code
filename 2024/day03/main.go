package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"
)

func GetTestInputPartOne() string {
	return "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
}

func GetTestInputPartTwo() string {
	return "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
}

func GetInput() string {
	lines, err := utils.ReadInputAsString("day03")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// line:=GetInput()
	linePartOne := GetTestInputPartOne()
	linePartTwo := GetTestInputPartTwo()

	fmt.Println("Day 03, Part 1 Answer: ", partOne(linePartOne))
	fmt.Println("Day 03, Part 2 Answer: ", partTwo(linePartTwo))
}

func partOne(line string) int {
	total := checkMemory(line)

	return total
}

func partTwo(line string) int {
	total := 0

	// split at don't()
	parts := strings.Split(line, "don't()")

	for i, current := range parts {
		// starting with do()
		if i == 0 {
			total += checkMemory(current)
			continue
		}

		if i != 0 {
			// split at do()
			splitted := strings.Split(current, "do()")
			if len(splitted) > 1 {
				for j := 1; j < len(splitted); j++ {
					total += checkMemory(splitted[j])
				}
			}
		}
	}

	return total
}

func checkMemory(input string) int {
	total := 0
	re := regexp.MustCompile(`(?i)mul\([0-9]{1,3},[0-9]{1,3}\)`)
	bites := re.FindAll([]byte(input), -1)

	validMemory := []string{}
	for _, byteSlice := range bites {
		validMemory = append(validMemory, string(byteSlice))
	}

	for _, raw := range validMemory {
		valid := strings.Trim(raw, "mul()")
		numbers := utils.StringToIntArr(strings.Split(valid, ","))

		total += numbers[0] * numbers[1]
	}

	return total
}
