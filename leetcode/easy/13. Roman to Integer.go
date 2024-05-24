package main

// func main() {

// 	fmt.Println(romanToInt("MCMXCIV"))
// }

func romanToInt(s string) int {

	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0

	len := len(s)

	for i := len - 1; i != -1; i-- {

		currentEl := m[string(s[i])]

		if i == 0 {
			result += currentEl
			break
		}

		nextEl := m[string(s[i-1])]

		if currentEl > nextEl {
			result += currentEl - nextEl
			i--
		} else {
			result += currentEl
		}
	}

	return result
}
