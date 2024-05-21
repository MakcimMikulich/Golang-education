package main

import "fmt"

func main () {

	fmt.Println(Decode("GA-DE-RY-PO-LU-KI"))
}


func Encode(str string) string {



	return str
}

func Decode(str string) string {

	newStr := []rune(str)

	for i, v := range str {
		fmt.Println(v)
		newLetter := v
		switch newLetter {
		case 'G':
			newLetter = 'A'
		case 'A':
			newLetter = 'G'
		case 'D':
			newLetter = 'E'
		case 'E':
			newLetter = 'D'
		case 'R':
			newLetter = 'R'
		case 'Y':
			newLetter = 'Y'
		case 'P':
			newLetter = 'O'
		case 'O':
			newLetter = 'P'
		case 'L':
			newLetter = 'U'
		case 'U':
			newLetter = 'L'
		case 'K':
			newLetter = 'I'
		case 'I':
			newLetter = 'K'	
		}

		newStr[i] = newLetter
	}

	return string(newStr)
}