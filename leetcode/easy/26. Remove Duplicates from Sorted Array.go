package main

// func main() {

// 	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
// }

func removeDuplicates(nums []int) int {

	index := 1

	for i, v := range nums {
		if i == 0 {
			continue
		}

		if v > nums[i-1] {
			nums[index] = v
			index++
		}
	}

	return index
}
