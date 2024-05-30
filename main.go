package main

import "fmt"

func main() {

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))

}

func searchInsert(nums []int, target int) int {

	start := 0
	end := len(nums)

	for {
		middle := (end - start) / 2

		if target == nums[middle] {
			return middle
		}

		if target > nums[middle] {
			start = middle
		}

		if target < nums[middle] {
			end = middle
		}
	}

}
