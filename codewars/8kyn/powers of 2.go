package main

import (
	"fmt"
)

func main () {
	fmt.Println(PowersOfTwo(5))
}

func PowersOfTwo(n int) []uint64 {

	result := []uint64{}

 for i := 0; i <= n; i++ {
	result[i] = 1 << uint64(i)
 }

 return result
}