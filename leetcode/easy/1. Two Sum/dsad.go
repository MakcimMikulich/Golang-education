package main

import (
	"fmt"
)

func main() {

	fmt.Println(twoSum([]int{7, 2, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {

	tmp := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			tmp = nums[i] + nums[j]

			if tmp == target {
				return []int{nums[i], nums[j]}
			}
		}
	}

	return []int{}
}
