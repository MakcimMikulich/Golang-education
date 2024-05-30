package main

// func main() {

// 	fmt.Println(strStr("sadbutsad", "sad"))

// }

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	for index, v := range haystack {
		if v == rune(needle[0]) {

			if len(needle) == 1 {
				return index
			}

			for i := 1; i < len(needle); i++ {

				if len(haystack)-index < len(needle) {
					return -1
				}

				if haystack[index+i] != needle[i] {
					break
				}
				if i == len(needle)-1 {

					return index

				}

			}

		}
	}

	return -1
}
