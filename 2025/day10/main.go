package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
	"slices"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
		"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
		"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day10")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 10, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 10, Part 2 Answer: ", partTwo(lines))
}

type Machine struct {
	lights  string
	buttons [][]int
	joltage []int
}

func partOne(lines []string) int {
	total := len(lines)
	machines := getMachines(lines)

	for _, machine := range machines {
		min := 0
		// test btns
		for _, button := range machine.buttons {
			// fmt.Println("button: ", button)
			x := pressBtns(button, machine)
			if min == 0 {
				min = x
			}
			min = int(math.Min(float64(min), float64(x)))
		}
		fmt.Println("min: ", min)
	}

	// fmt.Println(machines)
	return total
}

func pressBtns(button []int, machine Machine) int {
	total := 1
	initialIndicators := strings.Split(strings.Repeat(".", len(machine.lights)), "")

	fmt.Println("BUTTON: ", button)
	for i := 0; i < len(machine.buttons); i++ {
		btn := machine.buttons[i]
		total++

		if slices.Compare(btn, button) == 0 {
			continue
		}
		fmt.Println("BTN: ", btn)

		for _, idx := range btn {
			switch initialIndicators[idx] {
			case ".":
				initialIndicators[idx] = "#"
			case "#":
				initialIndicators[idx] = "."
			}
		}

		// fmt.Println("btn: ", btn)
		if strings.Join(initialIndicators, "") == machine.lights {
			fmt.Println("initialIndicators: ", strings.Join(initialIndicators, ""))
			fmt.Println("machine.lights: ", machine.lights)
			break
		}
	}
	return total
}

func getMachines(lines []string) []Machine {
	machines := make([]Machine, 0, len(lines))
	for _, line := range lines {
		firstPart := strings.Split(line, "]")
		secondPart := strings.Split(firstPart[1], "{")

		lights := strings.Split(firstPart[0], "[")[1]

		btnStrSection := strings.Trim(secondPart[0], " ")
		btnStrArrSection := strings.Split(btnStrSection, " ")

		buttons := make([][]int, 0, len(btnStrArrSection))

		for _, str := range btnStrArrSection {
			trimed := strings.Trim(str, "()")
			splitted := strings.Split(trimed, ",")

			button := utils.StringToIntArr(splitted)
			buttons = append(buttons, button)
		}

		joltageSection := strings.Split(secondPart[1], "}")[0]
		joltage := utils.StringToIntArr(strings.Split(joltageSection, ","))

		machine := Machine{
			lights,
			buttons,
			joltage,
		}

		machines = append(machines, machine)
	}

	return machines
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
