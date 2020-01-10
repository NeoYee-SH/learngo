package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	done func()
}

func doWorker(i int, w worker)  {
	for n :=range w.in {
		fmt.Printf("worker %d get %c\n", i, n)
		w.done()
	}
}

func createWorker(i int, wg *sync.WaitGroup)worker  {
	w := worker{
		in:   make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(i, w)
	return w
}

func channelDemo()  {

	var wg sync.WaitGroup
	var workers  [10]worker
	for i:=0; i<10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i,worker := range workers{
		worker.in <- 'a' + i
	}

	for i,worker := range workers{
		worker.in <- 'A' + i
	}

	wg.Wait()

}
func main() {
	channelDemo()
}
