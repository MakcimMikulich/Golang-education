package main

// func main() {

// 	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
// }

func longestCommonPrefix(strs []string) string {
	prefix := ""

	minStr := strs[0]
	for _, str := range strs {
		if len(str) < len(minStr) {
			minStr = str
		}
	}

	for i := 0; i < len(minStr); i++ {
		letter := minStr[i]

		for j := 0; j < len(strs); j++ {

			if letter != strs[j][i] {
				return prefix
			}
		}

		prefix += string(letter)

	}

	return prefix
}
