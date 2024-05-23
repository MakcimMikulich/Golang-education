package main

import (
	"fmt"
)

func main() {

	fmt.Println(SequenceSum(2, 6, 2))
}

func SequenceSum(start, end, step int) int {
	result := 0

	for i := start; i <= end; i += step {
		result += i
	}

	return result
}
