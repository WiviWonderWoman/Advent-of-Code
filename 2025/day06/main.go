package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

func GetTestInput() []string {
	return []string{"123 328 51 64", "45 64 387 23", "6 98 215 314", "* + * +"}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day06")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 06, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 06, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	sortedProblems := [][]string{}

	sortedProblems = append(sortedProblems, make([]string, 0, len(lines[0])))
	for _, raw := range lines {
		line := strings.Fields(raw)
		for i, str := range line {
			if len(sortedProblems) <= i {
				sortedProblems = append(sortedProblems, make([]string, 0, len(lines[0])))
			}
			sortedProblems[i] = append(sortedProblems[i], str)
		}
	}
	// fmt.Println("sortedProblems: ", sortedProblems)
	return calculateProblems(sortedProblems)
}

func calculateProblems(sortedProblems [][]string) int {
	total := 0

	for _, problem := range sortedProblems {
		method, numbers := problem[len(problem)-1], problem[:len(problem)-1]

		switch method {
		case "*":
			total += utils.MultiplyNumbers(utils.StringToIntArr(numbers))
		case "+":
			total += utils.AddNumbers(utils.StringToIntArr(numbers))
		}
	}

	return total
}

func partTwo(lines []string) int {
	splitAndCount := [][]string{}
	max := 0

	for _, raw := range lines {
		line := strings.Fields(raw)
		sanitized := []string{}

		for _, str := range line {
			max = int(math.Max(float64(max), float64(len(str))))
			sanitized = append(sanitized, str)
		}
		splitAndCount = append(splitAndCount, sanitized)
	}

	filled := fillAndAlign(splitAndCount, max)
	preSorted := sortFilledAndAligned(filled, len(lines[0]))

	sorted := [][]string{}
	for _, problem := range preSorted {
		sorted = append(sorted, getProblem(problem))
	}

	return calculateProblems(sorted)
}

func fillAndAlign(input [][]string, max int) [][]string {
	output := [][]string{}

	for i, line := range input {
		for j, rawStr := range line {
			filler := " "
			str := strings.Trim(rawStr, " ")

			if strings.Contains(str, "+") || strings.Contains(str, "*") {
				filler = str
			}

			if len(str) < max {
				filler = strings.Repeat(filler, max-len(str))

				if j%2 == 0 || j == 0 {
					input[i][j] = filler + str
				} else {
					input[i][j] = str + filler
				}
			}
		}
		output = append(output, input[i])
	}

	return output
}

func sortFilledAndAligned(input [][]string, length int) [][]string {
	output := [][]string{}
	output = append(output, make([]string, 0, length))
	for _, line := range input {
		for j, str := range line {
			if len(output) <= j {
				output = append(output, make([]string, 0, length))
			}

			output[j] = append(output[j], str)
		}
	}
	return output
}

func getProblem(input []string) []string {
	splitted := splittIntoCephalopodProblem(input)

	return joinCephalopodNumbers(splitted)
}

func splittIntoCephalopodProblem(input []string) [][]string {
	output := [][]string{}
	output = append(output, make([]string, 0, len(input)))
	for _, line := range input {
		for j, str := range line {
			if len(output) <= j {
				output = append(output, make([]string, 0, len(line)))
			}

			output[j] = append(output[j], string(str))
		}
	}

	return output
}

func joinCephalopodNumbers(input [][]string) []string {
	output := []string{}
	for i, rawLine := range input {
		line := strings.Fields(strings.Join(rawLine, " "))
		method, numbers := line[len(line)-1], line[:len(line)-1]

		numbers = strings.Fields(strings.Join(numbers, ""))
		if len(numbers) > 0 {
			output = append(output, strings.Join(numbers, ""))
		}

		if i == len(input)-1 {
			output = append(output, method)
		}
	}

	return output
}
