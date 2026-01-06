package main

import (
	"adventofcode/utils"
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func GetTestInput() []string {
	return []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}
}

func GetInput() []string {
	lines, err := utils.ReadInput("day08")
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	// lines:=GetInput()
	lines := GetTestInput()

	fmt.Println("Day 08, Part 1 Answer: ", partOne(lines, 1000))
	fmt.Println("Day 08, Part 2 Answer: ", partTwo(lines))
}

type EuclideanPoint []float64

type Box struct {
	Str            string
	Idx            int
	EuclideanPoint EuclideanPoint
}

func (ep EuclideanPoint) EuclideanDistance(p EuclideanPoint) float64 {
	var total float64 = 0

	for i, x_i := range ep {
		// using Abs since the value could be negative but we require the magnitude
		diff := math.Abs(x_i - p[i])
		total += diff * diff
	}

	return math.Sqrt(total)
}

func partOne(lines []string, laps int) int {
	total := 0
	boxes := getBoxes(lines)
	min := float64(0)

	connections := [][]int{}
	var connection []int

	idx := 0
	for idx <= laps {
		min, connection = getMinAndConnection(boxes, min)
		connection = slices.Compact(connection)
		connections = append(connections, connection)
		idx = len(connections)

	}

	newConnections := [][]int{}
	for i, c := range connections {
		if i == 0 {
			newConnections = checkConnections(c, connections)
		}
		checkConnections(c, newConnections)
	}
	connections = newConnections

	slices.SortFunc(connections, func(a, b []int) int {
		return cmp.Compare(len(b), len(a))
	})

	total = len(connections[0]) * len(connections[1]) * len(connections[2])

	return total
}

func checkConnections(comp []int, connections [][]int) [][]int {
	newConnections := connections

	for i, con := range connections {
		newCon := con
		if slices.Contains(con, comp[0]) && slices.Contains(con, comp[1]) {
			// fmt.Println("BOTH comp: ", comp, " con: ", con)
			return newConnections
		} else if slices.Contains(con, comp[0]) && !slices.Contains(con, comp[1]) {
			// fmt.Println("FIRST comp: ", comp, " con: ", con)
			newConnections[i] = append(newCon, comp[1])
			return newConnections
		} else if slices.Contains(con, comp[1]) && !slices.Contains(con, comp[0]) {
			// fmt.Println("LAST comp: ", comp, " con: ", con)
			newConnections[i] = append(newCon, comp[0])
			return newConnections
		}
	}

	return newConnections
}

func getMinAndConnection(boxes []Box, limit float64) (float64, []int) {
	min := float64(0)

	connection := []int{}
	for _, box := range boxes {
		current, con := compareBoxes(box, boxes, min, limit)
		if min == 0 || current < min {
			min = current
			connection = con
		}
	}

	return min, connection
}

func compareBoxes(comp Box, boxes []Box, min float64, limit float64) (float64, []int) {
	low := min
	connection := []int{}
	for _, box := range boxes {
		if box.Idx == comp.Idx {
			continue
		}

		current := comp.EuclideanPoint.EuclideanDistance(box.EuclideanPoint)

		if low == 0 || current < low && current > limit {
			low = current
			connection = []int{box.Idx, comp.Idx}
		}
	}
	slices.Sort(connection)
	return low, connection
}

func getBoxes(lines []string) []Box {
	boxes := []Box{}
	for i, line := range lines {
		ep := EuclideanPoint{}
		box := Box{Str: line, Idx: i}

		splitted := strings.Split(line, ",")
		for _, str := range splitted {
			number, _ := strconv.ParseFloat(str, 64)
			ep = append(ep, number)
		}

		box.EuclideanPoint = ep
		boxes = append(boxes, box)
	}

	return boxes
}

func partTwo(lines []string) int {
	total := len(lines)

	return total
}
