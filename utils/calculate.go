package utils

func MultiplyNumbers(input []int) int {
	product := 0
	for i := 0; i < len(input); i++ {
		if i == 0 {
			product = input[i]
			continue
		}
		product = product * input[i]
	}

	return product
}

func AddNumbers(input []int) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum = sum + input[i]
	}

	return sum
}
