package main

import (
	"fmt"
	"unicode/utf8"
)

func lengthOfLongestSubstring(s string) int {

	fmt.Println("count char=", utf8.RuneCountInString(s))
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

func main() {

	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("bbbbb"))
	//fmt.Println(lengthOfLongestSubstring("pwwkew"))
	//fmt.Println(lengthOfLongestSubstring("hello"))
	//fmt.Println(lengthOfLongestSubstring("yihuaiyuan"))
	fmt.Println(lengthOfLongestSubstring("你好4好3兄ac啊!好兄弟"))

	/*m1 := map[string]string {
		"name":"Bob",
		"sex":"man",
		"country":"china",
	}
	m2 := make(map[string]int)
	var m3 map[string]string
	fmt.Println(m1)//map[country:china name:Bob sex:man]
	fmt.Println(m2)//map[]
	fmt.Println(m3)//map[]

	for i,v := range m1 {
		fmt.Println(i, v)
	}

	sex := m1["sex"]
	fmt.Println(sex) // man
	school := m1["school"]
	fmt.Println(school)// ''

	if name, ok := m1["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("name not exist")
	}
	if class, ok := m1["class"]; ok {
		fmt.Println(class)
	} else {
		fmt.Println("class not exist")
	}

	delete(m1, "name")
	if name, ok := m1["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("name not exist")
	}

	fmt.Println(len(m1))*/

}
