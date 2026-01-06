package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
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
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 07, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 07, Part 2 Answer: ", partTwo(lines))
}

type Equation struct {
	answer  int
	numbers []int
}

func partOne(lines []string) int {
	total := 0

	equations := getEquations(lines)

	for _, equation := range equations {
		numbers := equation.numbers
		if len(numbers) == 1 {
			fmt.Println(numbers)
		}
		// add
		addAll := utils.AddNumbers(numbers)
		// multiply
		multipyAll := utils.MultiplyNumbers(numbers)

		if addAll == equation.answer || multipyAll == equation.answer {
			fmt.Println("CORRECT:", equation)
			total += equation.answer
			continue
		}

		sum := 0
		if len(numbers) > 2 {
			indencies := make([]int, 0, len(numbers))
			for i := range numbers {
				indencies = []int{i}

				methods := getMethods(numbers, indencies)

				current := calculateProblem(methods, numbers)

				if current == equation.answer {
					fmt.Println("CORRECT:", equation)
					sum = equation.answer
					break
				}
			}

			if sum != 0 {
				total += sum
				continue
			}
		}

	}

	return total
}

func calculateProblem(methods []string, numbers []int) int {
	current := numbers[0]
	for i, method := range methods {
		if method == "*" {
			current = current * numbers[i+1]
		}

		if method == "+" {
			current = current + numbers[i+1]
		}
	}

	return current
}

func getMethods(numbers []int, idx []int) []string {
	problem := make([]string, 0, len(numbers)-1)
	for i := range numbers {
		if i == len(numbers)-1 {
			break
		}
		if slices.Contains(idx, i) { // i == idx {
			problem = append(problem, "*")
		} else {
			problem = append(problem, "+")
		}
	}

	return problem
}

// func getMethods(numbers []int, idx int) []string {
// 	problem := make([]string, 0, len(numbers)-1)
// 	for i := range numbers {
// 		if i == len(numbers)-1 {
// 			break
// 		}
// 		if i == idx {
// 			problem = append(problem, "*")
// 		} else {
// 			problem = append(problem, "+")
// 		}
// 	}

// 	return problem
// }

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
