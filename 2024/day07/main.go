package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"math/bits"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day07")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	lines := GetInput()

	fmt.Println("Day 07, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 07, Part 2 Answer: ", partTwo(lines))
}

type Equation struct {
	answer  int
	numbers []int
}

// NOT pretty bur it works
func partOne(lines []string) int {
	total := 0

	equations := getEquations(lines)

	for _, equation := range equations {
		// add
		addAll := utils.AddNumbers(equation.numbers)
		// multiply
		multipyAll := utils.MultiplyNumbers(equation.numbers)

		if addAll == equation.answer || multipyAll == equation.answer {
			fmt.Println("SAME CORRECT:", equation)
			total += equation.answer
			continue
		}

		if len(equation.numbers) > 2 {
			// possible combinations
			possibilities := int(math.Pow(2, float64(len(equation.numbers)-1))) - 1

			// initial value
			current := equation.numbers[0]

			binaryArr := make([]string, 0, possibilities)
			for possible := range possibilities {
				// binary representation of number
				binaryStr := fmt.Sprintf("%.16b", possible)

				// cut excessive leading zeros
				if len(binaryStr) > len(equation.numbers)-1 {
					rm := bits.LeadingZeros16(uint16(possibilities))
					binaryStr = strings.Join(strings.Split(binaryStr, "")[rm:], "")
				}

				binaryArr = append(binaryArr, binaryStr)
			}

			for _, binaryStr := range binaryArr {
				for j, char := range binaryStr {
					switch string(char) {
					case "0":
						// if "0" add
						current = current + equation.numbers[j+1]
					case "1":
						// if "*" multiply
						current = current * equation.numbers[j+1]
					}
				}

				// YAY! stop this nonsense
				if current == equation.answer {
					break
				}

				// reset
				current = equation.numbers[0]
			}

			if current == equation.answer {
				fmt.Println("CORRECT:", equation)
				total += equation.answer
				continue
			}
		}
	}

	return total
}

func getEquations(lines []string) []Equation {
	equations := make([]Equation, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		equations = append(
			equations,
			Equation{
				answer:  utils.StringToInt(parts[0]),
				numbers: utils.StringToIntArr(strings.Fields(parts[1])),
			},
		)
	}

	return equations
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
