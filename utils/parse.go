package utils

import "strconv"

func StringToFloat64(input string) float64 {
	float, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}

	return float
}

func StringToInt(input string) int {
	integer, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return integer
}

func StringToIntArr(input []string) []int {
	output := make([]int, 0, len(input))
	for _, in := range input {
		integer, err := strconv.Atoi(in)
		if err != nil {
			panic(err)
		}
		output = append(output, integer)
	}

	return output
}
