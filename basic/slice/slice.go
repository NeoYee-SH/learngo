package main

import "fmt"

func resetSlice(s []int) {
	s[0] = 100
}
func printSlice(s []int) {
	fmt.Printf("%v, len(s)=%d,cap(s)=%d\n", s, len(s), cap(s))
}

func main() {

	//arr := [...]int{0,1,2,3,4,5,6,7}
	//fmt.Println(arr) //[0 1 2 3 4 5 6 7]
	//s1 := arr[2:6]//2,3,4,5
	//resetSlice(s1)
	//fmt.Println(arr)//[0 1 100 3 4 5 6 7]
	//s2 := arr[3:5]//3,4
	//resetSlice(s2)
	//fmt.Println(arr)//[0 1 100 100 4 5 6 7]
	//
	//arr1 := [...]int{0,1,2,3,4,5,6,7}
	//ss1 := arr1[3:5]//3,4[,5,6,7]
	//ss2 := ss1[3:5]//6,7
	//fmt.Println(ss1)//3,4
	//fmt.Println(ss2)//6,7

	var s1 []int
	for i := 0; i < 20; i++ {
		s1 = append(s1, i*8)
	}
	s2 := []int{1, 3, 4, 6, 8}
	s3 := make([]int, 16)
	s4 := make([]int, 10, 16)

	printSlice(s1) //[0 8 16 24 32 40 48 56 64 72 80 88 96 104 112 120 128 136 144 152], len(s)=20,cap(s)=32
	printSlice(s2) //[1 3 4 6 8], len(s)=5,cap(s)=5
	printSlice(s3) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len(s)=16,cap(s)=16
	printSlice(s4) //[0 0 0 0 0 0 0 0 0 0], len(s)=10,cap(s)=16

	copy(s1, s2)
	printSlice(s1) //[1 3 4 6 8 40 48 56 64 72 80 88 96 104 112 120 128 136 144 152], len(s)=20,cap(s)=32
	copy(s2, s3)
	printSlice(s2) //[0 0 0 0 0], len(s)=5,cap(s)=5

	//删除s1的第5个元素
	s1 = append(s1[:4], s1[5:]...)
	printSlice(s1) //[1 3 4 6 40 48 56 64 72 80 88 96 104 112 120 128 136 144 152], len(s)=19,cap(s)=32

	front := s1[0]
	s1 = s1[1:]
	fmt.Println("front is", front) //front is 1
	printSlice(s1)                 //[3 4 6 40 48 56 64 72 80 88 96 104 112 120 128 136 144 152], len(s)=18,cap(s)=31

	back := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	fmt.Println("back is", back) //back is 152
	printSlice(s1)               //[3 4 6 40 48 56 64 72 80 88 96 104 112 120 128 136 144], len(s)=17,cap(s)=31

}
