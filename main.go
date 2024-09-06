package main

import "fmt"

func main() {
	// fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))

	name := "Makcim"

	for _, v := range name {
		fmt.Println(rune(v))
		fmt.Println(string(v))
	}
}
