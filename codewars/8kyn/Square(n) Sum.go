package main

func SquareSum(numbers []int) int {
	result := 0

	for _, value := range numbers {
		result += value * value
	}

	return result
}
