package main

import (
	"adventofcode2025/utils"
	"fmt"
	"slices"
	"strings"
)

func GetPartOneTestInput() []string {
	return []string{
		"aaa: you hhh",
		"you: bbb ccc",
		"bbb: ddd eee",
		"ccc: ddd eee fff",
		"ddd: ggg",
		"eee: out",
		"fff: out",
		"ggg: out",
		"hhh: ccc fff iii",
		"iii: out",
	}
}

func GetPartTwoTestInput() []string {
	return []string{
		"svr: aaa bbb",
		"aaa: fft",
		"fft: ccc",
		"bbb: tty",
		"tty: ccc",
		"ccc: ddd eee",
		"ddd: hub",
		"hub: fff",
		"eee: dac",
		"dac: fff",
		"fff: ggg hhh",
		"ggg: out",
		"hhh: out",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day11")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines := GetInput()
	// lines := GetPartOneTestInput()
	lines := GetPartTwoTestInput()

	// fmt.Println("Day 11, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 11, Part 2 Answer: ", partTwo(lines))
}

func partOne(lines []string) int {
	total := 0

	youIdx := findDeviceIdx(lines, "you: ")

	total = youToOut(youIdx, lines)
	return total
}

func youToOut(index int, lines []string) int {
	if index == -1 {
		return 0
	}
	total := 0

	line := strings.Split(lines[index], ": ")
	outputs := strings.Fields(line[1])

	for _, output := range outputs {
		if output == "out" {
			total++
			continue
		}
		idx := findDeviceIdx(lines, output+":")
		total += youToOut(idx, lines)
	}

	return total
}

func findDeviceIdx(lines []string, input string) int {
	return slices.IndexFunc(lines, func(line string) bool {
		return strings.HasPrefix(line, input)
	})
}

// * test input - OK
// ! real input - infinite loop
func partTwo(lines []string) int {
	total := 0

	svrIdx := findDeviceIdx(lines, "svr: ")
	fmt.Println("svrIdx: ", svrIdx)

	total = svrToOut(svrIdx, lines, false, false)
	return total
}

func svrToOut(index int, lines []string, dacPassed bool, fftPassed bool) int {
	if index == -1 {
		return 0
	}
	total := 0

	line := strings.Split(lines[index], ": ")
	outputs := strings.Fields(line[1])

	dac := dacPassed
	fft := fftPassed

	device := strings.TrimSpace(line[0])

	switch device {
	case "dac":
		fmt.Println("DAC: ", index, lines[index])
		dac = true
	case "fft":
		fmt.Println("FFT: ", index, lines[index])
		fft = true
	}

	isValid := dac && fft

	for _, output := range outputs {
		if output == "out" {
			fmt.Println("VALID? : ", isValid)
			if isValid {
				total++
				dac = false
				fft = false
				continue
			}
		}
		idx := findDeviceIdx(lines, output+":")
		total += svrToOut(idx, lines, dac, fft)
	}

	return total
}
