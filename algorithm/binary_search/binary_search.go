package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 5, 6, 7, 9, 11, 13, 16, 19, 23, 27, 32, 44, 47, 56, 132, 155, 231}
	item := 56

	ret, ok := binarySearch(list, item)
	fmt.Println(ret, ok)
}

func binarySearch(list []int, item int) (int, bool) {
	start := 0
	end := len(list) - 1

	for start <= end {
		mid := (start + end) / 2
		if list[mid] > item {
			end = mid - 1
		} else if list[mid] < item {
			start = mid + 1
		} else {
			return mid, true
		}

	}
	return 0, false
}
