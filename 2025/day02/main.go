package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetTestInput() []string {
	raw := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	input := strings.Split(raw, ",")
	return input
}

func GetInput() []string {
	raw, err := utils.ReadInputAsString("day02")
	if err != nil {
		panic(err)
	}

	input := strings.Split(raw, ",")
	return input
}

func main() {
	lines := GetInput()
	// lines := GetTestInput()

	fmt.Println("Day 02, Part 1 Answer: ", dayTwo(lines, 1))
	fmt.Println("Day 02, Part 2 Answer: ", dayTwo(lines, 2))
}

func dayTwo(lines []string, part int) int {
	total := 0
	scopes := getScopes(lines)
	maxLength := getMaxLength(lines)

	primeNumbers := make([]int, 0, maxLength)
	if part == 2 {
		for number := 2; number < maxLength; number++ {
			if isPrime(number) {
				primeNumbers = append(primeNumbers, number)
			}
		}
	} else {
		primeNumbers = []int{2}
	}

	for _, scope := range scopes {
		total += checkNumbersInScope(scope, primeNumbers)
	}

	return total
}

func checkNumbersInScope(scope Scope, primeNumbers []int) int {
	total := 0
	for i := scope.from; i < scope.to+1; i++ {
		total += divideByPrimeNumbers(strconv.Itoa(i), primeNumbers)
	}

	return total
}

func divideByPrimeNumbers(input string, primeNumbers []int) int {
	total := 0

	for _, prime := range primeNumbers {
		if prime > len(input) {
			break
		}

		if len(input)%prime == 0 {
			res := findRepeatedSequence(input, prime)
			if res > 0 {
				total += res
				return total
			}
		}
	}

	return total
}

func findRepeatedSequence(input string, divisor int) int {
	total := 0
	quotient := len(input) / divisor

	initial := input[0:quotient]
	repeats := strings.Count(input, initial)

	if repeats == divisor {
		add := utils.StringToInt(input)
		total += add
	}

	return total
}

type Scope struct {
	from int
	to   int
}

func getScopes(input []string) []Scope {
	scopes := make([]Scope, 0, len(input))

	for _, str := range input {
		split := strings.Split(str, "-")
		scopes = append(scopes, Scope{
			from: int(utils.StringToFloat64(split[0])),
			to:   int(utils.StringToFloat64(split[1])),
		})
	}

	return scopes
}

func getMaxLength(input []string) int {
	max := 0
	for _, str := range input {
		split := strings.Split(str, "-")
		max = int(math.Max(float64(max), math.Max(float64(len(split[0])), float64(len(split[1])))))
	}
	return max
}

func isPrime(input int) bool {
	for i := 2; i < input; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}
