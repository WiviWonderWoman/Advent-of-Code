package main

import (
	"adventofcode2025/utils"
	"fmt"
)

func GetTestInput() []string {
	return []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day01")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	lines := GetInput()

	fmt.Println("Day 1, Part 1 Answer:", partOne(lines))
	fmt.Println("Day 1, Part 2 Answer:", partTwo(lines))
}

func partOne(lines []string) int {
	value := 50
	zeroes := 0

	for _, line := range lines {
		direction := line[0]
		number := utils.StringToInt((line[1:]))

		switch string(direction) {
		case "L":
			value -= number
		case "R":
			value += number

		}

		if value < 0 {
			value = 100 + (value % 100)
		}

		if value > 99 {
			value = value % 100
		}

		if value == 0 {
			zeroes++
		}
	}
	return zeroes
}

type Dial struct {
	value    int
	zeroes   int
	rotation int
}

func partTwo(lines []string) int {
	current := Dial{
		value:    50,
		zeroes:   0,
		rotation: 0,
	}

	for _, line := range lines {
		direction := line[0]
		current.rotation = utils.StringToInt(line[1:])

		switch string(direction) {
		case "L":
			current = decreaseValue(current)
		case "R":
			current = increaseValue(current)

		}
	}

	return current.zeroes
}

func increaseValue(input Dial) Dial {
	for range input.rotation {
		if input.value == 0 {
			input.zeroes++
		}

		if input.value == 99 {
			input.value = 0
		} else {
			input.value++
		}
	}
	return input
}

func decreaseValue(input Dial) Dial {
	for range input.rotation {
		if input.value == 0 {
			input.zeroes++
		}

		if input.value == 0 {
			input.value = 99
		} else {
			input.value--
		}
	}

	return input
}
