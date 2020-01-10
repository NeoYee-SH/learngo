package main

import "fmt"

/*func quickSort(list []int) {

	if len(list) < 2 {
		return
	}
	pivot := list[0]
	head, end := 0, len(list)-1

	for i := 1; i <= end; {
		if list[i] > pivot {
			list[i], list[end] = list[end], list[i]
			end--
		} else {
			list[i], list[head] = list[head], list[i]
			head++
			i++
		}
	}
	quickSort(list[:head])
	quickSort(list[head+1:])
	return

}*/

func quickSort(list []int) {

	if len(list) < 2 {
		return
	}
	head, end := 0, len(list)-1
	pivot := list[0]
	for i := 1; head < end; {
		if list[i] > pivot {
			list[i], list[end] = list[end], list[i]
			end--
		} else {
			list[i], list[head] = list[head], list[i]
			head++
			i++
		}
	}
	quickSort(list[:head])
	quickSort(list[head+1:])
}
func main() {
	list := []int{6, 1, 4, 5, 9, 7, 3, 54, 21, 12, 23}
	quickSort(list)
	fmt.Println(list)
}
