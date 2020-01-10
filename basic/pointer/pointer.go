package main

import "fmt"

func p(pa *string) {
	*pa = "Hello Go !"
}
func main() {
	a := "Hello World !"
	fmt.Println(a) //Hello World !
	p(&a)
	fmt.Println(a) //Hello Go !
}
