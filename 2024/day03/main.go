package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"
)

func GetTestInput() string {
	return "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
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
	line := GetTestInput()

	fmt.Println("Day 03, Part 1 Answer: ", partOne(line))
	fmt.Println("Day 03, Part 2 Answer: ", partTwo(line))
}

func partOne(line string) int {
	total := 0
	re := regexp.MustCompile(`(?i)mul\([0-9]{1,3},[0-9]{1,3}\)`)
	bites := re.FindAll([]byte(line), -1)

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

func partTwo(line string) int {
	total := len(line)
	// get index for all do() and don't()
	// only mul(X,Y) after do() befor don't() are valid...
	// starting with don't()

	return total
}
