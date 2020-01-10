package mystring

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello Go !"
	for _, v := range strings.Fields(s) {
		fmt.Println(v)
	}
	fmt.Println(LengthOfLongestSubstring("what are you 弄啥咧？"))

	s1 := []string{"H", "e", "l", "l", "o"}
	fmt.Println(strings.Join(s1, ","))
}

func LengthOfLongestSubstring(s string) int {
	start, maxLength := 0, 0
	lastOccurred := make(map[rune]int)

	for i, v := range []rune(s) {
		if pos, ok := lastOccurred[v]; ok && pos >= start {
			start = pos + 1
		}
		lastOccurred[v] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}
	return maxLength
}
