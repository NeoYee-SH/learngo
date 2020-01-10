package main

import "fmt"

func findSmallest(list []int) int {
	minIndex := 0
	min := list[minIndex]

	for key, value := range list {
		if value < min {
			min = value
			minIndex = key
		}
	}
	return minIndex
}

func selectionSort(list []int) []int {
	var ret []int
	var minIndex int
	for len(list) > 0 {
		minIndex = findSmallest(list)
		ret = append(ret, list[minIndex])
		list = append(list[:minIndex], list[minIndex+1:]...)
	}
	return ret
}

func main() {
	list := []int{3, 6, 8, 2, 4, 9, 5}
	ret := selectionSort(list)
	fmt.Println(ret)
}
