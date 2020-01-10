package main

import "fmt"

func resetArr(arr [5]int) {
	arr[0] = 100
}

func resetParr(parr *[5]int) {
	parr[0] = 100
}
func main() {

	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{1, 2, 3, 4}

	fmt.Println(arr1, arr2, arr3)

	for i := 0; i < len(arr2); i++ {
		println(arr2[i])
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println(arr2)
	resetArr(arr2)
	fmt.Println(arr2)
	resetParr(&arr2)
	fmt.Println(arr2)

}
