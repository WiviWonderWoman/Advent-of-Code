package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines, err := utils.ReadInput("day05")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 05, Part 1 Answer: ", partOne(lines))
	fmt.Println("Day 05, Part 2 Answer: ", partTwo(lines))
}

type Range struct {
	start float64
	end   float64
}

func partOne(lines []string) int {
	total := 0
	idx := slices.Index(lines, "")

	ranges := getRanges(lines[0:idx])

	inventory := getInventory(lines[idx:])

	for _, id := range inventory {
		if checkFresh(ranges, id) {
			total++
		}
	}

	return total
}

func getRanges(lines []string) []Range {
	ranges := []Range{}
	for _, line := range lines {
		split := strings.Split(line, "-")
		start, _ := strconv.ParseFloat(split[0], 64)
		end, _ := strconv.ParseFloat(split[1], 64)

		ranges = append(ranges, Range{start, end})
	}

	return ranges
}

func getInventory(lines []string) []int {
	ids := []int{}

	for _, line := range lines {
		if line != "" {
			id := utils.StringToInt(line)
			ids = append(ids, id)
		}
	}

	return ids
}

func checkFresh(ranges []Range, id int) bool {
	for _, r := range ranges {
		if id >= int(r.start) && id <= int(r.end) {
			return true
		}
	}
	return false
}

func partTwo(lines []string) int {
	total := 0

	idx := slices.Index(lines, "")
	ranges := getRanges(lines[0:idx])

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.start-b.start == 0 {
			return int(a.end - b.end)
		}
		return int(a.start - b.start)
	})

	reducedRanges := getCombinedRanges(ranges)

	for _, r := range reducedRanges {
		total += int(r.end) - int(r.start)
		total++
	}

	return total
}

func getCombinedRanges(input []Range) []Range {
	output := []Range{}

	for _, r := range input {
		found := false
		start := r.start
		end := r.end

		for i := 0; i < len(output); i++ {
			combined := output[i]

			if start >= combined.start && end <= combined.end {
				found = true
				break
			} else if start <= combined.end && start >= combined.start {
				start = combined.start
				end = math.Max(end, combined.end)

				output[i] = Range{start, end}
				found = true
				i = 0
			}
		}

		if !found {
			output = append(output, Range{start, end})
		}
	}

	return output
}
