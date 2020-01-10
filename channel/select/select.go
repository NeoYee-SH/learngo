package main

import "fmt"

func main() {

	var c1, c2 chan int

	//非阻塞
	select {
	case n := <-c1:
		fmt.Printf("Received value:%d from c1\n", n)
	case n := <-c2:
		fmt.Printf("Received value:%d from c2\n", n)
	default:
		fmt.Printf("No value received\n")
	}
}
