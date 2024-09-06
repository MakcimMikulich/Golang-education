package main

import (
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	result := ""

	for i, v := range numbers {
		strNum := strconv.Itoa(int(v))
		switch i {
		case 0:
			result += "(" + strNum
		case 2:
			result += strNum + ")"
		case 3:
			result += " " + strNum
		case 6:
			result += "-" + strNum
		default:
			result += strNum
		}

	}
	return result
}
