package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day03")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 03, Part 1 Answer: ", dayThree(lines, 1))
	fmt.Println("Day 03, Part 2 Answer: ", dayThree(lines, 11))
}

func dayThree(lines []string, turns int) int {
	total := 0

	for _, line := range lines {

		bank := []int{}
		for _, str := range strings.Split(line, "") {
			bank = append(bank, utils.StringToInt(str))
		}

		total += findHigest(turns, bank, 0, 0)
	}

	return total
}

func findHigest(
	turn int,
	bank []int,
	index int,
	value int,
) int {
	combo := ""
	if value != 0 {
		combo = strconv.Itoa(int(value))
	}

	current := 0
	for i := index; i < len(bank)-turn; i++ {
		if bank[i] > current {
			index = i
			current = int(math.Max(float64(current), float64(bank[i])))
		}
	}

	combo += strconv.Itoa(int(current))

	if turn > 0 {
		return findHigest(turn-1, bank, index+1, utils.StringToInt(combo))
	}

	return int(utils.StringToFloat64(combo))
}
