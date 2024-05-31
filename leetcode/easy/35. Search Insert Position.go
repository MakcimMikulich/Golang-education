package main

import "fmt"

// func main() {

// 	fmt.Println(searchInsert([]int{1, 2, 3, 5, 6}, 5))

// }

func searchInsert(nums []int, target int) int {

	start := 0
	end := len(nums)

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	if nums[0] > target {
		return 0
	}

	for i := 1; i < 5; i++ {
		middle := (end + start) / 2

		if target == nums[middle] {
			return middle
		}

		if target > nums[middle] && target < nums[middle+1] {
			return middle + 1
		}

		if target > nums[middle] {
			start = middle + 1
		}

		if target < nums[middle] {
			end = middle - 1
		}
		fmt.Println(middle, start, end)
	}
	return 5
}
