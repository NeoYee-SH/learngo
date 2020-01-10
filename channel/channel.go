package main

import (
	"fmt"
	"time"
)

func worker(i int, c chan int)  {
	//for {
	//	if n,ok := <-c; ok {
	//		fmt.Printf("worker %d get %c\n", i, n)
	//	}
	//}
	for n :=range c {
		fmt.Printf("worker %d get %c\n", i, n)
	}
}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go worker(i, c)
	return c
}
func channelDemo()  {
	var channel [10]chan<- int

	for i:=0; i<10; i++ {
		channel[i] = createWorker(i)
	}

	for i:=0; i<10; i++ {
		channel[i] <- 'a' + i
	}

	for i:=0; i<10; i++ {
		channel[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)//buffer为3
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	c <- 'f'
	c <- 'g'
	c <- 'h'
	close(c)
}

func main() {
	channelDemo()
	//bufferedChannel()
	time.Sleep(time.Millisecond)
}
