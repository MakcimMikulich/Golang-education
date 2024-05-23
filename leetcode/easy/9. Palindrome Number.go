package main

import (
	"strconv"
)

// func main() {

// 	fmt.Println(isPalindrome(923229))
// }

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
